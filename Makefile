format:
	@find . -name \*.go | xargs goimports -d
verify:
	@staticcheck ./... && go vet ./...
test:
	@go test ./... -coverprofile=./test_results/cover.out && go tool cover -html=./test_results/cover.out -o ./test_results/cover.html
serve:
	@go run ./main.go
