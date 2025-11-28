package blobservices

import (
	"github.com/squarecloudofc/sdk-api-go/rest"
	"github.com/squarecloudofc/sdk-api-go/squarecloud"
)

type ObjectService interface {
	Upload()
	List(config *squarecloud.ObjectSearchParameters) (*squarecloud.APIResponse[squarecloud.ListObject], error)
	Delete(objectId string) (*squarecloud.APIResponse[any], error)
}

type AccountService interface {
	Metrics() (*squarecloud.APIResponse[squarecloud.BlobMetrics], error)
}

type ObjectServiceImpl struct {
	Client rest.Client
}

type AccountServiceImpl struct {
	Client rest.Client
}

var (
	_ ObjectService  = (*ObjectServiceImpl)(nil)
	_ AccountService = (*AccountServiceImpl)(nil)
)
