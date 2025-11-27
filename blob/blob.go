package blob

import (
	blobservices "github.com/squarecloudofc/sdk-api-go/blob/services"
	"github.com/squarecloudofc/sdk-api-go/rest"
)

type blobImpl struct {
	Objects blobservices.ObjectService
	Account blobservices.AccountService
}

func New(token string, opts ...rest.ConfigOpt) blobImpl {
	httpClient := rest.NewClient(token, rest.WithURL(rest.BLOBAPIURL))

	return blobImpl{
		Objects: &blobservices.ObjectServiceImpl{Client: httpClient},
		Account: &blobservices.AccountServiceImpl{Client: httpClient},
	}
}
