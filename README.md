# meraki-dashboard-go-client
Meraki Dashboard API Client generated with go swagger

```
swagger generate client \
    -f https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json \
    -a meraki-dasboard-go-client
```

Usage example can be found in meraki_dashboard_api_client_test.go
```
EXPORT MERAKI_DASHBOARD_API_KEY=<YOUR API KEY>
go test -v
```

```
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
```
