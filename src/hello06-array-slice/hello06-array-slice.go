package hello06_array_slice

import "fmt"

// 数组切片简单使用（注意：切片后的数组包含源数组的指针,如果切片后数组再被切片且范围扩大（就是再切片的范围扩大到新数组外但在原数组内），这个时候扩大的范围切的就是原数组而不是新数组），详情看下图
// 0,1,2,3,4,5,6,7  原数组（这个时候数组当前开始位置ptr=0，数组的长度len和初始容量cap都是等于8）
//     2,3,4,5      切片2：6后的值（这个时候切片数组相对于原数组当前开始位置ptr=2，数组的当前长度len=4，原数组的容量cap=8）
//     0,1,2,3,4,5  切片2：6后的值的下标（注意：4和5下标对应的是原数组6和7的值）

// 通过上图我们可以知道Go里面的数组you3个标识，ptr数组当前开始位置，len数组当前长度，cap数组初始容量
// 所以在切片的时候我们最大可以切到cap数组初始容量的位置
func sliceTest() {
	a1 := [...]int{1, 2, 3, 4, 5, 6, 76, 7}
	// 获取数组的容量（注意：这个和数组的长度是有区别的，具体请看上图）
	fmt.Println("", cap(a1))
	// 切片（就是第2个开始到第5个结束,不包含第6个）
	fmt.Println("", a1[2:6])
	// 切片（就是第0个到第5个结束，不包含第6个）
	fmt.Println("", a1[:6])
	// 切片（就是第2个到最后一个）
	fmt.Println("", a1[2:])
	// 切片（这个是全部）
	fmt.Println("", a1[:])

	// a2是a1切片后的数组（注意：这里a2 其实是包含a1数组的指针的）
	a2 := a1[2:5]
	// 注意：这个切片范围在a2里面是没有的，所以切的是原数组a1的片
	a3 := a2[1:5]
	fmt.Println(a2, a3)

	// 往数组里面添加元素，这个是在往切片数组里面添加元素（注意：如果添加的元素没有超过原数组的最大容量cap，就是在替换原数组的元素；如果超过了就会产生一个新的原数组）
	a4 := append(a2, 10)

	// 定义一个数组切片（因为没有指定数组长度所以就是数组切片）
	var a5 []int
	// 定义一个有值的数组切片（它也是没有指定数组长度所以就是数组切片）
	a6 := []int{2, 34, 5, 6, 77}
	// 定义一个指定切片长度和数组长度都是16的数组切片（注意：指定切片长度一定要用make函数）
	a7 := make([]int, 16)
	// 定义一个指定切片长度是16，数组容量是32的数组切片
	a8 := make([]int, 16, 32)
	// 复制合并两个数组或数组切片形成一个新的原数组
	a9 := copy(a7, a8)

	// 删除a6数组里面下标为2的元素（原理：先切片a6数组里面0到1的元素，再将a6数组下标为3到最后一个元素，全部添加到前一个切片。注意：...三个点表示拆分数组）
	a10 := append(a6[:2], a6[3:]...)
	// 删除a6数组的第0个元素
	a11 := a6[1:]
	// 删除a6数组的最后一个元素
	a12 := a6[:len(a6)-1]
	fmt.Println("", a4, a5, a6, a7, a8, a9, a10, a11, a12)
}
