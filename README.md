# Go Command Line Urban Dictionary
Just a basic go command line example that find a definition of a word from Urban Dictionary API and displays the first result to the user.

`	url := fmt.Sprintf("http://api.urbandictionary.com/v0/define?term=%s", *wordPtr)`

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
Word: YOLO
Defenition: An overused acronym for "[You only live once]." There is an [exception] for those who believe in [reincarnation] or are cats.
Example: Examples:
—I got up.
#[yolo]
—I inhaled.
#yolo
—I exhaled.
#yolo
—I [took a shit].
#yolo
—I [flushed] the toilet.
#yolo
—I stretched.
#yolo
```

