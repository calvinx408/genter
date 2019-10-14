# genter
go

git rebase -i HEAD~4

go 源码学习：https://www.zhihu.com/question/327615791/answer/756625130

网关：https://blog.51cto.com/11976981/2372330

https://github.com/whjstc/openbilibili-go-common-1

类百度文库
https://github.com/TruthHun/DocHub

go  内存分配
https://www.cnblogs.com/qcrao-2018/p/10520785.html#mcache
https://www.toutiao.com/a6719703688031502861/

内存优化之struct对齐
https://www.jianshu.com/p/d314db5b3378

https://my.oschina.net/u/2950272/blog/1829197

有个工具可以检查我忘记了，不知道是不是golangci-lint

关于Golang同一struct中field的书写顺序不同内存分配大小也会不同，也是常说的内存对齐。主要表现如下：struct内field内存分配是以4B为基础，超过4B时必须独占。

示例
type A1 struct {
    a bool
    b uint32
    c bool
    d uint32
    e uint8
    f uint32
    g uint8
}
计算一下A1所需要占用的内存：

首先第1个4B中放入a，a是bool型，占用1B，剩余3B
这时看b是uint32，占用4B，剩余3B放不下，所以offset到下一个4B空间，这时我们会发现3B没有放东西，被浪费了
依次往下，A1要占用28B的空间
根据1，2两个步骤很容易看出，有很多浪费空间。
优化：

type A2 struct {
    a bool
    c bool
    e uint8
    g uint8
    b uint32
    d uint32
    f uint32
}
首先第1个4B中放入a，a是bool型，占用1B，剩余3B
c是bool，占用1B，放入后剩余2B
d是uint8，占用1B，放入后剩余1B
依次往下
这样会使内存使用率高很多。


golangci-lint
https://github.com/golangci/golangci-lint