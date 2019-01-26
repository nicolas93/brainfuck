package main



import ("fmt"
		"os"
		"bufio")



func interpret(code string){
	field := make([]int, 30000)
	stack := make([]int, len(code))
	s_i := 0
	f_i := 0
	reader := bufio.NewReader(os.Stdin)
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
				fmt.Println(",")
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
			fmt.Println("default")
		}
	if 1 == 11{ // debug
		fmt.Println(string(code[i]))
		text, _ := reader.ReadString('\n')
		fmt.Println(text)}
	}
}



func main() {
//	interpret("+++++ +++++[>++[>+++<-]<-]>>+++++.")
	interpret("++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.")
}