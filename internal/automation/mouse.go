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

// Move performs an instant mouse movement to the specified coordinates
func Move(x, y int) error {
	// Validate input coordinates
	if x < 0 {
		return fmt.Errorf("x coordinate must be non-negative, got %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate must be non-negative, got %d", y)
	}

	// Move mouse to coordinates instantly
	robotgo.Move(x, y)

	return nil
}

// SmoothMove performs a smooth animated mouse movement to the specified coordinates
func SmoothMove(x, y int, duration float64) error {
	// Validate input coordinates
	if x < 0 {
		return fmt.Errorf("x coordinate must be non-negative, got %d", x)
	}
	if y < 0 {
		return fmt.Errorf("y coordinate must be non-negative, got %d", y)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive, got %f", duration)
	}

	// Move mouse to coordinates with smooth animation
	robotgo.MoveSmooth(x, y, duration)

	return nil
}
