package cmd

import (
	"github.com/mzfarshad/music_store_api/internal"
	"github.com/spf13/cobra"
	"log"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin related commands",
}

var addAdminCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new admin",
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetString("email")
		pass, _ := cmd.Flags().GetString("password")
		name, _ := cmd.Flags().GetString("name")

		container, err := internal.NewContainer()
		if err != nil {
			log.Fatalln(err)
		}
		err = container.Cli.AdminCli.CreateAdmin(email, name, pass)
		if err != nil {
			log.Fatalf("failed create admin: %v", err)
		}
		log.Println("Admin created successfully")
	},
}

func init() {
	addAdminCmd.Flags().String("name", "", "admin name(required)")
	addAdminCmd.Flags().String("email", "", "admin email(required)")
	addAdminCmd.Flags().String("password", "", "admin password(required)")

	mustMarkFlagRequired(addAdminCmd, "name")
	mustMarkFlagRequired(addAdminCmd, "email")
	mustMarkFlagRequired(addAdminCmd, "password")

	adminCmd.AddCommand(addAdminCmd)
	rootCmd.AddCommand(adminCmd)
}

func mustMarkFlagRequired(cmd *cobra.Command, name string) {
	if err := cmd.MarkFlagRequired(name); err != nil {
		log.Fatalf("cmd %s: failed to mark flag %s as required: %v", cmd.Name(), name, err)
	}
}
