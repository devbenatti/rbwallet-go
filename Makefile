test-coverage:
	- go test -coverprofile ./report/cover.out ./... && go tool cover -html=./report/cover.out

logs:
	- docker-compose logs -f app