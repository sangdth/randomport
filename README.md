# RandomPort - Terminal Port Generator

<img width="498" alt="Screenshot 2025-04-01 at 23 58 03" src="https://github.com/user-attachments/assets/07b4e2e4-f6fc-4898-8cd3-39afaa8706ed" />


## Introduction

RandomPort is a terminal-based application that generates random port numbers within a specified range. It was designed to:

- Serve as a learning project for Go programming
- Provide a practical solution for generating random ports so I can use to expose my PostgresQL instances.

### Features

- Generates random ports within a configurable range (default: 55000-55999)
- Terminal UI with:
  - Min, max, and result fields
  - Navigation between fields
  - Regeneration of random port
- Input validation for port numbers

## Installation

### Prerequisites

- Go 1.24.1 or higher

### Installation Methods

#### Using go install

```bash
go install github.com/sangdth/randomport@latest
```

#### Building from Source

```bash
git clone https://github.com/sangdth/randomport.git
cd randomport
go build -o randomport
```

### Verification

After installation, verify it works by running:

```bash
randomport
```

## Usage

### Basic Usage

1. Run the application:

   ```bash
   randomport
   ```

2. Use arrow keys to navigate between fields
3. Enter minimum and maximum port values
4. Press Enter to generate a random port
5. Press 'r' to regenerate a new port

### Examples

Generate a port between 5000 and 6000:

```bash
randomport
```

Then enter:

- Min: 5000
- Max: 6000

## Contributing

We welcome contributions! Here's how to get started:

### Development Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/sangdth/randomport.git
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

### Code Guidelines

- Follow Go's official coding standards
- Write clear, concise comments
- Include tests for new features

### Pull Request Process

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a pull request

### Reporting Issues

Please report any issues through GitHub Issues with:

- Detailed description
- Steps to reproduce
- Expected vs actual behavior
