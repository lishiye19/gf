package main

import (
    "fmt"
    "gitee.com/johng/gf/g/os/gfpool"
    "os"
    "time"
)

func main() {
    for {
        time.Sleep(time.Second)
        if f, err := gfpool.Open("/home/john/temp/log.log", os.O_RDONLY, 0666, 60000000*1000); err == nil {
            s, _ := f.Stat()
            fmt.Println(f.Name(), s.Name())
        } else {
            fmt.Println(err)
        }
    }
}