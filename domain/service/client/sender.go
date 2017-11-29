package client

import (
	"context"
	"net/http"
	"net/url"

	"encoding/json"
	"io"
	"path"

	"os"

	"bytes"

	"github.com/pkg/errors"
	"github.com/smith-30/conoha-cli/domain/model"
	"go.uber.org/zap"
)

const (
	POST = "POST"
	GET  = "GET"
)

type (
	ConohaClient struct {
		URL        *url.URL
		HTTPClient *http.Client

		UserName string
		Password string

		TenantID string
		ServerID string

		Logger *zap.Logger
	}
)

func NewConohaClient(l *zap.Logger) (*ConohaClient, error) {
	u := os.Getenv("API_HOST")
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", u)
	}

	return &ConohaClient{
		URL:        parsedURL,
		HTTPClient: &http.Client{},

		UserName: os.Getenv("API_USER"),
		Password: os.Getenv("API_PASSWORD"),

		TenantID: os.Getenv("TENANT_ID"),
		ServerID: os.Getenv("SERVER_ID"),

		Logger: l,
	}, nil
}

func (c *ConohaClient) Auth(ctx context.Context) (*model.AuthRes, error) {
	path := "tokens"

	ar := model.NewAuthReq(c.UserName, c.Password, c.TenantID)
	body, err := json.Marshal(ar)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, POST, path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed HTTPClient.Do")
	}

	// Check status code here…
	var r model.AuthRes
	if err := decodeBody(res, &r); err != nil {
		return nil, errors.Wrapf(err, "failed decode %v", res)
	}

	return &r, nil
}

func (c *ConohaClient) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *ConohaClient) Boot(ctx context.Context, method, spath string, body io.Reader) error {
	path :=

	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)

	req, err := http.NewRequest(POST, u.String(), body)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")

	return nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
