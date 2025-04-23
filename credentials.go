package ciutils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

// LoadTLSCredentials load certificate for client from file
func LoadTLSCredentials(tlsCertPath string) (config *tls.Config, err error) {
	pemServerCA, err := os.ReadFile(tlsCertPath)
	if err != nil {
		return
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		err = fmt.Errorf("failed to add server CA's certificate")
		return
	}

	config = &tls.Config{
		RootCAs: certPool,
	}
	return
}

// GetTLSCredentials get certificate for client from server
func GetTLSCredentials(domain string) (config *tls.Config, err error) {

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	var certificate *x509.Certificate
	var req *http.Request
	url := "https://" + domain
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.TLS != nil {
		certificates := resp.TLS.PeerCertificates
		if len(certificates) > 0 {
			certificate = certificates[0]
		}
	}

	// ###########################################################
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)

	_, err = certificate.Verify(x509.VerifyOptions{Roots: certPool})
	if err != nil {
		return nil, err
	}
	config = &tls.Config{
		RootCAs: certPool,
	}
	return
}
