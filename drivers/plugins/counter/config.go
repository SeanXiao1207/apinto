package counter

import (
	"github.com/eolinker/apinto/drivers"
	"github.com/eolinker/apinto/drivers/plugins/counter/matcher"
	"github.com/eolinker/apinto/drivers/plugins/counter/separator"
	"github.com/eolinker/eosc"
)

type Config struct {
	Key   string         `json:"key" label:"格式化Key" required:"true"`
	Cache eosc.RequireId `json:"cache" label:"缓存计数器" skill:"github.com/eolinker/apinto/resources.resources.ICache" required:"false"`
	//Counter     eosc.RequireId       `json:"counter" label:"计数器" skill:"github.com/eolinker/apinto/drivers/counter.counter.IClient" required:"false"`
	//CountPusher eosc.RequireId       `json:"counterPusher" label:"计数推送器" skill:"github.com/eolinker/apinto/drivers/counter.counter.ICountPusher" required:"false"`
	Match Match                `json:"match" label:"响应匹配规则"`
	Count *separator.CountRule `json:"count" label:"计数规则"`
	//CountMode   string               `json:"count_mode" label:"计数模式" enum:"local,redis"`
}

type Match struct {
	Params      []*matcher.MatchParam `json:"params" label:"匹配参数列表"`
	StatusCodes []int                 `json:"status_codes" label:"匹配响应状态码列表"`
	Type        string                `json:"type" label:"匹配类型" enum:"json"`
}

func (m *Match) GenerateHandler() []matcher.IMatcher {
	return []matcher.IMatcher{
		matcher.NewStatusCodeMatcher(m.StatusCodes),
		matcher.NewJsonMatcher(m.Params),
	}
}

func Create(id, name string, conf *Config, workers map[eosc.RequireId]eosc.IWorker) (eosc.IWorker, error) {
	ct, err := separator.GetCounter(conf.Count)
	if err != nil {
		return nil, err
	}
	bc := &executor{
		WorkerBase:       drivers.Worker(id, name),
		matchers:         conf.Match.GenerateHandler(),
		separatorCounter: ct,
		counters:         eosc.BuildUntyped[string, ICounter](),
		keyGenerate:      newKeyGenerate(conf.Key),
		cacheID:          string(conf.Cache),
	}

	return bc, nil
}
