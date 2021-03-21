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

serverTestSuite:
	ginkgo bootstrap

serverMocks:
	(cd server/repository/interfaces/ && go generate)

serverUnitTests:
	(cd server/repository/interfaces/ && go generate)
	ginkgo -r -skipPackage=apiTests --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress

serverApiTests:
	ginkgo --randomizeAllSpecs --failOnPending --cover --trace --race --progress server/apiTests/

serverAllTests:
	ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --compilers=2

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
	docker run --rm -d --name testdb --network bridge -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:latest

db:
	docker run -d --name kdb --restart unless-stopped --network bridge -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:latest

dependencies:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/mysql
	go get
	(cd client/ && npm install)
	echo "** DON'T FORGET THE .ENV FILE **"
