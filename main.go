import (
	"fmt"
	"strconv"
	"strings"
)

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number  > maxRomanNumber{
		return strconv.Itoa(number)
	}
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M" },
		{900,  "CM"},
		{500,  "D" },
		{100,  "C" },
		{90,   "XC"},
		{50,   "L" },
		{40,   "XL"},
		{10,   "X" },
		{9,    "IX"},
		{5,    "V" },
		{4,    "IV"},
		{1,    "I" },
	}
	var roman strings.Builder
	for _, conversion := range conversions{
		for number >= conversion.value{
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}
	
	return roman.String()
}

package main

import (
	"fmt"
)

func main() {
	x, e := GetRoman("MMMCDXXXLVIII")
	
	if e == true{
		fmt.Println(x)
	}else{
		fmt.Println("error!")
	}
}

func isClearRoman(s string) bool{
	for _, r := range s	{
		if (r == 'I' || // 3 :    1
			r == 'V' || // 1 :    5 
			r == 'X' || // 3 :   10
			r == 'L' || // 1 :   50
			r == 'C' || // 3 :  100
			r == 'D' || // 1 :  500
			r == 'M'){  // 3 : 1000
				continue
			}else{
				return false
			}
	}
	return true
}

type Roman struct{
	count int
	start bool
	end   bool
	flex  bool
}

func newRoman(flex bool) Roman{
	return Roman{0, false, false, flex}
}

func GetRoman(s string) (int, bool){
	if !isClearRoman(s){
		return 0, false
	}
	
	I := newRoman(false)
	V := newRoman(false)
	X := newRoman(true)
	L := newRoman(false)
	C := newRoman(true)
	D := newRoman(false)
	M := newRoman(true)
		
	z := 0
	
	for i := len(s) - 1; i >= 0; i-- {
		switch(s[i]){
			case 'I':
				if I.end || I.count == 3 || (V.end && I.count == 1) || (X.count > 1) || (V.start && X.start){
					return 0, false
				}
				I.start = true
				I.count++			
								
				if V.start{
					V.end = true
					z--
				}
				
				if X.start {
					if X.count == 1{
						X.flex = false
					}else{
						X.end = true
					}
					z--
				}
				if !V.start && !X.start{
					z++
				}
				
			case 'V':
				if V.end || V.count == 1 || X.start {
					return 0, false
				}
				V.start = true
				V.count++			
				z = z + 5
				
				if I.start{
					I.end = true
				}
			case 'X':
				if X.end || X.count == 3{
					return 0, false
				}
				X.start = true
				X.count++
				
				if L.start{
					z = z - 10
				}else{
					z = z + 10
				}
				
				if I.start{
					I.end = true
				}
				if V.start{
					V.end = true
				}
				if X.count == 2 {
					X.flex = false
				}
				
			case 'L':
				if L.end || L.count == 1 {
					return 0, false
				}
				L.start = true
				L.count++			
				z = z + 50
				
				if I.start{
					I.end = true
				}
				if V.start{
					V.end = true
				}
				if X.start{
					X.end = true
				}
			case 'C':
				if C.end || C.count == 3{
					return 0, false
				}
				C.start = true
				C.count++
			
				if D.start{
					z = z - 100
				}else{
					z = z + 100
				}
				
				if I.start{
					I.end = true
				}
				if V.start{
					V.end = true
				}
				if C.count == 2 {
					C.flex = false
				}
					
			case 'D':
				if D.end || D.count == 1 {
					return 0, false
				}
				D.start = true
				D.count++			
				z = z + 500
				
				if I.start{
					I.end = true
				}
				if V.start{
					V.end = true
				}
				if X.start{
					X.end = true
				}
			case 'M':
				if M.end || M.count == 3{
					return 0, false
				}
				M.start = true
				M.count++
			
				z = z + 1000
				
				if I.start{
					I.end = true
				}
				if V.start{
					V.end = true
				}
				if M.count == 2 {
					M.flex = false
				}
			}
		
		fmt.Println(i, z, ", I:{", I.count, I.start, I.end,
		"}, V:{", V.count, V.start, V.end,
		"}, X:{", X.count, X.start, X.end,
		"}, L:{", L.count, L.start, L.end,
		"}, C:{", C.count, C.start, C.end,
		"}, D:{", D.count, D.start, D.end,
		"}")
	}
	return z, true
	
}
