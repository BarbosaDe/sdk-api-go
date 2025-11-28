package blobservices

import (
	"net/http"
	"net/url"

	"github.com/squarecloudofc/sdk-api-go/rest"
	"github.com/squarecloudofc/sdk-api-go/squarecloud"
)

type ListResponse = squarecloud.APIResponse[squarecloud.ListObject]

func (o *ObjectServiceImpl) List(config *squarecloud.ObjectSearchParameters) (*ListResponse, error) {
	var r ListResponse
	var queryParams url.Values
	baseURL := rest.EndpointBlobObject()

	endpoint, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	if config != nil {
		queryParams = url.Values{}
		if config.Prefix != "" {
			queryParams.Add("prefix", config.Prefix)
		}

		if config.ContinuationToken != "" {
			queryParams.Add("continuationToken", config.ContinuationToken)
		}

		endpoint.RawQuery = queryParams.Encode()
	}

	err = o.Client.Request(http.MethodGet, endpoint.String(), nil, &r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
