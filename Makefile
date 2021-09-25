# Load enviroment variables
include ./internal/config/.env

# Export enviroment variables to commands
export

# Variables
protosFolder=pkg/authiny_grpc/*.proto
docsFolder=docs/grpc
go_cover_file=coverage.out

help:: ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | sort | fgrep -v fgrep | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate_grpc_docs:: ## Command for generate the docs of grpc files
	@ protoc --doc_out=$(docsFolder) --doc_opt=html,index.html $(protosFolder)

generate_grpc:: ## Command for generate the pb.go files
	@ protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		$(protosFolder)

dev:: ## Run go Application with watcher
	@ fuser -k 6000/tcp
	@ gow -c run .

test:: ## Do the tests in go
	@ go test -race -coverprofile $(go_cover_file) ./...

cover:: test ## See coverage of tests, see more in https://go.dev/blog/cover
	@ go tool cover -func=$(go_cover_file)

cover-html:: test ## See of the coverage of the code on your default navigator
	@ go tool cover -html=$(go_cover_file)

docs-open:: ## Open documentation
	@ xdg-open docs/grpc/index.html

update-toolbox:: ## Open documentation
	@ go get -u github.com/yggbrazil/go-toolbox@latest
