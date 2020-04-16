BINARY='graphviz.wasm'
all: build

build:
	GOOS=js GOARCH=wasm go build -o ${BINARY} ./cmd/graphviz

build-prod:
	GOOS=js GOARCH=wasm go build -o ${BINARY} -ldflags "-s -w" ./cmd/graphviz
