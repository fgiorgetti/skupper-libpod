# Based on https://storage.googleapis.com/libpod-master-releases/swagger-v4.0.3.yaml
# A few changes have been done to the original spec due to some issues, like:
# https://github.com/containers/podman/issues/13092
#
# * IdResponse type (removed as not used and causing inconsistencies with IDResponse)
# * LibpodImageSummary (conflicts with ImageSummary)
#   - Removed x-go-name: ImageSummary
# * Mount type
#   - Added Destination (string)
#   - Added Options ([]string)
#
# Podman V5 update
# .definitions.InspectContainerConfig.properties.StopSignal (removed)
# 
prepare() {
    [[ -d v4 ]] && return 1
    mkdir v4
    cd v4
    go mod init github.com/fgiorgetti/skupper-libpod/v4 || true
    go get github.com/go-openapi/errors@v0.20.3 \
           github.com/go-openapi/runtime@v0.24.1 \
           github.com/go-openapi/strfmt@v0.21.3 \
           github.com/go-openapi/swag@v0.21.1 \
           github.com/go-openapi/validate@v0.22.0
    cd ../
    return 0
}

generate() {
    go generate
    cd v4
    go mod tidy
}

prepare || exit 0
generate
