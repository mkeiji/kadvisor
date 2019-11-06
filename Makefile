run:
	(cd client/ && pub run webdev build)
	go build
	go run kadvisor

build:
	(cd client/ && pub run webdev build)
	go build

dockerimg:
	make build
	docker build -t mgkeiji/kadvisor .

testdb:
	docker run --rm -d --name test -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7