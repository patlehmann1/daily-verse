# daily-verse

A simple CLI tool that displays Bible verses in your terminal.

## Features

- Zero dependencies - single Go binary
- Verses embedded in the binary
- Clean, plain text output (pipe-friendly)
- Random verse by default
- Consistent daily verse with `--daily` flag
- Filter by book or testament

## Installation

### Via Homebrew

```bash
brew install daily-verse
```

### From Source

```bash
go install github.com/yourusername/daily-verse@latest
```

## Usage

```bash
# Random verse
daily-verse

# Same verse all day
daily-verse --daily

# Filter by book
daily-verse --book john

# Filter by testament
daily-verse --testament old
daily-verse --testament new

# Combine filters
daily-verse --daily --testament new --book romans

# Show version
daily-verse --version

# Show help
daily-verse --help
```

## Examples

```bash
# Pipe to other commands
daily-verse | cowsay

# Save to file
daily-verse --daily >> daily-log.txt

# Display in message of the day
daily-verse --daily > /etc/motd
```

## License

MIT License - see LICENSE file for details.
