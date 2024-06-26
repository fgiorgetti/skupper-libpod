all: generate-client

generate-client:
	./scripts/libpod-generate-v4.sh

force-generate-client:
	FORCE=true ./scripts/libpod-generate-v4.sh

clean:
	rm -rf ./v4

