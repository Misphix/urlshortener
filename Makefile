.PHONY: all start stop clean test deploy

all:
	go build

start:
	docker start mysql-test
	docker start redis-test

stop:
	docker stop mysql-test
	docker stop redis-test

env:
	docker run -p 3306:3306 --name mysql-test -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=test -e MYSQL_PASSWORD=test -e MYSQL_DATABASE=short_url -d mysql:latest
	docker run -p 6379:6379 --name redis-test -d redis

clean:
	rm -f urlshortener
	rm -f *.out

test:
	go test -v -covermode=count -coverprofile=test.out ./...
	go tool cover -html=test.out

deploy:
	docker-compose build
	docker-compose up