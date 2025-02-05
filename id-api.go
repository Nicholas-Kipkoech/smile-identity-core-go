package smileidentity

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type SignatureInfo struct {
	Signature string `json:"signature"`
	Timestamp any    `json:"timestamp"`
}

type IdInfo struct {
	BusinessType string      `json:"business_type,omitempty"`
	Country      string      `json:"country,omitempty"`
	DOB          string      `json:"dob,omitempty"`
	Entered      interface{} `json:"entered,omitempty"` // Can be boolean or string
	FirstName    string      `json:"first_name,omitempty"`
	IDType       string      `json:"id_type,omitempty"`
	LastName     string      `json:"last_name,omitempty"`
	IDNumber     string      `json:"id_number,omitempty"`
	MiddleName   string      `json:"middle_name,omitempty"`
	PhoneNumber  string      `json:"phone_number,omitempty"`
}

type VerificationRequest struct {
	PartnerID     string        `json:"partner_id"`
	PartnerParams PartnerParams `json:"partner_params"`
	CallbackURL   string        `json:"callback_url,omitempty"`
	IdInfo        IdInfo        `json:",inline"`
	SignatureInfo SignatureInfo `json:",inline"`
}

type IDApi struct {
	partnerID string
	apiKey    string
	baseURL   string
}

func newIDApi(patnerID, apiKey string, sidServer int) *IDApi {
	url := MapServerUri(sidServer)
	return &IDApi{
		partnerID: patnerID,
		apiKey:    apiKey,
		baseURL:   fmt.Sprintf("http://%s", url),
	}
}

func validateInfo(idInfo IdInfo) error {
	if idInfo.IDNumber == "" {
		return errors.New("please provide an id_number in the id_info payload ")
	}
	return nil
}

func (api *IDApi) createVerificationRequest(idInfo IdInfo, partnerParams PartnerParams, callbackUrl string) VerificationRequest {
	signature, timestamp, err := GenerateSignature(api.partnerID, api.apiKey)
	if err != nil {
		fmt.Println("Error generating signature", err)
		return VerificationRequest{}
	}
	return VerificationRequest{
		PartnerID:     api.partnerID,
		PartnerParams: partnerParams,
		CallbackURL:   callbackUrl,
		IdInfo:        idInfo,
		SignatureInfo: SignatureInfo{Signature: signature, Timestamp: timestamp},
	}

}

func (api *IDApi) submitAsyncJob(partnerParams PartnerParams, idInfo IdInfo, callbackUrl string) (map[string]interface{}, error) {
	if err := validateInfo(idInfo); err != nil {
		return nil, err
	}
	request := api.createVerificationRequest(idInfo, partnerParams, "")
	endpoint := "/async_id_verification"
	return api.makeRequest(endpoint, request)
}

func (api *IDApi) pollJobStatus(partnerParams PartnerParams, maxRetries int, timeout time.Duration, returnHistory, returnImages bool) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"user_id":     partnerParams.UserID,
		"job_id":      partnerParams.JobID,
		"partner_id":  api.partnerID,
		"history":     returnHistory,
		"image_links": returnImages,
	}
	for retries := 0; retries < maxRetries; retries++ {
		response, err := api.makeRequest("/job_status", data)
		if err == nil && response["job_complete"].(bool) {
			return response, nil
		}
		time.Sleep(timeout)
	}
	return nil, errors.New("max retries reached")
}

func (api *IDApi) makeRequest(endpoint string, data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}
	resp, err := http.Post(api.baseURL+endpoint, "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil

}
