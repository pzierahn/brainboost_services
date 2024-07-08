package openai

import (
	"context"
	"github.com/pzierahn/chatbot_services/llm"
	"github.com/sashabaranov/go-openai"
)

func (client *Client) CreateEmbedding(ctx context.Context, req *llm.EmbeddingRequest) (*llm.EmbeddingResponse, error) {
	resp, err := client.client.CreateEmbeddings(
		ctx,
		openai.EmbeddingRequestStrings{
			Model: client.EmbeddingModel,
			Input: req.Inputs,
			User:  req.UserId,
		},
	)
	if err != nil {
		return nil, err
	}

	results := &llm.EmbeddingResponse{
		Embeddings: make([][]float32, len(resp.Data)),
		Tokens:     uint32(resp.Usage.PromptTokens),
	}

	for idx, item := range resp.Data {
		results.Embeddings[idx] = item.Embedding
	}

	return results, nil
}

func (client *Client) GetEmbeddingDimension() int {
	switch client.EmbeddingModel {
	case LargeEmbedding3:
		return DimensionModelLarge
	case SmallEmbedding3:
		return DimensionModelSmall
	default:
		return 0
	}
}
