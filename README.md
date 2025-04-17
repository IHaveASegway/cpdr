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

- Go 1.16 or higher (install from [golang.org/dl](https://golang.org/dl/) if needed)
- Git

## Installation Steps

```bash
# Clone the repository
git clone https://github.com/yourusername/cpdr.git
cd cpdr
```

### Set local cpdr executable

```bash
export cpdr="$(pwd)/cpdr"
```

### Add alias to your shell configuration

```bash
echo "alias cpdrtest='$(pwd)/cpdr'" >> ~/.zshrc
```

### Activate the alias

```bash
source ~/.zshrc
```
