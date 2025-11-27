package blobservices

import (
	"github.com/squarecloudofc/sdk-api-go/rest"
)

type ObjectService interface {
	Upload()
	List()
	Delete()
}

type AccountService interface {
	Metrics()
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
