package automation

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// Click performs a mouse click at the specified coordinates
func Click(x, y int) error {
	// Validate input coordinates
	if x < 0 {
		return fmt.Errorf("x coordinate must be non-negative, got %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate must be non-negative, got %d", y)
	}

	// Move mouse to coordinates and click
	robotgo.Move(x, y)
	robotgo.Click()

	return nil
}

// GetPosition returns the current mouse cursor position
func GetPosition() (int, int) {
	x, y := robotgo.Location()
	return x, y
}
