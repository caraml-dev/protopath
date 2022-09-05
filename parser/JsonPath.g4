grammar JsonPath;


jsonpath: root ;

root    :  root ('.' fieldAccess)+
        |  ROOT_IDENTIFIER fieldAccess
        ;

fieldAccess    : IDENTIFIER arrayQualifier
                | IDENTIFIER '.' arrayQualifier 
                | IDENTIFIER
                 ;

localFieldAccess      : LOCAL_IDENTIFIER fieldAccess
                        | localFieldAccess ('.' fieldAccess)+
                        ;

arrayQualifier : '[*]'                 # GetAll
                | '[' intSequence ']'  # GetByIndices
                | arrayRange           # GetByRange
                | '[?(' filterOperation ')]' # GetByQuery
                | '[@.length-' INT ']'      # GetByIndexBackward
                ;

arrayRange : '[:]'                     # AllRange
            | '[' INT+ COLON INT+ ']'   # GivenStartEndIdxRange
            | '[' COLON INT+ ']'        # GivenOnlyEndIdxRange
            | '[' INT+ COLON ']'        # GivenOnlyStartIdxRange
            ;

intSequence: INT (',' INT)* ;


queryFieldValue       : localFieldAccess        # QueryFieldValueLocalAccess
                        | root                  # QueryFieldValueRootAccess
                        | dbl                   # QueryFieldDoubleValue
                        | STRING                # QueryFieldStringValue
                        | bool_type             # QueryFieldBoolValue
                        ;

filterOperation : queryExpr ;

queryExpr : queryExpr (logical_op=('&&'| '||') queryExpr)+                                   # QueryWithLogicalOp                                                            
           | queryFieldValue                                                                  # QueryExistence
           | queryFieldValue comparator_op=('>'|'>='|'<'|'<='|'=='|'!=') queryFieldValue    # QueryWithComparator
           ;

dbl : INT+ '.' INT+
    | '.' INT+
    | INT+
    ;

bool_type : 'true' | 'false' ;

IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
STRING : '\'' IDENTIFIER '\'' ;
INT         : '0' | [1-9][0-9]* ;
WS  :   [ \t\n\r]+ -> skip ;
COLON: ':' ;
ROOT_IDENTIFIER: '$.' ;
LOCAL_IDENTIFIER: '@.' ;

GREATER: '>' ;
GREATER_EQ: '>=' ;
LESS: '<' ;
LESS_EQ: '<=' ;
EQ: '==' ;
NEQ: '!=' ;
AND: '&&';
OR: '||' ;