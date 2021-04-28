package models

// AuthKey contains API Key base user auth info
type AuthKey struct {
	ID         int    `json:"id"`
	KeyID      string `json:"key_id"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
	UserID     string `json:"userID"`
}
