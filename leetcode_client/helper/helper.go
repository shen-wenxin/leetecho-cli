package helper

import (
	"context"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/shurcooL/graphql"

	leetcode_client "github.com/CallanBi/leetecho-cli/leetcode_client"
)

type Helper struct {
	Credit        *leetcode_client.Credit
	BaseURI       *leetcode_client.EndpointURI
	HTTPClient    *http.Client
	GraphQLClient *graphql.Client
}

var helper *Helper = nil

/**
 * Get the singleton instance of Helper
 */
func GetHelper(credit *leetcode_client.Credit, baseURI *leetcode_client.EndpointURI) *Helper {
	if helper != nil {
		return helper
	}
	HTTPClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// disable redirect
			return http.ErrUseLastResponse
		},
	}

	GraphQlClient := graphql.NewClient(baseURI.Graphql, HTTPClient)

	helper = &Helper{
		Credit:        credit,
		BaseURI:       baseURI,
		HTTPClient:    HTTPClient,
		GraphQLClient: GraphQlClient,
	}

	return helper
}

func (helper *Helper) SetCredit(credit *leetcode_client.Credit) {
	helper.Credit = credit
}

func (helper *Helper) SetBaseURI(baseURI *leetcode_client.EndpointURI) {
	helper.BaseURI = baseURI
}

func ParseCookie(cookies []string, key string) string {
	if cookies == nil {
		return ""
	}

	pattern := key + "=(.+?);"

	for _, cookie := range cookies {
		r, err := regexp.Compile(pattern)

		if err != nil {
			continue
		}

		matches := r.FindStringSubmatch(cookie)

		if len(matches) > 0 {
			return matches[1]
		}
	}

	return ""
}

func (helper *Helper) SwitchEndPoint(endPoint leetcode_client.EndPoint) {
	switch endPoint {
	case leetcode_client.US:
		helper.SetBaseURI(leetcode_client.BaseURI.US)
	case leetcode_client.CN:
		helper.SetBaseURI(leetcode_client.BaseURI.CN)
	}
}

type HTTPRequestParam struct {
	Method                  string
	URL                     string
	Referer                 string
	ResolveWithFullResponse bool
	Form                    map[string][]string
	Body                    string
	Header                  map[string]string
}

func (helper *Helper) HTTPRequest(param HTTPRequestParam) (*http.Response, error) {
	req, err := http.NewRequest(param.Method, param.URL, nil)

	if err != nil {
		return nil, err
	}

	if helper.Credit != nil {
		req.Header.Set("Cookie", `LEETCODE_SESSION=`+helper.Credit.Session+`;csrftoken=`+helper.Credit.CSRFToken)
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	req.Header.Set("X-CSRFToken", helper.Credit.CSRFToken)

	if param.Referer != "" {
		req.Header.Set("Referer", param.Referer)
	} else {
		req.Header.Set("Referer", helper.BaseURI.Base)
	}

	for k, v := range param.Header {
		req.Header.Set(k, v)
	}

	if param.Form != nil {
		req.PostForm = param.Form
	}

	if param.Body != "" {
		req.Body = ioutil.NopCloser(strings.NewReader(param.Body))
	}

	return helper.HTTPClient.Do(req)
}

type GraphqlRequestParam struct {
	Origin  string
	Referer string
	/**
	* Query should be a pointer to struct that corresponds to the GraphQL schema
	 */
	Query     interface{}
	Variables *map[string]interface{}
}

type GraphqlRequestType string

const (
	QUERY    GraphqlRequestType = "query"
	MUTATION GraphqlRequestType = "mutation"
)

func (helper *Helper) GraphqlRequest(reqType GraphqlRequestType, param GraphqlRequestParam) (res *interface{}, reqErr error) {
	if reqType == "query" {
		err := helper.GraphQLClient.Query(context.Background(), param.Query, *param.Variables)
		if err != nil {
			reqErr = err
		}
	} else if reqType == "mutation" {
		err := helper.GraphQLClient.Mutate(context.Background(), param.Query, *param.Variables)
		if err != nil {
			reqErr = err
		}
	}
	res = &param.Query
	return
}
