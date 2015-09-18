package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
    "os"
)

func getFbShares(url string) (s string) {
    s = url
    return
}

func handler(w http.ResponseWriter, r *http.Request) {

    urlParam := strings.Join(r.URL.Query()["url"], "")
    url := "http://api.facebook.com/restserver.php?method=links.getStats&urls=" + urlParam
    response, err := http.Get(url)

    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
    }

    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    fmt.Print(getFbShares("foobar"))
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}