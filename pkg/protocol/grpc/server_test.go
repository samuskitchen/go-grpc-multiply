package grpc

import (
	"context"
	v1 "go-grpc-multiply/pkg/api/v1"
	"testing"
)

func TestRunServer(t *testing.T) {
	type args struct {
		ctx   context.Context
		v1API v1.MultiplyServiceServer
		port  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error runs gRPC service",
			args: args{
				ctx:   nil,
				v1API: nil,
				port:  "nil",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunServer(tt.args.ctx, tt.args.v1API, tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("RunServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
