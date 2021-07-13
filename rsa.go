package WeWorkFinanceSDK

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
)

// RSAEncrypt RSA加密
func RSAEncrypt(publicKey []byte, origData []byte) ([]byte, error) {
	pubInterface, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RSAEncryptBase64 RSA加密后base64encode
func RSAEncryptBase64(publicKey []byte, origData []byte) (string, error) {
	encryptedData, err := RSAEncrypt(publicKey, origData)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// RSADecrypt RSA解密
func RSADecrypt(privateKey []byte, ciphertext []byte) ([]byte, error) {
	priv, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RSADecryptBase64 base64decode后RSA解密
func RSADecryptBase64(privateKey []byte, cryptoText string) ([]byte, error) {
	encryptedData, _ := base64.StdEncoding.DecodeString(cryptoText)
	return RSADecrypt(privateKey, encryptedData)
}

// RSASignWithSHA1 RSA-SHA1签名
func RSASignWithSHA1(privateKey []byte, data []byte) ([]byte, error) {
	priv, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	//priv := privInterface.(*rsa.PrivateKey)
	hash := Sha1Sum(data)
	return rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, hash)
}

// RSASignWithSHA1Base64 RSA-SHA1签名后做base64encode
func RSASignWithSHA1Base64(privateKey []byte, data []byte) (string, error) {
	sign, err := RSASignWithSHA1(privateKey, data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

// RSAVerifySignWithSha1 RSA-SHA1签名验证
func RSAVerifySignWithSha1(publicKey []byte, origData []byte, signData string) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	pubInterface, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	hash := Sha1Sum(origData)
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hash, sign)
}

// Sha1Sum SHA1
func Sha1Sum(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
