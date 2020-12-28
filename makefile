sql:
	docker-compose exec db psql -U app -d weight_tracker
build:
	docker-compose exec app go build
	docker-compose restart app
