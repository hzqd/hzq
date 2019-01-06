package main
import "fmt"
func main() {
	/*本堂课内容：
	一.变量，常量：
		1.变量：
			(1)概念：变量是为存储特定类型的值而提供给内存位置的名称，在Go中声明变量有多种语法
			(2)使用：
				step1：变量的定义，也叫声明（开辟内存）
				step2：变量的访问，赋值或取值
			(3)定义：
				*可以使用类型推断
				*可以使用简短定义
		2.常量：
			(1)概念：常量广义概念是‘不变化的量’ ——> 在计算机程序运行时，不会被程序修改的量
			(2)常数：固定的数值
					100，"abc",3.14
			(3)定义：
					显式类型定义：
						const 常量名 数据类型 = 赋值	//不使用类型推断
					隐式类型定义：
						const 常量名 = 赋值
			(4)常量名的惯用法：
					所有字母都大写 --> 表示公有的常量
					若想表示为私有 --> 一般在常量名前+小写字母c
		3.注意事项：
				(1)Go语言是静态语言，要求变量的类型和赋值的类型一致
				(2)变量名不能冲突 ==> 同一作用域内不能冲突
				(3)简短定义方式： {左边的变量名至少有一个是新的->未被定义过的
								只能定义局部变量，不能定义全局变量}
				(4)变量的零值，也称默认值
	二.数据类型：
		基本数据类型：数值型(整数、浮点、复数)，布尔型(true、false)，字符串
		复合数据类型：数组，切片，map，结构体，指针，接口，函数...
	三.运算符： ...
	*/
/*关于fmt.Print*
	print()
	println() --> print + line
	printf() --> print + format
*/
	a := 10
	fmt.Printf("%T,%d \n",a,a)
	b := 20
	fmt.Printf("%t \n",a > b)
	c := 24.68
	fmt.Printf("%f \n",c)
	var d,e,f=100,200,300
	fmt.Println("d:",d,"e:",e,"f:",f)
	fmt.Printf("d: %d e: %d f: %d\n",d,e,f)
	fmt.Println(c,d,e,f,d<e,e>f)
	var h,i,j,k="你好",5050,3.14159265358979,"まどか"
	fmt.Println(h,i,j,k)
//局部变量：函数内部定义的变量
//全局变量：函数外部，整个作用域
//零值 ——> 默认值：
	var l int ; var m float64 ; var n string ; var o []int //切片slice，引用类型nil
	fmt.Println(l,m,n,o) ; fmt.Println(o == nil)
//空白字符，用于舍弃数值：
	p,q,_ := 1,2,3	//"_"只能写，不能读；用于接收函数的返回值
	fmt.Println(p,q)
//定义常量：
	const PATH string = "https:www.baidu.com"
	const PI = 3.14
//定义一组常量：
	const r,s,t = 666,3.1416,"hi"
	const (
		u = 1.23
		v = 4.56
		w = 7.89
	)
	//常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		r1 byte = 233
		r2
		r3
		r4;r5;r6
	)
	fmt.Print(r6)
//枚举类型：一组相关数值的数据，使用常量组作为枚举类型
	const(
		spring = 1
		summer = 2
		autumn = 3
		winter = 4
	)
	/*iota：特殊常量，可以认为是一个被编译器修改的常量，可以被用作枚举值
		也可以认为是Go底层自动维护的一个计数器
			当使用const定义一个常量组的时候，初始化一个iota，初始值为0
				每当在常量组中定义一个常量，该计数器就累加1
			直到下一个const出现，计数器清零
	 */
	const (
		g = iota
		x = iota
		y = iota
		z = iota
	)
	fmt.Println(z)
}