package smileidentity

// IMAGE_TYPE represents the type of image submitted in a job request.
type IMAGE_TYPE int

const (
	// SELFIE_IMAGE_FILE Selfie image in .png or .jpg file format
	SELFIE_IMAGE_FILE IMAGE_TYPE = iota
	// ID_CARD_IMAGE_FILE ID card image in .png or .jpg file format
	ID_CARD_IMAGE_FILE
	// SELFIE_IMAGE_BASE64 Base64 encoded selfie image (.png or .jpg)
	SELFIE_IMAGE_BASE64
	// ID_CARD_IMAGE_BASE64 Base64 encoded ID card image (.png or .jpg)
	ID_CARD_IMAGE_BASE64
	// LIVENESS_IMAGE_FILE Liveness image in .png or .jpg file format
	LIVENESS_IMAGE_FILE
	// ID_CARD_BACK_IMAGE_FILE Back of ID card image in .png or .jpg file format
	ID_CARD_BACK_IMAGE_FILE
	// LIVENESS_IMAGE_BASE64 Base64 encoded liveness image (.jpg or .png)
	LIVENESS_IMAGE_BASE64
	// ID_CARD_BACK_IMAGE_BASE64 Base64 encoded back of ID card image (.jpg or .png)
	ID_CARD_BACK_IMAGE_BASE64
)

// JOB_TYPE represents the type of verification job.
type JOB_TYPE int

const (
	// BIOMETRIC_KYC Verify the ID information of your users using facial biometrics
	BIOMETRIC_KYC JOB_TYPE = iota
	// SMART_SELFIE_AUTHENTICATION Used to identify your existing users.
	SMART_SELFIE_AUTHENTICATION
	//SMART_SELFIE_REGISTRATION Used to verify and register a user for future authentication.
	SMART_SELFIE_REGISTRATION
	// Verifies identity information of a person with their personal, information and ID number from one of our supported ID Types.
	BASIC_KYC
	//  ENHANCED_KYC query the Identity Information for an individual using their, ID number from one of our supported.
	ENHANCED_KYC
	// DOCUMENT_VERIFICATION Detailed user information retrieved from the ID issuing authority.
	DOCUMENT_VERIFICATION
	//  BUSINESS_VERIFICATION Verify the authenticity of Document IDs of your users and confirm it belongs to the user using facial biometrics.
	BUSINESS_VERIFICATION
	// Updates the photo on file for an enrolled user
	UPDATE_PHOTO
	// Compares document verification to an id check
	COMPARE_USER_INFO
	// ENHANCED_DOCUMENT_VERIFICATION Verifies user selfie with info retrieved from the ID issuing authority
	ENHANCED_DOCUMENT_VERIFICATION
)

var sidServerMapping = map[int]string{
	0: "testapi.smileidentity.com/v1",
	1: "api.smileidentity.com/v1",
}
