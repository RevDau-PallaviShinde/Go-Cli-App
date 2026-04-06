/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/RevDau-PallaviShinde/Go-Cli-App/api"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the REST API server",
	Long:  `Start the REST API server with Swagger UI on port [IP_ADDRESS].`,
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
