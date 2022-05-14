# rdapq
![Build](https://github.com/kadonnelly13/rdapq/actions/workflows/build.yml/badge.svg)

rdapq is a simple RDAP querying tool meant to display useful details for security and network defenders.  This tool returns a subset of information from RDAP records in the CLI and for full information you can output the response data to a local JSONfile for more detailed inspection.

## Build
```bash
# Get dependencies
go get -u

# build
go build rdapq.go
```

## Usage

Basic domain query

```bash
./rdapq -domain=example.com

(+) Querying RDAP Service for domain:   example.com
(+) Finding authoritative RDAP Service URL for TLD: com
(+) Service URL for 'com' TLD: https://rdap.verisign.com/com/v1/
(*) https://rdap.verisign.com/com/v1/domain/example.com

RDAP Query Results
---------------------------------------------------------------

Domain: EXAMPLE.COM
RDAP Data Source: https://rdap.verisign.com/com/v1/domain/EXAMPLE.COM
LDH Name: EXAMPLE.COM
Unicode Name: 

Nameservers:

        LDH Name: A.IANA-SERVERS.NET
        Unicode Name: 
        Status: []
        IP Addresses
                IPv4:
                IPv6:

        LDH Name: B.IANA-SERVERS.NET
        Unicode Name: 
        Status: []
        IP Addresses
                IPv4:
                IPv6:

Domain Statuses

        Status:         client delete prohibited

        Status:         client transfer prohibited

        Status:         client update prohibited

Latest DNS Events

        Action:         registration
        Date:           1995-08-14T04:00:00Z

        Action:         expiration
        Date:           2022-08-13T04:00:00Z

        Action:         last changed
        Date:           2021-08-14T07:01:44Z

        Action:         last update of RDAP database
        Date:           2022-05-14T21:43:13Z

Notices

        Title:          Terms of Use
        Description:    Service subject to Terms of Use.
        Links:          https://www.verisign.com/domain-names/registration-data-access-protocol/terms-service/index.xhtml

        Title:          Status Codes
        Description:    For more information on domain status codes, please visit https://icann.org/epp
        Links:          https://icann.org/epp

        Title:          RDDS Inaccuracy Complaint Form
        Description:    URL of the ICANN RDDS Inaccuracy Complaint Form: https://icann.org/wicf
        Links:          https://icann.org/wicf

($) Query Completed

```

Basic IPv4 query

```bash
./rdapq -ipv4=93.184.216.34

(+) Querying RDAP Service for IPv4 address:             93.184.216.34
(+) Finding Authoritative Service URL for Range:        93.0.0.0/8
(+) Service URL for CIDR Range '93.0.0.0/8':            https://rdap.db.ripe.net/

RDAP Query Results
---------------------------------------------------------------
IP Range:               93.184.216.0 - 93.184.216.255
IP Address Name:        EDGECAST-NETBLK-03
IP Address Type:        65
Start Address Range:    93.184.216.0
End Address Range:      93.184.216.255
Parent Handle:          93.184.208.0 - 93.184.223.255

Statuses

Latest Events

        Action:         last changed
        Date:           2012-06-22T21:48:41Z

Notices

        Title:          Filtered
        Description:    This output has been filtered.

        Title:          Source
        Description:    Objects returned came from source

        Title:          Terms and Conditions
        Description:    This is the RIPE Database query service. The objects are in RDAP format.
        Link:           http://www.ripe.net/db/support/db-terms-conditions.pdf

($) Query Completed

```

Saving full output to local JSON file

`rdapq -domain=example.com -output=./example-results.json`

## What is RDAP?

RDAP (Registration Data Access Protocol) is a new protocol for registration data which will eventually replace the WHOIS protocol. More can be learned by reading the following ICANN webpage and associated RFC's.

- https://www.icann.org/rdap
- [RFC-7480 / HTTP Usage in the Registration Data Access Protocol (RDAP)](https://datatracker.ietf.org/doc/html/rfc7480)
- [RFC-7481 / Security Services for the Registration Data Access Protocol (RDAP)](https://datatracker.ietf.org/doc/html/rfc7481)
- [RFC-9082 / Registration Data Access Protocol (RDAP) Query Format](https://datatracker.ietf.org/doc/html/rfc9082)
- [RFC-9083 / JSON Responses for the Registration Data Access Protocol (RDAP)](https://datatracker.ietf.org/doc/html/rfc9083)
- [RFC-9224 / Finding the Authoritative Registration Data Access Protocol (RDAP) Service](https://datatracker.ietf.org/doc/html/rfc9224)

## To-Do
- [ ] vCard output printing
- [ ] IPv6 lookup
- [ ] Subdomain handling
- [ ] Input file handling
