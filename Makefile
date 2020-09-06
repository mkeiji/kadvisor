run:
	(cd client/ && npm run build)
	go build
	go run kadvisor

runServer:
	go run main.go

formatServer:
	(cd server/prettier/ && go run prettier.go -path ../)

runClient:
	(cd client/ && npm start)

formatClient:
	(cd client/ && npm run format)

debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient

runDebug:
	(cd client/ && npm run build)
	go build -gcflags "-N -l"
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./kadvisor

build:
	(cd client/ && npm install && npm run build)
	go build

dockerimg:
	make build
	docker build -t mgkeiji/kadvisor .

testdb:
	docker run --rm -d --name test -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7

db:
	docker run -d --name kdb --restart unless-stopped -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7

dependencies:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/jinzhu/gorm
	go get
	(cd client/ && npm install)
	echo "** DON'T FORGET THE .ENV FILE **"
