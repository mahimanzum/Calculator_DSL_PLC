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
	return true
}
func write_message(index int, s string, length int) {
	fmt.Print(raw_code + "\n")
	for i := 0; i < index-length; i += 1 {
		fmt.Print(" ")
	}
	fmt.Print("^\n")
	fmt.Print(error_message[s])
	os.Exit(1)
}

func lexar(code string) []token {
	var a []token
	prev := ""
	id := 1
	for string_id, val := range code {
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
			if prev == "times" || prev == "plus" || prev == "power" || prev == "divide" || prev == "minus" {
				a = append(a, token{name: token_table[prev], value: 0, index: id, start_id: string_id - len(prev), end_id: string_id})
				id += 1
				prev = ""
			} else {
				if len(prev) > 0 && check_valid(prev) {
					a = append(a, token{name: "Number", value: romanToInt(prev), index: id, start_id: string_id - len(prev), end_id: string_id})
					id += 1
					prev = ""
				} else if len(prev) > 0 {
					write_message(string_id, "lexical_error", len(prev))
					return []token{{name: "error", value: 0}}
				}
			}
		} else if val == '$' && len(prev) == 0 {
			a = append(a, token{name: "end_token", value: 0, index: id})
			id += 1
			return a
		} else {
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

var idx int = 0
var universal_lexed []token
var current_token token
var raw_code string

func lex() token {
	if universal_lexed[idx].name == "end_token" {
		return token{name: "end_token", value: 0, start_id: len(raw_code) - 1, end_id: len(raw_code)}
	} else {
		next_token := universal_lexed[idx]
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
			return term
		}
	}
}

//term -> factor [ ('*' | '/') factor ]*
func parse_term() int {
	factor := parse_factor()
	for {
		next_token := peak()
		if next_token.name == "times_token" {
			lex()
			factor = factor * parse_factor()
		} else if next_token.name == "divide_token" {
			cur_tok := next_token
			lex()
			val := parse_factor()
			//fmt.Println("comes here val = ", val)
			if factor%val != 0 {
				write_message(cur_tok.start_id, "zero_error", 0)
			}
			factor = factor / val
		} else {
			return factor
		}
	}
}

//factor -> base [ '^' exponent ]*
func parse_factor() int {
	base := parse_base()
	var exp int
	exp = 1
	for {
		next_token := peak()
		if next_token.name == "power_token" {
			lex()
			exp = parse_exponent()
			return int(math.Pow(float64(base), float64(exp)))
		} else {
			return base
		}
	}
}

//base -> number| '(' expr ')'
func parse_base() int {
	next_token := lex()
	var value int
	if next_token.name == "left_bracket" {
		value = parse_expr()
		next_token = lex()
		if next_token.name != "right_bracket" {
			write_message(next_token.end_id, "syntax_error", 0)
		}
	} else {
		value = parse_number()
	}
	return value
}

//exponent -> base| [ '^' exponent ]*
func parse_exponent() int {
	base := parse_base()
	var exp int
	exp = 1.00
	for {
		next_token := peak()
		if next_token.name == "power_token" {
			lex()
			exp = parse_exponent()
			return int(math.Pow(float64(base), float64(exp)))
		} else {
			return base
		}
	}
}

func parse_number() int {
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
	//raw_code = "III plus {IV times II} power II]"
	raw_code = os.Args[1]
	universal_lexed = lexar(raw_code + " $")
	//fmt.Print(lex())
	val := Roman(parse_expr())
	final_token := lex()
	if final_token.name != "end_token" {
		fmt.Println(final_token)
		write_message(final_token.start_id, "syntax_error", 0)
	}
	fmt.Println(val)
	os.Exit(0)
}
