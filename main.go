package main

import (
	"context"
	"os"
	"time"

	"github.com/truescotian/oapi-codegen-example/internal/cmd"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ret := cmd.Execute(ctx)
	os.Exit(ret)
}
