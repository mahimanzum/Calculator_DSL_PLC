# Calculator_DSL_PLC
A sample calculator Module with Go lang with own lexar, grammar parser


My Grammar

expr -> term [ ('+' | '-') term ]*
term -> factor [ ('*' | '/') factor ]*
factor -> base [ '^' exponent ]*
base -> number| '(' expr ')'
exponent -> number| '(' expr ')'
