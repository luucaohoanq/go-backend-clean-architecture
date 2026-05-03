# Biến số cơ bản
BINARY_NAME=main
MAIN_PATH=./cmd/main.go
DOCKER_COMPOSE=docker-compose.yaml
GOLANGCI_LINT_VERSION=v2.12.1

.PHONY: all build test clean run dev docker-up docker-down lint fmt help install-tools tidy test-cov db-up

## help: Hiển thị danh sách các lệnh
help:
	@echo "Cách dùng: make [target]"
	@echo ""
	@echo "Các lệnh khả dụng:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


## install-tools: Cài đặt các công cụ bổ trợ cho dev (Air, Linter, Formatter)
install-tools:
	@echo "Installing tools..."
	go install github.com/air-verse/air@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/segmentio/golines@latest
# 	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)


## dev: Chạy app với Live Reload (Dùng Air) - Khuyên dùng khi code
dev:
	air

## run: Chạy app trực tiếp (Không reload)
run:
	go run $(MAIN_PATH)

## build: Build binary cho máy hiện tại
build:
	go build -o ./tmp/$(BINARY_NAME) $(MAIN_PATH)

## test: Chạy toàn bộ Unit Test nhanh
test:
	go test -v -short ./...

## test-cov: Chạy test và xuất báo cáo coverage (Giống JaCoCo)
test-cov:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

## fmt: Format code cực chuẩn (gofumpt)
fmt:
	gofumpt -l -w .
	golines . -w --max-len=120 --reformat-tags

## lint: Kiểm tra lỗi code tĩnh (golangci-lint)
lint:
	golangci-lint run

## tidy: Dọn dẹp và cập nhật go.mod (Giống Maven Update Project)
tidy:
	go mod tidy
	go mod verify

db-up:
	docker-compose up -d mongodb

docker-up:
	docker-compose up -d --build

docker-down:
	docker-compose down

## clean: Xóa các file tạm, file build
clean:
	go clean
	rm -rf ./tmp
	rm -f coverage.out