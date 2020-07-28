package main

import (
	"github.com/hunkeelin/klinpki"
	"os"
)

func main() {
	ca, err := klinpki.GenerateCaCertificate(klinpki.GenerateCaCertificateInput{
		EmailAddresses: []string{"foo@klin-pro.com"},
		MaxDays:        7200,
		RsaBits:        4096,
		Organization:   "klin-pro",
		DNSNames:       []string{"test1.klin-pro.com", "test1.ca.klin-pro.com"},
	})
	if err != nil {
		panic(err)
	}
	client, err := klinpki.GenerateCertificateSigningRequest(klinpki.GenerateCertificateSigningRequestInput{
		EmailAddresses:     []string{"devops@klin-pro.com"},
		RsaBits:            2048,
		Province:           []string{"CA"},
		Locality:           []string{"SF"},
		Organization:       []string{"klin-pro"},
		OrganizationalUnit: []string{"IT"},
		CommonName:         "test1",
		DNSNames:           []string{"test1-client.prod.klin-pro.com"},
	})
	if err != nil {
		panic(err)
	}
	service, err := klinpki.GenerateCertificateSigningRequest(klinpki.GenerateCertificateSigningRequestInput{
		EmailAddresses:     []string{"devops@klin-pro.com"},
		RsaBits:            2048,
		Province:           []string{"CA"},
		Locality:           []string{"SF"},
		Organization:       []string{"klin-pro"},
		OrganizationalUnit: []string{"IT"},
		CommonName:         "test1",
		DNSNames:           []string{"test1-service.prod.klin-pro.com", "test1.klin-pro.com"},
	})
	if err != nil {
		panic(err)
	}
	clientCert, err := klinpki.SignCsr(klinpki.SignCsrInput{
		IsCa:      false,
		CaCert:    ca.Cert,
		CaKey:     ca.Key,
		Csr:       client.Csr,
		ValidDays: 3650,
	})
	if err != nil {
		panic(err)
	}
	serviceCert, err := klinpki.SignCsr(klinpki.SignCsrInput{
		IsCa:      false,
		CaCert:    ca.Cert,
		CaKey:     ca.Key,
		Csr:       service.Csr,
		ValidDays: 3650,
	})
	if err != nil {
		panic(err)
	}

	// Writing client cert
	clientCertFile, err := os.Create("client.crt")
	if err != nil {
		panic(err)
	}
	defer clientCertFile.Close()
	_, err = clientCertFile.Write(clientCert.Cert)
	if err != nil {
		panic(err)
	}
	// Writing client key
	clientKeyFile, err := os.Create("client.key")
	if err != nil {
		panic(err)
	}
	defer clientKeyFile.Close()
	_, err = clientKeyFile.Write(client.Key)
	if err != nil {
		panic(err)
	}

	// Writing service cert
	serviceCertFile, err := os.Create("service.crt")
	if err != nil {
		panic(err)
	}
	defer clientCertFile.Close()
	_, err = serviceCertFile.Write(serviceCert.Cert)
	if err != nil {
		panic(err)
	}

	// Writing service key
	serviceKeyFile, err := os.Create("service.key")
	if err != nil {
		panic(err)
	}
	defer clientKeyFile.Close()
	_, err = serviceKeyFile.Write(service.Key)
	if err != nil {
		panic(err)
	}

	// Writing ca cert
	caCertFile, err := os.Create("ca.crt")
	if err != nil {
		panic(err)
	}
	defer clientCertFile.Close()
	_, err = caCertFile.Write(ca.Cert)
	if err != nil {
		panic(err)
	}

	// Writing ca key

	caKeyFile, err := os.Create("ca.key")
	if err != nil {
		panic(err)
	}
	defer clientKeyFile.Close()
	_, err = caKeyFile.Write(ca.Key)
	if err != nil {
		panic(err)
	}

}
