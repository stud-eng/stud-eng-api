test: ### test用
	docker-compose -f ./database/docker-compose.yaml down -v
	docker-compose -f ./database/docker-compose.yaml up -d
	go test -v -cover ./internal/...
	docker-compose -f ./database/docker-compose.yaml down -v

test-db: ### test用
	docker-compose -f ./database/docker-compose.yaml down -v
	docker-compose -f ./database/docker-compose.yaml up -d

test-go: ### test用
	go test -v -cover ./internal/...

