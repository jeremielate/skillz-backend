package main

import (
	"fmt"
	"testing"
)

func TestAssets(t *testing.T) {
	fmt.Println("list assets")
	list := AssetNames()
	if len(list) == 0 {
		t.Fail()
	}
	for _, v := range list {
		fmt.Println(v)
	}
}
