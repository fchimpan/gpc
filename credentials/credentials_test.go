package credentials

import (
	"reflect"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	type args struct {
		sectionName string
		credPath    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Credentials
		wantErr bool
	}{
		{
			name: "load default section",
			args: args{
				sectionName: "default",
				credPath:    "./test/credentials",
			},
			want: &Credentials{
				ConfluenceAPIToken: "default",
				ConfluenceAEmail:   "default",
			},
			wantErr: false,
		},
		{
			name: "insufficient fields",
			args: args{
				sectionName: "custom",
				credPath:    "./test/credentials",
			},
			want: &Credentials{
				ConfluenceAEmail: "custom",
			},
			wantErr: false,
		},
		{
			name: "section does not exist",
			args: args{
				sectionName: "section",
				credPath:    "./test/credentials",
			},
			want:    &Credentials{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCredentials(tt.args.credPath, tt.args.sectionName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}
