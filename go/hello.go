package main

import (
    "os"
    "fmt"
    "io"
    "strings"
    )
     
func main() {
    // 打开原始文件
    file, err := os.Open("../utils/daemon.json")
    if err != nil {
        fmt.Println("open file err :", err)
        return
    }
    defer file.Close()
    // 定义接收文件读取的字节数组
    var buf [128]byte
    var content []byte
    for {
        n, err := file.Read(buf[:])
        if err == io.EOF {
            // 读取结束
            break
        }
        if err != nil {
            fmt.Println("read file err ", err)
            return
        }
        content = append(content, buf[:n]...)
    }
    
    str := string(content);
    str = strings.Replace(str,"#{自己的ip最后尾数}","12",-1);
    fmt.Println(str)
}