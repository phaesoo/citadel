package models

// AuthKey contains API Key base user auth info
type AuthKey struct {
	ID         int    `json:"id"`
	KeyID      string `json:"key_id"`
	PublicPem  string `json:"public_pem"`
	PrivatePem string `json:"private_pem"`
	UserID     string `json:"userID"`
}

func (m *AuthKey) Equal(target AuthKey) bool {
	return m.KeyID == target.KeyID &&
		m.PublicPem == target.PublicPem &&
		m.PrivatePem == target.PrivatePem &&
		m.UserID == target.UserID
}
