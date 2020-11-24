GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_TEST=$(GO_CMD) test

TEST?=$$(go list ./...)

HOSTNAME=windtunnel.io
NAMESPACE=atlassian
NAME=bamboo
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_amd64

default: install

build:
	$(GO_BUILD) -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test:
	$(GO_TEST) -i $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 $(GO_TEST) $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 $(GO_TEST) $(TEST) -v $(TESTARGS) -timeout 120m
