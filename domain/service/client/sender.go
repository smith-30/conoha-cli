package client

import (
	"context"
	"net/http"
	"net/url"

	"encoding/json"
	"io"
	"path"

	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	POST = "POST"
	GET  = "GET"
)

type (
	ConohaClient struct {
		AuthURL    *url.URL
		ApiURL     *url.URL
		HTTPClient *http.Client

		UserName string
		Password string

		TenantID string
		ServerID string

		Logger *zap.Logger
	}
)

func NewConohaClient(l *zap.Logger) (*ConohaClient, error) {
	aah := os.Getenv("API_AUTH_HOST")
	authParsed, err := url.ParseRequestURI(aah)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", aah)
	}

	ah := os.Getenv("API_HOST")
	apiParsed, err := url.ParseRequestURI(ah)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse url: %s", ah)
	}

	return &ConohaClient{
		AuthURL:    authParsed,
		ApiURL:     apiParsed,
		HTTPClient: &http.Client{},

		UserName: os.Getenv("API_USER"),
		Password: os.Getenv("API_PASSWORD"),

		TenantID: os.Getenv("TENANT_ID"),
		ServerID: os.Getenv("SERVER_ID"),

		Logger: l,
	}, nil
}

func (c *ConohaClient) newRequest(ctx context.Context, method, spath string, body io.Reader, isAuth bool) (*http.Request, error) {
	var u url.URL

	if isAuth {
		u = *c.AuthURL
		u.Path = path.Join(c.AuthURL.Path, spath)
	} else {
		u = *c.ApiURL
		u.Path = path.Join(c.ApiURL.Path, spath)
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
