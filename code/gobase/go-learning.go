package main

import (
	"fmt"
	"math"
	//"reflect"
	"runtime"
	"time"
	"sync"
)

func add_(x int, y int) int {
	return x + y
}

func swap_(x, y string) (string, string) {
	return y, x
}

func named_value_return(tmp int) (a, b int) {
	a = tmp * 3
	b = tmp / 3
	return
}

// basic data type
var (
	isOk  bool
	myStr string
	by    byte
	n_32  rune
	n_64  int64
	n     int
	un    uint
	f_32  float32
	f_64  float64
	c_64  complex64
	c_128 complex128
)

// something about variable
var test bool
var str string

func variable() {
	var i, j int = 1, 2
	//i = 4
	var right, yes = true, "no"
	fmt.Println(i, j, test, str, right, yes)
}

/*
func type_conversion() {
	var i int = 4
	var f float32
	var s string

	// method 1
	var ff float32 = float32(i)
	var ss string = string(i)

	// method 2
	ii := 5
	tt := float32(ii)

	// all of these method are all the obvious conversion
}
*/

func constant() {
	const Pi = 3.1415926
	const word = "hello word"
	const ch = 'c'
	const b = false
}

// something about loop

func sum_to_n_for(n int) {
	var sum int
	for i := 0; i < n; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func sum_to_n_while(n int) int {
	i := 1
	sum := 0
	for i < n {
		sum += i
	}
	return sum
}

func loop_if(n int) {
	var i int
	if i <= n {
		fmt.Println(math.Exp(float64(n)))
	} else {
		fmt.Printf("i > n")
	}
	if j := 3; j < n {
		i++
	}
	fmt.Println(i)
}

// switch
func platform() {
	fmt.Print("Go on ")
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("GNU/Linux")
	default:
		fmt.Printf("%s", os)
	}

	// switch without condition just like clearable if-else
	var i int
	switch {
	case i < 4:
		fmt.Printf("i < 4")
	case i < 8:
		fmt.Printf("i < 8")
	default:
		fmt.Printf("None of above")
	}
}

/*
func test_func() {
	//fmt.Println(add_(1, 1))
	// short variable method := only working in the func, out of the func you have to use =
	a := "hello"
	b := "john"
	fmt.Println(a, b)
	a, b = swap_(a, b)
	fmt.Println(a, b)
	fmt.Println(named_value_return(15))
	variable()
}*/

func defer_test() {
	// the order of defer execute just like C++ the order of desconstructor.
	// and the defer will parse the argument by real time
	var j int = 1
	defer fmt.Println("--> ", func() int { return j * j }()) // ouput 1 rather than 4
	j++

	defer fmt.Printf("!!")
	defer fmt.Printf("Word")
	fmt.Printf("Hello ")

	fmt.Println("Starting...")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Printf("All done! ")
}

func pointers() {
	var ptr *int // empty pointer's value is nil
	fmt.Println(ptr)

	// pointer in Go not include pointer calculation
	i, j := 10, 100
	fmt.Printf("i = %v\n", i)
	fmt.Printf("j = %v\n", j)
	p := &i
	fmt.Printf("After p := &i, *p = %v\n", *p)
	*p = 1000
	fmt.Printf("After *p = 1000, i = %v\n", i)
	p = &j
	fmt.Printf("After p = &j, j = %v\n", j)
}

// something about struct
type Test struct {
	// variable must use upper letter
	X int
	Y float64
	Z bool
}

func struct_test() {
	fmt.Println(Test{1, 2.00089, true})
	T := Test{1, 2.00089, true}
	fmt.Println(T.X, T.Y, T.Z)
	T.X = 0
	T.Y = 2.01
	T.Z = false
	fmt.Println(T.X, T.Y, T.Z)

	// struct pointer
	p := &T
	p.X = 1000
	fmt.Println(T.X, T.Y, T.Z)

	// struct syntax
	var (
		T3 = Test{X: 200000, Y: 1.00000001, Z: true}
		T1 = Test{X: 20}
		T2 = Test{X: 2000, Y: 1.00001}
		T4 = Test{}
		T5 = &Test{}
	)
	fmt.Println(T1, T2, T3, T4, T5)
}

func array() {
	var num [3]int // the number 3(the length) is a part of array num
	// var num [3]int{0, 1, 2}  // remember the initialization is {} not ()
	for i := 0; i < 3; i++ {
		num[i] = i
	}
	fmt.Println(num)                    // output a array num [0 1 2]
	fmt.Println(num[0], num[1], num[2]) // just output single value like 0, 1, 2

	var str [2]string
	str[0] = "hello"
	str[1] = "Go word!"
	fmt.Println(str)
	fmt.Println(str[0], str[1])

	tmp := [6]int{1, 2, 3, 4, 5}
	fmt.Println(tmp)
}

func slice_test() {
	// slice just a reference of array, so when we modified some value in the slice and the value of array will be modified too.
	var array [10]int
	for i := 1; i < 10; i += 2 {
		array[i] = i
	}
	fmt.Println(array)
	var s []int = array[1:4]
	// s := array[1:4]
	fmt.Println(s)

	str := [3]string{"h3ll0", "G0", "w0rd"}
	s1 := str[0:2]
	s2 := str[1:]
	// this means :
	//		  0      1
	// s1	h3ll0	G0
	// s2           G0	 w0rd
	//				 0    1
	fmt.Println(s1)
	fmt.Println(s2)
	s1[0] = "hello"
	s2[1] = "word"
	// this means :
	//        0     1
	// s1   hello   G0
	// s2           G0   word
	//				0      1
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(str)
	// so str show this:
	// str  hello   G0   word
}

func slice_literals() {
	flag := []bool{true, false, false, true}
	fmt.Println(flag)

	str := []string{"hello", "word"}
	fmt.Println(str)

	num := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(num)
	num = append(num, 7, 8, 9)
	fmt.Println(num)

	s1 := num[:]
	fmt.Printf("len(s1) = %v, cap(s1) = %v, s1 = %v\n", len(s1), cap(s1), s1)

	s2 := num[0:0]
	fmt.Printf("len(s2) = %v, cap(s2) = %v, s2 = %v\n", len(s2), cap(s2), s2)

	s3 := num[1:5]
	fmt.Printf("len(s3) = %v, cap(s3) = %v, s3 = %v\n", len(s3), cap(s3), s3)

	s4 := num[4:6]
	fmt.Printf("len(s4) = %v, cap(s4) = %v, s4 = %v\n", len(s4), cap(s4), s4)

	//d_2 := [][]int{1, 2, 3, 4, 5, 6, 8}
	//d_ := make([]string, "0")
	//fmt.Println(d_)

	// nil slice
	var nil_slice []int
	fmt.Printf("nil_slice = %v, len(nil_slice) = %v, cap(nil_slice) = %v", nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("\nError here!")
	}

	// use make to create dynamic array
	new_slice := make([]int, 2, 5) // if you not specific third argument, it will not show you cap of this slice
	//new_slice := make([]int, 2) // it will show you [0 0]
	fmt.Println(new_slice, len(new_slice), cap(new_slice))

	new_slice1 := new_slice[0:1]
	fmt.Println(new_slice1)

	// slice of slice
	ss := [][]string{
		[]string{"slice", "of", "slice"},
		[]string{"slice", "of", "slice"},
		[]string{"slice", "of", "slice"},
		[]string{"slice", "of", "slice"},
	}
	fmt.Println(ss, len(ss), cap(ss))

	ss[0][0] = "X"
	ss[0][1] = "X"
	ss[1][0] = "X"
	ss[1][1] = "X"
	fmt.Println(ss, len(ss), cap(ss))

	// range for
	for i, v := range num {
		fmt.Printf("i = %d, v = %d\n", i, v)
	}
	// if you just want one argument in range, just ignore the another one,  or you can use _ to subsititude it
	for i, _ := range num {
		fmt.Println("i = ", i)
	}
	for _, v := range num {
		fmt.Println("v = ", v)
	}
	//for i := range num {}
}

type T struct {
	i int
	b bool
}

// map:
// map[key_type]value_type
func maptaion() {
	var m map[string]T
	m = make(map[string]T)
	m["hello"] = T{
		1, false,
	}
	fmt.Println("Before modified, m = ", m)
	fmt.Println(m["hello"])
	m["hello"] = T{
		0, true,
	}
	fmt.Println("After modified, m = ", m)
	fmt.Println(m["hello"])

	var m1 = map[int]string{
		1: "hello",
		2: "word",
	}
	fmt.Println(m1)

	/*
		v, ok := m1["hello"]
		if v == nil {
			fmt.Println("Error!\n")
		} else {
			fmt.Println("v = \n", v)
		}*/

	var m2 = map[int]T{
		0: {1, true},
	}
	fmt.Println(m2)
}

// the function just a value, but it's hard to comperhension for me
func functor(fn func(int, int) int) int {
	return fn(3, 4) * 3
}

func closure() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type P struct {
	X, Y int
}

// method
// method is function, and belows equivalent this: func Product(p P) int {}
// cannot declare a method for the builtin type in Go and position in the different package
func (p P) Product() int {
	return p.X * p.Y
}

func (p *P) Interaction() {
	p.X = p.X + p.Y
	p.Y = p.X
}

// equivalent Interaction() above
//func Interaction(p *P) {}

type I interface {
	M()
}

type T_ struct {
	S string
}

func (t_ T_) M() {
	fmt.Println(t_.S)
}

// Goruntine
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Go channel
func sum_chan(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func chan_buffer() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	//ch <- 3
	//ch <- 4
	fmt.Println(ch)
	fmt.Println(ch)
	//fmt.Println(ch)
	//fmt.Println(ch)
}

func chan_close(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch) // only you ensure you needn't transfer data to chan then close chan by sender
}

func select_test(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func select_default() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		// from here I just have one question: tick and boom not a chan type, why did they can use <- ?
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// struct go tag
// tag need reflect to read
// of course key:value also can represent tag
// Tag always use on marshaling
// Ref: https://juejin.im/post/5b338516f265da596e4ceb9e
type Tag struct {
	f1 string "f one"
	f2 string `two:"2"` // you can use Lookup and Get to fetch the value of this key-value pair
	f3 string `f three`
}

// callback function
func callback(i int, f func(int, int)) {
	f(i, 2)
}

func pro(i, j int) {
	fmt.Println("i * j = ", i*j)
}

// unnamed fucntion
//func(x, y int) int {
//	return x + y
//}

/*
func(x, y int) int {
    return x + y
}() // () represent call this function directionly
*/

/*
// calculate from 1 to 1000000
func() {
    sum := 0
    for i := 1; i <= 1e6; i++ {
	sum += i
    }
}()
*/

// more details about go struct please visit https://juejin.im/post/5ca2f37ce51d4502a27f0539
type M struct {
	Name string
	Age  int
}

// Mutex Lock & Semaphore
// Semaphore
var (
	sema := make(chan struct{}, 1)
	balance int
)

func Deposit-sema(amount int) {
	sema<-struct{}{}
	balance += amount
	<-sema
}

func Balance-sema() int {
	sema<-struct{}{}
	b := balance
	<-sema
	return b
}

// Mutex Lock
var (
	mu sync.Mutex
	balance_ int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	amount += balance_
}

func Balance() int {
	mu.Lock()
	//b := balance_
	//mu.Unlock()
	//return b
	defer mu.Unlock()
	return balance_
}



func main() {
	var m1 *M
	m1 = &M{}
	m1.Name = "Lisa"
	fmt.Println("m1.Name = ", m1.Name)

	var m2 = new(M)
	m2.Age = 20
	fmt.Println("m2.Age = ", m2.Age)

	/*
		// unnamed func
		fv := func(x, y int) int {
			fmt.Println(x + y)
			return x + y
		}
		fv(3, 4)
		/*
			//callback
			callback(1, pro)
			/*
			/*
				// tag & reflect
				t := reflect.TypeOf(Tag{})
				f1, _ := t.FieldByName("f1")
				fmt.Printf("%q\n", f1.Tag)
				f2, _ := t.FieldByName("f2")
				fmt.Println(f2.Tag)
				v, ok := f2.Tag.Lookup("two")
				fmt.Printf("%s, %t\n", v, ok)
				f3, _ := t.FieldByName("f3")
				fmt.Printf("%q\n", f3.Tag)
				/*
				// chan
				//select_default()
				/*
					c := make(chan int)
					quit := make(chan int)
					go func() {
						for i := 0; i < 10; i++ {
							fmt.Println(<-c)
						}
						quit <- 0
					}()
					select_test(c, quit)
	*/

	/*

		ch := make(chan int, 10)
		go chan_close(cap(ch), ch)
		for i := range ch {
			fmt.Println(i)
		}*/
	//chan_buffer()
	/*
		s := []int{1, 2, 3, 4, 5, 6, 7}
		c := make(chan int)
		go sum_chan(s[:len(s)/2], c)
		go sum_chan(s[len(s)/2:], c)
		x, y := <-c, <-c
		fmt.Println(x, y, x+y)
	*/
	/*
		go say("word") // the result of every exec is different
			say("hello")
			/*
				var i interface{} = "hello"
				//t := i.(T)
				t, ok := i.(string)
				fmt.Printf("%v %T\n", t, ok)
				/*
				//interface_test
				/*var i I = T_{"interface"}
				i.M()

					p := P{2, 6}
					p.Interaction()
					fmt.Println(p.Product())
					/*
					//closure
					/*
						a, b := closure(), closure()
						for i := 0; i < 10; i++ {
							fmt.Println(
								a(i),
								b(-i),
							)
						}*/
	/* function
	functor_ := func(x, y int) int {
		return x * y
	}
	fmt.Println(functor_(1, 2))
	fmt.Println(functor(functor_))
	*/
	//maptaion()
	//slice_literals()
	/*
		com := []struct {
			i int
			f float64
			b bool
			s string
		}{
			{1, 1.0, true, "1"},
			{0, 0.0, false, "0"},
		}
		fmt.Println(com)
	*/
	//slice_test()
	//array()
	//sum_to_n(10)
	//loop_if(9)
	//platform()
	//defer_test()
	//pointers()
	//struct_test()
}
