# Desktop Automation CLI

A beautiful and powerful CLI tool for desktop automation tasks including mouse control, keyboard input, and screen interactions. Built with Go and the Cobra CLI framework for an intuitive command-line experience.

## Features

- **Mouse Control**: Click at specific coordinates and move cursor to any screen position
- **Keyboard Input**: Type text at the current cursor position with full Unicode support
- **Cross-platform**: Works on Windows, macOS, and Linux
- **Modern CLI**: Built with Cobra framework for excellent user experience
- **Comprehensive Error Handling**: Clear error messages and input validation
- **Extensible Architecture**: Easy to add new automation commands

## Project Structure

```
desktop-automation/
├── cmd/
│   └── desktop-automation/    # Main CLI entry point (Cobra root command)
│       └── main.go
├── internal/
│   ├── automation/           # RobotGo automation backend
│   │   └── automation.go
│   ├── commands/            # Cobra CLI commands
│   │   ├── commands.go      # Package documentation
│   │   ├── click.go         # Click command implementation
│   │   ├── type.go          # Type command implementation
│   │   └── move.go          # Move command implementation
│   └── ui/                  # TUI components (Bubble Tea)
│       └── ui.go
├── bin/                     # Built binaries
├── go.mod                   # Go module definition
├── .gitignore              # Git ignore rules
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

### Basic Usage

```bash
./bin/desktop-automation [command] [flags]
```

### Available Commands

- `click` - Click at a specific screen coordinate
- `type` - Type text at current cursor position
- `move` - Move the mouse cursor to coordinates
- `help` - Help about any command
- `completion` - Generate autocompletion script

### Quick Examples

```bash
# Click at coordinates (100, 200)
./bin/desktop-automation click 100 200

# Type a message
./bin/desktop-automation type "Hello, World!"

# Move cursor to center of 1920x1080 screen
./bin/desktop-automation move 960 540

# Get help for any command
./bin/desktop-automation click --help
```

## Command Reference

### click

Click at a specific screen coordinate.

**Usage:**
```bash
desktop-automation click [x] [y]
```

**Description:**
Simulates a mouse click at the specified x and y coordinates. Coordinates are measured in pixels from the top-left corner of the screen (0,0).

**Examples:**
```bash
# Click at coordinates (100, 200)
desktop-automation click 100 200

# Click at the top-left corner
desktop-automation click 0 0

# Click at coordinates (500, 300)
desktop-automation click 500 300
```

**Arguments:**
- `x` - X coordinate (must be a non-negative integer)
- `y` - Y coordinate (must be a non-negative integer)

### type

Type text at the current cursor position.

**Usage:**
```bash
desktop-automation type [text]
```

**Description:**
Simulates keyboard input by typing the specified text at the current cursor position. Supports Unicode characters and special symbols.

**Examples:**
```bash
# Type a simple message
desktop-automation type "Hello, World!"

# Type text with spaces (use quotes)
desktop-automation type "This is a test message"

# Type a single word
desktop-automation type hello

# Type text with special characters
desktop-automation type "user@example.com"
```

**Arguments:**
- `text` - The text to type (supports multiple words when quoted)

### move

Move the mouse cursor to specific screen coordinates.

**Usage:**
```bash
desktop-automation move [x] [y]
```

**Description:**
Moves the mouse cursor to the specified x and y coordinates. The cursor will smoothly move to the target position without clicking.

**Examples:**
```bash
# Move cursor to coordinates (100, 200)
desktop-automation move 100 200

# Move cursor to the top-left corner
desktop-automation move 0 0

# Move cursor to the center of a 1920x1080 screen
desktop-automation move 960 540
```

**Arguments:**
- `x` - X coordinate (must be a non-negative integer)
- `y` - Y coordinate (must be a non-negative integer)

## Error Handling

The CLI includes comprehensive error handling and input validation:

```bash
# Missing arguments
$ desktop-automation click
Error: accepts 2 arg(s), received 0

# Invalid coordinate format
$ desktop-automation click abc 123
Error: invalid x coordinate 'abc': must be a number

# Negative coordinates
$ desktop-automation move -- -5 10
Error: x coordinate must be non-negative, got -5

# Empty text
$ desktop-automation type ""
Error: text cannot be empty
```

## Development

### Available Tasks

Using the Task runner:

```bash
task build    # Build the application
task run      # Build and run the application
task clean    # Clean build artifacts
task test     # Run tests
task fmt      # Format Go code
task mod      # Download and tidy dependencies
```

### Manual Commands

```bash
# Build
go build -o bin/desktop-automation ./cmd/desktop-automation

# Run tests
go test ./...

# Format code
go fmt ./...

# Vet code
go vet ./...
```

### Adding New Commands

1. Create a new file in `internal/commands/` (e.g., `newcommand.go`)
2. Implement the command using Cobra patterns
3. Add the command to the root command in `cmd/desktop-automation/main.go`
4. Include proper error handling and validation
5. Add usage examples and help text

## Implementation Status

**Current Status:** Command stubs with full validation and error handling

**TODO:**
- [ ] Integrate with actual automation backend (`internal/automation`)
- [ ] Implement mouse clicking functionality
- [ ] Implement keyboard typing functionality  
- [ ] Implement mouse movement functionality
- [ ] Add configuration options (click delay, typing speed, etc.)
- [ ] Add screen boundary validation
- [ ] Add support for different mouse buttons
- [ ] Add support for keyboard shortcuts and special keys

## Dependencies

- **[Cobra](https://github.com/spf13/cobra)** - Modern CLI framework
- **[RobotGo](https://github.com/go-vgo/robotgo)** - Cross-platform desktop automation
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Terminal UI framework

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Run `task fmt` and `task test`
6. Submit a pull request

## License

MIT License