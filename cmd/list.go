/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/RevDau-PallaviShinde/Go-Cli-App/handlers"
	"github.com/spf13/cobra"
)

var status string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List my all tasks`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := handlers.ListTasks(status, priority)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, task := range tasks {
			fmt.Println("ID:", task.ID)
			fmt.Println("Title:", task.Title)
			fmt.Println("Description:", task.Description)
			fmt.Println("Status:", task.Status)
			fmt.Println("Priority:", task.Priority)

			if task.DueDate != "" {
				fmt.Println("Due Date:", task.DueDate)
			}

			fmt.Println("---------------------------")
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&status, "status", "", "Filter tasks by status (pending, in-progress, completed)")
	listCmd.Flags().StringVar(&priority, "priority", "", "Filter tasks by priority (low, medium, high)")
}
