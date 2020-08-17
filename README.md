# Go SQLite3

This repo contains an implementation of a read-only SQLite3 client written in Go. It can be used as a drop-in, API-compatible, replacement for [`mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3).

This implementation is designed for an extremely limited subset of SQLite3 use-cases, in order to limit the complexity of this client. Specifically:

1. Only `SELECT` operations are supported, specifically `SELECT` operations operating on a single table, optionally with a single `WHERE` clause. No write operations are supported, nor planned to be supported at any point in the future.
1. Only Ubuntu 16.04, Ubuntu 18.04 and macOS platforms are supported. Portability is not a goal of this project.
1. Most SQLite3 client configuration is not supported.
1. Only WAL journaling mode is supported. DBs in legacy journaling mode will produce errors when queried.
1. Since this client is read-only, it means that it cannot support journal recovery which means that it will not be able to respond to queries until a full-featured SQLite client performs the recovery procedure.
1. Temporary, in-memory, and multi-file DBs (see: `ATTACH` / `DETACH`) are no supported.
1. ...TBA

In exchange for these limitations, this client can be used in native Go applications without CGo.

This library is meant to be used in scenarios where you have separate reader and writer processes, where the reader processes are bottlenecked by CGo calls. This library would be used in the readers, while the writers continue to use `mattn/go-sqlite3` (or another full-featured SQLite3 library).

## Architecture

This implementation of this SQLite3 client is meant to parallel the real C-based implementation that modularizes the various components. For a high-level overriew, see the [SQLite architecture docs](https://www.sqlite.org/arch.html).

The implementation was inspired by [SQLite Database System Design and Implementation (2015)](https://books.google.com/books?id=OEJ1CQAAQBAJ).

## TODO

- [ ] Add barebones database/sql implementation
- [ ] Add logic for parsing SELECT SQL queries into bytecode
- [ ] Add tree module for reading page content into B+ trees
- [ ] Add VM module for basic scans

At this point, we have a workable POC that can perform queries e2e.

- [ ] Add support for reading indexes into B-trees
- [ ] Add support for using indexes in queries in the VM module
- [ ] Add WAL support
- [ ] Clean up the codebase + add stats
- [ ] Hook up a real suite of tests and benchmarks
- [ ] Gameday using all examples from the DML ledger, against mattn/go-sqlite3
- [ ] Gameday via Argus

## Query Patterns

Checkpointing by the readers should be disabled (`PRAGMA wal_autocheckpoint = 0`), so we do not need to support that. There are some other SQLite query patterns used by the reflectors or executive, but we don't need to support them since those patterns are not needed the readers.

We will need to support `immutable` mode, ...

In terms of URI patterns, we'll need to support:

- `file:$FILE_NAME?_journal_mode=wal&mode=$MODE`
  - `MODE` is one of [`ro`, `rwc`] though the latter is only used for testing (see below)
- `file:$FILE_NAME?immutable=true`

### GetRowByKey

`GetRowByKey` is the same as `GetRowsByKeyPrefix`, except that it includes all of the keys in the PK and therefore will only return 0 or 1 results. Therefore, `GetRowsByKeyPrefix` is a superset of `GetRowByKey`.

```sql
SELECT *
FROM $ldbTableName
WHERE
  $pkCol1 = ?
  AND $pkCol2 = ?
  -- ...
```

### GetRowsByKeyPrefix

```sql
SELECT *
FROM $ldbTableName
```

```sql
SELECT *
FROM $ldbTableName
WHERE
  $pkCol1 = ?
  AND $pkCol2 = ?
  -- ...
```

### GetLedgerLatency

```sql
SELECT timestamp
FROM _ldb_last_update
-- "ledger"
WHERE name=?
```

### FetchSeqFromDB

```sql
SELECT seq
FROM _ldb_seq
WHERE id = 1
```

### Ping

```sql
SELECT seq
FROM _ldb_seq
-- 1
WHERE id = ?
```

### getPrimaryKey

```sql
SELECT name, type
-- ldbTable parameter
FROM pragma_table_info(?)
WHERE pk > 0
ORDER BY pk ASC
```

```sql
SELECT * FROM $ldbTable LIMIT 1
```

### TestUtils (EnsureLdbInitialized, etc.)

The TestUtils require that we support creating tables and writing to them.

In this case, let's just re-implement the testing libraries to use an in-memory LDB of some kind.

Alternatively, we could implement an exclusive writer in this package that assumes it is the only reader/writer to the DB and where ACID compliance (f.e. WALing) isn't necessary. This would allow us to make a number of simplifying assumptions.

### Edge Cases

- PKs being dropped and re-created. Seems like there is some logic to handle execution errors and refresh the PK cache.
