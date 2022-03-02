package util

import (
	"reflect"
	"testing"
)


var stubInspiroClient InspiroClient = *NewInspiroClient(
	InspiroConfig{
	API_url: "myapi",
	Logger: nil,
	Backup_image_link: "myBackup",
	},
)


func TestNewBloopyHttpClient(t *testing.T) {
	type args struct {
		inspiro *InspiroClient
	}
	tests := []struct {
		name string
		args args
		want *BloopyHttp
	}{
		{
			name: "Constructs",
			args: args{inspiro: &stubInspiroClient},
			want: &BloopyHttp{inspiro_api: &stubInspiroClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBloopyHttpClient(tt.args.inspiro); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBloopyHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
