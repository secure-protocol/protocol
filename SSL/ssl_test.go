package SSL

import (
	"context"
	"fmt"
	"github.com/secure-protocol/protocol/proto/signer"
	"google.golang.org/grpc"
	"net"
	"testing"
)

type fakeSigner struct {
	signer.UnimplementedSignerServer
}

func (f fakeSigner) Sign(ctx context.Context, in *signer.SignReq) (*signer.SignRes, error) {
	return &signer.SignRes{}, nil
}

func (f fakeSigner) Gen(ctx context.Context, in *signer.GenReq) (*signer.GenRes, error) {
	return &signer.GenRes{Address: []*signer.Address{{
		Type: 0,
		Bid:  "test",
		Addr: "test",
	}}}, nil
}

var caCrtPath = "../testSSL/ca.crt"
var signerCrtPath = "../testSSL/signer.crt"
var signerKeyPath = "../testSSL/signer.key"

var clientCrtPath = "../testSSL/client.crt"
var clientKeyPath = "../testSSL/client.key"

var target = "127.0.0.1:9999"

func startFakeSignerServerWithSSL() error {
	creds, err := MakeServerCredentials(signerCrtPath, signerKeyPath, caCrtPath)
	if err != nil {
		return err
	}

	signerServer := fakeSigner{signer.UnimplementedSignerServer{}}

	lis, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	signer.RegisterSignerServer(grpcServer, &signerServer)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			fmt.Println(err)
		}
	}()
	return nil
}

func startSSLClient() error {
	creds, err := MakeClientCredentials(clientCrtPath, clientKeyPath, caCrtPath, "signer")
	if err != nil {
		return err
	}

	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}
	defer conn.Close()

	client := signer.NewSignerClient(conn)
	ctx := context.Background()
	signRes, err := client.Sign(ctx, &signer.SignReq{})
	if err != nil {
		return err
	}
	fmt.Println(signRes)

	gen, err := client.Gen(ctx, &signer.GenReq{})
	if err != nil {
		return err
	}
	fmt.Println(gen)
	return nil
}

func TestSSL(t *testing.T) {
	err := startFakeSignerServerWithSSL()
	if err != nil {
		t.Fatal(err)
	}
	err = startSSLClient()
	if err != nil {
		t.Fatal(err)
	}
}
