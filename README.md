# Pki
[![CircleCI](https://circleci.com/gh/hunkeelin/klinpki.svg?style=shield)](https://circleci.com/gh/hunkeelin/klinpki)
[![Go Report Card](https://goreportcard.com/badge/github.com/hunkeelin/klinpki)](https://goreportcard.com/report/github.com/hunkeelin/klinpki)
[![GoDoc](https://godoc.org/github.com/hunkeelin/klinpki?status.svg)](https://godoc.org/github.com/hunkeelin/klinpki)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hunkeelin/klinpki/master/LICENSE)

## Inspiration
Cert generate via `openssl` could be convoluted. The defaults are also insecure at version 1. This library provide an easy way to generate and sign certificates. With secure defaults. 


## Example
```
package main

import (
	"fmt"
	"testing"
)

func TestSignCsr(t *testing.T) {
	csr, err := GenerateCertificateSigningRequest(GenerateCertificateSigningRequestInput{
		EmailAddresses:     []string{"foo@klin-pro.com"},
		RsaBits:            4096,
		Province:           []string{"CA"},
		Locality:           []string{"SF"},
		Organization:       []string{"klin-pro"},
		OrganizationalUnit: []string{"IT"},
		CommonName:         "pii-vault",
		DNSNames:           []string{""},
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	ca, err := GenerateCaCertificate(GenerateCaCertificateInput{
		EmailAddresses: []string{"foo@klin-pro.com"},
		MaxDays:        30,
		RsaBits:        4096,
		Organization:   "klin-pro",
		DNSNames:       []string{"test1.klin-pro.com"},
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	cert, err := SignCsr(SignCsrInput{
		IsCa:      true,
		CaCert:    ca.Cert,
		CaKey:     ca.Key,
		Csr:       csr.Csr,
		ValidDays: 2,
	})
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(string(cert.Cert))
}
```
