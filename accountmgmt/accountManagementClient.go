package accountmgmt

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
)

type AccountManagementClient struct {
	*http.Client
}

func NewClient(clientId, clientSecret, accountId string) *AccountManagementClient {

	endpointParams := url.Values{}
	endpointParams.Add("resource", "urn:dtaccount:" + accountId)

	clientConfig := &clientcredentials.Config{
		ClientID: clientId,
		ClientSecret: clientSecret,
		TokenURL: "https://sso.dynatrace.com/sso/oauth2/token",
		Scopes: []string{"account-idm-read", "account-idm-write"},
		EndpointParams: endpointParams,
	}

	clientContext := context.Background()
	client := clientConfig.Client(clientContext)

	return &AccountManagementClient{client}
}
