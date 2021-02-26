package v1

import (
	"context"
	v1 "go-grpc-multiply/pkg/api/v1"
	"reflect"
	"testing"
)

func TestNewMultiplyServiceServer(t *testing.T) {
	tests := []struct {
		name string
		want v1.MultiplyServiceServer
	}{
		{
			name: "Constructor Successfully",
			want: &multiplyServiceServer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMultiplyServiceServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiplyServiceServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyServiceServer_MultiplyNumbers(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *v1.NumbersRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.NumberResponse
		wantErr bool
	}{
		{
			name: "Multiply Successfully",
			args: args{
				ctx: context.Background(),
				request: &v1.NumbersRequest{
					Api:     "v1",
					Number1: 5,
					Number2: 2,
				},
			},
			want: &v1.NumberResponse{
				Api:     "v1",
				Number1: 10,
			},
			wantErr: false,
		},
		{
			name: "Check Api Version Fail",
			args: args{
				ctx: context.Background(),
				request: &v1.NumbersRequest{
					Api:     "v2",
					Number1: 5,
					Number2: 2,
				},
			},
			want: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &multiplyServiceServer{}
			got, err := ms.MultiplyNumbers(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("MultiplyNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyNumbers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplyServiceServer_checkAPI(t *testing.T) {
	type args struct {
		api string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Version Ok",
			args: args{
				api: "v1",
			},
			wantErr: false,
		},
		{
			name: "Version Fail",
			args: args{
				api: "v2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &multiplyServiceServer{}
			if err := ms.checkAPI(tt.args.api); (err != nil) != tt.wantErr {
				t.Errorf("checkAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
