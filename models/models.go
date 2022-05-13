package models

import "time"

// Bootstrap Service Registry Data Structure
type BootstrapRegistry struct {
	Version     string       `json:"version"`
	Publication time.Time    `json:"publication"`
	Description string       `json:"description"`
	Services    [][][]string `json:"services"`
}

////////////////////////////////////////////////////////////////////////////////
// Standard Data Objects
// https://datatracker.ietf.org/doc/html/rfc9083#section-4

// Links Data Structure
// https://datatracker.ietf.org/doc/html/rfc9083#section-4.2
type Links struct {
	Value    string `json:"value"`
	Rel      string `json:"rel"`
	Href     string `json:"href"`
	HrefLang string `json:"hreflang"`
	Title    string `json:"title"`
	Media    string `json:"media"`
	Type     string `json:"type"`
}

// Notices Data Structure
// https://datatracker.ietf.org/doc/html/rfc9083#section-4.3
type Notices struct {
	Title       string   `json:"title"`
	Description []string `json:"decription"`
	Links       []Links  `json:"links"`
}

// Remarks Data Structure
// https://datatracker.ietf.org/doc/html/rfc9083#section-4.3
type Remarks struct {
	// Check RFC
	Title        string   `json:"title"`
	Descriptions []string `json:"description"`
	Type         string   `json:"type"`
}

// Events Data Structure
// https://datatracker.ietf.org/doc/html/rfc9083#section-4.5
type Events struct {
	EventAction string `json:"eventAction"`
	EventActor  string `json:"eventActor"`
	EventDate   string `json:"eventDate"`
	Links       Links  `json:"links"`
}

// Public IDs Data Structure
// JSON Responses for the Registration Data Access Protocol (RDAP)
type PublicIds struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

////////////////////////////////////////////////////////////////////////////////
// Standard Object Classes
// https://datatracker.ietf.org/doc/html/rfc9083#section-5

// Entity Object Class
// https://datatracker.ietf.org/doc/html/rfc9083#section-5.1
type Entity struct {
	ObjectClassName string        `json:"objectClassName"`
	Handle          string        `json:"handle"`
	VcardArray      []interface{} `json:"vcardArray"`
	Roles           []string      `json:"roles"`
	PublicIds       []PublicIds   `json:"publicIds"`
	Entities        []Entity      `json:"entities"`
	Remarks         []Remarks     `json:"remarks"`
	Links           []Links       `json:"links"`
	Events          []Events      `json:"events"`
	AsEventActor    []Events      `json:"asEventActor"`
	Status          []string      `json:"status"`
	Port43          string        `json:"port43"`
	Networks        []IPNetwork   `json:"networks"`
	Autnums         []Autonum     `json:"autnums"`
	RdapConformance []string      `json:"rdapConformance"`
	Notices         []Notices     `json:"notices"`
}

// Nameserver Object Class
// https://datatracker.ietf.org/doc/html/rfc9083#section-5.2
type Nameserver struct {
	ObjectClassName string `json:"objectClassName"`
	Handle          string `json:"handle"`
	LdhName         string `json:"ldhName"`
	UnicodeName     string `json:"unicodeName"`
	IPAddresses     struct {
		V6 []string `json:"v6"`
		V4 []string `json:"v4"`
	} `json:"ipAddresses"`
	Entities        []Entity  `json:"entities"`
	Status          []string  `json:"status"`
	Remarks         []Remarks `json:"remarks"`
	Links           []Links   `json:"links"`
	Port43          string    `json:"port43"`
	Events          []Events  `json:"events"`
	RdapConformance []string  `json:"rdapConformance"`
	Notices         []Notices `json:"notices"`
}

// Domain Object Class
// https://datatracker.ietf.org/doc/html/rfc9083#section-5.3
type Domain struct {
	ObjectClassName string `json:"objectClassName"`
	Handle          string `json:"handle"`
	LdhName         string `json:"ldhName"`
	UnicodeName     string `json:"unicodeName"`
	Variants        []string
	Nameservers     []Nameserver `json:"nameservers"`
	SecureDNS       struct {
		ZoneSigned       bool   `json:"zoneSigned"`
		DelegationSigned bool   `json:"delegationSigned"`
		MaxSigLife       string `json:"maxSigLife"`
		DSData           []struct {
			KeyTag     int      `json:"keyTag"`
			Algorithm  int      `json:"algorithm"`
			Digest     string   `json:"digest"`
			DigestType int      `json:"digestType"`
			Events     []Events `json:"events"`
			Links      []Links  `json:"links"`
		} `json:"dsData"`
		KeyData struct {
			Flags     string   `json:"flags"`
			Protocol  string   `json:"protocol"`
			PublicKey string   `json:"publicKey"`
			Algorithm string   `json:"algorithm"`
			Events    []Events `json:"events"`
			Links     []Links  `json:"links"`
		} `json:"keyData"`
	} `json:"secureDNS"`
	Entities        []Entity    `json:"entities"`
	Status          []string    `json:"status"`
	PublicIds       []PublicIds `json:"publicIds"`
	Remarks         []Remarks   `json:"remarks"`
	Links           []Links     `json:"links"`
	Port43          string      `json:"port43"`
	Events          []Events    `json:"events"`
	Networks        []IPNetwork `json:"networks"`
	RdapConformance []string    `json:"rdapConformance"`
	Notices         []Notices   `json:"notices"`
}

// IP Network Object Class
// https://datatracker.ietf.org/doc/html/rfc9083#section-5.4
type IPNetwork struct {
	ObjectClassName string    `json:"objectClassName"`
	Handle          string    `json:"handle"`
	StartAddress    string    `json:"startAddress"`
	EndAddress      string    `json:"endAddress"`
	IPVersion       string    `json:"ipVersion"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Country         string    `json:"country"`
	ParentHandle    string    `json:"parentHandle"`
	Status          []string  `json:"status"`
	Entities        []Entity  `json:"entities"`
	Remarks         []Remarks `json:"remarks"`
	Links           []Links   `json:"links"`
	Port43          string    `json:"port43"`
	Events          []Events  `json:"events"`
	RdapConformance []string  `json:"rdapConformance"`
	Notices         []Notices `json:"notices"`
}

// Autonomous System Number Object Class
// https://datatracker.ietf.org/doc/html/rfc9083#section-5.5
type Autonum struct {
	ObjectClassName string    `json:"objectClassName"`
	Handle          string    `json:"handle"`
	StartAutnum     string    `json:"startAutnum"`
	EndAutnum       string    `json:"endAutnum"`
	Name            string    `json:"name"`
	Type            string    `json:"type"`
	Status          []string  `json:"status"`
	Country         string    `json:"country"`
	Entities        []Entity  `json:"entities"`
	Remarks         []Remarks `json:"remarks"`
	Links           []Links   `json:"links"`
	Port43          string    `json:"port43"`
	Events          []Events  `json:"events"`
	RdapConformance []string  `json:"rdapConformance"`
	Notices         []Notices `json:"notices"`
}
