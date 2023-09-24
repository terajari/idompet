SOURCE=postgresql://postgres:1234@localhost:5432/idompet?sslmode=disable

createcontainer:
	docker run --name idompet -p 5432:5432 -e POSTGRES_PASSWORD=1234 -d postgres 

startcontainer:
	docker start idompet

stopcontainer:
	docker stop idompet

createdb:
	docker exec -it idompet createdb --username=postgres --owner=postgres idompet

dropdb:
	docker exec -it idompet dropdb --username=postgres idompet

newmigrate:
	migrate create -ext sql -dir database/migration -seq init-schema

migrateup:
	migrate -path database/migration -database "${SOURCE}" -verbose up

migratedown:
	migrate -path database/migration -database "${SOURCE}" -verbose down

sqlc:
	sqlc generate

.PHONY: createcountainer startcountainer stopcountainer createdb dropdb newmigrate migrateup migratedown sqlc