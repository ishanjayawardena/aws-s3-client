.PHONY: update_deps
update_deps:
	go get github.com/aws/aws-sdk-go@latest

.PHONY: build
build: update_deps
	go build -ldflags="-X 'main.Environment=eu-stg'" -o stg-s3-eu
	go build -ldflags="-X 'main.Environment=us-stg'" -o stg-s3-us
	go build -ldflags="-X 'main.Environment=eu-prod'" -o prod-s3-eu
	go build -ldflags="-X 'main.Environment=us-prod'" -o prod-s3-us

.PHONY: build
install: build
	cp stg-s3-* prod-s3-* $(GOPATH)/bin
	@echo "Installed binaries to $(GOPATH)"
	rm stg-s3-* prod-s3-*

.PHONY: clean
clean:
	@rm prod-s3-* stg-s3-* 2>/dev/null || true
	@rm $(GOPATH)/bin/prod-s3-* 2>/dev/null || true
	@rm $(GOPATH)/bin/stg-s3-* 2>/dev/null || true

