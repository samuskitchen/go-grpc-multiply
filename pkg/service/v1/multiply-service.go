package v1

import (
	"context"
	v1 "go-grpc-multiply/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// multiplyServiceServer is implementation of v1.MultiplyServiceServer proto interface
type multiplyServiceServer struct{}

// NewMultiplyServiceServer creates Multiply service
func NewMultiplyServiceServer() v1.MultiplyServiceServer {
	return &multiplyServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (ms *multiplyServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (ms *multiplyServiceServer) MultiplyNumbers(ctx context.Context, request *v1.NumbersRequest) (*v1.NumberResponse, error) {
	// check if the API version requested by client is supported by server
	if err := ms.checkAPI(request.Api); err != nil {
		return nil, err
	}

	result := request.Number1 * request.Number2

	return &v1.NumberResponse{
		Api:     apiVersion,
		Number1: result,
	}, nil
}
