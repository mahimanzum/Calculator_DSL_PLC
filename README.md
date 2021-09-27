# Calculator_DSL_PLC
A sample calculator Module with Go lang with own lexar, grammar parser


My Grammar
Learning while writing this grammar:
The more deeper the grammar rule the higher presedance it has over other operations
You want an operation right associative you make it right recursion

Exponentiation function was the most informative , usefull to implement

expr -> term [ ('+' | '-') term ]* <br />
term -> factor [ ('*' | '/') factor ]* <br />
factor -> base [ '^' exponent ]* <br />
base -> number| '(' expr ')' <br />
exponent -> base | [ '^' exponent ]* <br />

First My grammar was not right associative. And there were bugs. Then i introduced the peak function that solved the complexities.

Main thing I learned while doing this project is <br />
1. Define the Grammar and think it through <br />
2. Once you are confident with the grammar just convert it to code with no bugs. Don't think about grammar correctness while coding that part will come again while testing. Coding and thinking about the correctness will only introduce bugs in the code.