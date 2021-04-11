tidy:
	go mod tidy
	go fmt ./...

retrieve-install-script:
	curl -sfL https://get.k3s.io > k3s-install.sh

run: tidy
	go run ./...

build: tidy
	go build -o k3s-helper ./... 