package util_Joe

func xmlFile(path string) int {

    fmt.Printf("xmlpth: [%s]\n", path)
    xmlcontent := make([]byte, 102400)
    var v ATOM

    file, err := os.Open(path)
    if err != nil {
        fmt.Printf("Open File Err: %v", err)
        return -1
    }

    file.Read(xmlcontent)
    file.Close()

    //fmt.Printf("xmlContent ;[ %s ]\n", xmlcontent)
    //enc := xml.procInstEncoding(string(xmlcontent))

    //enc := procInstEncoding(string(xmlcontent))

    //if strings.ToLower(enc) != "utf-8" /* && enc != "UTF-8"  */ {
    //    fmt.Printf("[Notice] : This Docement is Not utf-8 character! \n")
    //}

    err = xml.Unmarshal(xmlcontent, &v)

    //fmt.Printf("XML ENCODING : %s\n",v.XMLName.Space)
    //fmt.Printf("XML ENCODING : %s\n",v.XMLENC)

    if err != nil {
        fmt.Printf("Err in Parse Xml File : %s \n%v\n", path,err)
        return -1
    }

    fmt.Print("id : %s\n", v.id)

    return 0

}

