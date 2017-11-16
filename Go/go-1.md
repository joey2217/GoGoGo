# Hello Go
这就像一个传统，在学习大部分语言之前，你先学会如何编写一个可以输出hello world的程序。

    helloworld.go

    package main
    import "fmt"
    func main() {
	fmt.Println("Hello world!")
    }

运行程序或者在对应目录下使用命令行

    go run hellowrold.go

输出:Hello world!

## 解析
1.Go程序是通过`package`来组织的，一般的Go程序中`package <pkgName>`这一行告诉我们当前文件属于哪个包，而包名`main`则告诉我们它是一个可独立运行的包，它在编译后会产生可执行文件。除了`main`包之外，其它的包最后都会生成`*.a`文件（也就是包文件）。每一个可独立运行的Go程序，必定包含一个`package main`，在这个`main`包中必定包含一个入口函数`main`，而这个函数既没有参数，也没有返回值。
2.import ：导入需要的包
3.func main()：func 关键字定义main()函数
4.fmt.Println():使用导入的fmt包中的Println()输出