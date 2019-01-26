package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func skip_loop(code string, i int) int {
	count := 0
	for j := i + 1; j < len(code); j++ {
		if code[j] == 93 {
			if count == 0 {
				return j
			} else {
				count--
			}
		}
		if code[j] == 91 {
			count++
		}
	}
	return len(code) - 1
}

func interpret(code string) {
	field := make([]int, 30000)
	stack := make([]int, len(code))
	s_i := 0
	f_i := 0
	reader := bufio.NewReader(os.Stdin)
	var b []byte = make([]byte, 1)
	for i := 0; i < len(code); i++ {
		switch code[i] {
		case 43: // +
			field[f_i]++
		case 45: // -
			field[f_i]--
		case 62: // >
			f_i = (f_i + 1) % len(code)
		case 60: // <
			f_i = (f_i - 1)
			if f_i < 0 {
				f_i = len(code) - f_i
			}
		case 44: // ,
			os.Stdin.Read(b)
			field[f_i] = int(b[0])
			if field[f_i] == 10 {
				field[f_i] = 0
			}
			//fmt.Println(field[f_i])
		case 46: // .
			fmt.Printf(string(field[f_i]))
		case 91: // [
			if field[f_i] == 0 {
				i = skip_loop(code, i)
			} else {
				stack[s_i] = i
				s_i++
			}
		case 93: // ]
			i = stack[s_i-1] - 1
			s_i--
		default:
			//fmt.Println("default")
		}
		if 1 == 11 { // debug
			fmt.Println(i)
			fmt.Println(string(code[i]))
			fmt.Println(stack)
			fmt.Println(field[0:10])
			reader.ReadString('\n')
			//fmt.Println(text)
		}
	}
}

func main() {
	var code_file string
	flag.StringVar(&code_file, "i", "", "Brainfuck source code file")
	var code_string string
	flag.StringVar(&code_string, "c", "", "Brainfuck source code inline")
	flag.Parse()

	if strings.Compare(code_file, "") != 0 {
		b, err := ioutil.ReadFile(code_file)
		if err != nil {
			fmt.Print(err)
		}
		interpret(string(b))
	} else if strings.Compare(code_string, "") != 0 {
		interpret(code_string)
	}

	//	interpret("+++++ +++++[>++[>+++<-]<-]>>+++++.")
	//	interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
	//	interpret(",[+++.,]")
}
