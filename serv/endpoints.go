package serv

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	CalcEndpoint endpoint.Endpoint
}

// MakeCalcEndpoint returns the response from our service "Calc"
func MakeCalcEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(calcRequest)
		res, err := srv.Calc(ctx, req.A, req.B)
		if err != nil {
			return calcResponse{res, err.Error()}, nil
		}
		return calcResponse{res, ""}, nil
	}
}

// Calc endpoint mapping
func (e Endpoints) Calc(ctx context.Context, a, b int) (int, error) {
	req := calcRequest{A: a, B: b}
	resp, err := e.CalcEndpoint(ctx, req)
	if err != nil {
		return 0, err
	}
	CalcResp := resp.(calcResponse)
	if CalcResp.Err != "" {
		return 0, errors.New(CalcResp.Err)
	}
	return CalcResp.Res, nil
}
