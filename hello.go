package hello

import (
    "fmt"
    "net/http"
    //"time"
    //"net"
    "io/ioutil"

    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
)

func init() {
    http.HandleFunc("/", handler)
}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
    req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
    req.Header.Add("Content-Type", via[0].Header.Get("Content-Type"))
    return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    // json data
    url := "https://developer-api.nest.com"

    //var netTransport = &http.Transport{
    //    Dial: (&net.Dialer{
    //        Timeout: 5 * time.Second,
    //    }).Dial,
    //    TLSHandshakeTimeout: 5 * time.Second,
    //}
	//
    //var netClient = &http.Client{
    //    Timeout: time.Second * 10,
    //    Transport: netTransport,
    //}

    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)

    client.CheckRedirect = checkRedirectFunc

    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Add("Authorization", `Bearer TOKEN_GOES_HERE`)
    req.Header.Add("Content-Type", "application/json")

    fmt.Fprint(w, "Request:", req)

    res, err := client.Do(req)

    if err != nil {
        panic(err.Error())
    }

    fmt.Fprint(w, "Response:", res)

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        panic(err.Error())
    }

    fmt.Fprint(w, "Test\n")

    fmt.Fprintf(w, "JSON1: %v", body)
    fmt.Fprintf(w, "JSON2: %+v", body)
    fmt.Fprintf(w, "JSON3: %s", body)
}

