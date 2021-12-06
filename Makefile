PWD       		= $(shell pwd)
UID           	= $(shell id -u $$USER)
GID           	= $(shell id -g $$USER)

# Docker images
.PHONY: build-golang-development
build-golang-development: 
	docker build -t fakestore-development --file dockerfiles/Dockerfile.development .

# copy dlv binary to host
.PHONY: copy-dvl-from-container
copy-dvl-from-container: build-golang-development
	docker run --rm -ti --entrypoint /bin/sleep -d --name dlv-copier fakestore-development 10 && sudo docker cp dlv-copier:/go/bin/dlv /usr/local/bin/

# Docker containers
.PHONY: debug-container-development
debug-container-development: build-golang-development
	docker run --rm -ti -v $(PWD)/src:/src -u $(UID):$(GID) --workdir /src --entrypoint /bin/bash fakestore-development

.PHONY: run-container-development
run-container-development: build-golang-development
	docker run --rm -ti -p 2345:2345 --security-opt seccomp=unconfined --workdir /src --name golang-dev-env fakestore-development

.PHONY: stop-container-development
stop-container-development: 
	docker stop golang-dev-env 

.PHONY: run-application
run-application: build-golang-development
	docker run --rm -ti --workdir /src --name golang-application --entrypoint go fakestore-development run cmd/main.go