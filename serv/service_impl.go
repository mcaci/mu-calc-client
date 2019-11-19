package serv

import (
	"context"

	"github.com/mcaci/mu-calc/calc/calc"
)

type calcService struct{}

// NewCalcService makes a new Service
func NewCalcService() Service {
	return calcService{}
}

func (calcService) Calc(ctx context.Context, a, b int) (int, error) {
	return calc.Calc(a, b), nil
}
