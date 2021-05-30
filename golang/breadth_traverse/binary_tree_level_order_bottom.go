/**
 * 描述：
	*
 */

package main

import (
	"fmt"
	"strings"
)

func main(){
	//s := "connect: connection refusedxxxx"
	var err error

	if err != nil || strings.Contains(err.Error(), "connect: connection refused"){
		fmt.Println("ok")
	}


}