LANTERN_ARTIFACT=lantern
LANTERN_PKG=target/pkg/$(LANTERN_ARTIFACT)

.PHONY: lantern-local
lantern-local:
	DOCKER_BUILDKIT=1 docker build -f build/docker/Dockerfile --target=lantern-build --rm -t $(LANTERN_ARTIFACT):0-SNAPSHOT $(DOCKER_IMAGE_TAG) . && \
	docker run -i -p 9090:9090 $(LANTERN_ARTIFACT):0-SNAPSHOT

.PHONY: clean
clean:
	rm -f $(LANTERN_ARTIFACT)
	rm -rf $(LANTERN_PKG)

.PHONY: compile
compile:
	go build -o lantern github.com/meroedu/lantern/cmd/lantern

all: clean compile