package rest

import (
	"context"
	"testing"
)

func TestRunServer(t *testing.T) {
	type args struct {
		ctx      context.Context
		grpcPort string
		httpPort string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error in http error",
			args: args{
				ctx:      context.Background(),
				grpcPort: "9000",
				httpPort: "nil",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunServer(tt.args.ctx, tt.args.grpcPort, tt.args.httpPort); (err != nil) != tt.wantErr {
				t.Errorf("RunServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
