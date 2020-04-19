GOCMD := go
GOTEST := ${GOCMD} test
GOTESTCI := ${GOTEST} -race -v -coverprofile=coverage.txt -covermode=atomic
test.ci:
	${GOTESTCI} ./...
