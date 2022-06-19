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
