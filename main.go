package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func makeRequest(domain string, outputFileName string, delay time.Duration, valid bool) error {
	if valid{
		url := fmt.Sprintf("https://shrewdeye.app/domains/%s.txt?valid=true", domain)

		// Introduce delay if -i flag is supplied
		time.Sleep(delay)
	
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			return nil
		}
	
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
	
		fmt.Printf("%s", string(body))
	
		// Save the result if -o flag is supplied
		if outputFileName != "" {
			err := appendToFile(outputFileName, body)
			if err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}
		}
	
		return nil
	} else{
		url := fmt.Sprintf("https://shrewdeye.app/domains/%s.txt", domain)

		// Introduce delay if -i flag is supplied
		time.Sleep(delay)
	
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()
	
		if response.StatusCode != http.StatusOK {
			return nil
		}
	
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
	
		fmt.Printf("%s", string(body))
	
		// Save the result if -o flag is supplied
		if outputFileName != "" {
			err := appendToFile(outputFileName, body)
			if err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}
		}
	
		return nil
	}

}

func appendToFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

func main() {

	domainFlag := flag.String("d", "", "domains to find subdomains for")
	inputFileFlag := flag.String("i", "", "file containing list of domains for subdomain discovery")
	validFlag := flag.Bool("v",false,"subdomains with DNS information only")
	outputFileFlag := flag.String("o", "", "file to write output to")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: \n\t%s  [flags]\n\n", os.Args[0])
        flag.PrintDefaults()
    }
    flag.Parse()


	if *domainFlag == "" && *inputFileFlag == "" {
		fmt.Println("Please provide either -d or -i flag")
		os.Exit(1)
	}

	// Set the delay to 1 second when -i flag is supplied
	var delay time.Duration
	if *inputFileFlag != "" {
		delay = time.Second
	}

	if *domainFlag != "" {
		// Make a request for a single domain
		err := makeRequest(*domainFlag, *outputFileFlag, delay, *validFlag)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Make requests for domains from the input file
		file, err := os.Open(*inputFileFlag)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := NewLineScanner(file)
		for scanner.Scan() {
			domain := scanner.Text()
			err := makeRequest(domain, *outputFileFlag, delay, *validFlag)
			if err != nil {
				fmt.Printf("Error for domain %s: %v\n", domain, err)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
	}
}

// NewLineScanner creates a new scanner that splits on newline characters.
func NewLineScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}
