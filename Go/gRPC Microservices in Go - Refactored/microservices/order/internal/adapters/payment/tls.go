package payment

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"path/filepath"

	"google.golang.org/grpc/credentials"
)

// ClientCredentials builds mTLS credentials: Order shows client-cert.pem and
// trusts only a Payment certificate signed by ca-cert.pem. The server name is
// checked against the host in PAYMENT_SERVICE_URL (for example "payment").
func ClientCredentials(certDir string) (credentials.TransportCredentials, error) {
	clientCert, err := tls.LoadX509KeyPair(
		filepath.Join(certDir, "client-cert.pem"),
		filepath.Join(certDir, "client-key.pem"),
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
		return nil, errors.New("failed to append CA cert")
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
	}), nil
}
