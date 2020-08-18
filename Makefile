
Q=@

build-dbs:
	$Q echo "sqlite3 version: $$(sqlite3 --version)"
	$Q mkdir -p tmp
	$Q rm tmp/empty.db
	$Q sqlite3 tmp/empty.db "VACUUM;"
	$Q rm tmp/simple.db
	$Q sqlite3 tmp/simple.db "CREATE table table1(column1 int); INSERT INTO table1 (column1) VALUES (123), (456);"
.PHONY: build-db

generate: gen
gen:
	$Q go generate ./...
.PHONY: gen generate
