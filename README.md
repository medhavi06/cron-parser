# Cron Expression Parser

### Prerequisites
- Go 1.21+

### Installation
```bash
# Clone the repository
git clone https://github.com/medhavi06/cron-parser.git

# Change to project directory
cd cron-parser

# Build the application
go build 
```

### Usage
```bash
# Run the parser
./cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"

# Expected output:
# minute        0 15 30 45
# hour          0
# day of month  1 15
# month         1 2 3 4 5 6 7 8 9 10 11 12
# day of week   1 2 3 4 5
# command       /usr/bin/find
```

## Running Tests
```bash
# Run unit tests
go test ./test/...