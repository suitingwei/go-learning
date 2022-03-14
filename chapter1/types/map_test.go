package types

import "testing"

func TestDeleteMapKeys(t *testing.T) {
	person := make(map[int]bool)

	person[1] = true
	person[2] = true
	person[3] = true
	person[4] = true
	person[5] = true

	t.Log(person)

	delete(person, 3)
	delete(person, 4)

	t.Log(person)
}
