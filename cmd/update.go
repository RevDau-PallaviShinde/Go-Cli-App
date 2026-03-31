/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

var (
	updateID       int
	updateTitle    string
	updateDesc     string
	updateDue      string
	updateStatus   string
	updatePriority string
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Task",
	Long:  `Update my Task`,
	Run: func(cmd *cobra.Command, args []string) {

		err := handlers.UpdateTask(
			updateID,
			updateTitle,
			updateDesc,
			updateDue,
			updateStatus,
			updatePriority,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVar(&updateID, "id", 0, "Task ID (required)")
	updateCmd.Flags().StringVar(&updateTitle, "title", "", "New title")
	updateCmd.Flags().StringVar(&updateDesc, "desc", "", "New description")
	updateCmd.Flags().StringVar(&updateDue, "due", "", "New due date (YYYY-MM-DD)")
	updateCmd.Flags().StringVar(&updateStatus, "status", "", "New status (pending, in-progress, completed)")
	updateCmd.Flags().StringVar(&updatePriority, "priority", "", "New priority (low, medium, high)")

	_ = updateCmd.MarkFlagRequired("id")

}
