package main

import (
    "encoding/xml"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func xmlFile(path string) {

    fmt.Printf("xmlpth: [%s]\n", path)

}

type

func writefile(data []byte, fpath string) {

    fout, err := os.Create(fpath)

    if err != nil {
        fmt.Printf("Create file : %v", err)
    }

    fout.Write(data)
    fout.Close()
}

func main() {

    flag.Parse()
    s := ""

    xmlPath := "/tmp/joe.xu.xmlFile.httpClient"

    if flag.NArg() < 1 {
        s = "http://blog.codingnow.com/atom.xml"
    } else {
        for i := 0; i < flag.NArg(); i++ {
            s += flag.Arg(i)
        }
    }

    fmt.Printf("Args : [%s]\n", s)

    res, err := http.Get(s)

    if err != nil {
        fmt.Printf("Get : %v\nexit!\n", err)
        os.Exit(1)
    }
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Printf("ReadAll : %v", err)
    }
    fmt.Printf("Get Success : [%s]  Save to [%s] \r\n", s, xmlPath)

    writefile(data, xmlPath)

    xmlFile(xmlPath)
}
