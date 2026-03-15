package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

func Encrypt(plainText string, key []byte) (string, error) {
	block, err := aes.NewCipher(key) /*creates an AES encryption machine using your secret key.
	Think of it like putting your key into a lock — this step initializes the lock */
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block) /*GCM adds an extra layer — it not only encrypts your data but
	also signs it. So when decrypting, if anyone tampered with the data even one single character,
	it will refuse to decrypt. Like a tamper-proof seal on a medicine bottle. */
	if err != nil {
		return "", err
	}
	nonce := make([]byte, aesGCM.NonceSize()) /*A nonce is a random number that is used only once.
	It ensures that even if you encrypt the same plaintext multiple times, you will get different
	ciphertexts each time. Think of it like adding a unique salt to each encryption to prevent patterns from emerging. */
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	cipherBytes := aesGCM.Seal(nonce, nonce, []byte(plainText), nil) /*aesGCM.Seal does the actual encryption.
		It:1.Takes our plain text
		2.Encrypts it using the key and nonce
	    3.Glues the nonce to the front of the result */
	return base64.StdEncoding.EncodeToString(cipherBytes), nil /* base64.StdEncoding.EncodeToString converts
	raw bytes into readable text. Raw encrypted bytes look like garbage characters — base64 turns them into
	safe readable characters*/

}

func Decrypt(cipherText string, key []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherText) /*First, we decode the base64 string back into raw bytes. */
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key) /*We create the AES cipher again using the same key. */
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block) /*We set up GCM mode again. */
	if err != nil {
		return "", err
	}
	nonceSize := aesGCM.NonceSize() /*We know that the nonce is at the beginning of the data, so we extract it. */
	if len(data) < nonceSize {
		return "", errors.New("cipher text is too short")
	}
	nonce, cipherBytes := data[:nonceSize], data[nonceSize:]     /*We split the data into the nonce and the actual encrypted bytes. */
	plainBytes, err := aesGCM.Open(nil, nonce, cipherBytes, nil) /*aesGCM.Open does the decryption. It takes the nonce and the encrypted bytes and returns the original plaintext. */
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil /*Finally, we convert the decrypted bytes back into a string and return it. */
}
