package interceptor

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/secure-protocol/protocol/proto/blockchain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
)

type fakeServer struct {
	blockchain.UnimplementedBlockchainServerServer
}

func (f *fakeServer) GenAddress(context.Context, *blockchain.GenAddressReq) (*blockchain.GenAddressRes, error) {
	return &blockchain.GenAddressRes{}, nil
}

var addr string
var pri *ecdsa.PrivateKey
var client blockchain.BlockchainServerClient

func init() {
	var err error
	pri, err = GenPri()
	if err != nil {
		panic(err)
	}
	addr = ":9091"
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			panic(err)
		}
		grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(GetVeryInterceptor(&pri.PublicKey, EthVerify)))
		blockchain.RegisterBlockchainServerServer(grpcServer, &fakeServer{})
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(GetSignInterceptor(pri, EthSign)))
	if err != nil {
		panic(err)
	}
	client = blockchain.NewBlockchainServerClient(conn)

	_, err = client.GenAddress(context.Background(), &blockchain.GenAddressReq{
		AccountType: blockchain.AccountType_Tron,
		Bid:         "test",
		Nums:        1,
	})
	if err != nil {
		panic(err)
	}
}

func BenchmarkSig(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_, err := client.GenAddress(context.Background(), &blockchain.GenAddressReq{
			AccountType: blockchain.AccountType_Tron,
			Bid:         "test",
			Nums:        int64(i),
		})
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSig(t *testing.T) {
	_, err := client.GenAddress(context.Background(), &blockchain.GenAddressReq{
		AccountType: blockchain.AccountType_Tron,
		Bid:         "test",
		Nums:        1,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func EcdsaSign(pri *ecdsa.PrivateKey, h []byte) ([]byte, error) {
	return ecdsa.SignASN1(rand.Reader, pri, h)
}

func TestSignWrong(t *testing.T) {
	pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(GetSignInterceptor(pri, EcdsaSign)))
	if err != nil {
		t.Fatal(err)
	}
	client := blockchain.NewBlockchainServerClient(conn)

	_, err = client.GenAddress(context.Background(), &blockchain.GenAddressReq{
		AccountType: blockchain.AccountType_Tron,
		Bid:         "test",
		Nums:        1,
	})
	if err == nil {
		t.Fatal("should be err")
	}
}
