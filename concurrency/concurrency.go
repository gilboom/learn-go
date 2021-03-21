package concurrency

import (
	"fmt"
)

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result, len(results))
	for i, url := range urls {
		u := url
		go func(i int) {
			fmt.Printf("write: %v \n", i)
			resultChannel <- result{u, wc(u)}
			fmt.Printf("write after: %v \n", i)
		}(i)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		fmt.Printf("get result: %+v \n", result)
		results[result.string] = result.bool
	}
	close(resultChannel)
	//fmt.Printf("check channel \n")
	//result := <- resultChannel
	//fmt.Printf("get result: %+v \n", result)
	//time.Sleep(time.Second * 10)
	return results
	//for i := 0 ; i < len(urls); i++ {
	//	result := <- resultChannel
	//	results[result.string] = result.bool
	//}
	//time.Sleep(2 * time.Second)
	//return results
}
