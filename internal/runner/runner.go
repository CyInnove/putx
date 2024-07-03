package runner

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/CyInnove/putx/config"
)

// Run starts the PUT test for each domain in the file.
func Run(filePath string, outputFile string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var wg sync.WaitGroup
	var output *os.File

	if outputFile != "" {
		output, err = os.Create(outputFile)
		if err != nil {
			fmt.Printf("Failed to create output file: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		wg.Add(1)
		go putTest(domain, output, &wg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	wg.Wait()
}

func putTest(domain string, output *os.File, wg *sync.WaitGroup) {
	defer wg.Done()

	url := domain + "/evil.txt"
	data := []byte("hello world")
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Failed to create request for %s\n", domain)
		return
	}

	client := config.GetHTTPClient()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request failed for %s\n", domain)
		return
	}
	defer resp.Body.Close()

	result := fmt.Sprintf("URL: %s - Response: %d\n", url, resp.StatusCode)
	fmt.Print(result)

	if output != nil {
		_, err := output.WriteString(result)
		if err != nil {
			fmt.Printf("Failed to write to output file\n")
		}
	}
}
