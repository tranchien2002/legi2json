package main

import (
	"encoding/json"
	"fmt"
	
)

func main() {
	fmt.Println("haha")
	bb, err := json.Marshal("Điều")

	if err != nil {
		fmt.Println("Opps =]] :", err)
		return
	}

	

	fmt.Println(string(bb))

}
