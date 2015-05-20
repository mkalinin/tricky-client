package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/sha1"  
)

import (
  "golang.org/x/crypto/pbkdf2"
)

const (
  SALT        = "123"
  ITERS_COUNT = 4096
  KEY_LENGTH  = 32
)

type Decryptor struct {
  aes cipher.Block
}

func NewDecryptor(key string) (*Decryptor, error) {
    // setting up decryptor
    secretKey := pbkdf2.Key([]byte(key), []byte(SALT), ITERS_COUNT, KEY_LENGTH, sha1.New)
    aes, err := aes.NewCipher(secretKey)
    return &Decryptor{aes}, err
}

func (d Decryptor) Decrypt(src []byte) []byte {
  // getting IV
  iv := src[:aes.BlockSize]
  stream := cipher.NewCFBDecrypter(d.aes, iv)

  // getting payload and decrypting it
  data := src[aes.BlockSize:]
  stream.XORKeyStream(data, data)

  return data
}
