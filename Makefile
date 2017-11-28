GIT_HASH := $(shell git log --pretty=format:%H -n 1)
GIT_TAG := $(shell git describe --always --tags --abbrev=0 | tail -c+2)
GIT_COMMIT := $(shell git rev-list v${GIT_TAG}..HEAD --count)
VERSION := ${GIT_TAG}.${GIT_COMMIT}
RELEASE := 1

.PHONY: default build prepare clean test rpm

default: clean prepare test build rpm

build: clean prepare
	mkdir -p build/root/usr/bin/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=${VERSION}-${RELEASE}" -o build/root/usr/bin/httpcat

prepare:
	go get .

clean:
	rm -rf build

test:
	go test

tar:
	tar -czvPf build/httpcat-${VERSION}-${RELEASE}.tar.gz -C build/root .

rpm:
	fpm -t rpm \
		-s "tar" \
		--description "httpcat is simple http server for debugging" \
		--url "https://github.com/AlexAkulov/httpcat" \
		--name "httpcat" \
		--version "${VERSION}" \
		--iteration "${RELEASE}" \
		-p build \
		build/httpcat-${VERSION}-${RELEASE}.tar.gz

deb:
	fpm -t deb \
		-s "tar" \
		--description "httpcat is simple http server for debugging" \
		--url "https://github.com/AlexAkulov/httpcat" \
		--name "httpcat" \
		--version "${VERSION}" \
		--iteration "${RELEASE}" \
		-p build \
		build/httpcat-${VERSION}-${RELEASE}.tar.gz
