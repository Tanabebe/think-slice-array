package main

import (
	"fmt"
	"reflect"
)

type A struct {
	array1 [2]int `filed:"structArray1"`
	array2 [4]int `filed:"structArray2"`
	array3 [4]int `filed:"structArray3"`
}

type S struct {
	slice1 []int `filed:"structSlice1"`
	slice2 []int `filed:"structSlice2"`
	slice3 []int `filed:"structSlice3"`
	slice4 []int `filed:"structSlice4"`
	slice5 []int `filed:"structSlice5"`
}

// reflectを使った場合にどうなるかを見てみる
func Result(i interface{}) {

	rv := reflect.ValueOf(i)
	rt := rv.Type()
	fmt.Println("type : ",rt)
	switch i.(type) {
	case S:
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			rf := field.Tag.Get("filed")
			ri := rv.FieldByIndex([]int{i})
			fmt.Printf("%v => length=%d capacity=%d value=%v address=%p \n", rf, ri.Len(), ri.Cap(), ri, &rv)
		}
	case A:
		for i := 0; i < rt.NumField(); i++ {
			field := rt.Field(i)
			rf := field.Tag.Get("filed")
			ri := rv.FieldByIndex([]int{i})
			fmt.Printf("%v => length=%d capacity=%d value=%v address=%p \n", rf, ri.Len(), ri.Cap(), ri, &rv)
		}
	default:
		fmt.Println("該当なし")
	}
}


func main() {
	var target1 = make([]string, 1000)
	fmt.Printf("target1 => length=%d capacity=%d value=%v address=%p \n", len(target1), cap(target1), target1, &target1)

	var target2 = make([]string, 0, 1000)
	fmt.Printf("target2 => length=%d capacity=%d value=%v address=%p \n", len(target2), cap(target2), target2, &target2)

	// 配列の宣言方法
	var array1 [2]int
	fmt.Printf("array1 => length=%d capacity=%d value=%v address=%p \n", len(array1), cap(array1), array1, &array1)

	array2 := [4]int{}
	fmt.Printf("array2 => length=%d capacity=%d value=%v address=%p \n", len(array2), cap(array2), array2, &array2)

	array3 := [4]int{1, 2, 3}
	fmt.Printf("array3 => length=%d capacity=%d value=%v address=%p \n", len(array3), cap(array3), array3, &array3)
	array3[3] = 4
	fmt.Printf("array3 => length=%d capacity=%d value=%v address=%p \n", len(array3), cap(array3), array3, &array3)
	// 配列にappendすることは出来ない
	// array3 = append(array3, 4)
	for _, v := range array3 {
		fmt.Println(v)
	}
	for i := 0; i < len(array3); i++ {
		fmt.Println(array3[i])
	}


	var slice1 []int
	fmt.Printf("slice1 => length=%d capacity=%d value=%v address=%p \n", len(slice1), cap(slice1), slice1, slice1)
	slice2 := []int{1, 2, 3, 4}
	fmt.Printf("slice2 => length=%d capacity=%d value=%v address=%p \n", len(slice2), cap(slice2), slice2, slice2)

	slice3 := make([]int, 10)
	fmt.Printf("slice3 => length=%d capacity=%d value=%v address=%p \n", len(slice3), cap(slice3), slice3, slice3)
	if slice3 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("is not nil")
		fmt.Println(slice3)

	}
	var slice4 []int
	fmt.Printf("slice4 => length=%d capacity=%d value=%v address=%p \n", len(slice4), cap(slice4), slice4, slice4)
	//// appendするとどうなるか
	slice4 = append(slice4, 2)
	fmt.Printf("slice4 => length=%d capacity=%d value=%v address=%p \n", len(slice4), cap(slice4), slice4, slice4)

	// makeを使う
	slice5 := make([]int, 0, 1)
	fmt.Printf("slice5 => length=%d capacity=%d value=%v address=%p \n", len(slice5), cap(slice5), slice5, slice5)
	slice5 = append(slice5, 1)
	// この時点ではcapacityを超えてこないのでアドレスは変わらない
	fmt.Printf("slice5 => length=%d capacity=%d value=%v address=%p \n", len(slice5), cap(slice5), slice5, slice5)
	slice5 = append(slice5, 2)
	// capacityを超えるのでアドレスが変わる
	fmt.Printf("slice5 => length=%d capacity=%d value=%v address=%p \n", len(slice5), cap(slice5), slice5, slice5)
	slice5 = append(slice5, 3)
	// capacityを超えるのでまたアドレスが変わる、capacityは以前確保していたcapacity^2で増加していく
	fmt.Printf("slice5 => length=%d capacity=%d value=%v address=%p \n", len(slice5), cap(slice5), slice5, slice5)

	//// 試しにこんなデータがあったとする
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}
	// sliceのlen, capacityの指定なして、dataを入れていく
	// こうやりたくなる
	var slice6 []int
	// ついでにnilになるか見てみる
	if slice6 == nil {
		fmt.Println("slice 6 is nil")
	}
	for i, v := range data {
		slice6 = append(slice6, v)
		fmt.Printf("slice6 loop => capacity=%d  address=%p value=%v\n", cap(slice6), slice6, slice6[i])
	}
	// もしくはこう
	var slice7 []int

	for i := 0; i < len(data); i++ {
		slice7 = append(slice7, i)
		fmt.Printf("slicslice7e4 loop => capacity=%d  address=%p value=%v\n", cap(slice7), slice7, slice7[i])
	}

	// 上の状態だとcapacityを超えたら元のCapacity^2で倍増していく
	// このやり方だと都度メモリ上に領域を確保するのでNGだと考える
	// なのでこうする。決まりきったCapacityを確保する
	slice8 := make([]int, len(data))

	for i, v := range data {
		slice8[i] = v
		fmt.Printf("slice8 loop => capacity=%d  address=%p value=%v\n", cap(slice8), slice8, slice8[i])
	}

	slice9 := make([]int, len(data))
	// もしくはこう
	for i := 0; i < len(data); i++ {
		slice9[i] = data[i]
		fmt.Printf("slice9 loop => capacity=%d  address=%p value=%v\n", cap(slice9), slice9, slice9[i])
	}

	// sliceについて気になる点を試してみる
	thinkSlice1 := []int{10, 20}

	// capacityが増える前
	fmt.Printf("\nthinkSlice1 => capacity=%d  address=%p value=%v\n", cap(thinkSlice1), thinkSlice1, thinkSlice1)
	tmp := thinkSlice1 // 共有している
	tmp[0] = 1
	fmt.Printf("tmp  => capacity=%d  address=%p value=%v\n", cap(tmp), tmp, tmp)
	fmt.Printf("thinkSlice1 => capacity=%d  address=%p value=%v\n", cap(thinkSlice1), thinkSlice1, thinkSlice1)

	// capacityがここで増える
	thinkSlice1 = append(thinkSlice1, 30)
	// tmpが見ているアドレスは違うため影響が出るはず
	fmt.Printf("tmp  => capacity=%d  address=%p value=%v\n", cap(tmp), tmp, tmp)
	fmt.Printf("thinkSlice1 => capacity=%d  address=%p value=%v\n", cap(thinkSlice1), thinkSlice1, thinkSlice1)

	// なのでここでtmpを変えてもhogeは変更出来ないはず
	tmp[1] = 99
	fmt.Printf("tmp  => capacity=%d  address=%p value=%v\n", cap(tmp), tmp, tmp)
	fmt.Printf("thinkSlice1 => capacity=%d  address=%p value=%v\n", cap(thinkSlice1), thinkSlice1, thinkSlice1)

	hoge2 := make([]int, 3)
	hoge2[0] = 10
	hoge2[1] = 20
	// capacityが増えない場合
	fmt.Printf("\nhoge2 => capacity=%d  address=%p value=%v\n", cap(hoge2), hoge2, hoge2)
	tmp2 := hoge2 // 共有している？
	tmp2[0] = 1
	fmt.Printf("tmp2  => capacity=%d  address=%p value=%v\n", cap(tmp2), tmp2, tmp2)
	fmt.Printf("hoge2 => capacity=%d  address=%p value=%v\n", cap(hoge2), hoge2, hoge2)

	// capacityが増えない時
	hoge2[2] = 30
	// tmp2内とhoge2は同じ
	fmt.Printf("tmp2  => capacity=%d  address=%p value=%v\n", cap(tmp2), tmp2, tmp2)
	fmt.Printf("hoge2 => capacity=%d  address=%p value=%v\n", cap(hoge2), hoge2, hoge2)

	// なのでここでtmpを変えてもhogeは変更出来ないはず
	tmp2[1] = 99
	fmt.Printf("tmp2  => capacity=%d  address=%p value=%v\n", cap(tmp2), tmp2, tmp2)
	fmt.Printf("hoge2 => capacity=%d  address=%p value=%v\n", cap(hoge2), hoge2, hoge2)


	// reflectで配列とスライスを試してみる

	// structに内包したarrayを試す
	a := A{
		[2]int{},
		[4]int{},
		[4]int{1, 2, 3},
	}
	Result(a)

	// structに内包したsliceを試す
	s := S {
		slice2: []int{1, 2, 3, 4},
		slice3: make([]int, 10),
		slice5: make([]int, 0, 1),
	}

	// 拡張していくとどうなるか見てみる
	Result(s)
	s.slice4 = append(s.slice4, 2)
	Result(s)
	s.slice5 = append(s.slice5, 1)
	Result(s)
	s.slice5 = append(s.slice5, 2)
	Result(s)
	s.slice5 = append(s.slice5, 3)
	Result(s)
}
