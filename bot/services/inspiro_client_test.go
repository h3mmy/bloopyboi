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

func TestNewBloopyHttpClient(t *testing.T) {
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
			want: &InspiroClient{Inspiro_api: &stubInspiroService},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroHttpClient(tt.args.inspiro); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloopyHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBloopyClient(t *testing.T) {
	tests := []struct {
		name string
		want *InspiroClient
	}{
		// Replace with actual useful test after learning how to mock in golang
		{
			name: "Constructs basic",
			want: NewInspiroClient(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInspiroClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloopyClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
