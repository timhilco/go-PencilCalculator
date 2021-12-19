lexer grammar HilcoPencilLexer;


// *** Keywords 

// Control
IF : 'if' ;
THEN :  'then' ; 
ELSE :  'else';
CASE : 'case';
CASE_IS : 'is';
CASE_END : 'endcase';
SWITCH: 'switch';
SWITCH_END: 'endswitch';

// Logical 
AND: 'AND';
OR: 'OR';
NOT: 'NOT';


fragment DIGIT: '0' ..'9';
fragment LOWERCASELETTERS: 'a' ..'z';
fragment UPPERCASELETTERS: 'A' ..'Z';

INT: (DIGIT)+;

ID:
	LOWERCASELETTERS (
		LOWERCASELETTERS
		| UPPERCASELETTERS
		| DIGIT
		| '_'
		| COLON
		| '\\'
	)*;

STRING: '\'' ( ESC_SEQ | ~('\\' | '\''))* '\'';

fragment EXPONENT: ('e' | 'E') ('+' | '-')? ('0' ..'9')+;

fragment HEX_DIGIT: ('0' ..'9' | 'a' ..'f' | 'A' ..'F');

fragment ESC_SEQ:
	'\\' ('b' | 't' | 'n' | 'f' | 'r' | '"' | '\'' | '\\')
	| UNICODE_ESC
	| OCTAL_ESC;

fragment OCTAL_ESC:
	'\\' ('0' ..'3') ('0' ..'7') ('0' ..'7')
	| '\\' ('0' ..'7') ('0' ..'7')
	| '\\' ('0' ..'7');

fragment UNICODE_ESC:
	'\\' 'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT;

FLOAT: (DIGIT)+ '.' (DIGIT)*;

DATE:
	CURLYLBRACKET DIGIT DIGIT ('-' | '/') (
		(DIGIT DIGIT)
		| (
			(LOWERCASELETTERS | UPPERCASELETTERS) (
				LOWERCASELETTERS
				| UPPERCASELETTERS
			) (LOWERCASELETTERS | UPPERCASELETTERS)
		)
	) ('-' | '/') DIGIT DIGIT (DIGIT DIGIT)? CURLYRBRACKET;

ATFUNCTION:
	'@' (LOWERCASELETTERS | UPPERCASELETTERS) (
		LOWERCASELETTERS
		| UPPERCASELETTERS
		| DIGIT
	)*;
CLASSNAME:
	UPPERCASELETTERS (
		LOWERCASELETTERS
		| UPPERCASELETTERS
		| DIGIT
		| '_'
	)*;
PERCENT: (INT | FLOAT) '%';
//COMMENT : '\"' (~'\"')* '\"' {skip ();};

NEWLINE: '\r'? '\n' -> skip;

WS: (
		' '
		| '\t'
		| '\f'
		// handle newlines
		| (
			'\r\n' // DOS/Windows
		)
		| '\u0003'
	) -> skip;
// *************************************** Operators ***************************************

DOT: '.';
// BECOMES    : ':='  ;
COLON: ':';
SEMI: ';';
COMMA: ',';
EQUALS: '=';
LBRACKET: '[';
RBRACKET: ']';
CURLYLBRACKET: '{';
CURLYRBRACKET: '}';
DOTDOT: '..';
LPAREN: '(';
RPAREN: ')';
NOT_EQUALS: '~=';
LT: '<';
LTE: '<=';
GT: '>';
GTE: '>=';
PLUS: '+';
MINUS: '-';
TIMES: '*';
DIV: '/';
EXPONENTIATE: '^';
