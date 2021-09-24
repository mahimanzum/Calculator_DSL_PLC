package main

import "fmt"

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
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
		} else if roman[string(s[id])] <= roman[string(s[id-1])] {
			continue
		} else if _, ok := valids[s[id-1:id+1]]; ok {
			continue
		} else if _, ok := roman[string(val)]; !ok {

			return false
		} else {
			return false
		}
	}
	return true
}
func lexar(code string) []string {
	var a []string
	prev := ""
	//fmt.Println("comes")
	for _, val := range code {
		fmt.Println(prev)
		fmt.Println(a)
		if val == '(' || val == '{' || val == '[' {
			a = append(a, "left_bracket")
		} else if val == ')' || val == '}' || val == ']' {
			if len(prev) > 0 && check_valid(prev) {
				a = append(a, "Number")
				prev = ""
			}
			a = append(a, "right_bracket")
		} else if val == ' ' {
			if prev == "times" || prev == "plus" || prev == "power" || prev == "divide" || prev == "minus" {
				a = append(a, token_table[prev])
				prev = ""
			} else {
				if len(prev) > 0 && check_valid(prev) {
					a = append(a, "Number")
					prev = ""
				} else if len(prev) > 0 {
					fmt.Println("print from error", prev)
					return []string{"error"}
				}
			}
		} else if val == '$' && len(prev) == 0 {
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

//if val, ok := dict["foo"]; ok

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

func main() {
	/*
		for i := 5; i < 50; i++ {
			if i != romanToInt(Roman(i)) {
				fmt.Print(i)
			}
		}
	*/
	fmt.Println(check_valid("XI"))
	//fmt.Println(lexar("XI plus (X plus X)$"))
	fmt.Println(lexar("{MCMXCVIII divide III divide VI minus XI) divide X power II"))
	//fmt.Println(check_valid("XV"))
	//fmt.Println("hello world")
	//fmt.Printf("output = %v \n", romanToInt("VX"))
	//fmt.Printf("【input】:%v    【output】:%v\n", p.one, romanToInt(p.one))
	//fmt.Printf("\n\n\n")
}

/*
Grammar

expr -> term [ ('+' | '-') term ]*
term -> factor [ ('*' | '/') factor ]*
factor -> base [ '^' exponent ]*
base -> '(' expr ')' | identifier | number
exponent -> '(' expr ')' | identifier | number


*/
