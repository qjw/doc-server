.PHONY: build

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif


build: build_static

build_static:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./doc-server .

