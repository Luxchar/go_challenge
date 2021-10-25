package main

func QuadA(x,y int) {
	a := "o"
	b := "-"
	c := "|"
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(a)
	for i:=1;i<y-1;i++ {
		print(c)
		for i:=1;i<x;i++ {
			print(" ")	
		}
		println(c)
	}
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(a)
}

func QuadB(x,y int) {
	a := "/"
	aprime := "\\"
	b := "*"
	c := "*"
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
	for i:=1;i<y-1;i++ {
		print(c)
		for i:=1;i<x;i++ {
			print(" ")	
		}
		println(c)
	}
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
}

func QuadC(x,y int) {
	a := "A"
	aprime := "C"
	b := "B"
	c := "B"
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(a)
	for i:=1;i<y-1;i++ {
		print(c)
		for i:=1;i<x;i++ {
			print(" ")	
		}
		println(c)
	}
	print(aprime)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
}

func QuadD(x,y int) {
	a := "A"
	aprime := "C"
	b := "B"
	c := "B"
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
	for i:=1;i<y-1;i++ {
		print(c)
		for i:=1;i<x;i++ {
			print(" ")	
		}
		println(c)
	}
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
}

func QuadE(x,y int) {
	a := "A"
	aprime := "C"
	b := "B"
	c := "B"
	print(a)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(aprime)
	for i:=1;i<y-1;i++ {
		print(c)
		for i:=1;i<x;i++ {
			print(" ")	
		}
		println(c)
	}
	print(aprime)
	for i:=1;i<x;i++ {
		print(b)
	}
	println(a)
}

func main() {
	QuadA(5,5)
}
