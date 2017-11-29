build-cli:
	go build -v -o ./build/bin/conoha-cli -ldflags "-X main.revision=$(git rev-parse --short HEAD)" main.go

build-cron:
	go build -v -o ./build/bin/cron -ldflags "-X main.revision=$(git rev-parse --short HEAD)" cron.go