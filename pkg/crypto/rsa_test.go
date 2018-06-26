package crypto

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	// publicKeyDERFile is a RSA public key in DER format
	publicKeyDERFile string = "tests/public_key.der"
	// publicKeyPEMFile is a RSA public key in PEM format
	publicKeyPEMFile string = "tests/public_key.pem"
	// privateKeyPEMFile is a RSA public key in PEM format
	privateKeyPEMFile string = "tests/private_key.pem"
	// testDataFile which should be verified by the good signature
	testDataFile string = "tests/data"
	// signatureGoodFile is a good signature of testDataFile
	signatureGoodFile string = "tests/verify_rsa_pkcs15_sha256.signature"
	// signatureBadFile is a bad signature which does not work with testDataFile
	signatureBadFile string = "tests/verify_rsa_pkcs15_sha256.signature2"
)

var (
	// password is a PEM encrypted passphrase
	password = []byte{'k', 'e', 'i', 'n', 's'}
)

func TestLoadDERPublicKey(t *testing.T) {
	_, err := LoadPublicKeyFromFile(publicKeyDERFile)
	require.Error(t, err)
}

func TestLoadPEMPublicKey(t *testing.T) {
	_, err := LoadPublicKeyFromFile(publicKeyPEMFile)
	require.NoError(t, err)
}

func TestLoadPEMPrivateKey(t *testing.T) {
	_, err := LoadPrivateKeyFromFile(privateKeyPEMFile, password)
	require.NoError(t, err)
}

func TestSignData(t *testing.T) {
	privateKey, err := LoadPrivateKeyFromFile(privateKeyPEMFile, password)
	require.NoError(t, err)

	testData, err := ioutil.ReadFile(testDataFile)
	require.NoError(t, err)

	_, err = SignRsaSha256Pkcs1v15Signature(privateKey, testData)
	require.NoError(t, err)
}

func TestVerifyData(t *testing.T) {
	publicKey, err := LoadPublicKeyFromFile(publicKeyPEMFile)
	require.NoError(t, err)

	testData, err := ioutil.ReadFile(testDataFile)
	require.NoError(t, err)

	signatureGood, err := ioutil.ReadFile(signatureGoodFile)
	require.NoError(t, err)

	err = VerifyRsaSha256Pkcs1v15Signature(publicKey, testData, signatureGood)
	require.NoError(t, err)
}

func TestBadSignature(t *testing.T) {
	publicKey, err := LoadPublicKeyFromFile(publicKeyPEMFile)
	require.NoError(t, err)

	testData, err := ioutil.ReadFile(testDataFile)
	require.NoError(t, err)

	signatureBad, err := ioutil.ReadFile(signatureBadFile)
	require.NoError(t, err)

	err = VerifyRsaSha256Pkcs1v15Signature(publicKey, testData, signatureBad)
	require.Error(t, err)
}
