Brainfuck
===

## bfi
### Usage:
```
bfi -h
Usage of bfi:
  -c string
    	Brainfuck source code inline
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

### Installation

```
git clone https://github.com/nicolas93/brainfuck.git
cd brainfuck
go install ./bfi
```

You can also just use `go run bfi/bfi.go`.
