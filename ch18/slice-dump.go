package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Makemake", "Haumea"}
	dump("dwarfs", dwarfs)
	// appending
	dwarfs2 := append(dwarfs, "Orcus")
	dump("dwarfs2", dwarfs2)

	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs3", dwarfs3)

	//lets see modifying changes the original string
	dwarfs3[1] = "Pluto!"
	fmt.Println(dwarfs)
	fmt.Println(dwarfs2)
	fmt.Println(dwarfs3)

	// dump("dwarfs[1:2]", dwarfs[1:2])
}
