package smileidentity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"time"
)

// GenerateSignature generates a signature for the given input
func GenerateSignature(partnerID, apiKey string, timestamp ...interface{}) (string, string, error) {
	var isoTimestamp string

	// Handle optional timestamp
	if len(timestamp) == 0 {
		// Default to current timestamp if none is provided
		isoTimestamp = time.Now().UTC().Format(time.RFC3339)
	} else {
		switch t := timestamp[0].(type) {
		case int, int64:
			isoTimestamp = time.Unix(int64(t.(int)), 0).UTC().Format(time.RFC3339)
		case string:
			isoTimestamp = t
		default:
			return "", "", errors.New("invalid timestamp format")
		}
	}

	// Generate HMAC signature
	mac := hmac.New(sha256.New, []byte(apiKey))
	mac.Write([]byte(isoTimestamp))
	mac.Write([]byte(partnerID))
	mac.Write([]byte("sid_request"))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), isoTimestamp, nil

}

// Signature struct handles signature generation and verification
type Signature struct {
	PartnerID string
	APIKey    string
}

// NewSignature creates a new signature instance
func NewSignature(partnerID, apiKey string) *Signature {
	return &Signature{
		PartnerID: partnerID,
		APIKey:    apiKey,
	}
}

func (s *Signature) ConfirmSignature(timestamp interface{}, signature string) (bool, error) {
	sig, _, err := GenerateSignature(s.PartnerID, s.APIKey, timestamp)
	if err != nil {
		return false, err
	}
	return sig == signature, nil
}
