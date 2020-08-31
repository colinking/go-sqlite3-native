# Go SQLite3

This repo contains an implementation of a read-only SQLite3 client written in Go. It can be used as a drop-in, API-compatible, replacement for [`mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3).

The goal of this project is to provide an optimized Go-native implementation of the SQLite client for a specific subset of read-only SQLite behavior. It aims to parallel the high-level architecture of the real C-based SQLite3 client.

**It is strongly recommended that you do not use this library in production.**

## Compatibility

This library supports [GOOS](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) of `[darwin, linux]` and [GOARCH](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) of `[amd64]`.

This library supports queries of the following pattern:

```sql
SELECT <cols|*>
FROM <table|pragma_table_info(?)>
[WHERE <clause> [ AND|OR <clause>]?]
[ORDER BY <id>]
[LIMIT <n>]
```

It also supports the following SQLite features:

| Feature | Support |
| -- | ---- |
| WAL Journaling | ✅ Yes, except for checkpointing and recovery |
| Legacy (Rollback) Journaling | ❌ No |
| Pragmas | ❌ `pragma_table_info`, but no others |
| `WITHOUT ROWID` tables | ❌ |
| DB Type: File | ✅ |
| DB Type: `:memory:` | ❌ (PRs welcome) |
| DB Type: Temporary | ❌ |
| `ATTACH/DETACH` | ❌ |
| Handlers | ❌ |
| ... | ... |

Using the client for any unsupported features will lead to undefined behavior.

## Architecture

For a high-level overview of the real SQLite3 architecture, see the [technical design docs](https://www.sqlite.org/arch.html). This implementation was also inspired by [SQLite Database System Design and Implementation (2015)](https://books.google.com/books?id=OEJ1CQAAQBAJ).

## Query Patterns

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

Here is an example bytecode plan:

```sql
sqlite> explain select * from flagon2___gates where family=CAST("colin" AS BLOB) and name=CAST("gate" as BLOB);
addr  opcode         p1    p2    p3    p4             p5  comment
----  -------------  ----  ----  ----  -------------  --  -------------
0     Init           0     28    0                    00  Start at 28
1     OpenRead       0     1563  0     14             00  root=1563 iDb=0; flagon2___gates
2     OpenRead       1     1564  0     k(3,,,)        02  root=1564 iDb=0; sqlite_autoindex_flagon2___gates_1
3     String8        0     1     0     colin          00  r[1]='colin'
4     Cast           1     65    0                    00  affinity(r[1])
5     IsNull         1     27    0                    00  if r[1]==NULL goto 27
6     String8        0     2     0     gate           00  r[2]='gate'
7     Cast           2     65    0                    00  affinity(r[2])
8     IsNull         2     27    0                    00  if r[2]==NULL goto 27
9     SeekGE         1     27    1     2              00  key=r[1..2]
10    IdxGT          1     27    1     2              00  key=r[1..2]
11    DeferredSeek   1     0     0                    00  Move 0 to 1.rowid if needed
12    Column         1     0     3                    00  r[3]=flagon2___gates.family
13    Column         1     1     4                    00  r[4]=flagon2___gates.name
14    Column         0     2     5                    00  r[5]=flagon2___gates.description
15    Column         0     3     6                    00  r[6]=flagon2___gates.id_type
16    Column         0     4     7                    00  r[7]=flagon2___gates.tier_list_id
17    Column         0     5     8                    00  r[8]=flagon2___gates.rollout
18    Column         0     6     9                    00  r[9]=flagon2___gates.salt
19    Column         0     7     10                   00  r[10]=flagon2___gates.open
20    Column         0     8     11                   00  r[11]=flagon2___gates.archived
21    Column         0     9     12                   00  r[12]=flagon2___gates.archived_at
22    Column         0     10    13                   00  r[13]=flagon2___gates.user_id
23    Column         0     11    14                   00  r[14]=flagon2___gates.user_type
24    Column         0     12    15                   00  r[15]=flagon2___gates.created_at
25    Column         0     13    16                   00  r[16]=flagon2___gates.updated_at
26    ResultRow      3     14    0                    00  output=r[3..16]
27    Halt           0     0     0                    00
28    Transaction    0     0     90934  0              01  usesStmtJournal=0
29    Goto           0     1     0                    00
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

Here's an example bytecode plan:

```sql
sqlite> explain select * from flagon2___gates where family=CAST("colin" AS BLOB);
addr  opcode         p1    p2    p3    p4             p5  comment
----  -------------  ----  ----  ----  -------------  --  -------------
0     Init           0     26    0                    00  Start at 26
1     OpenRead       0     1563  0     14             00  root=1563 iDb=0; flagon2___gates
2     OpenRead       1     1564  0     k(3,,,)        02  root=1564 iDb=0; sqlite_autoindex_flagon2___gates_1
3     String8        0     1     0     colin          00  r[1]='colin'
4     Cast           1     65    0                    00  affinity(r[1])
5     IsNull         1     25    0                    00  if r[1]==NULL goto 25
6     SeekGE         1     25    1     1              00  key=r[1]
7       IdxGT          1     25    1     1              00  key=r[1]
8       DeferredSeek   1     0     0                    00  Move 0 to 1.rowid if needed
9       Column         1     0     2                    00  r[2]=flagon2___gates.family
10      Column         1     1     3                    00  r[3]=flagon2___gates.name
11      Column         0     2     4                    00  r[4]=flagon2___gates.description
12      Column         0     3     5                    00  r[5]=flagon2___gates.id_type
13      Column         0     4     6                    00  r[6]=flagon2___gates.tier_list_id
14      Column         0     5     7                    00  r[7]=flagon2___gates.rollout
15      Column         0     6     8                    00  r[8]=flagon2___gates.salt
16      Column         0     7     9                    00  r[9]=flagon2___gates.open
17      Column         0     8     10                   00  r[10]=flagon2___gates.archived
18      Column         0     9     11                   00  r[11]=flagon2___gates.archived_at
19      Column         0     10    12                   00  r[12]=flagon2___gates.user_id
20      Column         0     11    13                   00  r[13]=flagon2___gates.user_type
21      Column         0     12    14                   00  r[14]=flagon2___gates.created_at
22      Column         0     13    15                   00  r[15]=flagon2___gates.updated_at
23      ResultRow      2     14    0                    00  output=r[2..15]
24    Next           1     7     0                    00
25    Halt           0     0     0                    00
26    Transaction    0     0     90934  0              01  usesStmtJournal=0
27    Goto           0     1     0                    00
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
