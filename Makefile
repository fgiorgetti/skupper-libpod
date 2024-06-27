all: generate

install-swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.2

generate: install-swagger
	@./scripts/libpod-generate-v4.sh

clean:
	@rm -rf ./v4
