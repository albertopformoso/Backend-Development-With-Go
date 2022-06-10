DB_CONTIANER_MYSQL=mysql
DB_IMG_MYSQL=mysql:5.6.27

DB_CONTIANER_PSQL=pg
DB_IMG_PSQL=postgres:alpine

db-psql:
	docker rm -f ${DB_CONTIANER_PSQL}

	docker run \
		--name ${DB_CONTIANER_PSQL} \
		-d \
		--env-file ./.env \
		-p 5432:5432 \
		-v ${PWD}/initdb/init-db.sh:/docker-entrypoint-initdb.d/init-db.sh \
		${DB_IMG_PSQL}

db-remove-psql:
	docker stop ${DB_CONTIANER_PSQL}
	docker rm -f ${DB_CONTIANER_PSQL}
	docker rmi -f ${DB_IMG_PSQL}
	docker system prune -fa

db-mysql:
	docker rm -f ${DB_CONTIANER_MYSQL}

	docker run \
		--name ${DB_CONTIANER_MYSQL} \
		-d \
		--env-file ./.env \
		-p 3306:3306 \
		-v ${PWD}/initdb:/docker-entrypoint-initdb.d \
		${DB_IMG_MYSQL}

db-remove-mysql:
	docker stop ${DB_CONTIANER_MYSQL}
	docker rm -f ${DB_CONTIANER_MYSQL}
	docker rmi -f ${DB_IMG_MYSQL}
	docker system prune -fa