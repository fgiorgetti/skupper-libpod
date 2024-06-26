all: generate-v4

install-swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.2

generate-v4: install-swagger
	@./scripts/libpod-generate-v4.sh

clean:
	@rm -rf ./v4
