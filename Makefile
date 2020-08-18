
Q=@

build-dbs:
	$Q echo "sqlite3 version: $$(sqlite3 --version)"
	$Q mkdir -p tmp
	$Q sqlite3 tmp/empty.db "VACUUM;"
	$Q sqlite3 tmp/simple.db "CREATE table table1(column1 int); INSERT INTO table1 (column1) VALUES (123);"
.PHONY: build-db

generate:
	$Q go generate ./...
.PHONY: generate
