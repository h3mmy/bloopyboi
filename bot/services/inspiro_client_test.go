package services

import (
	"reflect"
	"testing"
)

var stubInspiroService InspiroService = *NewInspiroServiceWithConfig(
	InspiroConfig{
		API_url:           "myapi",
		Logger:            nil,
		Backup_image_link: "myBackup",
	},
)

func TestNewInspiroClient(t *testing.T) {
	type args struct {
		inspiro *InspiroService
	}
	tests := []struct {
		name string
		args args
		want *InspiroClient
	}{
		{
			name: "Constructs",
			args: args{inspiro: &stubInspiroService},
			want: &InspiroClient{inspiroService: &stubInspiroService},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroHttpClient(tt.args.inspiro); !reflect.DeepEqual(got.inspiroService, tt.want.inspiroService) {
				t.Errorf("NewBloopyHttpClient() = %v, want %v", got.inspiroService, tt.want.inspiroService)
			}
		})
	}
}

func TestGetsService(t *testing.T) {
	type args struct {
		inspiro *InspiroService
	}
	tests := []struct {
		name string
		args args
		want *InspiroService
	}{
		{
			name: "Constructs basic",
			args: args{inspiro: &stubInspiroService},
			want: &stubInspiroService,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroHttpClient(tt.args.inspiro).GetService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloopyClient().GetService() = %v, want %v", got, tt.want)
			}
		})
	}
}
