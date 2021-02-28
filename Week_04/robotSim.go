// 模拟机器人
package main

// robotSim ...
// 记录障碍点为map
func robotSim(commands []int, obstacles [][]int) int {
    type p struct {
        x int
        y int
    }
    res := p{0, 0}
    maxRes := 0
    dire := 0
    obsMap := make(map[p]bool, len(obstacles))
    for i := 0; i< len(obstacles); i++ {
        tmp := p{obstacles[i][0], obstacles[i][1]}
        obsMap[tmp] = true
    }
    for _, v := range commands {
        // 步数行走
        if v > 0 {
            // 一步一步走，每步判断是否障碍点
            for i := 1; i <= v; i++ {
                // 临时点
                tmp := res
                // 结果往前走
                switch dire {
                case 3: res.x -= 1
                case 0: res.y += 1
                case 1: res.x += 1
                case 2: res.y -= 1
                }
                // 如果结果遇到障碍，则结果回到临时点
                if _, ok := obsMap[res]; ok {
                    res = tmp
                    break
                } else { // 没遇到则计算
                    if (res.x*res.x + res.y*res.y) > maxRes {
                        maxRes = res.x*res.x + res.y*res.y
                    }
                }
            }
        }
        // 处理方向，0北，1东，2南，3西，取余简单循环
        if v == -1 {
            dire = (dire + 1) % 4
        }
        if v == -2 {
            dire = (dire + 4 - 1) % 4
        }
    }
    return maxRes
}

func main() {
    
}