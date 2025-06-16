package commands

import (
	"fmt"
	"strconv"

	"github.com/AlexTLDR/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// NewClickCmd creates a new click command
func NewClickCmd() *cobra.Command {
	var clickCmd = &cobra.Command{
		Use:   "click [x] [y]",
		Short: "Click at a specific screen coordinate",
		Long: `Click at a specific screen coordinate.

This command simulates a mouse click at the specified x and y coordinates on the screen.
The coordinates are measured in pixels from the top-left corner of the screen (0,0).`,
		Example: `  # Click at coordinates (100, 200)
  desktop-automation click 100 200

  # Click at the top-left corner
  desktop-automation click 0 0

  # Click at coordinates (500, 300)
  desktop-automation click 500 300`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runClick(args)
		},
	}

	return clickCmd
}

// runClick executes the click command
func runClick(args []string) error {
	// Parse x coordinate
	x, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid x coordinate '%s': must be a number", args[0])
	}

	// Parse y coordinate
	y, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("invalid y coordinate '%s': must be a number", args[1])
	}

	// Validate coordinates are non-negative
	if x < 0 {
		return fmt.Errorf("x coordinate must be non-negative, got %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate must be non-negative, got %d", y)
	}

	// Show current mouse position before clicking
	currentX, currentY := automation.GetPosition()
	fmt.Printf("Current mouse position: (%d, %d)\n", currentX, currentY)

	// Perform the click
	if err := automation.Click(x, y); err != nil {
		return fmt.Errorf("failed to click at (%d, %d): %v", x, y, err)
	}

	// Confirm success
	fmt.Printf("Successfully clicked at coordinates (%d, %d)\n", x, y)

	return nil
}
