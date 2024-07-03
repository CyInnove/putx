# putx

`putx` is a command-line tool written in Go for testing PUT request functionality against a list of URLs. It sends a PUT request with a sample payload to each URL specified in a file and logs the response status code.

## Features

- Sends PUT requests to multiple URLs concurrently.
- Logs the URL and response status code for each request.
- Supports skipping TLS certificate verification for testing purposes.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/CyInnove/putx.git
   ```

2. **Build the executable:**
   ```bash
   cd putx
   go build -o putx ./cmd/putx
   ```

## Usage

```bash
./putx -l <file> [-o <output>]
```

- `-l <file>`: Specifies the path to a file containing a list of base URLs to test. Each URL should be on a new line.
- `-o <output>`: Optional. Specifies the output file to write the results to.

Example:
```bash
./putx -l hosts.txt -o output.txt
```

## Configuration

The tool uses a default HTTP client configuration that skips TLS certificate verification. This can be adjusted in `config/config.go`.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
