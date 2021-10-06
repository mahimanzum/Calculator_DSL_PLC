# Part 1A Calculator_DSL_PLC
A sample calculator Module with Go lang with own lexar, grammar parser


My Grammar
Learnings while writing this grammar:
The more deeper the grammar rule the higher presedance it has over other operations
if You want an operation right associative you make it right recursion
Exponentiation function was the most informative , usefull to implement
Here is the Grammar. <br >

expr -> term [ ('+' | '-') term ]* <br />
term -> factor [ ('*' | '/') factor ]* <br />
factor -> base [ '^' exponent ]* <br />
base -> number| '(' expr ')' <br />
exponent -> base | [ '^' exponent ]* <br />

First My grammar was not right associative. And there were bugs. Then I changed the recursion from left recursion to right recursion and introduced the peak function that LR(1) parsers implement to see one token forward that solved the complexities.

Main thing I learned while doing this project is <br />
1. Define the Grammar and think it through (put 80% of the time here as it is the core of this) <br />
2. Once you are confident with the grammar just convert it to code with no bugs. Don't think about grammar correctness while coding, that part will come again while testing. Coding and thinking about the correctness will only introduce bugs in the code.

About Go lang:

Variable declaration
fixed and dynamic Array declaration
Function, Conditional, Map
These are enough to learn any languages and start with in a heart bit.

# Part2A
Very simple implementation of json parsing with python.
# Part2B
