package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/dashboard-api-go-swag/client"
	"github.com/dashboard-api-go-swag/client/organizations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

/*
This is a test using the meraki client.
*/
func TestClient(t *testing.T) {
	if os.Getenv("MERAKI_DASHBOARD_API_KEY") == "" {
		t.SkipNow()
	}
	params := organizations.NewGetOrganizationsParams()
	client := client.New(httptransport.New(client.DefaultHost, client.DefaultBasePath, nil), strfmt.Default)
	apiKeyHeaderAuth := httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", os.Getenv("MERAKI_DASHBOARD_API_KEY"))
	resp, _ := client.Organizations.GetOrganizations(params, apiKeyHeaderAuth)
	fmt.Println(resp.Payload[0].Name)
	assert.Greater(t, len(resp.Payload[0].Name), 1)
}
