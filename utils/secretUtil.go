// @Title  secretUtil.go
// @Description  数据加解密相关的工具类
// @Author  zhengbin  2020/7/23 16:49
// @Update  姓名（需要改）  2020/7/23 16:49
package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//var AESkey = []byte("lifeisshortandartislong.")

// @title    AESEncrypt
// @description   AES加密方法
// @auth      zhengbin             2020/7/23 16:49
// @param     origData        []byte         "原始数据"
// @param     origData        []byte         "密钥"
// @return    []byte         "加密后的数据"
func AESEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// @title    AESDecrypt
// @description   AES解密方法
// @auth      zhengbin             2020/7/23 16:55
// @param     origData        []byte         "加密数据"
// @param     origData        []byte         "密钥"
// @return    []byte         "原始数据"
func AESDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
