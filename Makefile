createcountainer:
	docker run --name idompet -p 5432:5432 -e POSTGRES_PASSWORD=1234 -d postgres 

startcountainer:
	docker start idompet

stopcountainer:
	docker stop idompet

createdb:
	docker exec -it idompet createdb --username=postgres --owner=postgres idompet

dropdb:
	docker exec -it idompet dropdb --username=postgres idompet

.PHONY: createcountainer startcountainer stopcountainer createdb dropdb