package main

import (
	"context"

	"github.com/google/go-github/v72/github"
)

type contextKey string

var clientKey = contextKey("client")

// WithClient returns a new context with the GitHub client set.
func WithClient(ctx context.Context, client *github.Client) context.Context {
	return context.WithValue(ctx, clientKey, client)
}

// ClientFromContext retrieves the GitHub client from the context.
func ClientFromContext(ctx context.Context) (*github.Client, bool) {
	client, ok := ctx.Value(clientKey).(*github.Client)
	return client, ok
}
