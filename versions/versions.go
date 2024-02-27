package versions

var (
	// Version is the service version
	// Set with: go build -ldflags="-X github.com/amaury95/protoc-gen-go-tag/versions.Version=$(git describe --tags 2>/dev/null)"
	Version string = "v0.0.0"

	// GitHash is the hash of git commit the service is built from
	// Set with: go build -ldflags="-X github.com/amaury95/protoc-gen-go-tag/versions.GitHash=$(git rev-parse --short HEAD)"
	GitHash string = ""

	// BuildTime build time in ISO-8601 format
	// Set with: go build -ldflags="-X github.com/amaury95/protoc-gen-go-tag/versions.BuildTime=$(date -u +'%Y-%m-%dT%H:%M:%SZ')"
	BuildTime string = ""
)
