package cmd

import "github.com/spf13/cobra"

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Box to get an Access Token. One off activity.",
	Long: `To access the Box API, you need to authenticate with your client_id. Your client_id 
	can be found under the development area within the Box website`,
	Run: func(cmd *cobra.Command, args []string) {

	}}

func init() {
	RootCmd.AddCommand(authCmd)
}
