package constants

import (
	"fmt"
	"net/http"
)

const (
	AccessView     = "view"
	AccessCreate   = "create"
	AccessUpdate   = "update"
	AccessPut      = "put"
	AccessApproval = "approval"
	AccessGet      = "get"
	AccessList     = "list"
	AccessDelete   = "delete"

	ConfigRoutes                 = "common.config-routes-key"
	ConfigPrefixRoutesBackoffice = "common.prefix-config-route-backoffice"

	MddwTokenKey       = "token-data"
	MddwUserData       = "user-data"
	MddwUserHandheld   = "user-handheld"
	MddwUserBackoffice = "user-backoffice"
	MddwUserAccess     = "user-access"
	MddwKeyRole        = "role-data"

	ConfigAppointment     = "config_appointment"
	ConfigFeeRegisPatient = "config_fee_registration_patient"

	TimeExpire    = 24
	timeoutNumber = 20000
	retryNumber   = 2
)

func generateCurlCommand(req *http.Request, reqBody []byte) string {
	curlCommand := fmt.Sprintf("curl -X %s '%s'", req.Method, req.URL.String())

	// Add headers
	for key, values := range req.Header {
		for _, value := range values {
			curlCommand += fmt.Sprintf(" -H '%s: %s'", key, value)
		}
	}

	// Add body if present
	if len(reqBody) > 0 {
		curlCommand += fmt.Sprintf(" -d '%s'", string(reqBody))
	}

	return curlCommand
}
