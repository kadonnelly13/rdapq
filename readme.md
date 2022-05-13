# rdapq

rdapq is a minimalistic RDAP querying tool meant to display useful details for security and network defenders.  This tool returns a subset of information from RDAP records in the CLI and for full information you can output the response data to a local JSONfile for more detailed inspection.

## Build
```bash
# Get dependencies
go get -u

# build
go build rdapq.go
```

## Usage

Basic domain query

`rdapq -domain=example.com`

Basic IPv4 query

`rdapq -ipv4=8.8.8.8`

Full output

`rdapq -domain=example.com -output=./example-results.json`

## What is RDAP?

RDAP (Registration Data Access Protocol) is a new protocol for registration data which will eventually replace the WHOIS protocol. More can be learned by reading the following ICANN webpage and associated RFC's.

- https://www.icann.org/rdap
- (RFC-7480 / HTTP Usage in the Registration Data Access Protocol (RDAP))[https://datatracker.ietf.org/doc/html/rfc7480]
- (RFC-7481 / Security Services for the Registration Data Access Protocol (RDAP))[https://datatracker.ietf.org/doc/html/rfc7481]
- (RFC-9082 / Registration Data Access Protocol (RDAP) Query Format)[https://datatracker.ietf.org/doc/html/rfc9082]
- (RFC-9083 / JSON Responses for the Registration Data Access Protocol (RDAP))[https://datatracker.ietf.org/doc/html/rfc9083]
- (RFC-9224 / Finding the Authoritative Registration Data Access Protocol (RDAP) Service)[https://datatracker.ietf.org/doc/html/rfc9224]

