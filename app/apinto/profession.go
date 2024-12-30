/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package main

import (
	"github.com/eolinker/eosc"
)

func ApintoProfession() []*eosc.ProfessionConfig {
	return []*eosc.ProfessionConfig{
		{
			Name:         "router",
			Label:        "路由",
			Desc:         "路由",
			Dependencies: []string{"service", "template", "transcode", "ai-provider"},
			AppendLabels: []string{"host", "service", "listen", "disable"},
			Drivers: []*eosc.DriverConfig{
				{
					Id:     "eolinker.com:apinto:http_router",
					Name:   "http",
					Label:  "http",
					Desc:   "http路由",
					Params: nil,
				},
				{
					Id:     "eolinker.com:apinto:grpc_router",
					Name:   "grpc",
					Label:  "grpc",
					Desc:   "grpc路由",
					Params: nil,
				},
				{
					Id:     "eolinker.com:apinto:dubbo2_router",
					Name:   "dubbo2",
					Label:  "dubbo2",
					Desc:   "dubbo2路由",
					Params: nil,
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "service",
			Label:        "服务",
			Desc:         "服务",
			Dependencies: []string{"discovery"},
			AppendLabels: []string{"discovery"},
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:service_http",
					Name:  "http",
					Label: "service",
					Desc:  "服务",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "template",
			Label:        "模版",
			Desc:         "模版",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:plugin_template",
					Name:  "plugin_template",
					Label: "插件模版",
					Desc:  "插件模版",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "discovery",
			Label:        "注册中心",
			Desc:         "注册中心",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:discovery_static",
					Name:  "static",
					Label: "静态服务发现",
					Desc:  "静态服务发现",
				}, {
					Id:    "eolinker.com:apinto:discovery_nacos",
					Name:  "nacos",
					Label: "nacos服务发现",
					Desc:  "nacos服务发现",
				}, {
					Id:    "eolinker.com:apinto:discovery_consul",
					Name:  "consul",
					Label: "consul服务发现",
					Desc:  "consul服务发现",
				}, {
					Id:    "eolinker.com:apinto:discovery_eureka",
					Name:  "eureka",
					Label: "eureka服务发现",
					Desc:  "eureka服务发现",
				}, {
					Id:    "eolinker.com:apinto:discovery_polaris",
					Name:  "polaris",
					Label: "北极星服务发现",
					Desc:  "北极星服务发现",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "app",
			Label:        "应用",
			Desc:         "应用",
			Dependencies: nil,
			AppendLabels: []string{"disable"},
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:app",
					Name:  "app",
					Label: "应用",
					Desc:  "应用",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		}, {
			Name:         "strategy",
			Label:        "策略",
			Desc:         "策略",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:strategy-limiting",
					Name:  "limiting",
					Label: "限流策略",
					Desc:  "限流策略",
				},
				{
					Id:    "eolinker.com:apinto:strategy-cache",
					Name:  "cache",
					Label: "缓存策略",
					Desc:  "缓存策略",
				},
				{
					Id:    "eolinker.com:apinto:strategy-grey",
					Name:  "grey",
					Label: "灰度策略",
					Desc:  "灰度策略",
				},
				{
					Id:    "eolinker.com:apinto:strategy-visit",
					Name:  "visit",
					Label: "访问策略",
					Desc:  "访问策略",
				},
				{
					Id:    "eolinker.com:apinto:strategy-fuse",
					Name:  "fuse",
					Label: "熔断策略",
					Desc:  "熔断策略",
				},
				{
					Id:    "eolinker.com:apinto:strategy-data_mask",
					Name:  "data_mask",
					Label: "数据脱敏策略",
					Desc:  "数据脱敏策略",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "counter",
			Label:        "计数器",
			Desc:         "计数器",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:http-counter",
					Name:  "http",
					Label: "http计数器",
					Desc:  "http计数器",
				},
				{
					Id:    "eolinker.com:apinto:nsq-counter",
					Name:  "nsq",
					Label: "nsq计数器",
					Desc:  "nsq计数器",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "output",
			Label:        "输出",
			Desc:         "输出",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:file_output",
					Name:  "file",
					Label: "文件输出",
					Desc:  "文件输出",
				},
				{
					Id:    "eolinker.com:apinto:nsqd",
					Name:  "nsqd",
					Label: "NSQ输出",
					Desc:  "NSQ输出",
				},
				{
					Id:    "eolinker.com:apinto:http_output",
					Name:  "http_output",
					Label: "http输出",
					Desc:  "http输出",
				},
				{
					Id:    "eolinker.com:apinto:syslog_output",
					Name:  "syslog_output",
					Label: "syslog输出",
					Desc:  "syslog输出",
				},
				{
					Id:    "eolinker.com:apinto:kafka_output",
					Name:  "kafka_output",
					Label: "kafka输出",
					Desc:  "kafka输出",
				},
				{
					Id:    "eolinker.com:apinto:redis",
					Name:  "redis",
					Label: "redis 集群",
					Desc:  "redis 集群",
				},
				{
					Id:    "eolinker.com:apinto:influxdbv2",
					Name:  "influxdbv2",
					Label: "influxdbv2输出",
					Desc:  "influxdbv2输出",
				},
				{
					Id:    "eolinker.com:apinto:prometheus_output",
					Name:  "prometheus",
					Label: "prometheus输出",
					Desc:  "prometheus输出",
				},
				{
					Id:    "eolinker.com:apinto:loki_output",
					Name:  "loki",
					Label: "loki输出",
					Desc:  "loki输出",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},

		{
			Name:         "certificate",
			Label:        "证书",
			Desc:         "证书",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:ssl-server",
					Name:  "server",
					Label: "证书",
					Desc:  "证书",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "transcode",
			Label:        "编码器",
			Desc:         "编码器",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:protobuf_transcode",
					Name:  "protobuf",
					Label: "protobuf编码器",
					Desc:  "protobuf编码器",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
		{
			Name:         "ai-provider",
			Label:        "AI服务提供者",
			Desc:         "AI服务提供者",
			Dependencies: nil,
			AppendLabels: nil,
			Drivers: []*eosc.DriverConfig{
				{
					Id:    "eolinker.com:apinto:openai",
					Name:  "openai",
					Label: "openAI",
					Desc:  "openAI",
				},
				{
					Id:    "eolinker.com:apinto:google", // 插件ID
					Name:  "google",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Google",
					Desc:  "Google",
				},
				{
					Id:    "eolinker.com:apinto:moonshot", // 插件ID
					Name:  "moonshot",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "moonshot",
					Desc:  "moonshot",
				},
				{
					Id:    "eolinker.com:apinto:tongyi", // 插件ID
					Name:  "tongyi",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "tongyi",
					Desc:  "tongyi",
				},
				{
					Id:    "eolinker.com:apinto:zhipuai", // 插件ID
					Name:  "zhipuai",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "zhipuai",
					Desc:  "zhipuai",
				},
				{
					Id:    "eolinker.com:apinto:fireworks", // 插件ID
					Name:  "fireworks",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "fireworks",
					Desc:  "fireworks",
				},
				{
					Id:    "eolinker.com:apinto:novita", // 插件ID
					Name:  "novita",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "novita",
					Desc:  "novita",
				},
				{
					Id:    "eolinker.com:apinto:mistralai", // 插件ID
					Name:  "mistralai",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "mistralai",
					Desc:  "mistralai",
				},
				{
					Id:    "eolinker.com:apinto:anthropic", // 插件ID
					Name:  "anthropic",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Anthropic",
					Desc:  "Anthropic",
				},
				{
					Id:    "eolinker.com:apinto:baichuan", // 插件ID
					Name:  "baichuan",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Baichuan",
					Desc:  "Baichuan",
				},
				{
					Id:    "eolinker.com:apinto:stepfun", // 插件ID
					Name:  "stepfun",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Stepfun",
					Desc:  "Stepfun",
				},
				{
					Id:    "eolinker.com:apinto:wenxin", // 插件ID
					Name:  "wenxin",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Wenxin",
					Desc:  "Wenxin",
				},
				{
					Id:    "eolinker.com:apinto:yi", // 插件ID
					Name:  "yi",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "yi",
					Desc:  "yi",
				},
				{
					Id:    "eolinker.com:apinto:perfxcloud", // 插件ID
					Name:  "perfxcloud",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Perfxcloud",
					Desc:  "Perfxcloud",
				},
				{
					Id:    "eolinker.com:apinto:cohere", // 插件ID
					Name:  "cohere",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Cohere",
					Desc:  "Cohere",
				},
				{
					Id:    "eolinker.com:apinto:deepseek", // 插件ID
					Name:  "deepseek",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Deepseek",
					Desc:  "Deepseek",
				},
				{
					Id:    "eolinker.com:apinto:hunyuan", // 插件ID
					Name:  "hunyuan",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Hunyuan",
					Desc:  "Hunyuan",
				},
				{
					Id:    "eolinker.com:apinto:openrouter", // 插件ID
					Name:  "openrouter",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "OpenRouter",
					Desc:  "OpenRouter",
				},
				{
					Id:    "eolinker.com:apinto:groq", // 插件ID
					Name:  "groq",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "GroqCloud",
					Desc:  "GroqCloud",
				},
				{
					Id:    "eolinker.com:apinto:upstage", // 插件ID
					Name:  "upstage",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "upstage",
					Desc:  "upstage",
				},
				{
					Id:    "eolinker.com:apinto:minimax", // 插件ID
					Name:  "minimax",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "minimax",
					Desc:  "minimax",
				},
				{
					Id:    "eolinker.com:apinto:chatglm", // 插件ID
					Name:  "chatglm",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "chatglm",
					Desc:  "chatglm",
				},
				{
					Id:    "eolinker.com:apinto:bedrock", // 插件ID
					Name:  "bedrock",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "AWS Bedrock",
					Desc:  "AWS Bedrock",
				},
				{
					Id:    "eolinker.com:apinto:spark", // 插件ID
					Name:  "spark",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "讯飞星火",
					Desc:  "讯飞星火",
				},
				{
					Id:    "eolinker.com:apinto:nvidia", // 插件ID
					Name:  "nvidia",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Nvidia",
					Desc:  "Nvidia",
				},
				{
					Id:    "eolinker.com:apinto:vertex_ai", // 插件ID
					Name:  "vertex_ai",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Vertex AI",
					Desc:  "Vertex AI",
				},
				{
					Id:    "eolinker.com:apinto:fakegpt", // 插件ID
					Name:  "fakegpt",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "Fake GPT",
					Desc:  "Fake GPT",
				}, {
					Id:    "eolinker.com:apinto:zhinao", // 插件ID
					Name:  "zhinao",                     // 驱动名称，应和定义文件的provider字段一致
					Label: "zhinao",
					Desc:  "zhinao",
				},
			},
			Mod: eosc.ProfessionConfig_Worker,
		},
	}
}
