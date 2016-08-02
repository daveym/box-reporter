package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd - The base command for the application.
var RootCmd = &cobra.Command{
	Use:   "box-reporter",
	Short: "Extract Events from Box",
	Long:  `Extract Events from Box`,
}

// Execute - Adds all defined commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	viper.Get("PublicKeyID")
	viper.Get("ClientID")
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	viper.SetConfigName("box-reporter")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Please make sure that box-reporter.yaml exists in the working directory.")
		os.Exit(-1)
	}
}
