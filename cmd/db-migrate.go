package cmd

import (
	"github.com/mzfarshad/music_store_api/internal/adapter/repository"
	"github.com/spf13/cobra"
	"log"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "auto migrate postgres database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := repository.NewPostgresConnection()
		if err != nil {
			log.Fatalf(err.Error())
		}
		if err := repository.Migrate(db); err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("Successfully done migration")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
