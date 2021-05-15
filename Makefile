.PHONY: all test

all:
	go build

start:
	docker start mysql-test

stop:
	docker stop mysql-test

mysql:
	docker run -p 3306:3306 --name mysql-test -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_USER=test -e MYSQL_PASSWORD=test -e MYSQL_DATABASE=short_url -d mysql:latest

clean:
	rm -f urlshorterner
	rm -f *.out

test:
	go test -v -covermode=count -coverprofile=test.out ./...
	go tool cover -html=test.out