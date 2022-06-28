package helper

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/shurcooL/graphql"
)

type Helper struct {
	Credit        *Credit
	BaseURI       *EndpointURI
	GraphQLClient *graphql.Client
}

var helper *Helper = nil

var httpClient *http.Client = nil

/**
 * Get the singleton instance of Helper
 */
func GetHelper(credit *Credit, baseURI *EndpointURI) *Helper {
	if helper != nil {
		return helper
	}

	GraphQlClient := graphql.NewClient(baseURI.Graphql, httpClient)

	helper = &Helper{
		Credit:        credit,
		BaseURI:       baseURI,
		GraphQLClient: GraphQlClient,
	}

	return helper
}

func GetHTTPClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				// disable redirect
				return http.ErrUseLastResponse
			},
		}
	}
	return httpClient
}

func (helper *Helper) SetCredit(credit *Credit) {
	helper.Credit = credit
}

func (helper *Helper) SetBaseURI(baseURI *EndpointURI) {
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

func GetEndPoint(endPoint EndPoint) *EndpointURI {
	switch endPoint {
	case US:
		return BaseURI.US
	case CN:
		return BaseURI.CN
	}
	return nil
}

func (helper *Helper) SwitchEndPoint(endPoint EndPoint) {
	switch endPoint {
	case US:
		helper.SetBaseURI(BaseURI.US)
	case CN:
		helper.SetBaseURI(BaseURI.CN)
	}
}

type HTTPRequestParam struct {
	Method                  string
	URL                     string
	Referer                 string
	ResolveWithFullResponse bool
	Form                    *map[string][]string
	Body                    string
	Header                  *map[string]string
}

func HTTPRequest(param *HTTPRequestParam) (*http.Response, error) {
	var (
		method string = param.Method
	)

	if method == "" {
		method = "GET"
	}

	form := url.Values{}

	if param.Form != nil {
		for key, value := range *param.Form {
			form.Add(key, value[0])
		}
	}

	req, err := http.NewRequest(method, param.URL, strings.NewReader(form.Encode()))

	if err != nil {
		return nil, err
	}

	if helper != nil {
		if helper.Credit != nil {
			req.Header.Set("Cookie", `LEETCODE_SESSION=`+helper.Credit.Session+`;csrftoken=`+helper.Credit.CSRFToken)
		}
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	if helper != nil {
		req.Header.Set("X-CSRFToken", helper.Credit.CSRFToken)
	}

	if param.Referer != "" {
		req.Header.Set("Referer", param.Referer)
	} else {
		if helper != nil {
			req.Header.Set("Referer", helper.BaseURI.Base)
		} else {
			req.Header.Set("Referer", BaseURI.CN.Base)
		}
	}

	if param.Header != nil {
		for k, v := range *param.Header {
			req.Header.Set(k, v)
		}
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.PostForm = form
	}

	if param.Body != "" {
		req.Body = ioutil.NopCloser(strings.NewReader(param.Body))
	}

	innerHttpClient := GetHTTPClient()

	return innerHttpClient.Do(req)
}

// HTTPParseHandler parses the response and returns the struct of the response and error
//  @param resp *http.Response
//  @param requestErr error
//  @return interface{}  the struct of the response
//  @return error  	 *ErrorResp{}
func HTTPParseHandler(resp *http.Response, requestErr error) (interface{}, error) {
	if requestErr != nil {
		return nil, requestErr
	}

	defer resp.Body.Close()

	var respJSONStruct interface{}

	if resp.StatusCode != http.StatusOK {
		decodeErr := DecodeResponseJSONBody(resp, &respJSONStruct)
		if decodeErr != nil {
			return nil, &ErrorResp{
				Status:  LeetechoStatus(resp.Status),
				Code:    LeetechoCode(resp.StatusCode),
				Message: LeetechoMessage(resp.Status),
			}
		}
		// type assertion
		if str, ok := respJSONStruct.(string); ok {
			return nil, &ErrorResp{
				Status:  LeetechoStatus(resp.Status),
				Code:    LeetechoCode(resp.StatusCode),
				Message: LeetechoMessage(str),
			}
		} else {
			return nil, &ErrorResp{
				Status:  LeetechoStatus(resp.Status),
				Code:    LeetechoCode(resp.StatusCode),
				Message: LeetechoMessage(resp.Status),
			}
		}
	} else {
		decodeErr := DecodeResponseJSONBody(resp, &respJSONStruct)
		if decodeErr != nil {
			return nil, &ErrorResp{
				Status:  DECODE_JSON_ERROR_STATUS,
				Code:    DECODE_JSON_ERROR_CODE,
				Message: DECODE_JSON_ERROR_MESSAGE,
			}
		}
		return respJSONStruct, nil
	}
}

func WrappedHTTPRequest(param *HTTPRequestParam) (rawResp *http.Response, respJSON interface{}, err error) {
	resp, err := HTTPRequest(param)
	respJSON, err = HTTPParseHandler(resp, err)
	rawResp = resp
	return
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
