package moonshot

import (
	"encoding/json"

	"github.com/eolinker/apinto/convert"
	ai_provider "github.com/eolinker/apinto/drivers/ai-provider"
	"github.com/eolinker/eosc"
	"github.com/eolinker/eosc/eocontext"
	http_context "github.com/eolinker/eosc/eocontext/http-context"
)

var (
	modelModes = map[string]IModelMode{
		ai_provider.ModeChat.String(): NewChat(),
	}
)

type IModelMode interface {
	Endpoint() string
	convert.IConverter
}

type Chat struct {
	endPoint string
}

func NewChat() *Chat {
	return &Chat{
		endPoint: "/v1/chat/completions",
	}
}

func (c *Chat) Endpoint() string {
	return c.endPoint
}

func (c *Chat) RequestConvert(ctx eocontext.EoContext, extender map[string]interface{}) error {
	httpContext, err := http_context.Assert(ctx)
	if err != nil {
		return err
	}
	body, err := httpContext.Proxy().Body().RawBody()
	if err != nil {
		return err
	}
	// 设置转发地址
	httpContext.Proxy().URI().SetPath(c.endPoint)
	baseCfg := eosc.NewBase[ai_provider.ClientRequest]()
	err = json.Unmarshal(body, baseCfg)
	if err != nil {
		return err
	}
	messages := make([]Message, 0, len(baseCfg.Config.Messages)+1)
	for _, m := range baseCfg.Config.Messages {
		messages = append(messages, Message{
			Role:    m.Role,
			Content: m.Content,
		})
	}
	baseCfg.SetAppend("messages", messages)
	for k, v := range extender {
		baseCfg.SetAppend(k, v)
	}
	body, err = json.Marshal(baseCfg)
	if err != nil {
		return err
	}
	httpContext.Proxy().Body().SetRaw("application/json", body)
	return nil
}

func (c *Chat) ResponseConvert(ctx eocontext.EoContext) error {
	httpContext, err := http_context.Assert(ctx)
	if err != nil {
		return err
	}
	body := httpContext.Response().GetBody()
	data := eosc.NewBase[Response]()
	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	// HTTP Status Codes for Moonshot API
	// Status Code | Type                | Error Message
	// ------------|---------------------|-------------------------------------
	// 200         | Success             | Request was successful.
	// 400         | Client Error        | Invalid request parameters (invalid_request_error).
	// 401         | Authentication Error | Invalid API key (invalid_key).
	// 403         | Forbidden           | Access denied (forbidden_error).
	// 404         | Not Found           | Resource not found (not_found_error).
	// 429         | Rate Limit Exceeded | Too many requests (rate_limit_error).
	// 500         | Server Error        | Internal server error (server_error).
	// 503         | Service Unavailable  | Service is temporarily unavailable (service_unavailable).
	switch httpContext.Response().StatusCode() {
	case 200:
		// Calculate the token consumption for a successful request.
		usage := data.Config.Usage
		ai_provider.SetAIStatusNormal(ctx)
		ai_provider.SetAIModelInputToken(ctx, usage.PromptTokens)
		ai_provider.SetAIModelOutputToken(ctx, usage.CompletionTokens)
		ai_provider.SetAIModelTotalToken(ctx, usage.TotalTokens)
	case 400:
		// Handle the bad request error.
		ai_provider.SetAIStatusInvalidRequest(ctx)
	case 401:
		// 过期和无效的API密钥
		ai_provider.SetAIStatusInvalid(ctx)
	case 429:
		switch data.Config.Error.Type {
		case "exceeded_current_quota_error":
			// Handle the insufficient quota error.
			ai_provider.SetAIStatusQuotaExhausted(ctx)
		case "engine_overloaded_error", "rate_limit_reached_error":
			// Handle the rate limit error.
			ai_provider.SetAIStatusExceeded(ctx)
		}
	default:
		ai_provider.SetAIStatusInvalidRequest(ctx)
	}

	responseBody := &ai_provider.ClientResponse{}
	if len(data.Config.Choices) > 0 {
		msg := data.Config.Choices[0]
		responseBody.Message = ai_provider.Message{
			Role:    msg.Message.Role,
			Content: msg.Message.Content,
		}
		responseBody.FinishReason = msg.FinishReason
	} else {
		responseBody.Code = -1
		responseBody.Error = data.Config.Error.Message
	}
	body, err = json.Marshal(responseBody)
	if err != nil {
		return err
	}
	httpContext.Response().SetBody(body)
	return nil
}
