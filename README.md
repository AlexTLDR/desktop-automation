# Desktop Automation CLI

A Go-based command-line interface for desktop automation tasks including mouse movements, keyboard input, screen capture, and window management.

## Features

- Mouse automation (clicks, movements)
- Keyboard input simulation
- Screen capture functionality
- Interactive TUI components using Bubble Tea
- Cross-platform desktop automation using RobotGo

## Project Structure

```
desktop-automation/
├── cmd/
│   └── desktop-automation/    # Main entry point
│       └── main.go
├── internal/
│   ├── automation/           # RobotGo wrappers
│   │   └── automation.go
│   ├── commands/            # CLI commands
│   │   └── commands.go
│   └── ui/                  # Bubble Tea TUI components
│       └── ui.go
├── go.mod                   # Go module definition
├── .gitignore              # Git ignore file
├── Taskfile.yml            # Task runner configuration
└── README.md               # This file
```

## Prerequisites

- Go 1.21 or later
- Task runner (optional, for using Taskfile)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AlexTLDR/desktop-automation.git
   cd desktop-automation
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   task build
   # or
   go build -o bin/desktop-automation ./cmd/desktop-automation
   ```

## Usage

### Basic Commands

```bash
# Show version
./bin/desktop-automation version

# Mouse automation
./bin/desktop-automation mouse click

# Keyboard automation
./bin/desktop-automation keyboard type "Hello, World!"

# Screen capture
./bin/desktop-automation screen capture
```

### Development

Available tasks (using Task runner):

```bash
task build    # Build the application
task run      # Build and run the application
task clean    # Clean build artifacts
task test     # Run tests
task fmt      # Format Go code
task mod      # Download and tidy dependencies
```

## Dependencies

- [RobotGo](https://github.com/go-vgo/robotgo) - Cross-platform desktop automation
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework

## License

MIT License