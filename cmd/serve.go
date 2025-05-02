package cmd

import (
	"github.com/mzfarshad/music_store_api/internal"
	"github.com/mzfarshad/music_store_api/rest/api"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long: `Start the web server to serve the application.

This command initializes the application with all necessary dependencies
and starts the web server using the Fiber framework. The server listens
on the port specified in the application configuration.

Examples:
  app serve

Ensure that the configuration file is set up correctly before starting the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		container, err := internal.InjectDependencies()
		if err != nil {
			log.Fatalln(err)
		}

		if err := api.Serve(container); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
