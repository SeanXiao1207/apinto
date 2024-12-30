package novita

import (
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/eolinker/eosc/eocontext"

	"github.com/eolinker/apinto/convert"
	ai_provider "github.com/eolinker/apinto/drivers/ai-provider"
	http_context "github.com/eolinker/apinto/node/http-context"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

var (
	defaultConfig = `{
		"frequency_penalty": "",
		"max_tokens": 512,
		"presence_penalty": "",
		"response_format": "",
		"temperature": "",
		"top_p": "0.1"
	}`
	successBody = []byte(`{
		"messages": [
			{
				"content": "Hello, how can I help you?",
				"role": "user"
			}
		]
	}`)
	failBody = []byte(`{
		"messages": [
			{
				"content": "Hello, how can I help you?",
				"role": "yyy"
			}
		],"top_p":"0.0"
	}`)
)

func validNormalFunc(ctx eocontext.EoContext) bool {
	fmt.Printf("input token: %d\n", ai_provider.GetAIModelInputToken(ctx))
	fmt.Printf("output token: %d\n", ai_provider.GetAIModelOutputToken(ctx))
	fmt.Printf("total token: %d\n", ai_provider.GetAIModelTotalToken(ctx))
	if ai_provider.GetAIModelInputToken(ctx) <= 0 {
		return false
	}
	if ai_provider.GetAIModelOutputToken(ctx) <= 0 {
		return false
	}
	return ai_provider.GetAIModelTotalToken(ctx) > 0
}

// TestSentTo tests the end-to-end execution of the novita integration.
func TestSentTo(t *testing.T) {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Test data for different scenarios
	testData := []struct {
		name       string
		apiKey     string
		wantStatus string
		body       []byte
		validFunc  func(ctx eocontext.EoContext) bool
	}{
		{
			name:       "success",
			apiKey:     os.Getenv("ValidKey"),
			wantStatus: ai_provider.StatusNormal,
			body:       successBody,
			validFunc:  validNormalFunc,
		},
		{
			name:       "invalid request",
			apiKey:     os.Getenv("ValidKey"),
			wantStatus: ai_provider.StatusInvalidRequest,
			body:       failBody,
		},
		{
			name:       "invalid key",
			apiKey:     os.Getenv("InvalidKey"),
			wantStatus: ai_provider.StatusInvalid,
		},
		{
			name:       "expired key",
			apiKey:     os.Getenv("ExpiredKey"),
			wantStatus: ai_provider.StatusInvalid,
		},
	}

	// Run tests for each scenario
	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			if err := runTest(data.apiKey, data.body, data.wantStatus, data.validFunc); err != nil {
				t.Fatalf("Test failed: %v", err)
			}
		})
	}
}

// runTest handles a single test case
func runTest(apiKey string, requestBody []byte, wantStatus string, validFunc func(ctx eocontext.EoContext) bool) error {
	cfg := &Config{
		APIKey:       apiKey,
		Organization: "",
	}

	// Create the worker
	worker, err := Create("novita", "novita", cfg, nil)
	if err != nil {
		return fmt.Errorf("failed to create worker: %w", err)
	}

	// Get the handler
	handler, ok := worker.(convert.IConverterDriver)
	if !ok {
		return fmt.Errorf("worker does not implement IConverterDriver")
	}

	// Default to success body if no body is provided
	if len(requestBody) == 0 {
		requestBody = successBody
	}

	// Mock HTTP context
	ctx := createMockHttpContext("/xxx/xxx", nil, nil, requestBody)

	// Execute the conversion process
	err = executeConverter(ctx, handler, "meta-llama/llama-3.1-8b-instruct", "https://api.novita.ai")
	if err != nil {
		return fmt.Errorf("failed to execute conversion process: %w", err)
	}

	// Check the status
	if ai_provider.GetAIStatus(ctx) != wantStatus {
		return fmt.Errorf("unexpected status: got %s, expected %s", ai_provider.GetAIStatus(ctx), wantStatus)
	}
	if validFunc != nil {
		if validFunc(ctx) {
			return nil
		}
		return fmt.Errorf("execute validFunc failed")
	}

	return nil
}

// executeConverter handles the full flow of a conversion process.
func executeConverter(ctx *http_context.HttpContext, handler convert.IConverterDriver, model string, baseUrl string) error {
	// Balance handler setup
	balanceHandler, err := ai_provider.NewBalanceHandler("test", baseUrl, 30*time.Second)
	if err != nil {
		return fmt.Errorf("failed to create balance handler: %w", err)
	}

	// Get model function
	fn, has := handler.GetModel(model)
	if !has {
		return fmt.Errorf("model %s not found", model)
	}

	// Generate config
	extender, err := fn(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to generate config: %w", err)
	}

	// Get converter
	converter, has := handler.GetConverter(model)
	if !has {
		return fmt.Errorf("converter for model %s not found", model)
	}

	// Convert request
	if err := converter.RequestConvert(ctx, extender); err != nil {
		return fmt.Errorf("request conversion failed: %w", err)
	}

	// Select node via balance handler
	node, _, err := balanceHandler.Select(ctx)
	if err != nil {
		return fmt.Errorf("node selection failed: %w", err)
	}

	// Send request to the node
	if err := ctx.SendTo(balanceHandler.Scheme(), node, balanceHandler.TimeOut()); err != nil {
		return fmt.Errorf("failed to send request to node: %w", err)
	}

	// Convert response
	if err := converter.ResponseConvert(ctx); err != nil {
		return fmt.Errorf("response conversion failed: %w", err)
	}

	return nil
}

// createMockHttpContext creates a mock fasthttp.RequestCtx and wraps it with HttpContext.
func createMockHttpContext(rawURL string, headers map[string]string, query url.Values, body []byte) *http_context.HttpContext {
	req := fasthttp.AcquireRequest()
	u := fasthttp.AcquireURI()

	// Set request URI and path
	uri, _ := url.Parse(rawURL)
	u.SetPath(uri.Path)
	u.SetScheme(uri.Scheme)
	u.SetHost(uri.Host)
	u.SetQueryString(uri.RawQuery)
	req.SetURI(u)
	req.Header.SetMethod("POST")

	// Set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.SetBody(body)

	// Create HttpContext
	return http_context.NewContext(&fasthttp.RequestCtx{
		Request:  *req,
		Response: fasthttp.Response{},
	}, 8099)
}
