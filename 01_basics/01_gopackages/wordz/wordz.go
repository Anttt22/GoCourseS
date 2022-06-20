package wordz

//reusable package because nami in not main
// and as a resutl can not be runing without import
import (
	"crypto/rand"
	"math/big"
)

var Hello = "This is package wordz"
var Prefix = "Random word: "
var Words = []string{"one", "two", "'three", "four", "five"}

func Random() string {
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64())
}

func get(index int64) string {
	return Prefix + Words[index]
}
