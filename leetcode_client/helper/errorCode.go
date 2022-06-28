package helper

import "fmt"

// Error implement the customized error function for all leetecho-cli business
//  @receiver e the customized error in leetecho-cli business
//  @return string the customized error message
func (e *ErrorResp) Error() string {
	return fmt.Sprintf(`
Error Status: %s;
Error Code: %d;
Error Message: %s;`, e.Status, e.Code, e.Message)
}

func GetErrorCodeMessage(code LeetechoCode) LeetechoMessage {
	switch code {
	case OK_CODE:
		return LeetechoMessage(OK_MESSAGE)
	case REDIRECT_CODE:
		return LeetechoMessage(REDIRECT_MESSAGE)
	case NOT_MODIFIED_CODE:
		return LeetechoMessage(NOT_MODIFIED_MESSAGE)
	case BAD_REQUEST_CODE:
		return LeetechoMessage(BAD_REQUEST_MESSAGE)
	case FORBIDDEN_CODE:
		return LeetechoMessage(FORBIDDEN_MESSAGE)
	case NOT_FOUND_CODE:
		return LeetechoMessage(NOT_FOUND_MESSAGE)
	case INTERNAL_SERVER_ERROR_CODE:
		return LeetechoMessage(INTERNAL_SERVER_ERROR_MESSAGE)
	case BAD_GATEWAY_CODE:
		return LeetechoMessage(BAD_GATEWAY_MESSAGE)
	case SERVICE_UNAVAILABLE_CODE:
		return LeetechoMessage(SERVICE_UNAVAILABLE_MESSAGE)
	case UNKNOWN_CODE:
		return LeetechoMessage(UNKNOWN_MESSAGE)
	case NOT_LOGIN_CODE:
		return LeetechoMessage(NOT_LOGIN_MESSAGE)
	case REQUEST_PARAMS_ERROR_CODE:
		return LeetechoMessage(REQUEST_PARAMS_ERROR_MESSAGE)
	case NOT_A_REPO_CODE:
		return LeetechoMessage(NOT_A_REPO_MESSAGE)
	case NO_AC_SUBMISSIONS_CODE:
		return LeetechoMessage(NO_AC_SUBMISSIONS_MESSAGE)
	case NO_NOTES_CODE:
		return LeetechoMessage(NO_NOTES_MESSAGE)
	case REPO_CONNECTION_ERROR_CODE:
		return LeetechoMessage(REPO_CONNECTION_ERROR_MESSAGE)
	case REPO_PUSH_ERROR_CODE:
		return LeetechoMessage(REPO_PUSH_ERROR_MESSAGE)
	case NO_USER_CONFIG_CODE:
		return LeetechoMessage(NO_USER_CONFIG_MESSAGE)
	case REPO_INIT_ERROR_CODE:
		return LeetechoMessage(REPO_INIT_ERROR_MESSAGE)
	default:
		return LeetechoMessage(UNKNOWN_MESSAGE)
	}
}
