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
SELECT: 'SELECT';
FROM: 'FROM';
WHERE: 'WHERE';
ORDER: 'ORDER';
BY: 'BY';
ASC: 'ASC';
DESC: 'DESC';
LIMIT: 'LIMIT';

// Expressions
AND: 'AND';
EQUAL: '=';
GREATER: '>';

// Identifiers
IDENTIFIER: 'todo'; // TODO
PLACEHOLDER: '?';

// Literals
INT_LITERAL: '0'; // TODO
STRING_LITERAL: 'todo'; // TODO

// Syntax
COMMA: ',';
LPAREN: '(';
RPAREN: ')';

/* Rules */

start: expression EOF;

expression
  : SELECT
  ;
