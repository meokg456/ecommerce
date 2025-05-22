sync:
	deck gateway sync --kong-addr http://localhost:8001 kong.yaml

docker-build:
	docker build -t internal-gateway:latest .

docker-run:
	docker run --name internal-gateway --env-file .env -d -p 8000:8000 internal-gateway:latest