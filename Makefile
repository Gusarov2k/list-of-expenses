MIN_COVERAGE = 60
test:
	go test ./... -v -race -timeout 120s -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt \
	| grep total | sed 's/\%//g' | awk '{err=0}{if ($$3 < $(MIN_COVERAGE)) {printf "=== FAIL: Coverage failed at %.2f%%\n", $$3; err=1}} END {exit err}'
lint:
	golangci-lint run --deadline=5m -v
s:
	go run cmd/api/api.go
up:
	docker-compose -f docker/docker-compose.yml up -d --build

down:
	docker-compose -f docker/docker-compose.yml down
