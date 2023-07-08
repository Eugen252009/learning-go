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
	fmt.Printf("domain,hasMX,hasSPF,spfRecord,hasDMARK,dmarkRecord\n")
	for scanner.Scan() {
		checkdomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not Read from Input: %v\n", err)
	}
}

func checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARK bool
	var spfRecord, dmarkRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarkRecords, err := net.LookupTXT("_dmark." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)

	}
	for _, record := range dmarkRecords {
		if strings.HasPrefix(record, "v=DMARKC1") {
			hasDMARK = true
			dmarkRecord = record
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v\n", domain, hasMX, hasSPF, spfRecord, hasDMARK, dmarkRecord)
}
