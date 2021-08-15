
build-and-push-docker-image: docker-build docker-push

docker-build:
	docker build -t iakunin/cloudflare-cache-warmup:latest .

docker-push:
	docker push iakunin/cloudflare-cache-warmup:latest
