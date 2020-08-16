# Go SQLite3

This repo contains an implementation of a read-only SQLite3 client written in Go. It can be used as a drop-in, API-compatible, replacement for [`mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3).

This implementation is designed for an extremely limited subset of SQLite3 use-cases, in order to limit the complexity of this client. Specifically:

1. Only `SELECT` operations are supported, specifically `SELECT` operations operating on a single table, optionally with a single `WHERE` clause. No write operations are supported, nor planned to be supported at any point in the future.
2. Only Ubuntu 16.04, Ubuntu 18.04 and macOS platforms are supported. Portability is not a goal of this project.
3. Most SQLite3 client configuration is not supported.
4. ...TBA

In exchange for these limitations, this client can be used in native Go applications without CGo.

This library is meant to be used in scenarios where you have separate reader and writer processes, where the reader processes are bottlenecked by CGo calls. This library would be used in the readers, while the writers continue to use `mattn/go-sqlite3` (or another full-featured SQLite3 library).
