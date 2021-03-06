package main

import (
	"fmt"
)


// 有一堆桃子，猴子第一天吃了其中的一半，
// 并再多吃了一个!以后每天猴子都吃其中的一半，然后再多吃一个。
// 当到第十天时，想再吃时(还没吃)，发现只有 1 个桃子了

/*
思路分析：
1、到第十天只有一个桃子了。

2、第九天有多少个桃子呢？ 第九天 =  （第十天桃子 + 1）* 2 ，因为第九天吃了一半加一个桃子，
所以就是   *2 +1

3、要求前一天得桃子其实就是下一天桃子数量 + 1 * 2 

4、规律：第 n 天的桃子数 peach(n) = (peach(n+1) + 1)  * 2
*/

// n 的范围是 1 - 10 之间

func peach( n int) int {
	if n > 10 || n < 1 {
		fmt.Println("输入天数不对")
		return 0
	}
	if n == 10 {
		return 1 
	} else {
		return (peach ( n + 1 ) + 1 ) * 2
	}
}

func main()  {
	fmt.Println("输入第九天吃的桃子数量：",peach(7))
}