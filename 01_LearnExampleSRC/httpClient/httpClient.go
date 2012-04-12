package main

import (
    "3lib/mahonia"
    "encoding/xml"
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strings"
    //"util_Joe"
)

type ATOM struct {
    XMLName      xml.Name `Name{"http://www.w3.org/2005/Atom", "feed"}`
    title        string   `xml:"title"`
    subtitle     string   `xml:"subtitle"`
    id           string   `xml:"id"`
    updated      string   `xml:"updated"`
    author_name  string   `xml:"author>name"`
    author_email string   `xml:"author>email"`
}

// COPY From /usr/local/go/src/pkg/encoding/xml/xml.go
// procInstEncoding parses the `encoding="..."` or `encoding='...'`
// value out of the provided string, returning "" if not found.
func procInstEncoding(s string) string {
    idx := strings.Index(s, "encoding=")
    if idx == -1 {
        return ""
    }
    v := s[idx+len("encoding="):]
    if v == "" {
        return ""
    }
    if v[0] != '\'' && v[0] != '"' {
        return ""
    }
    idx = strings.IndexRune(v[1:], rune(v[0]))
    if idx == -1 {
        return ""
    }
    return v[1 : idx+1]
}

func xmlFile(path string) int {
    //fmt.Printf("xmlpth: [%s]\n", path)
    xmlcontent := make([]byte, 102400)
    v := ATOM{updated: "testNone"}

    file, err := os.Open(path)
    if err != nil {
        fmt.Printf("Open File Err: %v", err)
        return -1
    }

    file.Read(xmlcontent)
    file.Close()

    enc := procInstEncoding(string(xmlcontent))

    if strings.ToLower(enc) != "utf-8" && enc != "UTF-8" {
        fmt.Printf("[Notice] : This Docement is Not utf-8 character! - [%s] \n", enc)
        fmt.Printf("[Notice] : Need Convert Character!\n")
        decode := mahonia.NewDecoder(string("gbk"))
        if decode == nil {
            log.Fatalf("Could not Create decoder ! ")
        }

        string8 := decode.ConvertString(string(xmlcontent))
        if len(string8) == 0 {
            fmt.Printf("Err in Parse Xml File\n")
        }
        utf8string := strings.Replace(string8, string(enc), "utf-8", len(enc))

        //fmt.Printf("[%100s]\n", utf8string)
        err = xml.Unmarshal([]byte(utf8string), &v)
        if err != nil {
            fmt.Printf("Err in Parse Xml File : %s \n%v\n", path, err)
            return -1
        }

        fmt.Printf("title : %q\n", v.title)

    } else {
        err = xml.Unmarshal(xmlcontent, &v)

        if err != nil {
            fmt.Printf("Err in Parse Xml File : %s \n%v\n", path, err)
            return -1
        }
    }

    fmt.Printf("id : %q\n", v.id)

    return 0
}

func writefile(data []byte, fpath string) int {

    fout, err := os.Create(fpath)

    if err != nil {
        fmt.Printf("Create file : %v", err)
        return -1
    }

    fout.Write(data)
    fout.Close()

    return 0
}

func main() {

    flag.Parse()
    s := ""

    xmlPath := "/tmp/joe_xu_xmlFile_httpClient.xml"

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
        return
    }
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Printf("ReadAll : %v", err)
    }
    //fmt.Printf("Get Success : [%s]  \nSave to [%s] \r\n", s, xmlPath)

    if writefile(data, xmlPath) != 0 {
        fmt.Printf("Can not Write Content To File!\n")
        return
    }

    if xmlFile(xmlPath) == 0 {
        fmt.Printf("Success\n")
    }

    return
}
