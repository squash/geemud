package main

import (
_    "net"
_    "net/http"
    "fmt"
    
    )
    
    
func main() {
// Set up interface. Basic text for now, maybe more interesting later.

    i:=CreateIface("test")
    i.Add("This is a test.\n")

    commandIn:=make (chan string)
    go getInput(i, commandIn)    
    for {
        select {
           case cin:= <- commandIn:
            fmt.Println("In select, got ", cin)
        }
    }
    
}

func getInput(i *Iface, c chan <- string) {
    for {
        in:=string(i.Read())
        fmt.Println("Got ", in)
        c <- in
    

    }
}
