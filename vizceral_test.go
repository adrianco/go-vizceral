// Tests for vizceral
package vizceral

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	fmt.Println("Simple read/write test")
	Write(ReadFile("simple_cleanup.json"))
	fmt.Println("Complex read test")
	v := ReadFile("complex_cleanup.json")
	fmt.Println(v.Name)
	for _, region := range v.Nodes {
		fmt.Println(region.Name)
		for _, node := range region.Nodes {
			fmt.Printf("\t%s.%s.%s\n", v.Name, region.Name, node.Name)
		}
	}
}
