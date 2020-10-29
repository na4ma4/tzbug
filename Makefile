.PHONY: test
test: artifacts/test/alpine-3.8 artifacts/test/alpine-3.12

artifacts/test/alpine-%: main.go Dockerfile
	-@mkdir -p "$(@D)"
	docker build --build-arg "GO_VERSION=1.15" --build-arg "ALPINE_VERSION=$(*)" -t timezone-bug:dev .
	docker run --rm timezone-bug:dev | tee "$(@)"
