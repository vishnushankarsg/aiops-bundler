install-dev:
	go install github.com/cosmtrek/air@latest
	go mod tidy

generate-environment:
	go run ./scripts/genenv

generate-aimiddleware-pkg:
	abigen --abi=./abi/aimiddleware.json --pkg=aimiddleware --out=./pkg/aimiddleware/bindings.go

fetch-wallet:
	go run ./scripts/fetchwallet

prod-private-mode:
	air -c .air.private-mode.toml

prod-searcher-mode:
	air -c .air.searcher-mode.toml

prod-reset-default-data-dir:
	rm -rf /tmp/aiops_bundler
