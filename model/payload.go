package model

type EncryptRequest struct {
	PlainText string `json: "plain_text"`
}
type EncryptResponse struct {
	CipherText string `json: "cipher_text"`
}
type DecrptRequest struct {
	CipherText string `json: "cipher_text"`
}
type DecrptResponse struct {
	PlainText string `json: "plain_text"`
}
