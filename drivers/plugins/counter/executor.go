package counter

import (
	"fmt"
	"reflect"
	"sync"

	scope_manager "github.com/eolinker/apinto/scope-manager"

	"github.com/eolinker/apinto/drivers/plugins/counter/matcher"

	"github.com/eolinker/apinto/resources"

	"github.com/eolinker/apinto/drivers/counter"

	"github.com/eolinker/apinto/drivers"
	"github.com/eolinker/apinto/drivers/plugins/counter/separator"

	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/eocontext"
	http_service "github.com/eolinker/eosc/eocontext/http-context"
)

var _ http_service.HttpFilter = (*executor)(nil)
var _ eocontext.IFilter = (*executor)(nil)

type executor struct {
	drivers.WorkerBase
	matchers         []matcher.IMatcher
	separatorCounter separator.ICounter
	counters         eosc.Untyped[string, counter.ICounter]
	cache            scope_manager.IProxyOutput[resources.ICache]
	client           scope_manager.IProxyOutput[counter.IClient]
	keyGenerate      IKeyGenerator
	once             sync.Once
}

func (b *executor) DoFilter(ctx eocontext.EoContext, next eocontext.IChain) (err error) {
	return http_service.DoHttpFilter(b, ctx, next)
}

func (b *executor) DoHttpFilter(ctx http_service.IHttpContext, next eocontext.IChain) error {
	b.once.Do(func() {
		b.cache = scope_manager.Auto[resources.ICache]("", "redis")
		b.client = scope_manager.Auto[counter.IClient]("", "counter")
	})

	key := b.keyGenerate.Key(ctx)
	ct, has := b.counters.Get(key)
	if !has {
		ct = NewRedisCounter(key, b.cache, b.client)
		b.counters.Set(key, ct)
	}
	var count int64 = 1
	var err error
	if !reflect.ValueOf(b.separatorCounter).IsNil() {
		separatorCounter := b.separatorCounter
		count, err = separatorCounter.Count(ctx)
		if err != nil {
			ctx.Response().SetStatus(400, "400")
			return fmt.Errorf("%s count error", separatorCounter.Name())
		}
		if count > separatorCounter.Max() {
			ctx.Response().SetStatus(403, "not allow")
			return fmt.Errorf("%s number exceed", separatorCounter.Name())
		} else if count == 0 {
			ctx.Response().SetStatus(400, "400")
			return fmt.Errorf("%s value is missing", separatorCounter.Name())
		}
	}

	err = ct.Lock(count)
	if err != nil {
		// 次数不足，直接返回
		return err
	}
	if next != nil {
		err = next.DoChain(ctx)
		if err != nil {
			// 转发失败，回滚次数
			return ct.RollBack(count)
			//return err
		}
	}
	match := true
	for _, matcher := range b.matchers {
		ok := matcher.Match(ctx)
		if !ok {
			match = false
			break
		}
	}
	if match {
		// 匹配，扣减次数
		return ct.Complete(count)
	}
	// 不匹配，回滚次数
	return ct.RollBack(count)
}

func (b *executor) Start() error {
	return nil
}

func (b *executor) Reset(conf interface{}, workers map[eosc.RequireId]eosc.IWorker) error {
	cfg, ok := conf.(*Config)
	if !ok {
		return fmt.Errorf("invalid config, driver: %s", Name)
	}
	ct, err := separator.GetCounter(cfg.Count)
	if err != nil {
		return err
	}
	b.keyGenerate = newKeyGenerate(cfg.Key)
	b.separatorCounter = ct
	b.matchers = cfg.Match.GenerateHandler()
	return nil
}

func (b *executor) Stop() error {
	return nil
}

func (b *executor) Destroy() {
	return
}

func (b *executor) CheckSkill(skill string) bool {
	return http_service.FilterSkillName == skill
}
