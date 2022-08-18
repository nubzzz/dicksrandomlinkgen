package main

import (
    "fmt"
    "net/http"
    "os"
    "bufio"
    "math/rand"
    "time"
    "log"
)

var links []string

func randomlink(w http.ResponseWriter, req *http.Request) {
    http.Redirect(w, req, links[rand.Intn(len(links))], 302)
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

func main() {
    rand.Seed(time.Now().Unix())
    
    log.Println("Attempting to open randomlinks.txt")
    file, err := os.Open("randomlinks.txt")
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    for sc.Scan() {
        links = append(links, sc.Text())
    }
    if err := sc.Err(); err != nil {
        log.Fatalln(err)
    }
    log.Println("Loaded links into RAM successfully")

    http.HandleFunc("/randomlink", randomlink)
    http.HandleFunc("/headers", headers)

    var port = ":8090"
    log.Println("Listening on port", port)
    http.ListenAndServe(port, Log(http.DefaultServeMux))
}
