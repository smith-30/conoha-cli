package subcmd

import (
	"github.com/smith-30/conoha-cli/domain/service/client"

	"context"

	"time"

	"go.uber.org/zap"
)

type Halt struct{}

func (f *Halt) Help() string {
	return "conoha-cli"
}

func (f *Halt) Run(args []string) int {
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

	err = cc.Halt(ctx, r)
	if err != nil {
		logger.Sugar().Fatalf("%v\n", err)
		return 1
	}

	return 0
}

func (f *Halt) Synopsis() string {
	return "conoha-cli"
}
