package main

import (
	"fmt"
	"strings"
	"time"

	nt "github.com/smirnoffkin/blockchain/network"
)

const (
	TO_UPPER = iota + 1
	TO_LOWER
)

const (
	ADDRESS = ":8080"
)

func main() {
	go nt.Listen(ADDRESS, handleServer)
	time.Sleep(500 * time.Millisecond)

	res := nt.Send(ADDRESS, &nt.Package{
		Option: TO_UPPER,
		Data:   "HEllo World",
	})
	fmt.Println(res.Data)

	res = nt.Send(ADDRESS, &nt.Package{
		Option: TO_LOWER,
		Data:   "HEllo World",
	})
	fmt.Println(res.Data)

}

func handleServer(conn nt.Conn, pack *nt.Package) {
	nt.Handle(TO_UPPER, conn, pack, handleToUpper)
	nt.Handle(TO_LOWER, conn, pack, handleToLower)
}

func handleToUpper(pack *nt.Package) string {
	return strings.ToUpper(pack.Data)
}
func handleToLower(pack *nt.Package) string {
	return strings.ToLower(pack.Data)
}
