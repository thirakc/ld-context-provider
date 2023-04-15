IMAGE?=kcskbcnd93.kcs:5000/aries-framework-go/agent-utils
TAG?=latest
DOCKER_FILE?=Dockerfile

test:
	go test -v -race -coverprofile cover.out ./...

out:
	go tool cover -html=cover.out

run:
	go run cmd/main.go

mongo:
	docker run -d -it --rm -p 27017:27017 mongo