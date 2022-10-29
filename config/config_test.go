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
				configFilePath: "./test/config1",
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
				configFilePath: "./test/config1",
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

func TestGetAllConfig(t *testing.T) {
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "get all section name",
			args: args{
				configFilePath: "./test/config1",
			},
			want:    []string{"config", "config name includes space"},
			wantErr: false,
		},
		{
			name: "config file is empty",
			args: args{
				configFilePath: "./test/config2",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllConfig(tt.args.configFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
