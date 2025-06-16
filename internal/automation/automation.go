package automation

// Package automation provides wrappers around robotgo for desktop automation tasks.
// This includes mouse movements, keyboard input, screen capture, and window management.

import (
	"github.com/go-vgo/robotgo"
)

// Mouse provides mouse automation functionality
type Mouse struct{}

// NewMouse creates a new Mouse instance
func NewMouse() *Mouse {
	return &Mouse{}
}

// Click performs a mouse click at the current position
func (m *Mouse) Click() {
	robotgo.Click()
}

// Move moves the mouse to the specified coordinates
func (m *Mouse) Move(x, y int) {
	robotgo.Move(x, y)
}

// Keyboard provides keyboard automation functionality
type Keyboard struct{}

// NewKeyboard creates a new Keyboard instance
func NewKeyboard() *Keyboard {
	return &Keyboard{}
}

// Type types the given text using the TypeString function
func (k *Keyboard) Type(text string) error {
	return TypeString(text)
}

// TypeWithDelay types the given text with a delay between characters
func (k *Keyboard) TypeWithDelay(text string, delayMs int) error {
	return TypeStringWithDelay(text, delayMs)
}

// Screen provides screen capture functionality
type Screen struct{}

// NewScreen creates a new Screen instance
func NewScreen() *Screen {
	return &Screen{}
}

// Capture captures the entire screen
func (s *Screen) Capture() robotgo.CBitmap {
	return robotgo.CaptureScreen()
}
