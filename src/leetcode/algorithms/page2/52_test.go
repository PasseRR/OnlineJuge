package page2

import (
	"fmt"
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	fmt.Printf("%v\n", totalNQueens(1))
	fmt.Printf("%v\n", totalNQueens(4))
}
