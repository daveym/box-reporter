package cmd

import (
	"fmt"

	"github.com/daveym/box-reporter-go/api"
	"github.com/daveym/box-reporter-go/boxapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Box to get an Access Token. One off activity.",
	Long: `To access the Box API, you need to authenticate with your client_id. Your client_id 
	can be found under the development area within the Box website`,
	Run: func(cmd *cobra.Command, args []string) {

		bc := &box.Client{}
		bc.SetPublicKeyID(viper.GetString("PublicKeyID"))
		bc.SetClientID(viper.GetString("ClientID"))
		bc.SetClaimSub(viper.GetString("ClaimSub"))

		msg := api.Authenticate(bc)

		fmt.Println(msg)

	}}

func init() {
	RootCmd.AddCommand(authCmd)
}
