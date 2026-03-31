/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

var deleteID int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Task",
	Long:  `Delete task by Id`,
	Run: func(cmd *cobra.Command, args []string) {

		err := handlers.DeleteTask(deleteID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVar(&deleteID, "id", 0, "Task ID (required)")
	_ = deleteCmd.MarkFlagRequired("id")

}
