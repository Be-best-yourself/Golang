package main

func main() {
	var name1 string
	var name2 string
	var name3 string
	name1, name2, name3 = "梨子", "叶子", "耗子"
	var name1point *string
	name1point = &name1
	if name1point == &name2 || &name2 == &name3 {
		println("恭喜，三个都不是傻逼，但是不存在的")

	} else {
		println("在座的除了强聪全是傻逼")
	}

}
