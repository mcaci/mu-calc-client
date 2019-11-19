package serv

import "context"

// Service provides and interface for calcing integers
type Service interface {
	Calc(ctx context.Context, a, b int) (int, error)
}
