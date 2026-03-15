package model

type EncryptRequest struct {
	PlainText string `json: "plain_text"`
}
type EncryptResponse struct {
	CipherText string `json: "cipher_text"`
}
type DecryptRequest struct {
	CipherText string `json: "cipher_text"`
}
type DecryptResponse struct {
	PlainText string `json: "plain_text"`
}
