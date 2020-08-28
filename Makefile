
Q=@

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
