package commands

// Package commands provides CLI command implementations for desktop automation.
// This package contains all the command-line interface logic for the desktop automation tool.

import (
	"fmt"

	"github.com/AlexTLDR/desktop-automation/internal/automation"
)

// Command represents a CLI command
type Command interface {
	Execute(args []string) error
	Help() string
}

// MouseCommand handles mouse-related automation commands
type MouseCommand struct {
	mouse *automation.Mouse
}

// NewMouseCommand creates a new mouse command
func NewMouseCommand() *MouseCommand {
	return &MouseCommand{
		mouse: automation.NewMouse(),
	}
}

// Execute runs the mouse command
func (mc *MouseCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mouse command requires subcommand")
	}

	switch args[0] {
	case "click":
		mc.mouse.Click()
		return nil
	default:
		return fmt.Errorf("unknown mouse command: %s", args[0])
	}
}

// Help returns help text for mouse commands
func (mc *MouseCommand) Help() string {
	return "Mouse automation commands:\n  click - Click at current position"
}

// KeyboardCommand handles keyboard-related automation commands
type KeyboardCommand struct {
	keyboard *automation.Keyboard
}

// NewKeyboardCommand creates a new keyboard command
func NewKeyboardCommand() *KeyboardCommand {
	return &KeyboardCommand{
		keyboard: automation.NewKeyboard(),
	}
}

// Execute runs the keyboard command
func (kc *KeyboardCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("keyboard command requires subcommand")
	}

	switch args[0] {
	case "type":
		if len(args) < 2 {
			return fmt.Errorf("type command requires text argument")
		}
		kc.keyboard.Type(args[1])
		return nil
	default:
		return fmt.Errorf("unknown keyboard command: %s", args[0])
	}
}

// Help returns help text for keyboard commands
func (kc *KeyboardCommand) Help() string {
	return "Keyboard automation commands:\n  type <text> - Type the specified text"
}

// ScreenCommand handles screen-related automation commands
type ScreenCommand struct {
	screen *automation.Screen
}

// NewScreenCommand creates a new screen command
func NewScreenCommand() *ScreenCommand {
	return &ScreenCommand{
		screen: automation.NewScreen(),
	}
}

// Execute runs the screen command
func (sc *ScreenCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("screen command requires subcommand")
	}

	switch args[0] {
	case "capture":
		bitmap := sc.screen.Capture()
		fmt.Printf("Screen captured: %dx%d\n", bitmap.Width, bitmap.Height)
		return nil
	default:
		return fmt.Errorf("unknown screen command: %s", args[0])
	}
}

// Help returns help text for screen commands
func (sc *ScreenCommand) Help() string {
	return "Screen automation commands:\n  capture - Capture the entire screen"
}

// PrintUsage prints general usage information
func PrintUsage() {
	fmt.Println("Desktop Automation CLI")
	fmt.Println("Usage: desktop-automation <command> [args...]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  mouse    - Mouse automation commands")
	fmt.Println("  keyboard - Keyboard automation commands")
	fmt.Println("  screen   - Screen automation commands")
	fmt.Println("  version  - Show version information")
	fmt.Println("  help     - Show this help message")
}

// PrintVersion prints version information
func PrintVersion() {
	fmt.Println("desktop-automation v0.1.0")
}
