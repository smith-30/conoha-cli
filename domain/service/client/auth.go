package client

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/smith-30/conoha-cli/domain/model"
)

func (c *ConohaClient) Auth(ctx context.Context) (*model.AuthRes, error) {
	path := "tokens"

	ar := model.NewAuthReq(c.UserName, c.Password, c.TenantID)
	body, err := json.Marshal(ar)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, POST, path, bytes.NewReader(body), true)
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
