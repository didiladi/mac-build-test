BINARY=mac-build-test

.PHONY: build clean test integration-test

build: clean
	@echo Build ${BINARY}
	@go build -o ${BINARY} ./...

clean:
	@echo Remove ${BINARY}
	@rm -f ${BINARY}

test: build
	@go test -tags=unit -v ./...

integration-test: build
	@go test -tags=cleanup -v ./...
	@go test -tags=integration -v ./...

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64
