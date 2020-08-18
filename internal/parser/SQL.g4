/*
 * This file declares the grammar supported by this library.
 *
 * This file can be regenerated with `go generate ./...` which will
 * use goyacc to regenerate the Go bindings for parsing SQL statements
 * according to this grammar.
 * 
 * You will need to install the antlr CLI if you are going to make changes
 * to the grammar in this file. To install on macOS, run: `brew install antlr`
 *
 * For more information on yacc syntax, see: http://dinosaur.compilertools.net/yacc
 */

grammar SQL;

/* Tokens */

// Keywords
Select: 'SELECT';
From: 'FROM';
Where: 'WHERE';
Order: 'ORDER';
By: 'BY';
Asc: 'ASC';
Desc: 'DESC';
Limit: 'LIMIT';
Star: '*';
Placeholder: '?';

// Expressions
And: 'AND';
Equal: '=';
Greater: '>';

// Literals
Number: [0-9]+;
Letter: [a-zA-Z_]+;
Identifier: Letter (Letter | Number)*;

// Syntax
Comma: ',';
LParen: '(';
RParen: ')';
Semicolon: ';';

// Pragmas
PragmaTableInfo: 'pragma_table_info';

// Ignore whitespace
WHITESPACE: [ \r\n\t]+ -> skip;

/* Rules */

start
  : expression (EOF | Semicolon) 
  ;

expression
  : selectExpression
  ;

selectExpression
  : Select args From table where? orderBy? limit?
  ;

table
  : Identifier
  | PragmaTableInfo LParen Placeholder RParen
  ;

args
  : Star
  | columns
  ;

columns
  : Identifier (Comma columns)?
  ;

where
  : Where clause (And clause)*
  ;

clause
  : Identifier Equal ( Number | Placeholder )
  | Identifier Greater ( Number | Placeholder )
  ;

orderBy
  : Order By Identifier (Asc | Desc)?
  ;

limit
  : Limit Number
  ;
