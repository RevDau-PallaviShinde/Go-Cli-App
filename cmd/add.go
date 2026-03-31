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
	title    string
	desc     string
	due      string
	priority string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `Add your new task here.`,
	Run: func(cmd *cobra.Command, args []string) {

		err := handlers.AddTask(title, desc, due, priority)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Task added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&title, "title", "", "Title of the task (required)")
	addCmd.Flags().StringVar(&desc, "desc", "", "Description of the task")
	addCmd.Flags().StringVar(&due, "due", "", "Due date (YYYY-MM-DD)")
	addCmd.Flags().StringVar(&priority, "priority", "", "Priority of the task")
	_ = addCmd.MarkFlagRequired("title")

}
