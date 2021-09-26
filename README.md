# Calculator_DSL_PLC
A sample calculator Module with Go lang with own lexar, grammar parser


My Grammar

  expr -> term [ ('+' | '-') term ]* <br />
  term -> factor [ ('*' | '/') factor ]* <br />
  factor -> base [ '^' exponent ]* <br />
  base -> number| '(' expr ')' <br />
  exponent -> number| '(' expr ')' <br />
