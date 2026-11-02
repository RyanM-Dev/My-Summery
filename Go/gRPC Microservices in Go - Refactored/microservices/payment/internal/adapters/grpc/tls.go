package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"path/filepath"

	"google.golang.org/grpc/credentials"
)

// ServerCredentials builds mTLS credentials: the server shows server-cert.pem
// and accepts only clients whose certificate is signed by ca-cert.pem.
func ServerCredentials(certDir string) (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair(
		filepath.Join(certDir, "server-cert.pem"),
		filepath.Join(certDir, "server-key.pem"),
	)
	if err != nil {
		return nil, err
	}

	caCert, err := os.ReadFile(filepath.Join(certDir, "ca-cert.pem"))
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("failed to append the CA certs")
	}

	return credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{serverCert},
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
	}), nil
}
