// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Verifiable digital credential.
type Credential struct {
	// Compact representation of the generated credential. Usually to be used as
	// a 'bearer' token to access a digital service or other form of protected resource.
	Token string `json:"token"`
}

// Request a new credential.
type CredentialRequest struct {
	// The principal that is the subject of the requested credential, required.
	Subject string `json:"subject"`
	// Recipients the credential is intended for, required.
	Audience []string `json:"audience"`
	// Set an expiration value for the credential.
	// A duration string is a signed sequence of decimal numbers, each with optional
	// fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units
	// are: "ns", "us" (or "µs"), "ms", "s", "m", "h"
	// Optional when generating a new token, defaults to "720h".
	Expiration string `json:"expiration"`
	// The time before which the credential MUST NOT be accepted for processing.
	// A duration string is a signed sequence of decimal numbers, each with optional
	// fraction and a unit suffix, such as "300ms", "1.5h" or "2h45m". Valid time units
	// are: "ns", "us" (or "µs"), "ms", "s", "m", "h"
	// Optional when generating a new token, defaults to "0s".
	NotBefore string `json:"notBefore"`
	// JSON-encoded custom claims to include in the credential payload.
	Payload string `json:"payload"`
}

// Request a new proof document.
type ProofRequest struct {
	// ID segment of the identifier instance to be used to generate the proof.
	ID string `json:"id"`
	// Data to use as proof content. Usually some form of challenge or data structure.
	Data string `json:"data"`
	// A string value that specifies the operational domain of a digital proof.
	Domain string `json:"domain"`
	// The specific intent for the proof, the reason why an entity created it.
	// Acts as a safeguard to prevent the proof from being misused for a purpose
	// other than the one it was intended for. For example, a proof can be used
	// for purposes of authentication, for asserting control of a Verifiable
	// Credential (assertionMethod), and several others.
	//
	// Common values include: authentication, assertionMethod, keyAgreement,
	// capabilityInvocation, capabilityDelegation.
	//
	// https://w3c-ccg.github.io/ld-proofs/#proof-purpose
	Purpose string `json:"purpose"`
}

// PublicKey represents a cryptographic key according to the "Linked Data
// Cryptographic Suites".
//
// https://w3c-ccg.github.io/ld-cryptosuite-registry/
type PublicKey struct {
	// Unique identifier for the key reference.
	ID string `json:"id"`
	// Cryptographic suite identifier.
	Kind string `json:"kind"`
	// Subject controlling the corresponding private key.
	Controller string `json:"controller"`
	// Public key value. Encoded in base64 as defined in RFC 4648
	Value string `json:"value"`
}

// The DateFormat enum list all the supported styles available
// when formatting Time values.
type DateFormat string

const (
	// Example: "2006-01-02T15:04:05Z07:00"
	DateFormatRfc3339 DateFormat = "RFC3339"
	// Example: "02 Jan 06 15:04 MST"
	DateFormatRfc822 DateFormat = "RFC822"
	// Returns the date as its UNIX timestamp value (in seconds)
	DateFormatUnix DateFormat = "UNIX"
)

var AllDateFormat = []DateFormat{
	DateFormatRfc3339,
	DateFormatRfc822,
	DateFormatUnix,
}

func (e DateFormat) IsValid() bool {
	switch e {
	case DateFormatRfc3339, DateFormatRfc822, DateFormatUnix:
		return true
	}
	return false
}

func (e DateFormat) String() string {
	return string(e)
}

func (e *DateFormat) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DateFormat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DateFormat", str)
	}
	return nil
}

func (e DateFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Supported standard for generating hash values. All hashes are
// displayed as hex encoded strings.
type Digest string

const (
	// Returns a 32 byte hash using the SHA256 standard.
	DigestSha2 Digest = "SHA2"
	// Returns a 32 byte hash using the SHA3-256 standard.
	DigestSha3 Digest = "SHA3"
)

var AllDigest = []Digest{
	DigestSha2,
	DigestSha3,
}

func (e Digest) IsValid() bool {
	switch e {
	case DigestSha2, DigestSha3:
		return true
	}
	return false
}

func (e Digest) String() string {
	return string(e)
}

func (e *Digest) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Digest(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Digest", str)
	}
	return nil
}

func (e Digest) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Supported representation formats for linked data (LD) documents
type DocumentMode string

const (
	// RDF dataset on the JSON-LD document, the algorithm used is "URDNA2015"
	// and the format "application/n-quads.
	//
	// https://json-ld.github.io/normalization/spec
	DocumentModeNormalized DocumentMode = "NORMALIZED"
	// Expanded JSON-LD document.
	//
	// http://www.w3.org/TR/json-ld-api/#expansion-algorithm
	DocumentModeExpanded DocumentMode = "EXPANDED"
)

var AllDocumentMode = []DocumentMode{
	DocumentModeNormalized,
	DocumentModeExpanded,
}

func (e DocumentMode) IsValid() bool {
	switch e {
	case DocumentModeNormalized, DocumentModeExpanded:
		return true
	}
	return false
}

func (e DocumentMode) String() string {
	return string(e)
}

func (e *DocumentMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DocumentMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DocumentMode", str)
	}
	return nil
}

func (e DocumentMode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
