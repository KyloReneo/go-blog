package config

import "github.com/kyloReneo/go-blog/config"

//Get config variables all over the project
func Get() config.Config {
	return configurations
}
