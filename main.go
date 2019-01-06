package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
// Read more: https://gobyexample.com/command-line-flags
import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type defenitionResponse struct {
	Definition string `json:"definition"`
	Word       string `json:"word"`
	Example    string `json:"example"`
}

type listResponse struct {
	List []defenitionResponse `json:"list"`
}

func main() {

	// Basic flag declarations are available for string,
	// integer, and boolean options. Here we declare a
	// string flag `word` with a default value `"foo"`
	// and a short description. This `flag.String` function
	// returns a string pointer (not a string value);
	// we'll see how to use this pointer below.
	wordPtr := flag.String("word", "foo", "a string")

	// Once all flags are declared, call `flag.Parse()`
	// to execute the command-line parsing.
	flag.Parse()

	// Here we'll just dump out the parsed options and
	// any trailing positional arguments. Note that we
	// need to dereference the pointers with e.g. `*wordPtr`
	// to get the actual option values.
	url := fmt.Sprintf("http://api.urbandictionary.com/v0/define?term=%s", *wordPtr)

	// Check for HTTP error on connection
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(errors.New("Failed to connect to Urban Dictionary API"))
	}

	// Check for HTTP status code
	if res.StatusCode != 200 {
		fmt.Println(fmt.Errorf("Failed with HTTP Code: %d", res.StatusCode))
	}

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to read body from response"))
	}

	// Parse response
	var response listResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to unmarshal JSON: %s", body))
	}

	firstDefenition := response.List[0]

	fmt.Printf("Word: %s\n", firstDefenition.Word)
	fmt.Printf("Defenition: %s\n", firstDefenition.Definition)
	fmt.Printf("Example: %s\n", firstDefenition.Example)

}
