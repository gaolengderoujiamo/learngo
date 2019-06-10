package deferfatal

import (
	"fmt"
	"log"
)

func test_deferfatal() {
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

func main() {
	test_deferfatal()
}
