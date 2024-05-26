.PHONY: run build run-bin clean docs test doc-api

run:
	go run ./cmd/main.go

test:
	go test -v -cover ./...

build:
	mkdir -p bin
	go build -o bin/dashboard ./cmd/main.go

run-bin:
	./bin/dashboard

doc-api:
	raml2html api/api.raml > api/api.html

clean:
	rm -rf bin

docker-build:
	docker build -t analytics/dashboard:latest .

docker-run:
	docker run -p 8080:8080 analytics/dashboard:latest

docs:
	~/go/bin/godoc -http=:6060
