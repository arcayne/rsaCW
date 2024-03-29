package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {

	// Generate RSA Keys
	joanPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	joanPublicKey := &joanPrivateKey.PublicKey

	hilariPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	hilarilPublicKey := &hilariPrivateKey.PublicKey

	fmt.Println("Private Key : ", joanPrivateKey)
	fmt.Println("Public key ", joanPublicKey)
	fmt.Println("Private Key : ", hilariPrivateKey)
	fmt.Println("Public key ", hilarilPublicKey)

	//Encrypt Miryan Message
	message := []byte("the code must be like a piece of music")
	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, hilarilPublicKey, message, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(message), ciphertext)
	fmt.Println()

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, joanPrivateKey, newhash, hashed, &opts)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("PSS Signature : %x\n", signature)

	// Decrypt Message
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, hilariPrivateKey, ciphertext, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)

	//Verify Signature
	err = rsa.VerifyPSS(joanPublicKey, newhash, hashed, signature, &opts)

	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify Signature successful")
	}

}