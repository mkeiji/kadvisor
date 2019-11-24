run:
	(cd client/ && pub run webdev build)
	go build
	go run kadvisor

debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient

runDebug:
	(cd client/ && pub run webdev build)
	go build -gcflags "-N -l"
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./kadvisor

build:
	(cd client/ && pub run build_runner build && pub run webdev build)
	go build

dockerimg:
	make build
	docker build -t mgkeiji/kadvisor .

testdb:
	docker run --rm -d --name test -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7
