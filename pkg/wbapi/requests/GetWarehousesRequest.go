package requests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gasuhwbab/wbbot/pkg/wbapi"
	"github.com/gasuhwbab/wbbot/pkg/wbapi/entities"
)

func GetWarehouses(ctx context.Context, wbc *wbapi.WbClient) (*entities.WarehouseList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/warehouses", wbc.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := entities.WarehouseList{}
	if err := sendRequest(wbc, req, &res.Warehouses); err != nil {
		return nil, err
	}
	return &res, nil
}
