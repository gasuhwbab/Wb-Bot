package requests

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gasuhwbab/wbbot/pkg/wbapi"
	"github.com/gasuhwbab/wbbot/pkg/wbapi/entities"
)

func GetWarehouseCoefficients(ctx context.Context,
	wbc *wbapi.WbClient, param int) (*entities.CoefficientList, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/acceptance/coefficients", wbc.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	req.URL.RawQuery = url.Values{
		"warehouseIDs": {fmt.Sprintf("%d", param)},
	}.Encode()
	res := entities.CoefficientList{}
	req = req.WithContext(ctx)
	if err := sendRequest(wbc, req, &res.Coefficients); err != nil {
		return nil, err
	}
	return &res, err
}
