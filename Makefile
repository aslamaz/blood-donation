dep-up:
	docker compose up -d

dep-down:
	docker compose down

run:
	go run .

migrate-up:
	cd migrations &&\
	export GOOSE_DRIVER=mysql &&\
	export GOOSE_DBSTRING=root:root@/blood_donation?parseTime=true &&\
	goose up

migrate-down:
	cd migrations &&\
	export GOOSE_DRIVER=mysql &&\
	export GOOSE_DBSTRING=root:root@/blood_donation?parseTime=true &&\
	goose down

migrate-new:
	cd migrations &&\
	goose create t sql 

install:
	go install github.com/pressly/goose/v3/cmd/goose@latest