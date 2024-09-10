package SSL

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"os"
)

func MakeServerCredentials(serverCrtPath, serverKeyPath, caCrtPath string) (credentials.TransportCredentials, error) {
	crt, err := tls.LoadX509KeyPair(serverCrtPath, serverKeyPath)
	if err != nil {
		return nil, err
	}
	crtPool := x509.NewCertPool()
	caCrt, err := os.ReadFile(caCrtPath)
	if err != nil {
		return nil, err
	}
	ok := crtPool.AppendCertsFromPEM(caCrt)
	if !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{crt},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    crtPool,
	}), nil
}

func MakeClientCredentials(clientCrtPath, clientKeyPath, caCrtPath, serverName string) (credentials.TransportCredentials, error) {
	crt, err := tls.LoadX509KeyPair(clientCrtPath, clientKeyPath)
	if err != nil {
		return nil, err
	}
	crtPool := x509.NewCertPool()
	caCrt, err := os.ReadFile(caCrtPath)
	if err != nil {
		return nil, err
	}
	ok := crtPool.AppendCertsFromPEM(caCrt)
	if !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{crt},
		ServerName:   serverName,
		RootCAs:      crtPool,
	}), nil
}
