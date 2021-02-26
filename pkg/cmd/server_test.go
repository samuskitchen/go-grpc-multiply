package cmd

import "testing"

func TestRunServer(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Error runs gRPC server and HTTP gateway",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunServer(); (err != nil) != tt.wantErr {
				t.Errorf("RunServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
