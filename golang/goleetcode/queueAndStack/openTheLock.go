/*
	你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。
	每个拨轮可以自由旋转：例如把 '9' 变为  '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
	锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
	列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
	字符串 target 代表可以解锁的数字，你需要给出最小的旋转次数，如果无论如何不能解锁，返回 -1
 */

/*
	思路： 典型的广度优先搜索，在一棵n叉树中广度搜索目标数字，搜索从起点"0000"出发，将其一步之内可能出现的"1000","0100","0010","0001"
            "9000","0900","0090","0009"数字加入队列，然后从队列中取出队首元素，将该元素一步之内可能出现的数字再次加入队列，以此类推，
            直到遇到目标数字；需要注意的是，为了避免数字的重复搜索，需要建立一个map字典，该字典以已经搜索过的数字为key，以达到该数字所用的步数为value
			当搜索的数字已经存在与字典时，则不进行任何操作，若搜索的数字不存在于字典时，则将该数字添加进入字典以及队列
 */

package main

import (
	"fmt"
	"strconv"
)

func openLock(deadends []string, target string) int{
	start := "0000"
	// 初始化map字典
	dict := make(map[string]int)
	// 遍历死亡数组，将出现在死亡数字中的数字全部添加进入字典，并初始化value为0
	for _, value := range deadends{
		// 若起点或者目标数字存在于死亡数字，则直接返回-1
		if value == start || value == target{
			return -1
		}
		dict[value] = 0
	}
	// 声明并初始化队列，将起点"0000"添加进入队列
	queue := []string{start}
	// 将起点已添加进入字典，value初始化0
	dict["0000"] = 0
	for len(queue) != 0{
		// 获得队列的队首元素
		cur := queue[0]
		// 删除队列的第一个元素
		queue = queue[1:]
		// 获得到达队首元素数字需要的步数
		step := dict[cur]

		for i := 0; i < 4; i ++{
			var down string
			var up string

			// 将数字串之一的数字进行int类型转换
			curInt, _ := strconv.Atoi(cur[i:i+1])
			// 对转换为int类型的数字进行加1操作，判断该数字是否小于9，若小于9，则直接将原数字进行加1即可，否则将原数字进行改0操作
			if curInt < 9{
				up = cur[0:i] + strconv.Itoa(curInt + 1) + cur[i+1:]
			}else{
				up = cur[0:i] + "0" + cur[i+1:]
			}

			// 对转换为int类型的数字进行减1操作，判断该数字是否大于0，若大于0，则直接将原数字进行减1即可，否则，将原数字进行改9操作
			if curInt > 0{
				down = cur[0:i] + strconv.Itoa(curInt - 1) + cur[i+1:]
			}else{
				down = cur[0:i] + "9" + cur[i+1:]
			}

			// 判断数字串进行加1操作后的数字是否存在于字典，若存在，不进行任何操作，否则将数字添加进入队列以及map字典
			if _, upExist := dict[up]; !upExist{
				queue = append(queue, up)
				dict[up] = step + 1
				// 判断数字是否等于target，若是，直接返回到达该数字所需要的步数
				if up == target{
					return dict[up]
				}
			}

			// 判断数字串进行减1操作后的数字是否存在于字典，若存在，不进行任何操作，否则将数字添加进入队列以及map字典
			if _, downExist := dict[down]; !downExist{
				queue = append(queue, down)
				dict[down] = step + 1
				// 判断数字是否等于target，若是，直接返回到达该数字所需要的步数。
				if down == target{
					return dict[down]
				}
			}
		}
	}
	return -1
}

func main() {
	// 测试用例
	deadends := []string{"0201","0101","0102","1212","2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
