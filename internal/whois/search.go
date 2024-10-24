package whois

import (
	"fmt"
	"os"

	"github.com/likexian/whois"
)

func Search(domain string) {
	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println(result)
}