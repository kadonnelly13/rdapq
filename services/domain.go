package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	m "github.com/kadonnelly13/rdapq/models"
)

// Query domain service
func GetDomainData(domain *string, registryURL string, outputLocation *string) {
	var authoritativeServerData m.Domain
	var secondaryRDAPServerData m.Domain

	TLD := parseDomain(*domain)
	URL := getAuthoritativeDomainServerURL(TLD, registryURL)
	domainURL := URL + "domain/" + *domain

	authoritativeServerData = queryAuthoritativeDomainServer(domainURL)

	prettyPrintDomainData(authoritativeServerData)

	// Check if links has a "related" HREF and query to return
	for _, link := range authoritativeServerData.Links {
		if link.Rel == "related" {
			fmt.Printf("\n\n\nAnother RDAP server found. Printing data...")
			secondaryRDAPServerResponseURL := link.Href

			secondaryRDAPServerData = queryAuthoritativeDomainServer(secondaryRDAPServerResponseURL)

			prettyPrintDomainData(secondaryRDAPServerData)
		}
	}

	// Save to file to output location
	if *outputLocation != "" {
		outputStart := []byte("[\n")

		outputDataAuthoritative, _ := json.MarshalIndent(authoritativeServerData, "", "\t")
		outputFile := append(outputStart, outputDataAuthoritative...)

		// Adding for secondary data
		outputMiddle := []byte(",\n")
		outputFile = append(outputFile, outputMiddle...)

		outputDataSecondary, _ := json.MarshalIndent(secondaryRDAPServerData, "", "\t")
		outputFile = append(outputFile, outputDataSecondary...)

		// Final
		outputEnd := []byte("\n]")
		outputFile = append(outputFile, outputEnd...)

		err := os.WriteFile(*outputLocation, outputFile, 0644)
		if err != nil {
			fmt.Printf("\n(!) Error creating output data file\n%v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n\n($) Query Completed\n\n")

}

// https://datatracker.ietf.org/doc/html/rfc9224#section-4
func getAuthoritativeDomainServerURL(TLD string, registryURL string) string {
	var bootstrapRegistryData m.BootstrapRegistry

	fmt.Printf("\n(+) Finding authoritative RDAP Service URL for TLD: %v", TLD)

	queryResponse, err := http.Get(registryURL)

	if err != nil {
		if os.IsTimeout(err) {
			fmt.Printf("\n(!) Timeout querying IANA RDAP service registry. Please try again.\n")
			os.Exit(1)
		}
		fmt.Printf("\n(!) Error querying IANA RDAP service registry:\n%v", err)
		os.Exit(1)
	}

	queryResponseBody, err := io.ReadAll(queryResponse.Body)
	queryResponse.Body.Close()

	if err != nil {
		fmt.Printf("\n(!) Error reading query response:\n%v", err)
		os.Exit(1)
	} else if queryResponse.StatusCode == 200 {
		err = json.Unmarshal(queryResponseBody, &bootstrapRegistryData)

		if err != nil {
			fmt.Printf("\n(!) Error un-marshalling query response body:\n%v", err)
			os.Exit(1)
		}

		// Loop through response data to find authoritative server's URL
		for _, service := range bootstrapRegistryData.Services {
			for _, serviceTLD := range service[0] {
				if serviceTLD == TLD {
					fmt.Printf("\n(+) Service URL for '%s' TLD: %v", serviceTLD, service[1][0])
					// Returning URL
					return service[1][0]
				}
			}
		}
	} else if queryResponse.StatusCode == 429 {
		// Querying too much 429 returned from IANA
		fmt.Printf("\n(!) Returned 429...Slow down there cowboy on the requests you are being throttled. Go take a lap around the neighboorhood before your next query.")
		os.Exit(1)
	} else {
		fmt.Printf("\n(!) Did not recieve \"200 OK\" status code from IANA query: %v", queryResponse.StatusCode)
		os.Exit(1)
	}

	return ""
}

func queryAuthoritativeDomainServer(RDAPServerURL string) m.Domain {
	var ResponseData m.Domain

	queryResponse, err := http.Get(RDAPServerURL)

	if err != nil {
		if os.IsTimeout(err) {
			fmt.Printf("\n(!) Timeout querying remote RDAP server. Please try again.\n")
			os.Exit(1)
		}
		fmt.Printf("\n(!) Error querying RDAP service URL:\n%v", err)
		os.Exit(1)
	}

	queryResponseBody, err := io.ReadAll(queryResponse.Body)
	queryResponse.Body.Close()

	if err != nil {
		fmt.Printf("\n(!) Error reading query response:\n%v", err)
		os.Exit(1)
	} else if queryResponse.StatusCode == 200 {
		err = json.Unmarshal(queryResponseBody, &ResponseData)

		if err != nil {
			fmt.Printf("\n(!) Error un-marshalling query response body:\n%v", err)
			os.Exit(1)
		}
		return ResponseData
	} else if queryResponse.StatusCode == 429 {
		// Querying too much 429 returned from IANA
		fmt.Printf("\n(!) Returned 429...Slow down there cowboy on the requests you are being throttled. Go take a lap around the neighboorhood before your next query.")
		os.Exit(1)
	} else {
		fmt.Printf("\n(!) Did not recieve \"200 OK\" status code from IANA query: %v", queryResponse.StatusCode)
		os.Exit(1)
	}

	return ResponseData
}

// Parse domain to get TLD
func parseDomain(domainName string) (TLD string) {
	re := regexp.MustCompile(`[[:alnum:]]+.*\.(?P<TLD>.*)`)
	parsedTLDMatch := re.FindStringSubmatch(domainName)
	TLDIndex := re.SubexpIndex("TLD")
	TLD = parsedTLDMatch[TLDIndex]

	return TLD
}

// Pretty print domain data
func prettyPrintDomainData(serverResponseData m.Domain) {
	fmt.Printf("\n\nRDAP Query Results")
	fmt.Printf("\n---------------------------------------------------------------")
	fmt.Printf("\n\nDomain: %v", serverResponseData.LdhName)
	fmt.Printf("\nRDAP Data Source: %v", serverResponseData.Links[0].Value)
	fmt.Printf("\nLDH Name: %v", serverResponseData.LdhName)
	fmt.Printf("\nUnicode Name: %v", serverResponseData.UnicodeName)

	// Printing Nameservers
	fmt.Printf("\n\nNameservers:")
	for _, nameserver := range serverResponseData.Nameservers {
		fmt.Printf("\n\n\tLDH Name: %v", nameserver.LdhName)
		fmt.Printf("\n\tUnicode Name: %v", nameserver.UnicodeName)
		fmt.Printf("\n\tStatus: %v", nameserver.Status)

		fmt.Printf("\n\tIP Addresses")
		fmt.Printf("\n\t\tIPv4:")
		for _, v4 := range nameserver.IPAddresses.V4 {
			fmt.Printf("\t\t%v", v4)
		}
		fmt.Printf("\n\t\tIPv6:")
		for _, v6 := range nameserver.IPAddresses.V6 {
			fmt.Printf("\n\t\t%v", v6)
		}
	}

	// Printing Statuses
	fmt.Printf("\n\nDomain Statuses")
	for _, status := range serverResponseData.Status {
		fmt.Printf("\n\n\tStatus:\t\t%v", status)
	}

	// Printing Events
	fmt.Printf("\n\nLatest DNS Events")
	for _, event := range serverResponseData.Events {
		fmt.Printf("\n\n\tAction:\t\t%v", event.EventAction)
		fmt.Printf("\n\tDate:\t\t%v", event.EventDate)
	}

	// Printing Notices
	fmt.Printf("\n\nNotices")
	for _, notice := range serverResponseData.Notices {
		fmt.Printf("\n\n\tTitle:\t\t%v", notice.Title)
		fmt.Printf("\n\tDescription:\t%v", notice.Descriptions[0])
		fmt.Printf("\n\tLinks:\t\t%v", notice.Links[0].Href)
	}

	// Printing Entities
	for _, entity := range serverResponseData.Entities {
		fmt.Printf("\n\nEntities:")
		fmt.Printf("\n\tType: %v", entity.PublicIds[0].Type)
		fmt.Printf("\n\tHandle: %v", entity.Handle)
		fmt.Printf("\n\tRole: %v", entity.Roles[0])
		fmt.Printf("\n\tvCard Data:")

		/*
			To-Do
			Better print out vCard data
		*/
		for _, vcard := range entity.VcardArray {
			if vcard != "vcard" {
				switch data := vcard.(type) {
				case []interface{}:
					for _, d := range data {
						fmt.Printf("\n\t\t%+v", d)
					}
				}
				fmt.Printf("\n")
			}
		}
	}
}
