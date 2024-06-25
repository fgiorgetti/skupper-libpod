all: generate-client

generate-client:
	./scripts/libpod-generate.sh

force-generate-client:
	FORCE=true ./scripts/libpod-generate.sh

clean:
	rm -rf ./client ./models

