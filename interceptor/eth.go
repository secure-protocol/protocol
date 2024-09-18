package interceptor

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

func EthSign(pri *ecdsa.PrivateKey, hash []byte) ([]byte, error) {
	return crypto.Sign(hash, pri)
}

func EthVerify(pub *ecdsa.PublicKey, hash []byte, sig []byte) bool {
	pubBytes := crypto.FromECDSAPub(pub)
	if len(sig) == 65 {
		sig = sig[:64]
	}
	return crypto.VerifySignature(pubBytes, hash, sig)
}

func HexToPri(hexStr string) (pri *ecdsa.PrivateKey, err error) {
	return crypto.HexToECDSA(hexStr)
}

func HexToPub(hexStr string) (pub *ecdsa.PublicKey, err error) {
	pubBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	return crypto.DecompressPubkey(pubBytes)
}

func PriToHex(pri *ecdsa.PrivateKey) string {
	return pri.D.Text(16)
}

func PubToHex(pub *ecdsa.PublicKey) string {
	pubBytes := crypto.CompressPubkey(pub)
	return hex.EncodeToString(pubBytes)
}

func GenPri() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}
