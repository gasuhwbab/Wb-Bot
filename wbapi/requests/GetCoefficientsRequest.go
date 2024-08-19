package requests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gasuhwbab/wbbot/wbapi"
	"github.com/gasuhwbab/wbbot/wbapi/entities"
)

func GetCoefficients(ctx context.Context, wbc *wbapi.WbClient) (*entities.CoefficientList,
	error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/acceptance/coefficients", wbc.BaseUrl), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := entities.CoefficientList{}
	if err := sendRequest(wbc, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
