.PHONY: default start stop start-backend stop-backend rebuild-backend start-frontend stop-frontend rebuild-frontend

default: start

# backend 
start-backend:
	@if [ $$(docker ps -a --filter "name=^/smart-mirror-backend$$" -q) ]; then \
		echo "Re-starting smart-mirror-backend..."; \
		docker start smart-mirror-backend; \
	else \
		echo "Building smart-mirror-backend..."; \
		docker build -t smart-mirror-backend ./server; \
		docker run -d --name smart-mirror-backend -p 8080:8080 smart-mirror-backend; \
	fi

stop-backend:
	docker stop smart-mirror-backend

rebuild-backend:
	docker stop smart-mirror-backend
	docker container rm smart-mirror-backend
	docker build -t smart-mirror-backend ./server
	docker run -d --name smart-mirror-backend -p 8080:8080 smart-mirror-backend

# frontend
start-frontend:
	@if [ $$(docker ps -a --filter "name=^/smart-mirror-frontend$$" -q) ]; then \
		echo "Re-starting smart-mirror-frontend..."; \
		docker start smart-mirror-frontend; \
	else \
		echo "Building smart-mirror-frontend..."; \
		docker build -t smart-mirror-frontend ./frontend; \
		docker run -d --name smart-mirror-frontend -p 80:80 smart-mirror-frontend; \
	fi

stop-frontend:
	docker stop smart-mirror-frontend

rebuild-frontend: 
	docker stop smart-mirror-frontend
	docker container rm smart-mirror-frontend
	docker build -t smart-mirror-frontend ./frontend
	docker run -d --name smart-mirror-frontend -p 80:80 smart-mirror-frontend

# start all
start: start-backend start-frontend

# stop all
stop: stop-backend stop-frontend

# rebuild all
rebuild: rebuild-backend rebuild-frontend