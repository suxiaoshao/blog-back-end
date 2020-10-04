package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//@brief:填充明文
func PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

//@brief:去除填充数据
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// 编码
func Encryption(sourceString string, keyString string) (string, error) {
	sourceByteList := []byte(sourceString)
	key := []byte(keyString + "1234567890123456")[0:16]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	sourceByteList = PKCS5Padding(sourceByteList, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	encrypted := make([]byte, len(sourceByteList))
	blockMode.CryptBlocks(encrypted, sourceByteList)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 解码
func Decryption(encrypted string, keyString string) (string, error) {
	key := []byte(keyString + "1234567890123456")[0:16]
	bytesPass, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	sourceString := make([]byte, len(bytesPass))
	blockMode.CryptBlocks(sourceString, bytesPass)
	sourceString = PKCS5UnPadding(sourceString)
	return string(sourceString), nil
}
