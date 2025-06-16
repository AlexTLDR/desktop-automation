package main

import (
	"fmt"
	"os"

	"github.com/AlexTLDR/desktop-automation/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "desktop-automation",
	Short: "Beautiful Desktop Automation CLI",
	Long:  "A beautiful and powerful CLI tool for desktop automation tasks including mouse control, keyboard input, and screen interactions.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Add all command stubs
	rootCmd.AddCommand(commands.NewClickCmd())
	rootCmd.AddCommand(commands.NewTypeCmd())
	rootCmd.AddCommand(commands.NewMoveCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
