SERVER=cmd/server/main.go
LIB=$(shell find . -type f -name '*.go' -not -path "cmd/*")
DOCKERTAG=m1w
BINARY=bin/m1w
SBINARY=bin/m1ws

$(BINARY): ${SERVER} ${LIB}
	go build ${LDFLAGS} -o ${BINARY} ${SERVER}

$(SBINARY): ${SERVER} ${LIB}
	CGO_ENABLED=0 go build --tags netgo -a -v -o ${SBINARY} ${SERVER}

.PHONY: docker
docker: ${SBINARY} Dockerfile
	docker build -t ${DOCKERTAG} .

.PHONY: clean
clean:
	rm -rf bin/

