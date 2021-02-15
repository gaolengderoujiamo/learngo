package main

import "learngo/demotest/effectiveGo/some/conversions"

func main() {
	dog0 := new(conversions.Dog)
	dog0.WangWang()
	(*conversions.Animal)(dog0).Run()
	(*conversions.Object)(dog0).Work()

	object := new(conversions.Object)
	object.Work()

	animal := (*conversions.Animal)(object)
	animal.Run()

	dog := (*conversions.Dog)(object)
	dog.WangWang()

}
