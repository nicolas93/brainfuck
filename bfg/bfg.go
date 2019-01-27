package main

import (
	"bufio"
	"fmt"
	"os"
)

func cleanup(code string) string {
	cleancode := ""
	for i := 0; i < len(code); i++ {
		if code[i] == 60 && i < len(code)-1 {
			if code[i+1] != 62 {
				cleancode += string(code[i])
			} else {
				i++
			}
		} else if code[i] == 62 && i < len(code)-1 {
			if code[i+1] != 60 {
				cleancode += string(code[i])
			} else {
				i++
			}
		} else {
			if i != len(code)-1 || code[i] == 46 {
				cleancode += string(code[i])
			}
		}
	}
	return cleancode
}

func generate(text string) string {
	code := ""
	lastc := 0
	for i := 0; i < len(text); i++ {
		if int(text[i]) > lastc {
			diff := int(text[i]) - lastc
			if diff > 25 {
				for j := 0; j < diff/10; j++ {
					code += "+"
				}
				code += "[>++++++++++<-]>"
				for j := 0; j < diff%10; j++ {
					code += "+"
				}
				code += ".<"
			} else {
				code += ">"
				for j := 0; j < diff; j++ {
					code += "+"
				}
				code += ".<"
			}
		} else {
			diff := lastc - int(text[i])
			if diff > 25 {
				for j := 0; j < diff/10; j++ {
					code += "+"
				}
				code += "[>----------<-]>"
				for j := 0; j < diff%10; j++ {
					code += "-"
				}
				code += ".<"
			} else {
				code += ">"
				for j := 0; j < diff; j++ {
					code += "-"
				}
				code += ".<"
			}
		}
		lastc = int(text[i])
	}
	return code
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(cleanup(generate(text)))
}
