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
	fmt.Printf("domain, hasMX, hasSPF, sprREcord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		CheckDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v\n", err)
	}
}

func CheckDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err1 := net.LookupTXT(domain)

	if err1 != nil {
		log.Printf("Error: %v\n", err1)
	}

	for _, rec := range txtRecords {
		if strings.HasPrefix(rec, "v=spf1") {
			hasSPF = true
			spfRecord = rec
			break
		}
	}

	txtDmarc, err2 := net.LookupTXT("_dmarc." + domain)

	if err2 != nil {
		log.Printf("Error: %v\n", err2)
	}

	for _, rec := range txtDmarc {
		if strings.HasPrefix(rec, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = rec
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
