package puppy

import "github.com/LvBH/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof"
}

func BigBark() string {
	return dog.WheGrownUp("WOOF!!!!")
}
