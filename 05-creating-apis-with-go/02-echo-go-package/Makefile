run:
	docker-compose -f ./docker/docker-compose.yml up -d --build

remove:
	docker-compose -f ./docker/docker-compose.yml down

clean:
	docker-compose -f ./docker/docker-compose.yml down
	docker system prune -fa