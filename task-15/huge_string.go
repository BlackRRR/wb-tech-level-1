package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

var justString string

// createHugeString - creates a huge string from random characters
func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func SomeFunc() {
	v := createHugeString(1 << 10)

	// justString = v[:100]

	// problems with characters other than UTF-8, so it's better to convert to runes
	//justString = string([]rune(v)[:100])

	// problem with base array memory size
	justString = string(append([]rune{}, []rune(v)[:100]...))
	// tmp := make([]rune, 100)
	// copy(tmp, []rune(v)[:100])
	// justString = string(tmp)

	fmt.Println("justString = ", justString)
	fmt.Println("cap(justString) ", cap([]rune(justString)))
	fmt.Printf("%p\n", &justString)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(justString))+len(justString))
	fmt.Println()
	fmt.Println("v = ", v)
	fmt.Println("cap(v) ", cap([]rune(v)))
	fmt.Printf("%p\n", &v)
	fmt.Println("Size in bytes ", int(unsafe.Sizeof(v))+len(v))

}

func main() {
	SomeFunc()
}
