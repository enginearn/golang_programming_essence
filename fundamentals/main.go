package main

import (
	"fmt"
	"runtime"
	"log"
	"os"
	"sync"

	// "golang.org/x/text/encoding/japanese"
)

func server(ch chan string) {
	defer close(ch)
	ch <- "hello"
	ch <- "world"
	ch <- "hoge"
}

// func FindUser(name string)(*User, error) {
// 	user, err := findUserFromList(users, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// User := struct {
// 	Name string
// 	Age int
// }

// func findUserFromList(users []Users, name string)(*User, error) {
// 	for _, user := range users {
// 		if user.Name == name {
// 			return &user, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("User not found: %s", name)
// }

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

	// 17. function call
	// users := []User{
	// 	{Name: "Bob", Age: 20},
	// 	{Name: "John", Age: 21},
	// }
	// user, _ := findUserFromList(users, "Bob")
	// user, err := FindUser("Bob")
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// 18. conditions
	if x == 24 {
		fmt.Println("x is 24")
	}

	// 19. conditions
	// if user, err := FindUser("Bob"); err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(user)
	// }

	// 20. conditions
	switch x {
	case 24:
		fmt.Println("x is 24")
	case 25:
		fmt.Println("x is 25")
	default:
		fmt.Println("x is not 24 nor 25")
	}

	// 21. loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 22. loop
	ii := 0
	for ii < 10 {
		fmt.Println(ii)
		ii++
	}

	// 23. map
	mp := map[string]int{"x": 10, "y": 20}
	fmt.Println(mp)

	// 24. slice
	for i, v := range []int{1, 2, 3} {
		fmt.Println(i, v)
	}

	// 25. label
	Loop:
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j == 2 {
					break Loop
				}
				fmt.Println(i, j)
			}
		}

	// 26. defer
	defer fmt.Println("defer")
	fmt.Println("not defer")

	// 27. defer
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("not defer")

	// 28. make
	ss := make([]int, 3)
	ss[0] = 1
	ss[1] = 2
	ss[2] = 3
	fmt.Println(ss)

	// 29. make
	mp2 := make(map[string]int)
	mp2["x"] = 10
	mp2["y"] = 20
	fmt.Println(mp2)

	// 30. make
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// }
	// fmt.Println(<-ch)

	// 31. dimensions
	var arr0 [2][3]int
	arr0[0][0] = 1
	arr0[0][1] = 2
	arr0[0][2] = 3
	arr0[1][0] = 4
	arr0[1][1] = 5
	arr0[1][2] = 6
	fmt.Println(arr0)

	// 32. dimensions
	arr1 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr1)

	// 33. dimensions slice
	arr2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(arr2)

	// 34. slice append
	ss2 := []int{1, 2, 3}
	ss2 = append(ss2, 4, 5, 6)
	fmt.Println(ss2)
	fmt.Printf("ss2: %p\n長さ: %d\n", &ss2, len(ss2))

	// 35. slice append
	a2 := make([]int, 0, len(ss2)/2)
	for i := 0; i < len(ss2); i++ {
		if i % 2 == 0 {
			a2 = append(a2, ss2[i])
		}
	}
	ss2 = a2
	fmt.Println(ss2)

	// 36. slice append
	// nn := 50
	// var a []int = make([]int, 0, nn)
	// a = append(a[:n], a[nn+1:]...)
	// fmt.Println(a)

	// 37. slice append
	// a = a[:n+copy(a[n:], a[nn+1:])]

	// 38. string
	str0 := "Hello, World"
	fmt.Println(str0)

	// 39. string
	str1 := "Hello "
	str1 += "World"
	fmt.Println(str1)

	// 40. string
	str2 := "Hello"
	fmt.Printf("%c", str2[0]) // H
	fmt.Printf("%c", str2[1]) // e
	fmt.Printf("%c", str2[2]) // l
	fmt.Printf("%c", str2[3]) // l
	fmt.Printf("%c", str2[4]) // o
	fmt.Println()

	// 41. string
	str3 := "Hello"
	bb := []byte(str3) // byte型のスライスに変換
	bb[0] = 'h' // byte型のスライスの要素を変更
	str3 = string(bb) // byte型のスライスをstring型に変換
	fmt.Println(str3)

	// 42. string
	str4 := "こんにちわ。世界。"
	fmt.Println(str4)
	rs := []rune(str4) // rune型のスライスに変換
	rs[4] = 'は' // rune型のスライスの要素を変更
	str4 = string(rs) // rune型のスライスをstring型に変換
	fmt.Println(str4)

	// 43. string
	{
		var b [5]byte
		b2 := b[:]
		b2[0] = 'H'
		b2[1] = 'e'
		b2[2] = 'l'
		b2[3] = 'l'
		b2[4] = 'o'
		fmt.Println(string(b2))
	}

	// 44. string
	var contents = `複数行の
	文章からなる
	テキストを
	扱うことが
	できます。
	`
	fmt.Println(contents)

	// 45. file open
	{
		f, err := os.Open("data.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		var b[512]byte
		n, err := f.Read(b[:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b[:n]))
	}

	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")

	// 46. sync
	{
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println(i)
			}(i)
		}
		wg.Wait()
	}

	fmt.Printf("========================================\n")
	// 47. sync
	{
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			v := i // この行を追加
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				fmt.Println(v)
			}(i)
		}
		wg.Wait()
	}

	// 48. sync
	// {
	// 	for _, tt := range tests {
	// 		wg.Add(1)
	// 		go func(tt *Test) {
	// 			defer wg.Done()
	// 			fmt.Println(tt.name)
	// 		}(&tt)
	// 	}
	// 	wg.Wait()
	// }

	println("========================================\n")

	// 49. sync
	{
		n := 0
		var mu sync.Mutex

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
		}()

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mu.Lock()
				n++
				mu.Unlock()
			}
		}()

		wg.Wait()
		fmt.Println(n)
	}

	// 50. channel
	{
		var s string
		ch := make(chan string)
		go server(ch)

		s = <-ch
		fmt.Println(s)

		s = <-ch
		fmt.Println(s)

		s = <-ch
		fmt.Println(s)
	}

}
