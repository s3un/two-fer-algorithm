
package twofer

import (
	"fmt"
	"github.com/s3un/two-fer-algorithm"
)

func ShareWith(string) {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)

	condition := name != ""
	
	if condition {
		fmt.Println("One for ", name, "one for me")
	} else {
		fmt.Println("One for you, one for me")
	}
}
