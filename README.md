Brainfuck
===
Several programs to handle the brainfuck programming language.
This project is developed for learning purposes and to provide some simple programs.


## bfi
Brainfuck interpreter

### Usage
```bash
bfi -h
Usage of bfi:
  -c string
    	Brainfuck source code inline
  -debug
    	Print any non-bf symbols for debugging purposes.
  -i string
    	Brainfuck source code file
```
#### Examples
* interpret a source code file:
	`bfi -i source.bf`
* directly interpret a piece of code:
	`bfi -c "+++++ +++++[>++++++<-]>+++++."`
* use caesar "encryption" :
	`echo "Brainfuck" | bfi -c ",[+++.,]"`
* use debugflag:
  `bfi -i bfsource/reverse-input.bf -debug`

### Installation

```bash
git clone https://github.com/nicolas93/brainfuck.git
cd brainfuck/bfi
go install ./bfi
```

You can also just use `go run bfi/bfi.go`.

## bfg
Brainfuck generator

### Usage
```bash
bfg
<ctrl> + <D> twice to exit insert mode

```

#### Examples
Text to brainfuck.
```bash
echo "abc" | go run bfg/bfg.go
(ctrl) + (D) twice to exit insert mode
```
```brainfuck
+++++++++[>++++++++++<-]>+++++++.+.+.<++++++++[>----------<-]>---------.
```

Brainfuck from textfile:
```bash
cat LICENSE | bfg
```

### Installation

```bash
git clone https://github.com/nicolas93/brainfuck.git
cd brainfuck/bfg
go install ./bfg
```
