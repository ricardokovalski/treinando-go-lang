package rate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arquivei/foundationkit/errors"
	"github.com/ricardokovalski/treinando-go-lang/src/adapters"
	"github.com/ricardokovalski/treinando-go-lang/src/services/payment/get_rate"
)

type Client struct {
	client adapters.HTTPInterface
	url    string
}

func GetClient(url string) *Client {
	return &Client{
		client: &http.Client{},
		url:    url,
	}
}

func (cli Client) GetRate(ctx context.Context) (get_rate.Rate, error) {
	const op = errors.Op("adapters.payment.rate.GetRate")
	var rate = get_rate.Rate{}

	endpoint := fmt.Sprintf("%s%s", cli.url, getRate)
	resp, err := adapters.HTTPRequest(cli.client, http.MethodGet, endpoint, http.Header{}, []byte{})
	if err != nil {
		return rate, errors.E(op, "error", fmt.Errorf("error to make http request: %v", err))
	}

	if err = decodeRate(resp.Body, &rate); err != nil {
		return rate, errors.E(op, "error", fmt.Errorf("error to decode body: %v", err))
	}

	defer resp.Body.Close()

	return rate, nil
}

func decodeRate(body io.ReadCloser, rate *get_rate.Rate) error {
	content, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, rate)
}
