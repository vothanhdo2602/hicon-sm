export ENV=development
export IMAGE_NAME=hicon
export VERSION=latest
export HUB=vothanhdo2602

server:
	go run cmd/*

run-compose:
	docker compose -f deployment/docker/docker-compose.yaml up -d

build-docker:
	docker build -t ${IMAGE_NAME}:${VERSION} . -f ./deployment/docker/Dockerfile -D

run-docker:
	docker run -p 7979:7979 --name hicon hicon

update-submodules:
	git submodule update --init --recursive && \
	git submodule foreach git checkout $(branch) && \
	git submodule foreach git pull origin $(branch)

proto:
	protoc --go_out=./hicon-sm --go-grpc_out=./hicon-sm ./hicon/hicon-sm/*.proto

publish:
	make build-docker && docker tag ${IMAGE_NAME}:${VERSION} ${HUB}/${IMAGE_NAME}:${VERSION} && docker push ${HUB}/${IMAGE_NAME}:${VERSION} && docker push ${HUB}/${IMAGE_NAME}:latest
