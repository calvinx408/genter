package main

import (
	"fmt"
	"github.com/smasher164/mem"
	"unsafe"
)

type Slice struct {
	Data unsafe.Pointer //万能指针类型
	len  int            //有效的长度
	cap  int            //有效的容量
}

//指针操作的偏移值
const TAG = 8

func main() {

	var s Slice
	//创建切片
	s.Create(5, 5, 1, 2, 3, 4, 5)
	s.Append(6, 7, 8)
	s.Append(9, 10, 11)
	//打印切片
	s.Print()

	fmt.Println("\n长度：", s.len)
	fmt.Println("容量：", s.cap)

	//根据下标获取元素
	value := s.GetdData(11)
	fmt.Println("值：", value)

	//根据元素获取下标
	index := s.Search(666)
	fmt.Println("下标：", index)

	//删除元素
	s.Delete(5)
	s.Delete(5)
	s.Print()
	fmt.Println()

	//指针偏移
	s.Insert(8, 666)
	s.Print()

	//slice销毁
	fmt.Println(s)
	s.Destroy()
	fmt.Println(s)

	//s1:=[]int{1,2,3,4,5}
	////将一个切片指向地址的值设置为nil  就会被gc回收
	//s1=nil
}

//Create(长度 容量 数据)
func (s *Slice) Create(l int, c int, Data ...int) {
	//如果数据为空返回
	if len(Data) == 0 {
		return
	}
	//长度小于0 容量小于0 长度大于容量 数据大于长度
	if l < 0 || c < 0 || l > c || len(Data) > l {
		return
	}
	//ulonglong unsigned long long  无符号的长长整型
	//开辟空间 存储数据
	//如果堆空间开辟失败 返回值为NULL 相当于nil 内存地址编号为0的空间
	s.Data = mem.Alloc(uint(c) * 8)
	s.len = l
	s.cap = c

	//转成可以计算的指针类型
	p := uintptr(s.Data)
	for _, v := range Data {
		//数据存储
		*(*int)(unsafe.Pointer(p)) = v
		//指针偏移
		p += TAG
		//p+=unsafe.Sizeof(1)
	}
}

//Print 打印切片
func (s *Slice) Print() {
	if s == nil {
		return
	}

	//将万能指针转成可以计算的指针
	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//获取内存中的数据
		fmt.Print(*(*int)(unsafe.Pointer(p)), " ")
		p += TAG
	}

}

//切片追加
func (s *Slice) Append(Data ...int) {
	if s == nil {
		return
	}
	if len(Data) == 0 {
		return
	}

	//如果添加的数据超出了容量
	if s.len+len(Data) > s.cap {
		//扩充容量
		//C.realloc(指针,字节大小)
		s.Data = mem.Alloc(uint(s.cap) * 2 * 8) //C.realloc(s.Data, C.ulonglong(s.cap)*2*8)
		//改变容量的值
		s.cap = s.cap * 2
	}

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//指针偏移
		p += TAG
	}

	//添加数据
	for _, v := range Data {
		*(*int)(unsafe.Pointer(p)) = v
		p += TAG
	}
	//更新有效数据（长度）
	s.len = s.len + len(Data)
}

//获取元素 GetData(下标)  返回值为int 元素
func (s *Slice) GetdData(index int) int {
	if s == nil || s.Data == nil {
		return 0
	}
	if index < 0 || index > s.len-1 {
		return 0
	}

	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	return *(*int)(unsafe.Pointer(p))
}

//查找元素 Search(元素)返回值为int 下标
func (s *Slice) Search(Data int) int {
	if s == nil || s.Data == nil {
		return -1
	}

	p := uintptr(s.Data)
	for i := 0; i < s.len; i++ {
		//查找数据  返回第一次元素出现的位置
		if *(*int)(unsafe.Pointer(p)) == Data {
			return i
		}
		//指针偏移
		p += TAG
	}
	return -1
}

//删除元素 Delete(下标)
func (s *Slice) Delete(index int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.len-1 {
		return
	}

	//将指针指向需要删除的下标位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}

	//用下一个指针对应的值为当前指针对应的值进行赋值
	temp := p
	for i := index; i < s.len; i++ {
		temp += TAG
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(temp))
		p += TAG
	}

	s.len--
}

//插入元素 Insert(下标 元素)
func (s *Slice) Insert(index int, Data int) {
	if s == nil || s.Data == nil {
		return
	}
	if index < 0 || index > s.len-1 {
		return
	}

	//如果插入数据是最后一个
	//if index == s.len-1 {
	//	p := uintptr(s.Data)
	//
	//	//for i := 0; i < s.len; i++ {
	//	//	p += TAG
	//	//}
	//	p += TAG * uintptr(s.len-1)
	//	*(*int)(unsafe.Pointer(p)) = Data
	//	s.len++
	//	return
	//}
	//调用追加方法
	if index == s.len-1 {
		s.Append(Data)
		return
	}

	//获取插入数据的位置
	p := uintptr(s.Data)
	for i := 0; i < index; i++ {
		p += TAG
	}
	//获取末尾的指针位置

	temp := uintptr(s.Data)
	temp += TAG * uintptr(s.len)

	//将后面数据依次向后移动
	for i := s.len; i > index; i-- {
		//用前一个数据为当前数据赋值
		*(*int)(unsafe.Pointer(temp)) = *(*int)(unsafe.Pointer(temp - TAG))
		temp -= TAG
	}

	//修改插入下标的数据
	*(*int)(unsafe.Pointer(p)) = Data
	s.len++
}

//销毁切片
func (s *Slice) Destroy() {
	//调用C语言  适释放堆空间
	mem.Free(s.Data)
	s.Data = nil
	s.len = 0
	s.cap = 0
	s = nil
}
