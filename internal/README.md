# File Structure

In top-down order, from what handles processing a SQL query to what performs low-level byte operations on the underlying DB file:

- [parser](./parser): implements the SQLite tokenizer and parser modules to process a SQL string into parse trees
- [compiler](./compiler): TODO: implements the SQLite compiler module to produce bytecode programs from parse trees
- [vm](./vm): implements the SQLite vm module to execute a bytecode program and produce results
- [tree](./tree): implements the SQLite tree module to traverse B and B+ trees
- [pager](./pager): implements the SQLite pager module to read pages from a DB file with ACID semantics
- [os](./os): TODO: implements the SQLite os module to offer an abstraction layer on top of OS syscalls
