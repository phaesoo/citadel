package models

// AuthKey contains API Key base user auth info
type AuthKey struct {
	ID         int    `json:"id"`
	KeyID      string `json:"key_id"`
	PublicKey  string `json:"public_pem"`
	PrivateKey string `json:"private_pem"`
	UserID     string `json:"userID"`
}
