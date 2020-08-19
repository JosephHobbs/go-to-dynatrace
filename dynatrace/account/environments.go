package account

import (
	"encoding/json"
	"io/ioutil"
)

const (
	environmentsUri string = "/environments"
)

type Tenant struct {
	Id string
	Name string
}

type ManagementZone struct {
	Id string
	Name string
	Parent string
}

type EnvironmentDetails struct {
	Tenants []Tenant `json:"tenantResources"`
	ManagementZones []ManagementZone `json:"managementZoneResources"`
}

func GetEnvironments(client *AcctMgmtClient) EnvironmentDetails{

	requestURL := GetBaseApiURL(client) + environmentsUri

	response, err := client.Get(requestURL)
	if err != nil {
		panic(err)
	}

	var environmentDetails EnvironmentDetails

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	json.Unmarshal(body, &environmentDetails)

	return environmentDetails
}
