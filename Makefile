
#Migrations

.PHONY: db-init
db-init:
	docker run --rm \
    	-v $(realpath ./db/migrations):/migrations \
    	migrate/migrate:v4.16.2 \
        	create \
        	-dir /migrations \
        	-ext .sql \
        	-seq -digits 3 \
        	init


.PHONY: pg
pg: 
	docker run --rm \
		--name=customerdb_v1 \
		-v $(abspath ./db/init/):/docker-entrypoint-initdb.d \
		-e POSTGRES_PASSWORD="postgres" \
		-d \
		-p 5433:5432 \
		postgres:15.3


.PHONY: pg-stop
pg-stop:
	docker stop customerdb_v1

.PHONY: clean-data
clean-data:
	sudo rm -rf ./db/data/

.PHONY: pg-up
pg-up:
	docker run --rm \
    -v $(realpath ./db/migrations):/migrations \
    migrate/migrate:v4.16.2 \
        -path=/migrations \
        -database postgres://customer:1@172.17.0.3:5432/customer?sslmode=disable \
        up