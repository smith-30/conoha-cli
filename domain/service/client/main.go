package client

import (
	"context"

	"github.com/smith-30/conoha-cli/domain/model"
)

type (
	Client interface {
		Auth(ctx context.Context) (model.AuthRes, error)
	}
)
