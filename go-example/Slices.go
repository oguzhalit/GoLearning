package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "i"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s,"d")
	s = append(s,"e","f")
	s = append(s,"g","t")
	fmt.Println("apd:",s)

	c := make([]string, len(s))
	fmt.Println(len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	l := s[2:5]
	fmt.Println("sl1:",l)

	l2 := s[:5]
	fmt.Println("sl2:",l2)

	l3 := s[2:]
	fmt.Println("sl3:",l3)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:",t)
	
	twoD := make([][]int,3)
	for i := 0; i<3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j :=0; j<innerlen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:",twoD)
	// Slices ile örneklere devam et
}
