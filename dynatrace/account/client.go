package account

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
)

const (
	baseURLPrefix string = "https://api.dynatrace.com/env/v1/accounts/"

	oauth2ResourceKey    string = "resource"
	oauth2ResourcePrefix string = "urn:dtaccount:"
	oauth2ScopeIDMRead   string = "account-idm-read"
	oauth2ScopeIDMWrite  string = "account-idm-write"
	oauth2TokenURL       string = "https://sso.dynatrace.com/sso/oauth2/token"
)

type AcctMgmtClient struct {
	*http.Client
	accountId string
}

func NewClient(clientId, clientSecret, accountId string) *AcctMgmtClient {

	endpointParams := url.Values{}
	endpointParams.Add(oauth2ResourceKey, oauth2ResourcePrefix + accountId)

	clientConfig := &clientcredentials.Config{
		ClientID: clientId,
		ClientSecret: clientSecret,
		TokenURL: oauth2TokenURL,
		Scopes: []string{oauth2ScopeIDMRead, oauth2ScopeIDMWrite},
		EndpointParams: endpointParams,
	}

	clientContext := context.Background()
	client := clientConfig.Client(clientContext)

	return &AcctMgmtClient{client, accountId}
}

func GetBaseApiURL(client *AcctMgmtClient) string{
	return baseURLPrefix + client.accountId
}
