# Go unsafe

## 指针
Go 的指针多了一些限制。但这也算是 Go 的成功之处：既可以享受指针带来的便利，又避免了指针的危险性。

限制一：`Go 的指针不能进行数学运算`。

来看一个简单的例子：
```go
a := 5
p := &a
p++
p = &a + 3
```
上面的代码将不能通过编译，会报编译错误：invalid operation，也就是说不能对指针做数学运算。

限制二：`不同类型的指针不能相互转换`。

例如下面这个简短的例子：
```go
func main() {
	a := int(100)
	var f *float64
	
	f = &a
}
// 也会报编译错误：
// cannot use &a (type *int) as type *float64 in assignment
```

限制三：`不同类型的指针不能使用 == 或 != 比较`。

只有在两个指针类型相同或者可以相互转换的情况下，才可以对两者进行比较。另外，指针可以通过 == 和 != 直接和 nil 作比较。

限制四：`不同类型的指针变量不能相互赋值`。

这一点同限制三。

## unsafe
前面所说的指针是类型安全的，但它有很多限制。Go 还有非类型安全的指针，这就是 unsafe 包提供的 unsafe.Pointer。在某些情况下，它会使代码更高效，当然，也更危险。

unsafe 包用于 Go 编译器，在编译阶段使用。从名字就可以看出来，它是不安全的，官方并不建议使用。我在用 unsafe 包的时候会有一种不舒服的感觉，可能这也是语言设计者的意图吧。

但是高阶的 Gopher，怎么能不会使用 unsafe 包呢？它可以绕过 Go 语言的类型系统，直接操作内存。例如，一般我们不能操作一个结构体的未导出成员，但是通过 unsafe 包就能做到。unsafe 包让我可以直接读写内存，还管你什么导出还是未导出。

## unsafe 实现原理
我们来看源码：
```go
type ArbitraryType int

type Pointer *ArbitraryType
```
从命名来看，Arbitrary 是任意的意思，也就是说 Pointer 可以指向任意类型，实际上它类似于 C 语言里的 void*。

unsafe 包还有其他三个函数：
```go
func Sizeof(x ArbitraryType) uintptr
func Offsetof(x ArbitraryType) uintptr
func Alignof(x ArbitraryType) uintptr
```
- Sizeof 返回类型 x 所占据的字节数，但不包含 x 所指向的内容的大小。例如，对于一个指针，函数返回的大小为 8 字节（64位机上），一个 slice 的大小则为 slice header 的大小。
- Offsetof 返回结构体成员在内存中的位置离结构体起始处的字节数，所传参数必须是结构体的成员。
- Alignof 返回 m，m 是指当类型进行内存对齐时，它分配到的内存地址能整除 m。

注意到以上三个函数返回的结果都是 `uintptr` 类型，这和 `unsafe.Pointer` 可以相互转换。三个函数都是在编译期间执行，它们的结果可以直接赋给 const 型变量。另外，因为三个函数执行的结果和操作系统、编译器相关，所以是不可移植的。

unsafe 包提供了 2 点重要的能力：

1. 任何类型的指针和 unsafe.Pointer 可以相互转换。
2. uintptr 类型和 unsafe.Pointer 可以相互转换。

pointer 不能直接进行数学运算，但可以把它转换成 uintptr，对 uintptr 类型进行数学运算，再转换成 pointer 类型。

还有一点要注意的是，uintptr 并没有指针的语义，意思就是 uintptr 所指向的对象会被 gc 无情地回收。而 unsafe.Pointer 有指针语义，可以保护它所指向的对象在“有用”的时候不会被垃圾回收。

## unsafe 如何使用

### 获取 slice 长度
slice header 的结构体定义：
```go
// runtime/slice.go
type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度 
    cap   int // 容量
}
```
调用 make 函数新建一个 slice，底层调用的是 `makeslice` 函数，返回的是 slice 结构体：
`func makeslice(et *_type, len, cap int) slice`
因此我们可以通过 unsafe.Pointer 和 uintptr 进行转换，得到 slice 的字段值。
```go
func main() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8))) 
	fmt.Println(Len, len(s)) // 9 9
	// +uintprt(8)表示unsafe.Pointer所占字节数。所以是得到len字段的地址。

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20
	// uintprt(16) 表示unsafe.Pointer + len 两个字段 所占字节数。所以是得到cap字段的地址。
}
```
Len，cap 的转换流程如下：

Len: &s => pointer => uintptr => pointer => *int => int
Cap: &s => pointer => uintptr => pointer => *int => int

### 获取 map 长度
再来看一下 map：
```go
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}
```
和 slice 不同的是，makemap 函数返回的是 hmap 的指针，注意是指针：
`func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap`

我们依然能通过 unsafe.Pointer 和 uintptr 进行转换，得到 hamp 字段的值，只不过，现在 count 变成二级指针了：
```go
func main() {
	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}
```
count 的转换过程：

&mp => pointer => **int => int


### Offsetof/Sizeof 获取成员偏移量
对于一个结构体，通过 offset 函数可以获取结构体成员的偏移量，进而获取成员的地址，读写该地址的内存，就可以达到改变成员值的目的。

这里有一个内存分配相关的事实：结构体会被分配一块连续的内存，结构体的地址也代表了第一个成员的地址。

我们来看一个例子：
```go
package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	age      int
	language string
}

func main() {
	p := Programmer{"wenran", 37, "go"}
	fmt.Println(p)

	name := (*string)(unsafe.Pointer(&p))
	*name = "zeyuan"

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	*lang = "Golang"
	fmt.Println(p)

	lang = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "C++"
	age := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(string(""))))
	*age = 18

	fmt.Println(p)
}
// 运行代码，输出：
// {wenran 37 go}
// {zeyuan 37 Golang}
// {zeyuan 18 C++}
```
- name 是结构体的第一个成员，因此可以直接将 &p 解析成 *string。对于结构体的私有成员，现在有办法可以通过 unsafe.Pointer 改变它的值了。
- 通过Offsetof() 函数，可以获取成员相对于结构体起始地址的偏移量。
- 通过 unsafe.Sizeof() 函数可以获取成员大小，进而计算出成员的地址，直接修改内存。

### string 和 slice 的相互转换
实现字符串和 bytes 切片之间的转换，要求是 zero-copy。想一下，一般的做法，都需要遍历字符串或 bytes 切片，再挨个赋值。
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	b := []byte{41, 42, 43, 44}
	fmt.Println(BytesToString(b))
	s := "abc"
	fmt.Println(StringToBytes(s))
}

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
	// return unsafe.String(&b[0], len(b))
}
```
直接用unsafe中提供的函数。

## 练习

Go中 uintptr和 unsafe.Pointer 的区别？

答：unsafe.Pointer 是通用指针类型，它不能参与计算，任何类型的指针都可以转化成 unsafe.Pointer，unsafe.Pointer 可以转化成任何类型的指针，uintptr 可以转换为 unsafe.Pointer，unsafe.Pointer 可以转换为 uintptr。uintptr 是指针运算的工具，但是它不能持有指针对象（意思就是它跟指针对象不能互相转换），unsafe.Pointer 是指针对象进行运算（也就是 uintptr）的桥梁。


输出是什么? 答：16，任何字符串都是16
```go
fmt.Printf("string size: %d\n", unsafe.Sizeof("Y4ergesdgergw4ge4saefgsadgf"))
```

## TBD

map 源码中的应用
在 map 源码中，mapaccess1、mapassign、mapdelete 函数中，需要定位 key 的位置，会先对 key 做哈希运算。

例如：

b := (*bmap)(unsafe.Pointer(uintptr(h.buckets) + (hash&m)*uintptr(t.bucketsize)))
h.buckets 是一个 unsafe.Pointer，将它转换成 uintptr，然后加上 (hash&m)*uintptr(t.bucketsize)，二者相加的结果再次转换成 unsafe.Pointer，最后，转换成 bmap 指针，得到 key 所落入的 bucket 位置。如果不熟悉这个公式，可以看看上一篇文章，浅显易懂。

上面举的例子相对简单，来看一个关于赋值的更难一点的例子：

// store new key/value at insert position
if t.indirectkey {
	kmem := newobject(t.key)
	*(*unsafe.Pointer)(insertk) = kmem
	insertk = kmem
}
if t.indirectvalue {
	vmem := newobject(t.elem)
	*(*unsafe.Pointer)(val) = vmem
}

typedmemmove(t.key, insertk, key)
这段代码是在找到了 key 要插入的位置后，进行“赋值”操作。insertk 和 val 分别表示 key 和 value 所要“放置”的地址。如果 t.indirectkey 为真，说明 bucket 中存储的是 key 的指针，因此需要将 insertk 看成指针的指针，这样才能将 bucket 中的相应位置的值设置成指向真实 key 的地址值，也就是说 key 存放的是指针。

下面这张图展示了设置 key 的全部操作：

map assign

obj 是真实的 key 存放的地方。第 4 号图，obj 表示执行完 typedmemmove 函数后，被成功赋值。