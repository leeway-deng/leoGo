# Setup the global params
GOPATH=$(CURDIR)
export GOPATH=$(CURDIR)

# Setup the build info
BUILD_OS=$(os)
BUILD_ARCH=$(arch)
#BUILD_VERSION=`git describe --tags`
BUILD_VERSION=$(tag)
BUILD_TIME=`date +%FT%T%z`
LDFLAGS=-ldflags "-X main.BuildVersion=${BUILD_VERSION} -X main.BuildTime=${BUILD_TIME}"
DIST_PATH=${GOPATH}/dist/leoGo-${BUILD_ARCH}-${BUILD_VERSION}

dep:
#	env GIT_TERMINAL_PROMPT=1 go get --insecure github.com/henrylee2cn/faygo

release: build
#	rm -Rf ${DIST_PATH}
#	mkdir -p ${DIST_PATH}
#	cp ${GOPATH}/bin/* ${DIST_PATH}
	rm -Rf ${DIST_PATH}／config
	mkdir -p ${DIST_PATH}/config
	cp ${GOPATH}/config/* ${DIST_PATH}/config

.PHONY: all clean install mac windows linux
all: install

install: dep
	go install leoGo

build: dep
ifeq ($(BUILD_OS), darwin)
	CGO_ENABLED=0 GOOS=${BUILD_OS} GOARCH=${BUILD_ARCH} go build ${LDFLAGS} -o ${DIST_PATH}/leoGo $(GOPATH)/src/*.go
else ifeq ($(BUILD_OS), windows)
	CGO_ENABLED=0 GOOS=${BUILD_OS} GOARCH=${BUILD_ARCH} go build ${LDFLAGS} -o ${DIST_PATH}/leoGo $(GOPATH)/src/*.go
else ifeq ($(BUILD_OS), linux)
	CGO_ENABLED=0 GOOS=${BUILD_OS} GOARCH=${BUILD_ARCH} go build ${LDFLAGS} -o ${DIST_PATH}/leoGo $(GOPATH)/src/*.go
else
	echo "the Built OS was unknown"
endif

clean:
ifeq ($(DIST_PATH), ${GOPATH}/dist/leoGo--)
	rm -Rf ${GOPATH}/dist
else
	rm -Rf ${DIST_PATH}
endif
	go clean -r -i

#build the target version
#make build tag=1.0.0 os=darwin arch=amd64
#make build tag=1.0.0 os=linux arch=amd64
#make build tag=1.0.0 os=windows arch=amd64

#build and release the target version
#make release tag=1.0.0 os=darwin arch=amd64
#make release tag=1.0.0 os=linux arch=amd64
#make release tag=1.0.0 os=windows arch=amd64

#clean the target version
#make clean tag=1.0.0 os=darwin arch=amd64
#make clean tag=1.0.0 os=linux arch=amd64
#make clean tag=1.0.0 os=windows arch=amd64

#clean all targets
#make clean

#run in background
#nohup ./dist/leoGo-amd64-1.0.0/leoGo &
#nohup /home/go/leoGo/leoGo &
#cd /home/go/leoGo;nohup ./leoGo &

#deploy
#scp -r /data/leoGo/* root@47.94.210.163:/data/leoGo
#scp -r ./dist/leoGo-amd64-1.0.0/* root@47.94.210.163:/home/go/leoGo

#test
#http://192.168.1.106:8081/apidoc/#//
#http://47.94.210.163/apidoc/#//