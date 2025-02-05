package smileidentity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OptionsParam struct {
	ReturnHistory bool
	ReturnImages  bool
}

func getJobStatus(partnerID, apiKey, url, userID, jobID string, options OptionsParam) (map[string]interface{}, error) {
	signature := NewSignature(partnerID, apiKey)
	generatedSignature, timestamp, err := GenerateSignature(partnerID, apiKey)
	if err != nil {
		return nil, err
	}
	requestBody := map[string]interface{}{
		"user_id":     userID,
		"job_id":      jobID,
		"partner_id":  partnerID,
		"history":     options.ReturnHistory,
		"image_links": options.ReturnImages,
		"signature":   generatedSignature,
		"timestamp":   timestamp,
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("https://%s/job_status", url), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(data, &responseData); err != nil {
		return nil, err
	}
	valid, err := signature.ConfirmSignature(responseData["timetstamp"].(string), responseData["signature"].(string))
	if err != nil || !valid {
		return nil, errors.New("unable to confirm validity of the job_status response")
	}

	return responseData, nil

}

type Utilities struct {
	PartnerID string
	APIKey    string
	URL       string
}

// NewUtilities creates a new Utilities instance
func NewUtilities(partnerID, apiKey, sidServer string) *Utilities {
	return &Utilities{
		PartnerID: partnerID,
		APIKey:    apiKey,
		URL:       MapServerUri(sidServer), // Using the imported function
	}
}

// GetJobStatus calls getJobStatus and returns the job status
func (u *Utilities) GetJobStatus(userID, jobID string, options OptionsParam) (map[string]interface{}, error) {
	return getJobStatus(u.PartnerID, u.APIKey, u.URL, userID, jobID, options)
}
