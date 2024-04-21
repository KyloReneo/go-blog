package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/kyloReneo/go-blog/config"

)

func Set() {
	viper.SetConfigName("config") //Name of config file
	viper.SetConfigType("yml")    //Extention of the config file
	viper.AddConfigPath("config") //Path of the config file

	//Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		//Handling errors while reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired.")
		} else {
			log.Fatal("Config file was found but another error was produced while reading the file...")
		}
	}

	var configs config.Config

	//Unmarshal all values to struct
	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("Unable to decode configs into struct")
	}
	fmt.Printf("Unmarshaled config file:\n%v", configs)
}
