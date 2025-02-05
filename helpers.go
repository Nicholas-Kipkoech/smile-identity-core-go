package smileidentity

import "fmt"

type SDKVersionInfo struct {
	SourceSDK        string
	SourceSDKVersion string
}

// MapServerUri converts a numeric key to a Smile server URI, or returns the original URI if not found in the map.
func MapServerUri(uriOrKey interface{}) string {
	switch v := uriOrKey.(type) {
	case int:
		if uri, exists := sidServerMapping[v]; exists {
			return uri
		}
	case string:
		return v
	}
	return ""
}

// PartnerParams defines the required parameters for each job.
type PartnerParams struct {
	UserID  string      `json:"user_id"`
	JobID   string      `json:"job_id"`
	JobType interface{} `json:"job_type"`
}

// ValidatePartnerParams validates the partner params like user_id, job_id and job_type
func validatePatnerParams(partnerParams PartnerParams) error {

	if (partnerParams == PartnerParams{}) {
		return fmt.Errorf("please ensure you send through patner params")
	}

	requiredFields := []string{"UserID", "JobID", "JobType"}

	for _, key := range requiredFields {
		switch key {
		case "UserID":
			if partnerParams.UserID == "" {
				return fmt.Errorf("please make sure that user_id is included in the partner params")
			}
		case "JobID":
			if partnerParams.JobID == "" {
				return fmt.Errorf("please make sure that job_id is included in the partner params")
			}
		case "JobType":
			if partnerParams.JobType == "" {
				return fmt.Errorf("please make sure that job_type is included in the partner params")
			}
		}
	}
	return nil
}
