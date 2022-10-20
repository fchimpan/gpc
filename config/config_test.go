package config

import (
	"reflect"
	"testing"
)

func TestGetconfig(t *testing.T) {
	type args struct {
		configFilePath string
		sectionName    string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "load config section",
			args: args{
				sectionName:    "config",
				configFilePath: "./test/config",
			},
			want: &Config{
				SpaceKey: "dummy",
				Domain:   "dummy",
				Parent:   "dummy",
			},
			wantErr: false,
		},
		{
			name: "config name includes space",
			args: args{
				sectionName:    "config name includes space",
				configFilePath: "./test/config",
			},
			want: &Config{
				SpaceKey: "dummy",
				Domain:   "dummy",
				Parent:   "dummy",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfig(tt.args.configFilePath, tt.args.sectionName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getconfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Getconfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
