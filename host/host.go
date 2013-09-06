package main

import (
	"fmt"
	"net/url"
)

func main() {
	gurl, err := url.Parse("http://www.baidu.com")
	fmt.Println(gurl.Host, err)
}
