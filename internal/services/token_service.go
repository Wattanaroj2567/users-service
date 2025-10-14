package services

import "context"

// TokenService abstracts JWT generation and validation.
type TokenService interface {
	Generate(ctx context.Context, subject string, claims map[string]any) (string, error)
	Invalidate(ctx context.Context, token string) error
}

// TODO: Provide concrete implementation once signing strategy is defined.
