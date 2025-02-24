package anthropic

import (
	"github.com/pzierahn/chatbot_services/llm"
	"strings"
)

const (
	ClaudeSonnet37 = "anthropic.claude-3-7-sonnet-20250219-v1:0"
	ClaudeSonnet35 = "anthropic.claude-3-5-sonnet-20241022-v2:0"
	ClaudeHaiku    = "anthropic.claude-3-haiku-20240307-v1:0"
)

var ModelCosts = map[string]llm.PricePer1000Tokens{
	"claude-v2": {
		Input:  0.008,
		Output: 0.024,
	},
	"claude-3-sonnet-28k-20240229": {
		Input:  0.003,
		Output: 0.015,
	},
	"claude-3-sonnet-200k-20240229": {
		Input:  0.003,
		Output: 0.015,
	},
	"claude-3-5-sonnet-20240620": {
		Input:  0.003,
		Output: 0.015,
	},
	"anthropic.claude-3-5-sonnet-20240620-v1:0": {
		Input:  0.003,
		Output: 0.015,
	},
	"anthropic.claude-3-5-sonnet-20241022-v2:0": {
		Input:  0.003,
		Output: 0.015,
	},
	"anthropic.claude-3-7-sonnet-20250219-v1:0": {
		Input:  0.003,
		Output: 0.015,
	},
	"claude-3-haiku-48k-20240307": {
		Input:  0.00025,
		Output: 0.00125,
	},
	"claude-3-opus-20240229": {
		Input:  0.01500,
		Output: 0.07500,
	},
}

func (client *Client) ProvidesModel(name string) bool {
	switch {
	case strings.HasPrefix(name, "anthropic."):
		return true
	default:
		return false
	}
}
