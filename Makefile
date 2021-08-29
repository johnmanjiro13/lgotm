.PHONY: test mockgen

test:
	go vet ./...
	go test -race -cover ./...

mockgen:
	mockgen -destination cmd/mock_cmd/mock_cmd.go github.com/johnmanjiro13/lgotm/cmd CustomSearchRepository
