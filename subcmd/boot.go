package subcmd

import (
	"github.com/smith-30/conoha-cli/domain/service/client"

	"context"

	"time"

	"go.uber.org/zap"
)

/** foo サブコマンド用の実装 **/
type Boot struct{}

func (f *Boot) Help() string {
	return "conoha-cli"
}

func (f *Boot) Run(args []string) int {
	logger, _ := zap.NewDevelopment()

	cc, err := client.NewConohaClient(logger)
	if err != nil {
		logger.Sugar().Fatalf("%v\n", err)
		return 1
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := cc.Auth(ctx)
	if err != nil {
		logger.Sugar().Fatalf("%v\n", err)
		return 1
	}

	err = cc.Boot(ctx, r)
	if err != nil {
		logger.Sugar().Fatalf("%v\n", err)
		return 1
	}

	return 0
}

func (f *Boot) Synopsis() string {
	return "conoha-cli"
}
