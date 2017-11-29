package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/smith-30/conoha-cli/domain/model"
)

func (c *ConohaClient) Boot(ctx context.Context, authRes *model.AuthRes) error {
	// {tenant_id}​/servers/​{server_id}​/action
	s := []string{c.TenantID, "servers", c.ServerID, "action"}

	b := model.NewBoot()
	body, err := json.Marshal(b)
	if err != nil {
		return errors.Wrapf(err, "failed marshal %v", b)
	}

	req, err := c.newRequest(ctx, POST, strings.Join(s, "/"), bytes.NewReader(body), false)
	if err != nil {
		return errors.Wrapf(err, "failed generate request %v", req)
	}

	req.Header.Set("X-Auth-Token", authRes.GetToken())

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed HTTPClient.Do")
	}

	if res.StatusCode != http.StatusAccepted {
		return errors.Wrapf(errors.New("Http request error"), " %d ", res.StatusCode)
	}

	c.Logger.Sugar().Info("StatusCode ", res.StatusCode)

	return nil
}
