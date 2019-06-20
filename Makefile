test:
	@go test -v -race

cover:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out

dep:
	@dep ensure -v

dep-init:
	echo ${GOPATH}
	@dep init

dep-status:
	@dep status 

dep-update:
	@dep ensure -update

.PHONY:	test
