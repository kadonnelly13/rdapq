package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	m "github.com/kadonnelly13/rdapq/models"
)

func GetIPv4Data(ipv4 *string, registryURL string, outputLocation *string) {
	fullIPv4 := getFullIPv4(*ipv4)
	URL := getAuthoritativeIPServerURL(fullIPv4, registryURL)
	ipv4URL := URL + "ip/" + *ipv4

	authoritativeServerData := queryAuthoritativeIPServer(ipv4URL)

	prettyPrintIPData(authoritativeServerData)

	// Save to file to output location
	if *outputLocation != "" {
		outputFile, _ := json.MarshalIndent(authoritativeServerData, "", " ")
		err := os.WriteFile(*outputLocation, outputFile, 0644)
		if err != nil {
			fmt.Printf("\n(!) Error writing data to output file:\n%v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\n\n($) Query Completed\n\n")
}

// https://datatracker.ietf.org/doc/html/rfc9224#section-4
func getAuthoritativeIPServerURL(fullIPv4 string, registryURL string) string {
	var bootstrapRegistryData m.BootstrapRegistry

	fmt.Printf("\n(+) Finding authoritative RDAP Service URL for IPv4: %v", fullIPv4)

	queryResponse, err := http.Get(registryURL)

	if err != nil {
		fmt.Printf("\n(!) Error querying IANA RDAP service registry:\n%v\n", err)
		os.Exit(1)
	}

	queryResponseBody, err := io.ReadAll(queryResponse.Body)
	queryResponse.Body.Close()

	if err != nil {
		fmt.Printf("\n(!) Error reading query response:\n%v", err)
		os.Exit(0)
	} else if queryResponse.StatusCode == 200 {
		err = json.Unmarshal(queryResponseBody, &bootstrapRegistryData)

		if err != nil {
			fmt.Printf("\n(!) Error un-marshalling query response body:\n%v\n", err)
			os.Exit(1)
		}

		// Loop through response data to find authoritative server's URL
		for _, service := range bootstrapRegistryData.Services {
			for _, serviceV4 := range service[0] {
				if serviceV4 == fullIPv4 {
					fmt.Printf("\n(+) Service URL for '%s' TLD: %v", serviceV4, service[1][0])
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

func queryAuthoritativeIPServer(RDAPServerURL string) m.IPNetwork {
	var ResponseData m.IPNetwork

	client := &http.Client{}

	requestRDAPServer, err := http.NewRequest("GET", RDAPServerURL, nil)

	if err != nil {
		fmt.Printf("\n(!) Error setting up new request to server:\n%v", err)
		os.Exit(1)
		return ResponseData
	}

	requestRDAPServer.Header.Add("Accept", "application/rdap+json")

	responseRDAPService, err := client.Do(requestRDAPServer)

	if err != nil {
		fmt.Printf("\n(!) Error sending new request to server:\n%v", err)
		os.Exit(1)
		return ResponseData
	}

	responseRDAPServerData, err := ioutil.ReadAll(responseRDAPService.Body)

	if err != nil {
		fmt.Printf("\n(!) Error reading new request to server:\n%v\n", err)
		os.Exit(1)
		return ResponseData
	}

	responseRDAPService.Body.Close()

	err = json.Unmarshal(responseRDAPServerData, &ResponseData)

	if err != nil {
		fmt.Printf("\n(!) Error unmarshalling response data:\n%v\n", err)
		os.Exit(1)
		return ResponseData
	}

	return ResponseData
}

// Parse IPv4 address to get /8 CIDR range
func getFullIPv4(ipv4 string) string {
	// Attach /8 CIDR range to IPv4 address
	ipv4 = ipv4 + "/8"
	_, ipv4Net, err := net.ParseCIDR(ipv4)
	if err != nil {
		fmt.Printf("\n(!) Error parsing IPv4 address IANA RDAP service registry:\n%v", err)
		os.Exit(1)
	}
	return ipv4Net.String()
}

// Pretty print ipv4 data
func prettyPrintIPData(serverResponseData m.IPNetwork) {
	fmt.Printf("\n\nRDAP Data for IPv4: %v", serverResponseData.Handle)
	fmt.Printf("\nIP Address Name: %v", serverResponseData.Name)
	fmt.Printf("\nIP Address Type: %v", serverResponseData.Type[0])

	// Printing entites
	fmt.Printf("\n\nEntites:")
	for _, entity := range serverResponseData.Entities {
		fmt.Printf("\n\n\tHandle: %v", entity.Handle)
		fmt.Printf("\n\tRole: %v", entity.Roles)

		/*
			To-Do
			Convert vCard interface{} to output
			fmt.Printf("\n\tvCard Data:")
			for _, vcard := range entity.VcardArray {
				fmt.Printf("\n\t\tName: %v", vcard)
			}
		*/
	}

	// Printing latest events
	fmt.Printf("\n\nLatest Events")
	for _, event := range serverResponseData.Events {
		fmt.Printf("\n\n\tAction:\t\t%v", event.EventAction)
		fmt.Printf("\n\tDate:\t\t%v", event.EventDate)
	}
}
