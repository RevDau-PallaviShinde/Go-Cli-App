/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

var fileName string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export tasks to a file",
	Long:  `Export all tasks to a CSV file.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := handlers.ExportTasksToCSV(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Tasks exported to", fileName)

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVar(&fileName, "file", "tasks.csv", "CSV file name")
}
