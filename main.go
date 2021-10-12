package main

import (
	"context"
	"fmt"

	"webscrap/funcao"
)

// "github.com/aws/aws-lambda-go/lambda"

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	return fmt.Sprintf("Hello %s!", name.Name), nil
}

func main() {
	//lambda.Start(HandleRequest)
	funcao.Teste()
	funcao.Testando_pacote()
}
