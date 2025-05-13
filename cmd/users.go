package cmd

import (
	"github.com/prs-io/openobserve-go/utils"

	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Users command",
	Long:  "A command-line interface for interacting with OpenObserve users.",
}

var ListUserCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := API.ListUser(orgID)
		if err != nil {
			cmd.Println("Error listing users:", err)
			return
		}
		// fmt.Printf("Response: %+v\n", response)
		var data [][]string
		for _, user := range response.Data {
			data = append(data, []string{
				user.Email,
				*user.FirstName,
				*user.LastName,
				string(user.Role),
			})
		}
		utils.Table([]string{"Email", "First Name", "Last Name", "Role"}, data)
	},
}

func init() {
	usersCmd.AddCommand(ListUserCmd)
}
