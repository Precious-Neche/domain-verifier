package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter email domain and we will check if it has a:\n valid domain\n Mx records\n SPF records\n sprRecord\n DMARC \n dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSpf, hasDmarc bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err == nil && len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSpf = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX: %t\n", hasMx)
	fmt.Printf("Has SPF: %t\n", hasSpf)
	fmt.Printf("SPF Record: %s\n", spfRecord)
	fmt.Printf("Has DMARC: %t\n", hasDmarc)
	fmt.Printf("DMARC Record: %s\n", dmarcRecord)
	fmt.Println("-----")

}
