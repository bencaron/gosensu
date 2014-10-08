
test:
	SENSU_SERVER_URL="http://localhost:8889" go test 

build:
	go build .
