package main

import (
	"errors"
	"fmt"
	"net/http"
)


var errRequestFailed = errors.New("Request failed")

func main() {
	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// dictionary.Search(baseWord)
	// dictionary.Delete(baseWord)

	// word, err := dictionary.Search(baseWord)

	// if err !=nil{
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(word)
	// }
	// err := dictionary.Update("asdfasd", "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var results = make(map[string]string) 

	urls := []string {
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _,url := range urls{
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results{
		fmt.Println(url, result)
	}
}

func hitURL(url string) error{
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil|| resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}
