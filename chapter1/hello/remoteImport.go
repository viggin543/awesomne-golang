package main

// a remote dependency
// in golang dependencies are just git repos.
import "github.com/spf13/viper"

func foo() {
	viper.New()
}
