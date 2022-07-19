package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"os"
)

func SavePem(filename, header string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	err = pem.Encode(file, &pem.Block{
		Type:  header,
		Bytes: content,
	})

	if err != nil {
		panic(err)
	}

	file.Close()
	return nil
}

func SaveRandomKey(publicName, privateName string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		panic(err)
	}

	err = SavePem(privateName, "RSA PRIVATE KEY", x509.MarshalPKCS1PrivateKey(privateKey))
	if err != nil {
		panic(err)
	}

	err = SavePem(publicName, "RSA PUBLIC KEY", x509.MarshalPKCS1PublicKey(&privateKey.PublicKey))
	if err != nil {
		panic(err)
	}
}

func Encrypt(filename string, data []byte) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	p, _ := pem.Decode(content)
	if p == nil {
		return "", err
	}

	publicKey, err := x509.ParsePKCS1PublicKey(p.Bytes)
	if err != nil {
		return "", err
	}

	encryptData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
	if err != nil {
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString(encryptData), nil
}

func Decrypt(filename, data string) ([]byte, error) {
	target, err := base64.RawStdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	p, _ := pem.Decode(content)
	if p == nil {
		return nil, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, target, nil)
}
