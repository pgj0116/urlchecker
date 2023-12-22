package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
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
	// var results = make(map[string]string) 

	// urls := []string {
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://instagram.com/",
	// 	"https://academy.nomadcoders.co/",
	// }
	// for _,url := range urls{
	// 	result := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// }
	// for url, result := range results{
	// 	fmt.Println(url, result)
	// }
	c:= make(chan string)
	people := [5]string {"A", "B", "C", "D", "E"}
	for _, person := range people{
		go isSexy(person, c)
	}
	for i:=0 ; i< len(people); i++{
		fmt.Println(<-c)
	}
	// go sexyCount("GJ")
	// go sexyCount("SH")
	// time.Sleep(time.Second * 5)
}

func hitURL(url string) error{
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil|| resp.StatusCode >= 400{
		return errRequestFailed
	}
	return nil
}

func sexyCount(person string) {
	for i:= 0; i <10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string)  {
	time.Sleep(time.Second * 5)
	c <- person +  " is sexy"
}