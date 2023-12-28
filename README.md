----
shrewdeye seamlessly integrates with the Shrewdeye API to uncover subdomains for a given domain or a list of domains. It streamlines the process of interacting with the API, making subdomain discovery effortless.
----
## Installation
shrewdeye requires Golang to install successfully. Run the following command to install:
```sh
go install github.com/omkar7505/shrewdeye@latest
```
## Usage:
```sh
shrewdeye [flags]
```
**Flags:**
```
-d - Single domain to query (required if -i is not used).
-i - File containing a list of domains to query (required if -d is not used).
-v - Only return subdomains with valid DNS information (default: false).
-o - File to write results to (default: print to console).
-h - Display help message.
```

## Example Usage:

Discover subdomains for a single domain:
```sh
shrewdeye -d google.com
```
Process multiple domains from a file:
```sh
shrewdeye -i domains.txt
```
Filter for valid DNS records and save results:
```sh
shrewdeye -d google.com -v -o results.txt
```
