package viper

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type CliOptions interface {
	String() string
	GetConfigType() string
}

func Viper(path string, configValue CliOptions) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				panic("config路径为空，请设置")
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(configValue.GetConfigType())
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&configValue); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&configValue); err != nil {
		fmt.Println(err)
	}
	return v
}
