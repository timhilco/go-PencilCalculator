grammar HilcoPencilParser;
import HilcoPencilLexer;
prog: statement;
statement:
	logical
	| ifStatement
	| caseStatement
	| switchStatement;
caseStatement: CASE statement CASE_IS caseList CASE_END;

caseList: caseItem (caseItem)*;
caseItem: base '::' statement SEMI;
switchStatement: SWITCH caseList SWITCH_END ;
ifStatement:
	IF statement THEN statement ( (ELSE ) statement | ());
logical:
	relationalExpression ((AND | OR) relationalExpression)*
	| NOT relationalExpression;
relationalExpression:
	addingExpression (
		(EQUALS | NOT_EQUALS | GT | GTE | LT | LTE) addingExpression
	)*;
addingExpression:
	multiplyingExpression ((PLUS | MINUS) multiplyingExpression)*;

multiplyingExpression: factor ((TIMES | DIV) factor)*;

factor: MINUS factor | base;

base:
	LPAREN statement RPAREN
	| name
	| PERCENT
	| INT
	| FLOAT
	| STRING
	| DATE
	| atFunction
	| dataAccessor
	| specialKeyword;

name: ID | CLASSNAME COLON ID;

atFunction: ATFUNCTION LPAREN argList RPAREN;

dataAccessor: CLASSNAME (fieldName)+;
fieldName:
	DOT (ID | CLASSNAME)
	| DOT ( ID | CLASSNAME) LPAREN argList RPAREN;

argList: statement (COMMA ( statement |))* |;
specialKeyword: 'today' | 'true' | 'false' | 'nil' | 'default';
