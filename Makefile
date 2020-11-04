run:
	(cd client/ && npm run build)
	go build
	go run kadvisor

runServer:
	go run main.go

formatServer:
	(cd server/prettier/ && go run prettier.go -path ../)

runClient:
	(cd client/ && nx serve)

runTestsServer:
	ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress

formatClient:
	(cd client/ && npm run format)

formatAll:
	(make formatServer && make formatClient)

debug:
	dlv debug --headless --listen=:2345 --api-version=2 --accept-multiclient

runDebug:
	(cd client/ && npm run build)
	go build -gcflags "-N -l"
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./kadvisor

build:
	(cd client/ && npm install && nx build)
	go build

dockerimg:
	make build
	docker build -t mgkeiji/kadvisor .

testdb:
	docker run --rm -d --name test --network bridge -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7

db:
	docker run -d --name kdb --restart unless-stopped --network bridge -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7

dependencies:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get
	(cd client/ && npm install)
	echo "** DON'T FORGET THE .ENV FILE **"
