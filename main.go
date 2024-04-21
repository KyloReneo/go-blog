package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/kyloReneo/go-blog/config"

)

func main() {

	configs := configSet()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"app name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080
}

func configSet() config.Config {

	viper.SetConfigName("config") //Name of config file
	viper.SetConfigType("yml")    //Extention of the config file
	viper.AddConfigPath("config") //Path of the config file

	//Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		//Handling errors while reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found; ignore error if desired.")
		} else {
			fmt.Println("Config file was found but another error was produced while reading the file...")
		}
	}

	var configs config.Config

	//Unmarshal all values to struct
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	fmt.Println(configs)
	return configs
}
