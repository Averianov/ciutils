package ciutils

import (
	"crypto/tls"
	"net/http"
	"time"

	sl "github.com/Averianov/cisystemlog"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

const (
	DIR_CACHE                    string = "./certs"
	DEBUG_LETS_ENCRYPT_DIRECTORY string = "https://acme-staging-v02.api.letsencrypt.org/directory"
)

type LetsEncryptManager struct {
	*autocert.Manager
	Domain    string
	Port      string
	DebugMode bool
}

func NewLetsEncryptManager(owner, domain, port string, debug bool) (cm *LetsEncryptManager) {
	return &LetsEncryptManager{
		&autocert.Manager{
			Cache:      autocert.DirCache(DIR_CACHE),
			Prompt:     autocert.AcceptTOS,
			Email:      owner,
			HostPolicy: autocert.HostWhitelist(domain),
		}, domain, port, debug}

}

func (cm *LetsEncryptManager) getSelfSignedOrLetsEncryptCert() func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
		hello.ServerName = cm.Domain
		return cm.GetCertificate(hello)
	}
}

func (cm *LetsEncryptManager) GetLetsEncryptConfig() (cfg *tls.Config) {
	cfg = cm.TLSConfig()
	cfg.GetCertificate = cm.getSelfSignedOrLetsEncryptCert()
	return
}

func (cm *LetsEncryptManager) GetLetsEncryptServer(handler http.Handler) (letsEncryptServer *http.Server) {

	if cm.DebugMode {
		sl.L.Debug("enables debug mode - using testing URL for LetsEncrypt get cert\n%s\n", DEBUG_LETS_ENCRYPT_DIRECTORY)
		cm.Client = &acme.Client{
			DirectoryURL: DEBUG_LETS_ENCRYPT_DIRECTORY,
		}
	}

	sl.L.Debug("http made on %s\n", cm.Domain+":"+cm.Port)
	letsEncryptServer = &http.Server{
		Addr: ":" + cm.Port,
		//Addr:         cm.Domain + ":" + cm.Port,
		Handler:      cm.HTTPHandler(handler),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    cm.GetLetsEncryptConfig(),
	}

	//go http.ListenAndServe(":http", cm.HTTPHandler(nil)) // for fix issue with LetsEncrypt get certs
	return
}
