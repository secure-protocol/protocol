package interceptor

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type SignFunc func(key *ecdsa.PrivateKey, digest []byte) ([]byte, error)
type VerifyFunc func(key *ecdsa.PublicKey, hash []byte, sig []byte) bool

func Sign(pri *ecdsa.PrivateKey, message proto.Message, signFunc SignFunc) (string, error) {
	marshal, err := proto.Marshal(message)
	if err != nil {
		return "", err
	}
	sum256 := sha256.Sum256(marshal)
	sign, err := signFunc(pri, sum256[:])
	//sign, err := pri.Sign(rand.Reader, sum256[:], nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func Verify(pub *ecdsa.PublicKey, signatureStr string, message proto.Message, verifyFunc VerifyFunc) error {
	signature, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return err
	}
	marshal, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	h := sha256.Sum256(marshal)
	ok := verifyFunc(pub, h[:], signature)
	if !ok {
		return errors.New("invalid signature")
	}
	return nil
}

const keySig = "sig"

func GetSignInterceptor(pri *ecdsa.PrivateKey, signFunc SignFunc) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		sig, err := Sign(pri, req.(proto.Message), signFunc)
		if err != nil {
			return err
		}
		ctx = metadata.AppendToOutgoingContext(ctx, keySig, sig)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func GetVeryInterceptor(pub *ecdsa.PublicKey, verifyFunc VerifyFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		v := metadata.ValueFromIncomingContext(ctx, keySig)
		if len(v) < 1 {
			return nil, errors.New("invalid signature")
		}
		err := Verify(pub, v[0], req.(proto.Message), verifyFunc)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
