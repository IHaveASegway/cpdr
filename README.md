# cpdr - Directory Contents to Clipboard

A command-line tool that copies directory structures and file contents to the clipboard.

## Features

- Copy directory structure to clipboard (`-s` flag)
- Copy both structure and file contents (default)
- Exclude directories/files using ignore patterns
- Control the depth of directory traversal
- Debug mode for troubleshooting

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Option 1: Clone and Build Manually

```bash
# Clone the repository
git clone https://github.com/yourusername/cpdr.git
cd cpdr

# Build the binary
go build -o cpdr
```

### Option 2: Install with Homebrew (macOS)

```bash
# Clone the repository
git clone https://github.com/yourusername/cpdr.git
cd cpdr

# Install with Homebrew formula
brew install --formula ./cpdr.rb
```

## Usage

```bash
# Copy directory structure only
cpdr -s /path/to/directory

# Copy with specified depth
cpdr -d 2 /path/to/directory

# Ignore specific directories
cpdr -i node_modules,vendor /path/to/directory

# Copy specific files
cpdr path/to/file1.txt path/to/file2.go

# Enable debug output
cpdr --debug /path/to/directory
```

### Command-line Options

- `-s, --structure`: Generate only directory structure
- `-i, --ignore`: Comma-separated list of patterns to ignore
- `-d, --depth`: Maximum depth for directory tree (-1 for no limit)
- `-f, --format`: Output format (text or json)
- `--debug`: Enable debug output

## Setting up an Alias

### Bash

Add to your `~/.bashrc` or `~/.bash_profile`:

```bash
# For manual installation
alias cpdr="/path/to/cpdr/cpdr"

# For Homebrew installation
# The alias is not needed as it's already in your PATH
```

Then apply the changes:
```bash
source ~/.bashrc  # or source ~/.bash_profile
```

### Zsh

Add to your `~/.zshrc`:

```bash
# For manual installation
alias cpdr="/path/to/cpdr/cpdr"

# For Homebrew installation
# The alias is not needed as it's already in your PATH
```

Then apply the changes:
```bash
source ~/.zshrc
```

## Examples

```bash
# Copy a project structure only
cpdr -s -i node_modules,build,dist ~/projects/myapp

# Copy an entire project with content
cpdr ~/projects/myapp

# Copy specific files
cpdr ~/projects/myapp/src/main.go ~/projects/myapp/README.md
```

```bash
cpdr . --structure --ignore .github,.secrets,.vscode,.git
```
