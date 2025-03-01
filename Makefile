.PHONY: default build-backend run-backend build-frontend run-frontend stop-backend stop-frontend

default: build-backend run-backend build-frontend run-frontend

# backend 
build-backend:
	docker build -t smart-mirror-backend ./server

run-backend:
	docker run -d --name smart-mirror-backend -p 8080:8080 smart-mirror-backend

stop-backend:
	docker stop smart-mirror-backend

# frontend
build-frontend:
	docker build -t smart-mirror-frontend ./frontend

run-frontend:
	docker run -d --name smart-mirror-frontend -p 80:80 smart-mirror-frontend

stop-frontend:
	docker stop smart-mirror-frontend