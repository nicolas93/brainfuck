package main



import ("fmt"
		"os"
		"flag"
		"strings"
		"io/ioutil"
		"bufio")



func interpret(code string){
	field := make([]int, 30000)
	stack := make([]int, len(code))
	s_i := 0
	f_i := 0
	reader := bufio.NewReader(os.Stdin)
	var b []byte = make([]byte, 1)
	for i:=0; i<len(code); i++{
		switch code[i]{
			case 43: // +
				field[f_i]++
			case 45: // -
				field[f_i]--
			case 62: // >
				f_i = (f_i +1) % len(code)
			case 60: // <
				f_i = (f_i -1)
				if f_i < 0 {
					f_i = len(code)-f_i
				}
			case 44: // ,
			    os.Stdin.Read(b)
			    field[f_i] = int(b[0])
			    if(field[f_i] == 10){
			    	field[f_i] = 0
			    }
			    //fmt.Println(field[f_i])
			case 46: // .
				fmt.Printf(string(field[f_i]))
			case 91: // [
				stack[s_i] = i 
				s_i++
			case 93: // ]
				if field[f_i]==0{
					s_i--
				}else{
					i = stack[s_i-1]
				}
		default:
			//fmt.Println("default")
		}
	if 1 == 11{ // debug
		fmt.Println(string(code[i]))
		text, _ := reader.ReadString('\n')
		fmt.Println(text)}
	}
}



func main() {
	var code_file string
	flag.StringVar(&code_file, "i", "", "Brainfuck source code file")
	var code_string string
	flag.StringVar(&code_string, "c", "", "Brainfuck source code inline")
	flag.Parse()

	if strings.Compare(code_file, "") != 0{
		b, err := ioutil.ReadFile(code_file)
    	if err != nil {
			fmt.Print(err)
   		}
   		interpret(string(b))
	}else if strings.Compare(code_string, "") != 0{
		interpret(code_string)
	}


//	interpret("+++++ +++++[>++[>+++<-]<-]>>+++++.")
//	interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
//	interpret(",[+++.,]")
}