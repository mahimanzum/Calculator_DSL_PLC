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

# Part2A and Part2B

Very simple implementation of json parsing with python. And the stock trade DSL was also straight forward. So I just had to follow instructions. For both of these I implemented different functions to create dsl and sql files. First 2 functions (make_dsl() and make_sql()) are for part2A. <br >

For part2B make_dsl_with_delete() and make_sql_with_delete() was implemented which were commented out for the previous part where I found the problem description was a bit ambigious. My grammar for the updated DSL was <br >

<stock_trade_requests> →  [delete]? ‘(' <trade> {‘,’ <trade>} ‘) for account' <acct_ident>’.’ <br >
<trade> →  <number> <stock_symbol> ‘shares’ (‘buy at max' | ‘sell at min') <number> <br >
<number> →  [1-9] {[0-9]} <br >
<stock_symbol> → 'AAPL'|'HP'|'IBM'|'AMZN'|'MSFT'|'GOOGL'|'INTC'|'CSCO'|'ORCL'|'QCOM' <br >
<acct_ident> →  ‘“‘alpha_char { alpha_char | digit | ’_’} ‘“‘ <br >
Note:  ‘“‘ is a “ surrounded by ‘ <br >

where i just added the "delete?" string which means some stock trades can also be deleted with this command. So my updated json was like this. <br >  

  {
	"user id" : "Hokie123",
	"buy" : [
        {"stock symbol" : "IBM", "shares" : 100, "at max" : 45},{"stock symbol" : "GOOGL", "shares" : 50, "at max" : 60}, {"stock symbol" : "AMZN", "shares" : 120, "at max" : 70}
],  "sell" : [{"stock symbol" : "ORCL", "shares" : 30, "at min" : 25},{"stock symbol" : "GOOGL", "shares" : 20, "at min" : 40} ]
, "delete" :{"sell":[{"stock symbol" : "ORCL", "shares" : 30, "at min" : 25}],"buy":[{"stock symbol" : "IBM", "shares" : 100, "at max" : 45} ]}

}
  
  
  
  
  




