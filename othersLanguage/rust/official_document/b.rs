fn main(){
//通常情况下，‘{}’会被任意变量内容所替换
//值内容会转化成字符串
	println!("{} days",31);
//不加后缀的话，31自动成为I32类型
//你可以添加后缀来改变31的类型

//下面有多种可选形式
//可以使用的位置参数
	println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

//可以使用赋值语句
	println!("{subject} {verb} {object}.",
		object="the lazy dog",
		subject="The quick brown fox",
		verb="jumps over");

//特殊的格式实现可以在后面加上`:`符号
	println!("{} of {:b} people know binary, the other half don't", 1, 2);

//println！会检查使用到的参数数量是否正确
	println!("My name is {0}, {0} {1}", "James", "Bond");

//创建一个包含`I32`类型结构体（structure）
	#[allow(dead_code)]
	struct Structure(i32);	
}
