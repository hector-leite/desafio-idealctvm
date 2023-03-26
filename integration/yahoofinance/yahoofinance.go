package yahoofinance

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hector-leite/desafio-idealctvm/domain/contract"
	"github.com/hector-leite/desafio-idealctvm/domain/entity"
	"github.com/hector-leite/desafio-idealctvm/infra/config"
	"github.com/hector-leite/desafio-idealctvm/integration/yahoofinance/viewmodel"
)

type YahooFinance struct {
	httpClient contract.HTTPClient
	config     config.YahooFinanceConfig
}

func NewYahooFinanceIntegration(
	httpClient contract.HTTPClient,
	config config.YahooFinanceConfig,
) contract.YahooFinance {
	return YahooFinance{
		httpClient: httpClient,
		config:     config,
	}
}

func (y YahooFinance) GetQuotesByAssets(assets ...string) ([]entity.Quote, error) {
	finalURL := fmt.Sprintf("%squote?region=US&lang=en&symbols=%s", y.config.APIURL, strings.Join(assets, "%2C"))
	req, err := http.NewRequest(http.MethodGet, finalURL, nil)
	if err != nil {
		return []entity.Quote{}, err
	}

	req.Header.Add("X-API-KEY", y.config.APIKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := y.httpClient.Do(req)
	if err != nil {
		return []entity.Quote{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []entity.Quote{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []entity.Quote{}, errors.New("error status code")
	}

	var response viewmodel.QuoteResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []entity.Quote{}, err
	}

	return response.Parse(), nil
}
