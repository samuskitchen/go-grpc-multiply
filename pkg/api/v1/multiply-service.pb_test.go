package v1

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"testing"
)

func TestNewMultiplyServiceClient(t *testing.T) {
	type args struct {
		cc grpc.ClientConnInterface
	}
	tests := []struct {
		name string
		args args
		want MultiplyServiceClient
	}{
		{
			name: "NewMultiplyServiceClient successfully",
			args: args{
				cc: nil,
			},
			want: &multiplyServiceClient{
				cc: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMultiplyServiceClient(tt.args.cc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiplyServiceClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberResponse_Descriptor(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		want1  []int
	}{
		{name: "successfully descriptor",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       10,
			},
			want:  file_multiply_service_proto_rawDescGZIP(),
			want1: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}
			got, got1 := nu.Descriptor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Descriptor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Descriptor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNumberResponse_GetApi(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "GetApi successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       0,
			},
			want: "v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}
			if got := x.GetApi(); got != tt.want {
				t.Errorf("GetApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberResponse_GetNumber1(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "GetNumber1 successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "",
				Number1:       1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}
			if got := x.GetNumber1(); got != tt.want {
				t.Errorf("GetNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumberResponse_ProtoMessage(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "verify version successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}

			if nu.Api != "v1" {
				t.Errorf("api = %v, want %v", "v1", nu.Api)
			}
		})
	}
}

func TestNumberResponse_Reset(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       10,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}

			if x.Api != "v1" {
				t.Errorf("api = %v, want %v", "v1", x.Api)
			}
		})
	}
}

func TestNumberResponse_String(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       0,
			},
			want: "api:\"v1\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumberResponse{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
			}
			if got := x.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersRequest_Descriptor(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
		want1  []int
	}{
		{
			name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "",
				Number1:       0,
				Number2:       0,
			},
			want:  file_multiply_service_proto_rawDescGZIP(),
			want1: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}
			got, got1 := nu.Descriptor()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Descriptor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Descriptor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNumbersRequest_GetApi(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			},
			want: "v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}
			if got := x.GetApi(); got != tt.want {
				t.Errorf("GetApi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersRequest_GetNumber1(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "GetNumber1 successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}
			if got := x.GetNumber1(); got != tt.want {
				t.Errorf("GetNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersRequest_GetNumber2(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			name: "GetNumber2 successfully",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}
			if got := x.GetNumber2(); got != tt.want {
				t.Errorf("GetNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumbersRequest_ProtoMessage(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nu := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}

			if nu.Api != "v1" {
				t.Errorf("api = %v, want %v", "v1", nu.Api)
			}
		})
	}
}

func TestNumbersRequest_Reset(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}

			if x.Api != "v1" {
				t.Errorf("api = %v, want %v", "v1", x.Api)
			}
		})
	}
}

func TestNumbersRequest_String(t *testing.T) {
	type fields struct {
		state         protoimpl.MessageState
		sizeCache     protoimpl.SizeCache
		unknownFields protoimpl.UnknownFields
		Api           string
		Number1       float32
		Number2       float32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "Ok",
			fields: fields{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			},
			want: "api:\"v1\" number1:5 number2:2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &NumbersRequest{
				state:         tt.fields.state,
				sizeCache:     tt.fields.sizeCache,
				unknownFields: tt.fields.unknownFields,
				Api:           tt.fields.Api,
				Number1:       tt.fields.Number1,
				Number2:       tt.fields.Number2,
			}
			if got := x.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnimplementedMultiplyServiceServer_MultiplyNumbers(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *NumbersRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *NumberResponse
		wantErr bool
	}{
		{name: "Error MultiplyNumbers", args: args{
			in0: context.Background(),
			in1: &NumbersRequest{
				state:         protoimpl.MessageState{},
				sizeCache:     0,
				unknownFields: nil,
				Api:           "v1",
				Number1:       5,
				Number2:       2,
			},
		},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			un := &UnimplementedMultiplyServiceServer{}
			got, err := un.MultiplyNumbers(tt.args.in0, tt.args.in1)
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

func Test_file_multiply_service_proto_rawDescGZIP(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			name: "Ok",
			want: file_multiply_service_proto_rawDescGZIP(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := file_multiply_service_proto_rawDescGZIP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("file_multiply_service_proto_rawDescGZIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
