DB_CONTIANER=pg
DB_IMG=postgres:alpine

db:
	docker rm -f ${DB_CONTIANER}

	docker run \
		--name ${DB_CONTIANER} \
		-d \
		--env-file ./.env \
		-p 5432:5432 \
		-v ${PWD}/initdb/init-db.sh:/docker-entrypoint-initdb.d/init-db.sh \
		${DB_IMG}

db-remove:
	docker stop ${DB_CONTIANER}
	docker rm -f ${DB_CONTIANER}
	docker rmi -f ${DB_IMG}
	docker system prune -fa