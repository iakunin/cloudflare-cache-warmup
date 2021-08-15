
build-and-push-docker-image: docker-build docker-push

docker-build:
	rm -f main && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . && \
	docker build -t iakunin/cloudflare-cache-warmup:latest . && \
	rm -f main

docker-push:
	docker push iakunin/cloudflare-cache-warmup:latest
