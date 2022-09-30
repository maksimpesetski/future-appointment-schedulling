
.PHONY: clean-build
clean-build:
	@rm -rf ${SERVICE_NAME}

.PHONY: build
build: clean-build
	@CGO_ENABLED=0 go build -o appointment-schedule ./cmd/main.go

.PHONY: local-service
local-service: build
local-service:
	./appointment-schedule
