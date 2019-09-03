package main

import (
	"fmt"
	"github.com/owenliang/awesomeProject"
)

func main()  {
	fmt.Println("开始攻击...")

	awesomeProject.Destroy("https://www.smzdm.com", 100, 10)
}