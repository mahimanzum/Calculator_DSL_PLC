package main

import (
	"fmt"
	"math"
	"os"
)

type token struct {
	name     string
	value    int
	index    int
	start_id int
	end_id   int
}

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var error_message = map[string]string{
	"lexical_error":  "Quid dicis? You offend Caesar with your sloppy lexical habits!",
	"syntax_error":   "Quid dicis? True Romans would not understand your syntax!",
	"zero_error":     "Quid dicis? Arab merchants haven't left for India yet!",
	"negative_error": "Quid dicis? Caesar demands positive thoughts!",
}
var valids = map[string]bool{
	"IV": true,
	"IX": true,
	"XL": true,
	"XC": true,
	"CD": true,
	"CM": true,
}

var token_table = map[string]string{
	"times":  "times_token",
	"plus":   "plus_token",
	"power":  "power_token",
	"divide": "divide_token",
	"minus":  "minus_token",
	"(":      "left_bracket",
	"[":      "left_bracket",
	"{":      "left_bracket",
	")":      "right_bracket",
	"}":      "right_bracket",
	"]":      "right_bracket",
}

func check_valid(s string) bool {
	//valids := [6]string{"IV", "IX", "XL", "XC", "CD", "CM"}
	//fmt.Println(s)
	if len(s) == 0 {
		return false
	}
	for id, val := range s {
		if id == 0 {
			continue
		} else if _, ok := roman[string(val)]; !ok {
			return false
		} else if roman[string(s[id])] <= roman[string(s[id-1])] {
			continue
		} else if _, ok := valids[s[id-1:id+1]]; ok {
			continue
		} else {
			return false
		}
	}
	//fmt.Println("true:", s)
	return true
}
func write_message(index int, s string, length int) {
	fmt.Println(raw_code)
	for i := 0; i < index-length; i += 1 {
		fmt.Print(" ")
	}
	fmt.Println("^ ")
	fmt.Println(error_message[s])
	os.Exit(1)
}

func lexar(code string) []token {
	var a []token
	prev := ""
	id := 1
	//fmt.Println("comes")
	for string_id, val := range code {
		//fmt.Println(prev, string_id, val)
		//fmt.Println(a)
		if val == '(' || val == '{' || val == '[' {
			a = append(a, token{name: "left_bracket", value: 0, index: id, start_id: string_id, end_id: string_id + 1})
			id += 1
		} else if val == ')' || val == '}' || val == ']' {
			if len(prev) > 0 && check_valid(prev) {
				a = append(a, token{name: "Number", value: romanToInt(prev), index: id, start_id: string_id - len(prev), end_id: string_id})
				id += 1
				prev = ""
			}
			a = append(a, token{name: "right_bracket", value: 0, index: id, start_id: string_id, end_id: string_id + 1})
			id += 1
		} else if val == ' ' {
			//fmt.Println("comes in 109")
			if prev == "times" || prev == "plus" || prev == "power" || prev == "divide" || prev == "minus" {
				a = append(a, token{name: token_table[prev], value: 0, index: id, start_id: string_id - len(prev), end_id: string_id})
				id += 1
				prev = ""
			} else {
				if len(prev) > 0 && check_valid(prev) {
					a = append(a, token{name: "Number", value: romanToInt(prev), index: id, start_id: string_id - len(prev), end_id: string_id})
					id += 1
					prev = ""
					//fmt.Println("comes in 118")
				} else if len(prev) > 0 {
					//fmt.Println("print from error", prev)
					write_message(string_id, "lexical_error", len(prev))
					return []token{{name: "error", value: 0}}
				}
			}
		} else if val == '$' && len(prev) == 0 {
			a = append(a, token{name: "end_token", value: 0, index: id})
			id += 1
			return a
		} else {
			//fmt.Println("comes:", string(val))
			prev = prev + string(val)
		}
	}
	return a

}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]
		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}

//if val, ok := dict["foo"]; ok
var idx int = 0
var universal_lexed []token
var current_token token
var raw_code string

func lex() token {
	if universal_lexed[idx].name == "end_token" {
		return token{name: "end_token", value: 0, start_id: len(raw_code) - 1, end_id: len(raw_code)}
	} else {

		next_token := universal_lexed[idx]
		//fmt.Println("############## Consumed ", idx, next_token)
		current_token = next_token
		idx = idx + 1
		return next_token
	}

}
func clear() {
	idx = 0
	universal_lexed = universal_lexed[:0]

}
func peak() token {
	if idx < len(universal_lexed) {
		return universal_lexed[idx]
	} else {
		fmt.Println("sorry final elemet reached")
		return universal_lexed[idx-1]
	}
}
func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

//expr -> term [ ('+' | '-') term ]*
func parse_expr() int {
	//fmt.Println("came in parse expression")
	term := parse_term()
	for {
		next_token := peak()
		if next_token.name == "plus_token" {
			lex()
			term = term + parse_term()
		} else if next_token.name == "minus_token" {
			temp := next_token
			lex()
			term = term - parse_term()
			if term < 0 {
				write_message(temp.start_id, "negative_error", 0)
			} else if term == 0 {
				write_message(temp.start_id, "zero_error", 0)
			}
		} else {
			//fmt.Println("line 185 , ", current_token)
			return term
		}
	}
}

//term -> factor [ ('*' | '/') factor ]*
func parse_term() int {
	//fmt.Println("came in parse term")
	factor := parse_factor()
	for {
		next_token := peak()
		if next_token.name == "times_token" {
			lex()
			factor = factor * parse_factor()
		} else if next_token.name == "divide_token" {
			lex()
			val := parse_factor()
			fmt.Println("comes here val = ", val)
			factor = factor / val
		} else {
			return factor
		}
	}
}

//factor -> base [ '^' exponent ]*
func parse_factor() int {
	//fmt.Println("came in parse factor")
	base := parse_base()
	var exp int
	exp = 1
	for {
		next_token := peak()
		if next_token.name == "power_token" {
			lex()
			exp = parse_exponent()
			//fmt.Println("comes 232 only base, exp ", base, exp)
			return int(math.Pow(float64(base), float64(exp)))
		} else {
			//fmt.Println("comes 234 only base ", base)
			return base
		}
	}
}

//base -> number| '(' expr ')'
func parse_base() int {
	//fmt.Println("came in parse base")
	next_token := lex()
	var value int
	if next_token.name == "left_bracket" {
		value = parse_expr()
		next_token = lex()
		if next_token.name != "right_bracket" {
			//fmt.Println("error in parsing base", next_token, current_token)
			write_message(next_token.end_id, "syntax_error", 0)
			//os.Exit(0)
		}
	} else {
		value = parse_number()
	}
	return value
}

//exponent -> base| [ '^' exponent ]*
func parse_exponent() int {
	//fmt.Println("came in parse exponent")
	base := parse_base()
	var exp int
	exp = 1.00
	for {
		next_token := peak()
		if next_token.name == "power_token" {
			lex()
			exp = parse_exponent()
			//fmt.Println("comes 270 base, exp = ", base, exp)
			return int(math.Pow(float64(base), float64(exp)))
		} else {
			//fmt.Println("comes 273 base ", base)
			return base
		}
	}
}

func parse_number() int {
	//fmt.Println("came in parse number")
	//next_token := lex()

	//fmt.Println("name = ", current_token.name, "value calculated", current_token.value)
	return current_token.value
}
func parse_code(code string) string {
	universal_lexed = lexar(code)
	val := Roman(parse_expr())
	clear()
	return val
}
func main() {

	//universal_lexed = lexar("[V minus {VI minus (III minus {II minus I]}])$") //1
	//universal_lexed = lexar("{MCMXCVIII divide III divide VI minus XI) divide X $") //X
	//lexical error
	//raw_code = "III plu {IV times II] power II"
	//raw_code = "I plus III minus VX times VI"

	//zero error
	//raw_code = "II times (I plus II minus III)"

	//negative error
	//raw_code = "II plus I times III minus VI"

	//syntax error
	raw_code = "III plus {IV times II power II"
	universal_lexed = lexar(raw_code + " $")

	//fmt.Println(universal_lexed)
	fmt.Println(Roman(parse_expr()))
	fmt.Print(lex())
	clear()

	//write_message(1, "exit from code called")
	//universal_lexed = lexar("{MCMXCVIII divide III divide VI minus XI) divide X power II $")
	//fmt.Println("value is ", Roman(parse_expr()))

	os.Exit(0)
}

/*
Grammar

expr -> term [ ('+' | '-') term ]*
term -> factor [ ('*' | '/') factor ]*
factor -> base [ '^' exponent ]*
base -> number| '(' expr ')'
exponent -> base | [ '^' exponent ]*

def parse_expr():
  term = parse_term()
  while 1:
    if match('+'):
      term = term + parse_term()
    elif match('-'):
      term = term - parse_term()
    else: return term

def parse_term():
  factor = parse_factor()
  while 1:
    if match('*'):
      factor = factor * parse_factor()
    elif match('/'):
      factor = factor / parse_factor()
    else: return factor

def parse_factor():
  if match('-'):
    negate = -1
  else: negate = 1
  if peek_digit():
    return negate * parse_number()
  if match('('):
    expr = parse_expr()
    if not match(')'): error...
    return negate * expr
  error...

def parse_number():
  num = 0
  while peek_digit():
    num = num * 10 + read_digit()
  return num
*/
/*
//fmt.Println(check_valid("XI"))
//fmt.Println(lexar("XI plus (X plus X)$"))
//universal_lexed = lexar("{MCMXCVIII divide III divide VI}$") //CXI
//universal_lexed = lexar("{MCMXCVIII divide III divide VI minus XI) divide X power II $") // I
//universal_lexed = lexar("III plus {IV times II] power II $") //LXVII
//universal_lexed = lexar("II power III power II $") //DXII

//fmt.Println(Roman(64))
//fmt.Println(check_valid("XV"))
//fmt.Println("hello world")
//fmt.Printf("output = %v \n", romanToInt("VX"))
//fmt.Printf("【input】:%v    【output】:%v\n", p.one, romanToInt(p.one))
//fmt.Printf("\n\n\n")
*/
