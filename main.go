package main

import (
	"crawler/fetcher"
	"fmt"
)

func main()  {
	context, _ := fetcher.Fetch("http://www.zhihu.com")
	fmt.Printf("%s", context)
}
