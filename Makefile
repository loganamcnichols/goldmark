.PHONY: test fuzz lint

lint:
	golangci-lint run -c .golangci.yml ./...

test:
	go test -coverprofile=profile.out -coverpkg=github.com/loganamcnichols/goldmark,github.com/loganamcnichols/goldmark/ast,github.com/loganamcnichols/goldmark/extension,github.com/loganamcnichols/goldmark/extension/ast,github.com/loganamcnichols/goldmark/parser,github.com/loganamcnichols/goldmark/renderer,github.com/loganamcnichols/goldmark/renderer/html,github.com/loganamcnichols/goldmark/text,github.com/loganamcnichols/goldmark/util ./...

cov: test
	go tool cover -html=profile.out

fuzz:
	cd ./fuzz && go test -fuzz=Fuzz
