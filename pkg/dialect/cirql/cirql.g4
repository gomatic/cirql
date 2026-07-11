// cirql's JSON pipeline grammar (ANTLR, compiled to src/grammar/cirql). The
// transform sublanguage (map/filter/reduce/sort/flatMap/limit/uniq + expressions
// + builtins) is executed by cirql-core; source stages (query/http/file/stdin)
// parse here and are executed by the sources sub-project. Numbers are unsigned
// in the lexer; negation is the unary-minus expression alt, so subtraction
// (`a - 5`) tokenizes correctly.
grammar cirql;

pipeline      : stage (PIPE stage)* EOF ;

stage         : sourceStage | transformStage ;

sourceStage   : queryStage | httpStage | fileStage | stdinStage ;
queryStage    : QUERY queryBody | queryBody ;
queryBody     : LBRACE selectionSet RBRACE ;
selectionSet  : field+ ;
field         : NAME (LPAREN arguments RPAREN)? (LBRACE selectionSet RBRACE)? ;
arguments     : argument (COMMA argument)* ;
argument      : NAME COLON argValue ;
httpStage     : HTTP (STRING | variable) ;
fileStage     : FILE STRING ;
stdinStage    : STDIN ;

transformStage: mapStage | filterStage | reduceStage
              | sortStage | flatMapStage | limitStage | uniqStage ;
mapStage      : MAP LBRACE mapping (COMMA mapping)* RBRACE ;
flatMapStage  : FLATMAP LBRACE mapping (COMMA mapping)* RBRACE ;
mapping       : fieldName COLON expr ;
filterStage   : FILTER expr ;
reduceStage   : REDUCE reduceOp (LPAREN expr RPAREN)? ;
reduceOp      : COUNT | SUM | MIN | MAX | AVG | FIRST | LAST | GROUP_BY | COLLECT ;
sortStage     : SORT expr (ASC | DESC)? ;
limitStage    : LIMIT INT ;
uniqStage     : UNIQ expr? ;

expr          : (NOT | MINUS) expr                  # UnaryExpr
              | expr (STAR | SLASH | PERCENT) expr  # MulExpr
              | expr (PLUS | MINUS) expr            # AddExpr
              | expr (EQ | NE | GE | LE | GT | LT) expr  # CmpExpr
              | expr AND expr                       # AndExpr
              | expr OR expr                        # OrExpr
              | LPAREN expr RPAREN                  # ParenExpr
              | funcCall                            # CallExpr
              | fieldAccess                         # FieldExpr
              | variable                            # VarExpr
              | literal                             # LitExpr
              ;

// The first segment follows the leading dot directly; every continuation
// segment carries its own dot (or is []). Requiring the continuation dot keeps
// a trailing keyword — e.g. the `desc` in `sort .name desc` — from being
// absorbed into the path now that keywords are valid field names.
fieldAccess   : DOT (fieldName | LBRACK RBRACK) pathSeg* | DOT ;
pathSeg       : DOT fieldName | LBRACK RBRACK ;
funcCall      : NAME LPAREN (expr (COMMA expr)*)? RPAREN ;
variable      : DOLLAR NAME ;

// fieldName is a JSON object key: an identifier or any keyword. Keywords are
// only reserved in their own stage/operator positions; as a field name (a path
// segment or a map output key) every keyword is a plain identifier, so common
// JSON fields like .count, .type, and map { first: ... } address correctly.
fieldName     : NAME | QUERY | HTTP | FILE | STDIN | MAP | FLATMAP | FILTER
              | REDUCE | SORT | LIMIT | UNIQ | COUNT | SUM | MIN | MAX | AVG
              | FIRST | LAST | GROUP_BY | COLLECT | ASC | DESC
              | TRUE | FALSE | NULL ;

argValue      : variable | literal | listLit | objectLit ;
objectLit     : LBRACE (NAME COLON argValue (COMMA NAME COLON argValue)*)? RBRACE ;
listLit       : LBRACK (argValue (COMMA argValue)*)? RBRACK ;
literal       : STRING | FLOAT | INT | TRUE | FALSE | NULL ;

QUERY:'query'; HTTP:'http'; FILE:'file'; STDIN:'stdin';
MAP:'map'; FLATMAP:'flatMap'; FILTER:'filter'; REDUCE:'reduce';
SORT:'sort'; LIMIT:'limit'; UNIQ:'uniq';
COUNT:'count'; SUM:'sum'; MIN:'min'; MAX:'max'; AVG:'avg';
FIRST:'first'; LAST:'last'; GROUP_BY:'group_by'; COLLECT:'collect';
ASC:'asc'; DESC:'desc'; TRUE:'true'; FALSE:'false'; NULL:'null';

PIPE:'|'; DOT:'.'; COMMA:','; COLON:':'; DOLLAR:'$';
LPAREN:'('; RPAREN:')'; LBRACE:'{'; RBRACE:'}'; LBRACK:'['; RBRACK:']';
OR:'||'; AND:'&&'; NOT:'!'; EQ:'=='; NE:'!='; GE:'>='; LE:'<='; GT:'>'; LT:'<';
PLUS:'+'; MINUS:'-'; STAR:'*'; SLASH:'/'; PERCENT:'%';

NAME   : [a-zA-Z] [a-zA-Z0-9_]* ;
FLOAT  : [0-9]+ '.' [0-9]+ ;
INT    : [0-9]+ ;
STRING : '"' ( ~["\\] | '\\' . )* '"' ;
WS     : [ \t\r\n]+ -> skip ;
