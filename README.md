# Go Command Line Urban Dictionary
Just a basic go command line example that find a definition of a word from Urban Dictionary API and displays the first result to the user.

## Start
```
➜  go-command-line-urban-dictionary git:(master) ✗ go build .
➜  go-command-line-urban-dictionary git:(master) ✗ ./go-command-line-urban-dictionary -h
Usage of ./go-command-line-urban-dictionary:
  -word string
        a string (default "foo")
```
use the `-word` flag to pass in the word
```
➜  go-command-line-urban-dictionary git:(master) ✗ ./go-command-line-urban-dictionary -word=yolo
word: yolo
➜  go-command-line-urban-dictionary git:(master) ✗
```

