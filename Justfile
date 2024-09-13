generate-sqlc:
    sqlc generate

create-migration NAME:
    migrate create -ext sql -dir db/migrations {{NAME}}

migrate-up DSN:
    migrate -path db/migrations -database "{{DSN}}?sslmode=disable" -verbose up

migrate-down DSN:
    migrate -path db/migrations -database "{{DSN}}?sslmode=disable" -verbose down