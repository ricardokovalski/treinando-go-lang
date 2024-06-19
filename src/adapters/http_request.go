package adapters

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/arquivei/foundationkit/errors"
)

type HTTPInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

// HTTPRequest make the request for a url
func HTTPRequest(client HTTPInterface, method, url string, headers http.Header, body []byte) (resp *http.Response, err error) {
	const op = errors.Op("adapters.httpRequest")

	req, err := newAuthorizedRequest(method, url, body)
	if err != nil {
		return resp, errors.E(op, fmt.Errorf("error to create new request: %v", err))
	}

	resp, err = doRequest(req, client)
	if err != nil {
		var responseBody string

		if resp != nil {
			respBody, _ := io.ReadAll(resp.Body)
			responseBody = string(respBody)

			resp.Body = io.NopCloser(bytes.NewBuffer(respBody))
		}
		return resp, errors.E(op, fmt.Errorf("error to do request: %v, responsebody:'%s'", err, responseBody))
	}

	return resp, nil
}

func newAuthorizedRequest(method, endpoint string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewReader([]byte(body)))
	if err != nil {
		return nil, err
	}

	//req.SetBasicAuth(secret, "")

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func doRequest(req *http.Request, cli HTTPInterface) (*http.Response, error) {
	const op = errors.Op("adapters.doRequest")
	resp, err := cli.Do(req)
	if err != nil {
		return resp, errors.E(op, err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return resp, errors.E(op, fmt.Errorf("status code: %d. Received unexpected status code", resp.StatusCode))
	}

	return resp, nil
}
