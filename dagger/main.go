package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		log.Fatalf("failed to connect to dagger: %v", err)
	}
	defer client.Close()

	container := client.Container().
		From("alpine:latest").
		WithExec([]string{"echo", "Hola Mundo"})

	out, err := container.Stdout(ctx)
	if err != nil {
		log.Fatalf("failed to execute container: %v", err)
	}

	fmt.Println(out)
}
