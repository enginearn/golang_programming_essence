package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())

	// 1. 型の宣言
	// should merge variable declaration with assignment on next line (S1021)
	// var a int
	// a = 1
	// fmt.Println(a)

	// 2. 型の宣言と初期化
	var b int = 2
	fmt.Println(b)

	// 3. 型推論
	c := 3
	fmt.Println(c)

	// 4. 複数の変数を宣言
	var d, e int = 4, 5
	fmt.Println(d, e)

	// 5. 複数の変数を宣言
	var (
		f int = 6
		g     = 7
	)
	fmt.Println(f, g)

	// 6. 複数の変数を宣言
	h, i := 8, 9
	fmt.Println(h, i)

	// 7. 複数の変数を宣言
	var (
		j, k int = 10, 11
		l, m     = 12, 13
	)
	fmt.Println(j, k, l, m)

	// 8. 複数の変数を宣言
	n, o := 14, 15
	p, q := 16, 17
	fmt.Println(n, o, p, q)

	// 9. 複数の変数を宣言
	var (
		r, s int = 18, 19
		t, u     = 20, 21
		v, w     = 22, 23
	)
	fmt.Println(r, s, t, u, v, w)

	// 10. 複数の変数を宣言
	x, y := 24, 25
	z, aa := 26, 27
	ab, ac := 28, 29
	fmt.Println(x, y, z, aa, ab, ac)

	// 11. キャスト
	var ad int = 30
	var ae float64 = float64(ad)
	fmt.Println(ae)

	// 12. キャスト
	af := 31
	ag := 2.0
	fmt.Println(float64(af) + ag)

	// 13. iota
	const (
		ai = iota
		aj
		ak
	)
	fmt.Println(ai, aj, ak)

	// 14. iota
	const (
		Apple = iota + iota // 0 + 0
		Banana 			// 1 + 1
		Cherry 			// 2 + 2
	)
	fmt.Println(Apple, Banana, Cherry)

	// 15. iota
	const (
		_ = iota // 0
		Almond   // 1
		_        // 2
		Apricot = iota // 3
	)
	fmt.Println(Almond, Apricot)

	// 16. iota
	type Fruit int
	type Animal int

	const (
		Apple2 Fruit = iota // Fruit型のApple Fruit(0)
		Banana2             // Fruit型のBanana Fruit(1)
		Cherry2             // Fruit型のCherry Fruit(2)
	)

	const (
		Dog Animal = iota // Animal型のDog Animal(0)
		Cat                // Animal型のCat Animal(1)
	)
	fmt.Println(Apple2, Banana2, Cherry2)
	fmt.Println(Dog, Cat)
}
