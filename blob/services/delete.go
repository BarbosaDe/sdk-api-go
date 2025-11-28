package blobservices

import (
	"encoding/json"
	"net/http"

	"github.com/squarecloudofc/sdk-api-go/rest"
	"github.com/squarecloudofc/sdk-api-go/squarecloud"
)

type payload struct {
	ID string `json:"object"`
}

func (o *ObjectServiceImpl) Delete(objectId string) (*squarecloud.APIResponse[any], error) {
	var r squarecloud.APIResponse[any]

	data := payload{ID: objectId}
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	err = o.Client.Request(
		http.MethodDelete,
		rest.EndpointBlobObject(),
		jsonData, &r,
		rest.WithHeader("Content-Type", "application/json"),
	)

	if err != nil {
		return nil, err
	}

	return &r, nil

}
