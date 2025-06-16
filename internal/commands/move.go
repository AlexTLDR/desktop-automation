package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AlexTLDR/desktop-automation/internal/automation"
	"github.com/spf13/cobra"
)

// NewMoveCmd creates a new move command
func NewMoveCmd() *cobra.Command {
	var smooth bool
	var duration float64

	var moveCmd = &cobra.Command{
		Use:   "move [x] [y]",
		Short: "Move the mouse cursor to coordinates",
		Long: `Move the mouse cursor to specific screen coordinates.

This command moves the mouse cursor to the specified x and y coordinates on the screen.
The coordinates are measured in pixels from the top-left corner of the screen (0,0).
The cursor will move to the target position without clicking.

Use the --smooth flag for animated movement with customizable duration.`,
		Example: `  # Move cursor to coordinates (800, 600) instantly
  desktop-automation move 800 600

  # Move cursor smoothly to coordinates (800, 600) over 1 second
  desktop-automation move --smooth 800 600

  # Move cursor smoothly to coordinates (800, 600) over 5 seconds
  desktop-automation move --smooth --duration 5.0 800 600

  # Move cursor to the top-left corner
  desktop-automation move 0 0

  # Move cursor to the center of a 1920x1080 screen
  desktop-automation move 960 540`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runMove(args, smooth, duration)
		},
	}

	// Add flags
	moveCmd.Flags().BoolVar(&smooth, "smooth", false, "Enable smooth animated movement")
	moveCmd.Flags().Float64Var(&duration, "duration", 1.0, "Duration in seconds for smooth movement (default: 1.0)")

	return moveCmd
}

// runMove executes the move command
func runMove(args []string, smooth bool, duration float64) error {
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

	// Validate duration for smooth movement
	if smooth && duration <= 0 {
		return fmt.Errorf("duration must be positive when using smooth movement, got %f", duration)
	}

	// Get current mouse position
	currentX, currentY := automation.GetPosition()
	fmt.Printf("Current position: (%d, %d)\n", currentX, currentY)
	fmt.Printf("Target position: (%d, %d)\n", x, y)

	// Check if we're already at the target position
	if currentX == x && currentY == y {
		fmt.Println("Already at target position!")
		return nil
	}

	// Perform the move operation
	if smooth {
		fmt.Printf("Moving smoothly over %.1f seconds...\n", duration)
		err = automation.SmoothMove(x, y, duration)
		if err != nil {
			return fmt.Errorf("failed to perform smooth move: %w", err)
		}
	} else {
		fmt.Println("Moving...")
		err = automation.Move(x, y)
		if err != nil {
			return fmt.Errorf("failed to perform move: %w", err)
		}
	}

	// Small delay to ensure the move is complete
	time.Sleep(100 * time.Millisecond)

	// Confirm final position
	finalX, finalY := automation.GetPosition()
	fmt.Printf("Final position: (%d, %d)\n", finalX, finalY)

	// Check if we reached the target position (allow for small tolerance)
	if abs(finalX-x) <= 1 && abs(finalY-y) <= 1 {
		fmt.Println("✓ Successfully moved to target position!")
	} else {
		fmt.Printf("⚠ Warning: Final position differs from target by (%d, %d) pixels\n",
			finalX-x, finalY-y)
	}

	return nil
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
