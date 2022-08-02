package helper

import "net/http"

var BaseURI *Base = &Base{
	US: &EndpointURI{
		Base:    "https://leetcode.com/",
		Login:   "https://leetcode.com/accounts/login/",
		Graphql: "https://leetcode.com/graphql",
	},
	CN: &EndpointURI{
		Base:    "https://leetcode.cn/",
		Login:   "https://leetcode.cn/accounts/login/",
		Graphql: "https://leetcode.cn/graphql",
	},
}

type EndPoint string

const (
	US EndPoint = "US"
	CN EndPoint = "CN"
)

type LeetechoStatus string
type LeetechoCode int
type LeetechoMessage string

var (
	OK_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusOK))
	OK_CODE    LeetechoCode    = LeetechoCode(http.StatusOK)
	OK_MESSAGE LeetechoMessage = LeetechoMessage("OK")

	REDIRECT_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusFound))
	REDIRECT_CODE    LeetechoCode    = LeetechoCode(http.StatusFound)
	REDIRECT_MESSAGE LeetechoMessage = LeetechoMessage("Redirect")

	NOT_MODIFIED_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusNotModified))
	NOT_MODIFIED_CODE    LeetechoCode    = LeetechoCode(http.StatusNotModified)
	NOT_MODIFIED_MESSAGE LeetechoMessage = LeetechoMessage("Not Modified")

	BAD_REQUEST_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusBadRequest))
	BAD_REQUEST_CODE    LeetechoCode    = LeetechoCode(http.StatusBadRequest)
	BAD_REQUEST_MESSAGE LeetechoMessage = LeetechoMessage("Bad Request")

	FORBIDDEN_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusForbidden))
	FORBIDDEN_CODE    LeetechoCode    = LeetechoCode(http.StatusForbidden)
	FORBIDDEN_MESSAGE LeetechoMessage = LeetechoMessage("Forbidden")

	NOT_FOUND_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusNotFound))
	NOT_FOUND_CODE    LeetechoCode    = LeetechoCode(http.StatusNotFound)
	NOT_FOUND_MESSAGE LeetechoMessage = LeetechoMessage("Not Found")

	INTERNAL_SERVER_ERROR_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusInternalServerError))
	INTERNAL_SERVER_ERROR_CODE    LeetechoCode    = LeetechoCode(http.StatusInternalServerError)
	INTERNAL_SERVER_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Internal Server Error")

	BAD_GATEWAY_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusBadGateway))
	BAD_GATEWAY_CODE    LeetechoCode    = LeetechoCode(http.StatusBadGateway)
	BAD_GATEWAY_MESSAGE LeetechoMessage = LeetechoMessage("Bad Gateway")

	SERVICE_UNAVAILABLE_STATUS  LeetechoStatus  = LeetechoStatus(http.StatusText(http.StatusServiceUnavailable))
	SERVICE_UNAVAILABLE_CODE    LeetechoCode    = LeetechoCode(http.StatusServiceUnavailable)
	SERVICE_UNAVAILABLE_MESSAGE LeetechoMessage = LeetechoMessage("Service Unavailable")

	UNKNOWN_ERROR   LeetechoStatus  = LeetechoStatus("Unknown Error")
	UNKNOWN_CODE    LeetechoCode    = LeetechoCode(1)
	UNKNOWN_MESSAGE LeetechoMessage = LeetechoMessage("Unknown Error")

	NOT_LOGIN_STATUS  LeetechoStatus  = LeetechoStatus("Not Login")
	NOT_LOGIN_CODE    LeetechoCode    = LeetechoCode(4000001)
	NOT_LOGIN_MESSAGE LeetechoMessage = LeetechoMessage("Not Login")

	REQUEST_PARAMS_ERROR_STATUS  LeetechoStatus  = LeetechoStatus("Request Params Error")
	REQUEST_PARAMS_ERROR_CODE    LeetechoCode    = LeetechoCode(4000002)
	REQUEST_PARAMS_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Request Params Error")

	NOT_A_REPO_STATUS  LeetechoStatus  = LeetechoStatus("Not a Repo")
	NOT_A_REPO_CODE    LeetechoCode    = LeetechoCode(4030003)
	NOT_A_REPO_MESSAGE LeetechoMessage = LeetechoMessage("Not a Repo")

	INVALID_CONFIG_STATUS  LeetechoStatus  = LeetechoStatus("Invalid Config")
	INVALID_CONFIG_CODE    LeetechoCode    = LeetechoCode(4030004)
	INVALID_CONFIG_MESSAGE LeetechoMessage = LeetechoMessage("Invalid Config")

	NO_AC_SUBMISSIONS_STATUS  LeetechoStatus  = LeetechoStatus("No AC Submissions")
	NO_AC_SUBMISSIONS_CODE    LeetechoCode    = LeetechoCode(5000001)
	NO_AC_SUBMISSIONS_MESSAGE LeetechoMessage = LeetechoMessage("No AC Submissions")

	NO_NOTES_STATUS  LeetechoStatus  = LeetechoStatus("No Notes")
	NO_NOTES_CODE    LeetechoCode    = LeetechoCode(5000002)
	NO_NOTES_MESSAGE LeetechoMessage = LeetechoMessage("No Notes")

	REPO_CONNECTION_ERROR_STATUS  LeetechoStatus  = LeetechoStatus("Repo Connection Error")
	REPO_CONNECTION_ERROR_CODE    LeetechoCode    = LeetechoCode(5000003)
	REPO_CONNECTION_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Repo Connection Error")

	REPO_PUSH_ERROR_STATUS  LeetechoStatus  = LeetechoStatus("Repo Push Error")
	REPO_PUSH_ERROR_CODE    LeetechoCode    = LeetechoCode(5000004)
	REPO_PUSH_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Repo Push Error")

	NO_USER_CONFIG_STATUS  LeetechoStatus  = LeetechoStatus("No User Config")
	NO_USER_CONFIG_CODE    LeetechoCode    = LeetechoCode(5000005)
	NO_USER_CONFIG_MESSAGE LeetechoMessage = LeetechoMessage("No User Config")

	REPO_INIT_ERROR_STATUS  LeetechoStatus  = LeetechoStatus("Repo Init Error")
	REPO_INIT_ERROR_CODE    LeetechoCode    = LeetechoCode(5000006)
	REPO_INIT_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Repo Init Error")

	DECODE_JSON_ERROR_STATUS  LeetechoStatus  = LeetechoStatus("Decode JSON Error")
	DECODE_JSON_ERROR_CODE    LeetechoCode    = LeetechoCode(5000007)
	DECODE_JSON_ERROR_MESSAGE LeetechoMessage = LeetechoMessage("Decode JSON Error")
)
