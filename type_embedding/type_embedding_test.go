package type_embedding

import (
	"fmt"
	"testing"
)

type B1 struct {
	Name string
}

func (b *B1) Hello() {
	fmt.Printf("Hello %s\n", b.Name)
}

type A1 struct {
	Name string
	*B1
}

func (a *A1) Hello() {

}

func TestTypeEmbedding(t *testing.T) {
	a := &A1{}

	a.Hello()
}
