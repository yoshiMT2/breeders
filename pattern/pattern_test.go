package pattern

import (
	"fmt"
	"testing"
)

func TestFuctory(t *testing.T) {
	factory := Product{}

	product := factory.New()

	fmt.Println(product.CreateAt.UTC())
}
