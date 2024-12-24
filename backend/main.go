package main

import (
	"context"
	"fmt"

	"github.com/Alkestian/api-practice/app"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("error starting server:", err)
	}
}