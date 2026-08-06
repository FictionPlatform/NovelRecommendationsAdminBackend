package encrypt

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"errors"
)

func AesEncrypt(v string, k []byte) (string, error) {
	value := []byte(v)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}

	blocksize := block.BlockSize()
	valueBytes := value

	fillsize := blocksize - len(valueBytes)%blocksize
	repeat := bytes.Repeat([]byte{byte(fillsize)}, fillsize)
	valueBytes = append(valueBytes, repeat...)

	result := make([]byte, len(valueBytes))

	temp := result
	for len(valueBytes) > 0 {
		block.Encrypt(temp, valueBytes[:blocksize])
		valueBytes = valueBytes[blocksize:]
		temp = temp[blocksize:]
	}
	return hex.EncodeToString(result), nil
}

func AesDecrypt(v string, k []byte) (string, error) {
	value, err := hex.DecodeString(v)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(k)
	if err != nil {
		return "", err
	}
	size := len(value)
	result := make([]byte, size)
	blocksize := block.BlockSize()
	if size == 0 || size%blocksize != 0 {
		return "", errors.New("invalid ciphertext length")
	}
	temp := result
	for len(value) > 0 {
		block.Decrypt(temp, value[:blocksize])
		value = value[blocksize:]
		temp = temp[blocksize:]
	}
	//校验 PKCS7 填充一致性，避免误剥数据
	padding := int(result[size-1])
	if padding == 0 || padding > blocksize || padding > size {
		return "", errors.New("invalid ciphertext padding")
	}
	for i := size - padding; i < size; i++ {
		if result[i] != byte(padding) {
			return "", errors.New("invalid ciphertext padding")
		}
	}
	result = result[:size-padding]
	return string(result), nil
}
