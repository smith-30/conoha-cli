build:
	go build -v -ldflags "-X main.revision=$(git rev-parse --short HEAD)"