package main

import "runtime/debug"

func main() {

}

//分两种处理
// 尝试 变出最长的F或T
func maxConsecutiveAnswers(answerKey string, k int) int {
	debug.SetGCPercent(-1)
	ans := 0
	var queueT []int
	firstT := 0
	kT := k
	lastTIndex := 0

	var queueF []int
	firstF := 0
	kF := k
	lastFIndex := 0

	for i, c := range answerKey {
		// T
		if c == 'F' {
			if kT == 0 {
				//可变的次数用完 取出队列头 入队
				// 从队列取出一个元素后 下一个T的位置在下一个
				lastTIndex = queueT[firstT] + 1
				// 出对
				firstT++
				// 入队
				queueT = append(queueT, i)
			} else {
				//次数还没用完 直接入队并扣件可变次数
				kT--
				queueT = append(queueT, i)
			}
		}
		if i-lastTIndex+1 > ans {
			ans = i - lastTIndex + 1
		}

		// F
		if c == 'T' {
			if kF == 0 {
				//可变的次数用完 取出队列头 入队
				// 从队列取出一个元素后 下一个T的位置在下一个
				lastFIndex = queueF[firstF] + 1
				// 出对
				firstF++
				// 入队
				queueF = append(queueF, i)
			} else {
				//次数还没用完 直接入队并扣件可变次数
				kF--
				queueF = append(queueF, i)
			}
		}
		if i-lastFIndex+1 > ans {
			ans = i - lastFIndex + 1
		}

	}

	return ans
}

//TTFTFFFTFTTTT
