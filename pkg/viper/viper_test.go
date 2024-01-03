package viper

import (
	"reflect"
	"testing"
)

type MockConfig struct {
	MockName string
}

func (c *MockConfig) String() string {
	return ""
}

func (c *MockConfig) GetConfigType() string {
	return ConfigTypeJson
}

var mockConfig = &MockConfig{}

func TestViper(t *testing.T) {
	type args struct {
		path        string
		configValue CliOptions
	}
	tests := []struct {
		name string
		args args
		want *MockConfig
	}{
		{
			name: "配置文件初始化",
			args: args{
				path:        "./mock_config.json",
				configValue: mockConfig,
			},
			want: &MockConfig{
				MockName: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Viper(tt.args.path, tt.args.configValue); !reflect.DeepEqual(tt.args.configValue, tt.want) {
				t.Errorf("config() = %v, want %v", tt.args.configValue, tt.want)
			}
		})
	}
}
