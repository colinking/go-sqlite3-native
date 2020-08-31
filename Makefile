Q=@

all: lint test

build-dbs:
	$Q echo "sqlite3 version: $$(sqlite3 --version)"
	$Q mkdir -p tmp
	$Q rm -f  tmp/empty.db
	$Q sqlite3 tmp/empty.db "PRAGMA journal_mode=WAL;"
	$Q rm -f  tmp/simple.db
	$Q sqlite3 tmp/simple.db "PRAGMA journal_mode=WAL; CREATE table table1(column1 int); INSERT INTO table1 (column1) VALUES (123), (456);"
	$Q rm -f tmp/pk.db
	$Q sqlite3 tmp/pk.db "PRAGMA journal_mode=WAL; CREATE table table1(key int PRIMARY KEY, value TEXT); INSERT INTO table1 (key, value) VALUES (123, 'hello world'), (456, 'goodbye world');"
	$Q rm -f tmp/source.db
	$Q sqlite3 tmp/source.db "PRAGMA journal_mode=WAL; create table core___source_id_write_key_mapping (write_key BLOB, source_id BLOB); insert into core___source_id_write_key_mapping (write_key, source_id) VALUES (CAST('wk-123' AS BLOB), CAST('source-123' AS BLOB)); insert into core___source_id_write_key_mapping (write_key, source_id) VALUES (CAST('wk-456' AS BLOB), CAST('source-456' AS BLOB)); insert into core___source_id_write_key_mapping (write_key, source_id) VALUES (CAST('wk-789' AS BLOB), CAST('source-789' AS BLOB));"
.PHONY: build-db

# You'll need:
#  - go get -u -a golang.org/x/tools/cmd/stringer
#  - brew install antlr
generate: gen
gen:
	$Q go generate ./...
.PHONY: gen generate

tmp/stage.db:
	$Qaws-okta exec stage-admin -- aws s3 cp s3://segment-ctlstore-snapshots-stage/snapshot.db.gz - | gzip --decompress > ./tmp/stage.db

# To install golangci-lint:
#   brew install golangci/tap/golangci-lint
#   brew upgrade golangci/tap/golangci-lint
lint:
	$Q golangci-lint run ./...
.PHONY: lint

test:
	$Q go test -race -v -count=1 ./...
.PHONY: test

# `make diff QUERY="select * from bar" DB=stage.db` will print out a colored diff of that query
# run on DB using sqlite3 and with the Go client.
diff: QUERY=select * from core___source_id_write_key_mapping;
diff: DB=stage.db
diff:
	$Q go run ./cmd/main.go query tmp/${DB} "${QUERY}" > tmp/go.out
	$Q sqlite3 tmp/${DB} "${QUERY}" > tmp/c.out
	$Q colordiff tmp/c.out tmp/go.out
.PHONY: diff

# `make search QUERY=foobar DB=stage.db` will print out the byte offsets of QUERY within DB.
# Accepts raw bytes (if using \x notation per byte) or ASCII strings.
search: QUERY=vb3bcHdvxEGjneRfwmJNzJ
search: DB=stage.db
search:
	$Q docker run -it --rm -v "$$(pwd):/binwalk" rjocoleman/binwalk -R=${QUERY} /binwalk/tmp/${DB}
.PHONY: search

dump: OFFSET=0x5DEBCC4
dump: DB=stage.db
dump:
	$Q docker run -it --rm -v "$$(pwd):/binwalk" rjocoleman/binwalk --hexdump --offset=${OFFSET} --length=4096 /binwalk/tmp/${DB}
.PHONY: dump

like: QUERY=3bcHdvxE
like: DB=stage.db
like:
	$Q sqlite3 ./tmp/${DB} "select ROWID, * from core___source_id_write_key_mapping where write_key LIKE '%${QUERY}%' OR source_id LIKE '%${QUERY}%';"
.PHONY: like
