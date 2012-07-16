package main

import (
       "net/http"
       "fmt"
       "io"
       "exp/html"
       "bytes"
)

func extractLinks(resp io.Reader) []string {
     var links = make([]string, 1000)
     anchorTag := []byte{'a'} 
     tkzer := html.NewTokenizer(resp)
     var more bool
     var value []byte
     var key []byte
     eof := false
     i := 0
     for {
                if eof == true {
                    break
                }
                switch tkzer.Next() { 
                case html.ErrorToken: 
                        if tkzer.Err() == io.EOF {
               		    eof = true
                        }
                case html.StartTagToken: 
                        tag, hasAttr := tkzer.TagName() 
                        if hasAttr && bytes.Equal(anchorTag, tag) {
                  	        more = true
                                for more==true {              
			   	    key, value, more = tkzer.TagAttr()
  				    if string(key) == "href" {
                                       more = false
                                       fmt.Printf("%d\n", string(value), len(links))
				       links[i] = string(value)
				       i++
                                    }
                                }
                        }
                }
        } 
     return links
}

func getUrl(url string) []string {
     resp, err := http.Get(url)
     defer resp.Body.Close()
     var links []string
     if err != nil {
     	panic("couldn't get the url")
     } else {
        links = extractLinks(resp.Body)
     }
     return links
}

func Crawl(url string) {
     links := getUrl(url)
     fmt.Printf("%v", links[50])
}


func main() {
     Crawl("http://www.bbc.co.uk/sport/0/football/")
}