package openai

import (
	"context"
	"fmt"
	"os"

	gpt "github.com/sashabaranov/go-openai"
)

var client *gpt.Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("Missing OPENAI_API_KEY")
	}
	client = gpt.NewClient(apiKey)
}

func Generate(raw, product, audience, tone string) (string, error) {

	prompt := fmt.Sprintf(`
You are an expert B2B content strategist at Forj Media. 
Turn the following raw messy founder notes into:

1) 5 viral LinkedIn hooks
2) 3 post outlines
3) 2 full posts

Ensure the tone matches: %s.
Product: %s.
Audience: %s.

Return ONLY valid JSON with this exact schema:

{
  "hooks": [...],
  "postOutlines": [...],
  "fullPosts": [...]
}

Raw notes:
%s
`, tone, product, audience, raw)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		gpt.ChatCompletionRequest{
			Model: gpt.GPT4oMini, // fast and cheap
			Messages: []gpt.ChatCompletionMessage{
				{Role: "user", Content: prompt},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
