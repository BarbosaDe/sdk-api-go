package blobservices

import (
	"net/http"

	"github.com/squarecloudofc/sdk-api-go/rest"
	"github.com/squarecloudofc/sdk-api-go/squarecloud"
)

func (o *AccountServiceImpl) Metrics() (*squarecloud.APIResponse[squarecloud.BlobMetrics], error) {
	var r squarecloud.APIResponse[squarecloud.BlobMetrics]
	err := o.Client.Request(http.MethodGet, rest.EndpointBlobMetrics(), nil, &r)

	if err != nil {
		return nil, err
	}

	return &r, nil
}
