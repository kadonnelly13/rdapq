package main

import (
	"flag"
	"fmt"

	s "github.com/kadonnelly13/rdapq/services"
)

const (
	RDAPServiceRegistryURL string = "https://data.iana.org/rdap/"
)

func main() {
	domain := flag.String("domain", "", "Enter FQDN")
	ipv4 := flag.String("ipv4", "", "Enter IPv4 address without CIDR range")
	outputLocation := flag.String("output", "", "Output results into JSON file at this location and filename\n(ex. -output=./test.json")
	flag.Parse()

	if *domain != "" && *ipv4 != "" {
		fmt.Printf("\n(!) You have provided too many flags. Choose to query on a domain or an IPv4 address.")
	} else if *domain != "" {
		fmt.Printf("\n(+) Querying RDAP Service for domain: %v", *domain)
		registryURL := RDAPServiceRegistryURL + "dns.json"
		s.GetDomainData(domain, registryURL, outputLocation)
	} else if *ipv4 != "" {
		fmt.Printf("\n(+) Querying RDAP Service for IPv4 address: %v", *ipv4)
		registryURL := RDAPServiceRegistryURL + "ipv4.json"
		s.GetIPv4Data(ipv4, registryURL, outputLocation)
	} else {
		fmt.Printf("\n(!) You have provided no search flags. Choose to query on a domain or an IPv4 address.")
		flag.PrintDefaults()
	}
}
