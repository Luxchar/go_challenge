package main

func ForEach(f func(int), a []int) {
	for i := range a {
		f(i)
	}
}
