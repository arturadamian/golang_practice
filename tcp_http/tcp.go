package main 

import (
       "fmt"
       "net"
       "bufio"
)
func main() {
     conn, x := net.Dial("tcp", "golang.org:80")
     fmt.Println(x)
     fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
     status, _ := bufio.NewReader(conn).ReadString('\n')
     fmt.Println(status)
}