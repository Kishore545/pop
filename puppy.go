package pop

import (
	"fmt"

	dol "github.com/Kishore545/Dol"
)

func Bark() string {
	return "Woof"
}

func Barks() string {
	return "Woof Woof Woof"
}

func Bigbark() string {
	return dol.WhenGrownUp(Bark())
}

func Bigbarks() string {
	return dol.WhenGrownUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}

func From12() {
	fmt.Println("I'm from version 1.2.0")
}

func From13() {
	fmt.Println("I'm from version 1.3.0")
}
