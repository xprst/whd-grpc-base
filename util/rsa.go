package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

/**
 * 默认使用的 RSA 加密私钥
 */
var DEFAULT_RSA_PRIVATE_KEY = []byte(`
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAL/m3nZB08ilDyCW
pCHFFgKIvBBWXeJnLsQWLZLAxkqlQXB8Hf81Kb2xYjfox6mA46lbUi7gkBs93JTz
DRnaXq9DnNRrU+FPQsAJ8Li/xNSP05iIuvjMATRj6pxS/t8IQLv2i5eWM8kp1mG+
JDuTuHIOSiUtHus1lVaDDZBbQh3rAgMBAAECgYBn4D2dP8a2/nnwxvozeW6PkppS
MZ4CVp4e8G5c2NK9RzTkAZtvMMTWdLVY1D13yFfzrYYP7+ixhkvnqKT30JedT+zH
FxVQJ4kK4aZnS9Gh1YFihrC5GMT5Ij+3+iExh4eP8l83tSS3uPwvJpGqpbF1M4Wq
rZWFxeGcJhikNv8wcQJBAPiwHenVQUIX5YzuGuVfhMnDoOSWwz97qtN4OPhaD3w/
JgNkgcBCqB9CQYe5qyJMkKMEX+QoFPiLdbdCNnwrCvMCQQDFi1Ip2zscOCLBS+Ym
8KltubehsXw1NeQmRDsPH5tJwJH5XWFqyl0iToBi+oHXfvRA/FXrpeyxWy1io6Mi
am8pAkEAjFQs/QbWJSqA4K53RNlKf+PBBVxBXrA069FaLGH9fPnRRHbRdKDoZ4Mm
oSTW+arErwhH5+HqO3nOehOF1TkgmwJAIM0LbYvLetoPW01BAAJB/8gwp5aS6zrx
kTEPJWm4HTzugBtzS4oigMnMI6M44BFieU/s7F32uVRMau6E7fgCUQJATn6r8eWZ
l3WDJWd+jC1bIDIy2f/VBc0ABxvhUgxX1TvhzzfCgPX9DPNOmrmmkgpW5rENeaxN
tc8jY8jiIeWKlA==
-----END PRIVATE KEY-----`)

/**
* 默认使用的 RSA 加密公钥
*/
var DEFAULT_RSA_PUBLIC_KEY = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/5t52QdPIpQ8glqQhxRYCiLwQ
Vl3iZy7EFi2SwMZKpUFwfB3/NSm9sWI36MepgOOpW1Iu4JAbPdyU8w0Z2l6vQ5zU
a1PhT0LACfC4v8TUj9OYiLr4zAE0Y+qcUv7fCEC79ouXljPJKdZhviQ7k7hyDkol
LR7rNZVWgw2QW0Id6wIDAQAB
-----END PUBLIC KEY-----`)

const DEFAULT_RSA_KEY_SIZE = 1024

type Cipher interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
}

type pkcsClient struct {
	pub  *rsa.PublicKey
	priv *rsa.PrivateKey
}

func NewPkcs(publicKey []byte, privateKey []byte) *pkcsClient {
	return &pkcsClient{
		pub:  genPubKey([]byte(publicKey)),
		priv: genPriKey([]byte(privateKey)),
	}
}

func (pc *pkcsClient) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pc.pub, plaintext)
}

func (pc *pkcsClient) Decrypt(cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, pc.priv, cipherText)
}

func genPubKey(publicKey []byte) *rsa.PublicKey {
	block, _ := pem.Decode(publicKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	return pub.(*rsa.PublicKey)
}

func genPriKey(privateKey []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(privateKey)
	var priKey *rsa.PrivateKey
	priKey, _ = x509.ParsePKCS1PrivateKey(block.Bytes)
	return priKey
}
