package some

import (
	"fmt"
	"strings"
	"testing"
)

func TestSome(t *testing.T) {
	s := "attachment; filename==5.2020-03-03.csv"
	ss := strings.Trim(s, "attachment; filename=")

	fmt.Println(ss)
}
