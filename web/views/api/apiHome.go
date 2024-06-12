package api

const ApiHomeJsonFile = "./web/views/api/apiHome.json"

type ApiHomeJSON struct {
	Title               string `json:"title"`
	About               string `json:"about"`
    APITableEndpointCol string `json:"api_table_endpoint_col"`
    APITableDescCol     string `json:"api_table_desc_col"`
	APIEndpoints        []struct {
		Endpoint string `json:"endpoint"`
		Desc     string `json:"desc"`
	} `json:"api_endpoints"`
}
