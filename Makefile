APP_NAME ?= app

.PHONY: go-install-air
go-install-air: ## Installs the air build reload system using 'go install'
	go install github.com/air-verse/air@latest

.PHONY: get-install-air
get-install-air: ## Installs the air build reload system using cUrl
	curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

.PHONY: go-install-templ
go-install-templ: ## Installs the templ Templating system for Go
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: get-install-tailwindcss
get-install-tailwindcss: ## Installs the tailwindcss cli
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
	chmod +x tailwindcss-linux-x64
	mv tailwindcss-linux-x64 tailwindcss

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	go test -race -v -timeout 30s ./...

.PHONY: tailwind-watch
tailwind-watch:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-watch
templ-watch:
	templ generate --watch

.PHONY: templ-generate
templ-generate:
	templ generate
	
.PHONY: dev
dev:
	go build -o ./tmp/main ./cmd/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go

.PHONY: docker-build
docker-build:
	docker-compose -f ./dev/docker-compose.yml build

.PHONY: docker-up
docker-up:
	docker-compose -f ./dev/docker-compose.yml up

.PHONY: docker-dev
docker-dev:
	docker-compose -f ./dev/docker-compose.yml -f ./dev/docker-compose.dev.yml up

.PHONY: docker-down
docker-down:
	docker-compose -f ./dev/docker-compose.yml down

.PHONY: docker-clean
docker-clean:
	docker-compose -f ./dev/docker-compose.yml down -v --rmi all