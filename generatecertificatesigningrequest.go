package klinpki

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
)

// GenerateCertificateSigningRequestInput the input for the function
type GenerateCertificateSigningRequestInput struct {
	RsaBits               int                      // RsaBits
	Country               []string                 // Country
	Province              []string                 // Province
	Locality              []string                 // Locality
	EmailAddresses        []string                 // EmailAddresses
	Organization          []string                 // Orginzation
	OrganizationalUnit    []string                 // OrginzationUnit
	CommonName            string                   // CommonName
	DNSNames              []string                 // DnsNames
	Password              string                   // Password
	RawCertificateRequest *x509.CertificateRequest // RawCertificateRequest
}

// GenerateCertificateSigningRequestOutput the output for the function
type GenerateCertificateSigningRequestOutput struct {
	Csr []byte // Csr
	Key []byte // Key
}

// GenerateCertificateSigningRequest the function that signs csr
func GenerateCertificateSigningRequest(g GenerateCertificateSigningRequestInput) (GenerateCertificateSigningRequestOutput, error) {

	priv, err := rsa.GenerateKey(rand.Reader, g.RsaBits)
	if err != nil {
		return GenerateCertificateSigningRequestOutput{}, err
	}
	subj := pkix.Name{
		CommonName:         g.CommonName,
		Country:            g.Country,
		Province:           g.Province,
		Locality:           g.Locality,
		OrganizationalUnit: g.OrganizationalUnit,
		Organization:       g.Organization,
	}
	var template x509.CertificateRequest
	template = x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
		DNSNames:           g.DNSNames,
		EmailAddresses:     g.EmailAddresses,
	}
	template.DNSNames = append(template.DNSNames, g.CommonName)
	if g.RawCertificateRequest != nil {
		template = *g.RawCertificateRequest
	}
	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, priv)
	if err != nil {
		return GenerateCertificateSigningRequestOutput{}, err
	}
	keyPem, err := pemBlockForKey(g.Password, priv)
	if err != nil {
		return GenerateCertificateSigningRequestOutput{}, err
	}
	return GenerateCertificateSigningRequestOutput{
		Csr: pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes}),
		Key: pem.EncodeToMemory(keyPem),
	}, nil
}
func pemBlockForKey(password string, priv interface{}) (*pem.Block, error) {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		var encryptedkey *pem.Block
		var err error
		if password != "" {
			encryptedkey, err = x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(k), []byte(password), x509.PEMCipherAES256)
			if err != nil {
				return &pem.Block{}, err
			}
		} else {
			encryptedkey = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
		}
		return encryptedkey, nil
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return &pem.Block{}, err
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}, nil
	default:
		return nil, fmt.Errorf("No valid digital signiture")
	}
}
