package main

import (
	"context"

	"github.com/google/go-github/v72/github"
)

type contextKey string

var clientKey = contextKey("client")

// withClient returns a new context with the GitHub client set.
func withClient(ctx context.Context, client *github.Client) context.Context {
	return context.WithValue(ctx, clientKey, client)
}

// clientFromContext retrieves the GitHub client from the context.
func clientFromContext(ctx context.Context) (*github.Client, bool) {
	client, ok := ctx.Value(clientKey).(*github.Client)
	return client, ok
}
