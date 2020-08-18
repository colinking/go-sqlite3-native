# File Structure

In top-down order, from what handles processing a SQL query to what performs low-level byte operations on the underlying DB file:

- [tokenizer](./tokenizer): implements the SQLite tokenizer module using goyacc
- [pager](./pager): implements the SQLite pager module to read pages from a DB file with ACID semantics
