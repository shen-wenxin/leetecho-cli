package leetcode_client

import (
	helper "github.com/CallanBi/leetecho-cli/leetcode_client/helper"
)

type LeetCodeClient struct {
	Session        string
	CSRFToken      string
	URIS           *helper.EndpointURI
	HelperInstance *helper.Helper
}

func Build(username string, password string, endpoint helper.EndPoint) (leetcode *LeetCodeClient, err error) {
	endpointURI := helper.GetEndPoint(endpoint)
	credit, LoginErr := Login(username, password, endpointURI)
	if err != nil {
		err = LoginErr
		return
	}

	helperInstance := helper.GetHelper(credit, endpointURI)
	helperInstance.SwitchEndPoint(endpoint)

	leetcode = &LeetCodeClient{
		Session:        credit.Session,
		CSRFToken:      credit.CSRFToken,
		URIS:           endpointURI,
		HelperInstance: helperInstance,
	}
	return

}

func Login(username string, password string, endpointURI *helper.EndpointURI) (credit *helper.Credit, err error) {
	res, firstLoginErr := helper.HTTPRequest(&helper.HTTPRequestParam{
		URL: endpointURI.Login,
	})

	if err != nil {
		credit = nil
		err = firstLoginErr
	}

	defer res.Body.Close() // avoid memory leak

	header := res.Header
	token := helper.ParseCookie(header["Set-Cookie"], "csrftoken")
	// LeetCode CN returns nil here, but it does not matter
	var tempCredit *helper.Credit = &helper.Credit{
		CSRFToken: token,
	}

	// then login
	realRes, realErr := helper.HTTPRequest(&helper.HTTPRequestParam{
		Method: "POST",
		URL:    endpointURI.Login,
		Form: &map[string][]string{
			"csrfmiddlewaretoken": {credit.CSRFToken},
			"login":               {username},
			"password":            {password},
		},
	})

	if realErr != nil {
		credit = nil
		err = realErr
	}
	defer realRes.Body.Close()

	session := helper.ParseCookie(realRes.Header["Set-Cookie"], "LEETCODE_SESSION")
	csrfToken := helper.ParseCookie(realRes.Header["Set-Cookie"], "csrftoken")

	tempCredit.Session = session
	tempCredit.CSRFToken = csrfToken

	credit = tempCredit
	return
}
