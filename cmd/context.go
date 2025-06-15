package cmd

import (
	"context"

	"github.com/google/go-github/v72/github"
)

type contextKey string

const clientKey contextKey = "client"

func getClientFromContext(ctx context.Context) (*github.Client, bool) {
	client, ok := ctx.Value(clientKey).(*github.Client)
	return client, ok
}
