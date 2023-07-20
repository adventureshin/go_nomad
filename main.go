package main

import (
    "net/http"
    "log"
    "fmt"
    "github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.fmkorea.com/index.php?mid=best&page=1"

func main(){
    getPages()
}


func getPages() int {
    res, err := http.Get(baseURL)
    checkErr(err)
    checkCode(res)

    defer res.Body.Close()

    doc, err := goquery.NewDocumentFromReader(res.Body)
    checkErr(err)
    checkCode(res)
    fmt.Println(doc)
    return 0
}

func checkErr(err error){
    if err != nil {
        log.Fatalln(err)
    }
}

func checkCode(res *http.Response){
    if res.StatusCode != 200 {
        log.Fatalln("Request failed with Status:", res.StatusCode)
    }
}
