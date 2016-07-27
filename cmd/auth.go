package cmd

import (
	"fmt"

	"github.com/daveym/box-reporter/api"
	"github.com/daveym/box-reporter/box"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Box to get an Access Token. One off activity.",
	Long: `To access the Box API, you need to authenticate with your client_id. Your client_id 
	can be found under the development area within the Box website`,
	Run: func(cmd *cobra.Command, args []string) {

		pc := &box.Client{}

		pc.SetClientID(viper.GetString("ClientID"))
		pc.SetAccessToken(viper.GetString("AccessToken"))

		msg := api.Authenticate(pc)

		viper.Set("ClientID", pc.GetClientID())
		viper.Set("AccessToken", pc.GetAccessToken())

		fmt.Println(msg)

	}}

func init() {
	RootCmd.AddCommand(authCmd)
}
