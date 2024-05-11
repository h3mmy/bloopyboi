package services

import (
	"reflect"
	"testing"
)

var stubInspiroClient InspiroClient = *NewInspiroClientWithConfig(
	InspiroConfig{
		API_url:           "myapi",
		Logger:            nil,
		Backup_image_link: "myBackup",
	},
)

func TestNewInspiroService(t *testing.T) {
	type args struct {
		inspiro *InspiroClient
	}
	tests := []struct {
		name string
		args args
		want *InspiroService
	}{
		{
			name: "Constructs",
			args: args{inspiro: &stubInspiroClient},
			want: &InspiroService{inspiroClient: &stubInspiroClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroService(tt.args.inspiro); !reflect.DeepEqual(got.inspiroClient, tt.want.inspiroClient) {
				t.Errorf("NewBloopyHttpClient() = %v, want %v", got.inspiroClient, tt.want.inspiroClient)
			}
		})
	}
}

func TestGetsClient(t *testing.T) {
	type args struct {
		inspiro *InspiroClient
	}
	tests := []struct {
		name string
		args args
		want *InspiroClient
	}{
		{
			name: "Constructs basic",
			args: args{inspiro: &stubInspiroClient},
			want: &stubInspiroClient,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroService(tt.args.inspiro).GetClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloopyClient().GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}
