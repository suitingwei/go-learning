package types

import (
	"fmt"
	"math/rand"
	"testing"
)

const matrixLength = 2048

func getSrcArray() [matrixLength][matrixLength]int {
	result := [matrixLength][matrixLength]int{}
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			result[i][j] = rand.Int()
		}
	}
	return result
}

func BenchmarkRowCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		src := getSrcArray()
		dest := [matrixLength][matrixLength]int{}

		rowCopy(src, dest)
	}
}
func BenchmarkColumnCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		src := getSrcArray()
		dest := [matrixLength][matrixLength]int{}

		rowCopy(src, dest)
	}
}

func rowCopy(src, dest [matrixLength][matrixLength]int) {
	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			dest[i][j] = src[i][j]
		}
	}
}
func columnCopy(src, dest [matrixLength][matrixLength]int) {
	for j := 0; j < matrixLength; j++ {
		for i := 0; i < matrixLength; i++ {
			dest[i][j] = src[i][j]
		}
	}
}

type User struct {
	Name string
	Age  int
}

type UserList []User

func TestAbc(t *testing.T) {
	user1 := User{Name: "jon"}
	user2 := User{Name: "alice"}
	user3 := User{Name: "bob"}

	users := []User{
		user1, user2, user3,
	}

	userList := UserList{}

	if userList == nil {
		fmt.Println("user list is nil")
	}

	fmt.Println("users", userList, users)
}
