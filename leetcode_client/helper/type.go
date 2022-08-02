package helper

type URI string

type Credit struct {
	Session   string
	CSRFToken string
}

type EndpointURI struct {
	Base    string
	Login   string
	Graphql string
}

type Base struct {
	US *EndpointURI
	CN *EndpointURI
}

// ErrorResp is the customized error for all leetecho-cli business
type ErrorResp struct {
	Status  LeetechoStatus  `json:"status"`
	Code    LeetechoCode    `json:"code"`
	Message LeetechoMessage `json:"message"`
}
