# Borf

A small, forth-like, language. 

## Build

	go build

## Run

	./borf

## Examples

```
; ./borf
» 2
ok.
» 3
ok.
» +
ok.
» 5
ok.
» +
ok.
» size
1
» show
{"10" → "Integral"}
» .
10
ok.
» +
err: could not push word #0: {"0x49e3a0" → "Procedure"} - stack is too small to Add upon; need: 2, got 0
» bye

Goodbye ☺
; 
```
