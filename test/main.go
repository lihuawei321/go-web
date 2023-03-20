package main

import (
	"fmt"
	"strconv"
)

//func main() {
//	c := make([]int, 0)
//	a := []int{1, 2, 3}
//	b := []int{4, 5, 6}
//
//	for _, v := range a {
//		c = append(c, v)
//	}
//	for _, v2 := range b {
//		c = append(c, v2)
//	}
//	fmt.Println("ccc", c)
//}

func main() {
	ExServerList := []string{"5748"}
	deleteExServerList(5748, ExServerList)

}
func deleteExServerList(serverId int, ExServerList []string) bool {
	for _, v := range ExServerList {
		if v == strconv.Itoa(serverId) {
			fmt.Println("vvvvvvvv", serverId)
			return true
		}
	}
	return false
}
