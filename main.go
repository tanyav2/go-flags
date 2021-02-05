package main

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	flagset := pflag.CommandLine
	flagset.SortFlags = false
	flagset.Int64("rotation-interval", 60*60*24, "Key rotation interval")
}

func main() {
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	ri := viper.GetInt64("rotation-interval")
	fmt.Printf("ri: %v\n", ri)
}
