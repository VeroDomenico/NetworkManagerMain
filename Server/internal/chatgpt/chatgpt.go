package chatgpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func generateHappyBirthday(name string, date string) (string, error) {
	client := openai.NewClient("your token")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "write a happy birthday message to" + name + "given the birthdate of" + date,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
