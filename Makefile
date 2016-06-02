build:
	docker run -ti --rm -v "$$PWD":/go/src/cas-server golang /bin/bash -c 'cd /go/src/cas-server && go get ./... && go build .'
