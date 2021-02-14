package config

var databaseYML = `
development:
  dialect: postgres 
  url: {{"{{"}} envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/{{.Name}}_development?sslmode=disable" {{"}}"}}

test:
  dialect: postgres
  url: {{"{{"}} envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/{{.Name}}_test?sslmode=disable" {{"}}"}}

production:
  pool: 20
  dialect: postgres
  url: {{"{{"}} envOr "DATABASE_URL" "postgres://postgres:postgres@127.0.0.1:5432/{{.Name}}_production?sslmode=disable" {{"}}"}}
`
