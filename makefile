.PHONY: build
build:
	docker build -t azuki774/remo-exporter -f build/dockerfile-exporter .
	docker build -t azuki774/remo-manager -f build/dockerfile-api .

.PHONY: run
run:
	docker-compose -f build/docker-compose.yml up -d

.PHONY: stop
stop:
	docker-compose -f build/docker-compose.yml down

.PHONY: test
test:
	gofmt -l -w .
	go test ./... -v -cover
