package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

//  sh -x run.sh  /Users/mac/project/go-quick-start/cmd/pac_man/game.txt /Users/mac/project/go-quick-start/cmd/pac_man/result.txt
func main() {
	readMapPath := os.Args[1]
	writePath := os.Args[2]
	gameInfoByte, err := ioutil.ReadFile(readMapPath)
	if err != nil {
		fmt.Println("Err:: ReadFile ", err)
		randomWriteResult(writePath)
		return
	}
	var gameInfo GameInfo
	var result ResultStep
	err = json.Unmarshal(gameInfoByte, &gameInfo)
	if err != nil {
		fmt.Println("Err:: gameInfo Unmarshal fail", err)
		randomWriteResult(writePath)
		return
	}
	result.Pacman = GetPacmanPath(gameInfo)
	result.Ghost = GetGhostPath(gameInfo)
	writeResult(writePath, result)
	return
}

func GetPacmanPath(gameInfo GameInfo) int {
	bigPacmanIndex := FindTwos(gameInfo.Map)
	tmpPath := 9999
	tmpMove := 0
	if len(gameInfo.Bonus) > 0 { // 吃奖励逻辑
		for _, value := range gameInfo.Bonus {
			move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Pacman.X, gameInfo.Owner.Pacman.Y, value.X, value.Y)
			if path < tmpPath {
				tmpPath = path
				tmpMove = move
			}
		}
		return tmpMove
	} else if len(bigPacmanIndex) > 0 { // 去吃大豆子

		// 优先级：吃大豆 > 吃虚弱状态的幽灵
		for _, value := range bigPacmanIndex {
			move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Pacman.X, gameInfo.Owner.Pacman.Y, value[0], value[1])
			if path < tmpPath {
				tmpPath = path
				tmpMove = move
			}
		}
		for _, player := range gameInfo.Group {
			if player.Ghost[0].S == 2 {
				move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Pacman.X, gameInfo.Owner.Pacman.Y, player.Ghost[0].X, player.Ghost[0].Y)
				if path < tmpPath {
					tmpPath = path
					tmpMove = move
				}
			}
			if player.Ghost[1].S == 2 {
				move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Pacman.X, gameInfo.Owner.Pacman.Y, player.Ghost[1].X, player.Ghost[1].Y)
				if path < tmpPath {
					tmpPath = path
					tmpMove = move
				}
			}
		}
		return tmpMove
	}
	return tmpMove
}

func GetGhostPath(gameInfo GameInfo) []int {
	// 幽灵人没有虚弱
	if gameInfo.Owner.Ghost[0].S != 2 {
		tmp1Path := 9999
		tmp1Move := Stay
		tmp1PlayerX := 0
		tmp1PlayerY := 0
		for _, player := range gameInfo.Group {
			move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Ghost[0].X, gameInfo.Owner.Ghost[0].Y, player.Pacman.X, player.Pacman.Y)
			if path < tmp1Path {
				tmp1Path = path
				tmp1Move = move
				tmp1PlayerX = player.Pacman.X
				tmp1PlayerY = player.Pacman.Y
			}
		}
		if tmp1Move == Stay {
			tmp1Move = rand.Intn(4)
		}

		tmp2Path := 9999
		tmp2Move := Stay
		for _, player := range gameInfo.Group {
			if tmp1PlayerX != player.Pacman.X && tmp1PlayerY != gameInfo.Owner.Ghost[1].Y {
				move, path := DijkstraFindPath(gameInfo.Map, gameInfo.Owner.Ghost[1].X, gameInfo.Owner.Ghost[1].Y, player.Pacman.X, player.Pacman.Y)
				if path < tmp2Path {
					tmp2Path = path
					tmp2Move = move
				}
			}
		}
		if tmp2Move == Stay {
			tmp2Move = rand.Intn(4)
		}
		return []int{tmp1Move, tmp2Move}
	} else {
		return []int{rand.Intn(4), rand.Intn(4)}
	}
}

func randomWriteResult(path string) {
	result := ResultStep{}
	rand.Seed(time.Now().UnixNano())
	result.Pacman = rand.Intn(4)
	result.SetGhostStep(rand.Intn(4), rand.Intn(4))
	writeResult(path, result)
}

func writeResult(path string, result ResultStep) {
	resultByte, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Err:: result Marshal fail", err)
		return
	}
	if err := ioutil.WriteFile(path, resultByte, 0666); err != nil {
		fmt.Println("Err:: write file fail", err)
	}
}

// -------- model ---------

type GameInfo struct {
	MatchId string       `json:"match_id"`
	Step    int          `json:"step"`  // 当前多少步，从0-199
	Bonus   []BonusInfo  `json:"bonus"` // 奖励位置
	Map     [][]int      `json:"map"`   // 地图信息，采用26 * 26的信息，0表示墙，1表示小豆，2表示大豆，3表示没有吃的
	Owner   PlayerInfo   `json:"owner"` // 自己所在阵营
	Group   []PlayerInfo `json:"group"` // 其他阵营信息
}

type PlayerInfo struct {
	Id             string      // 队伍名
	Score          int         // 得分
	EatenPacmanCnt int         // 小幽灵吃其他阵营吃豆人数量
	EatenGhostCnt  int         // 吃豆人吃其他阵营小幽灵数量
	EatenBonusCnt  int         // 吃豆人吃奖励数量
	Pacman         PacmanInfo  // 吃豆人信息
	Ghost          []GhostInfo // 小幽灵信息
	Idx            int         // 行动顺序，0表示第一个行动
}

type PacmanInfo struct {
	X int // x坐标，从0开始，到25
	Y int // y坐标，从0开始，到25
	S int // 状态，0表示正常，2表示死亡
	R int // 无敌状态剩余回合数
}

type GhostInfo struct {
	X int // x坐标，从0开始，到25
	Y int // y坐标，从0开始，到25
	S int // 状态，0表示正常，1表示死亡，2表示虚弱状态
	D int // 如果处于虚弱状态，剩余多少回合
	C int // 如果处于虚弱状态，本回合是否可以动，1表示可以，0表示不可以
}

type BonusInfo struct {
	X int `json:"x"` // x坐标，从0到25
	Y int `json:"y"` // y坐标，从0到25
	T int `json:"t"` // 奖励展示类型，可忽略
}

type ResultStep struct {
	Pacman int   `json:"pacman"` // 吃豆人移动
	Ghost  []int `json:"ghost"`  // 小幽灵移动
} // 0 表示向左移动，1 表示向上移动，2 表示向右移动，3 表示向下移动，其他数字表示不动。

func (self *ResultStep) SetGhostStep(ghost1, ghost2 int) {
	self.Ghost = append(self.Ghost, ghost1)
	self.Ghost = append(self.Ghost, ghost2)
}

//  -------------------

func FindTwos(arr [][]int) [][]int {
	var twos [][]int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 2 {
				twos = append(twos, []int{i, j})
			}
		}
	}

	return twos
}

// ------------ djikstra

const (
	Left int = iota
	Up
	Right
	Down
	Stay
)

type Point struct {
	X, Y int
	Cost int
}

func getNextStep(targetX, targetY, x, y int) int {
	if targetX == x-1 && y == targetY {
		return Up
	} else if targetX == x+1 && y == targetY {
		return Down
	} else if targetX == x && y-1 == targetY {
		return Left
	} else {
		return Right
	}
}

func DijkstraFindPath(grid [][]int, x, y, targetX, targetY int) (int, int) {
	start := Point{x, y, 0}
	target := Point{targetX, targetY, 0}

	shortestPath := findShortestPath(grid, start, target)
	if shortestPath != nil {
		fmt.Println("最短路径：")
		for _, p := range shortestPath {
			fmt.Printf("(%d, %d) -> ", p.X, p.Y)
		}
		if len(shortestPath) > 1 {
			return getNextStep(shortestPath[1].X, shortestPath[1].Y, x, y), len(shortestPath)
		}
		return Stay, 200
	} else {
		return Stay, 200
	}
}

func findShortestPath(grid [][]int, start, target Point) []Point {
	rows, cols := len(grid), len(grid[0])

	// 定义四个方向的偏移量：上、下、左、右
	directions := [4]Point{
		{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0},
	}

	visited := make(map[Point]bool)
	costs := make(map[Point]int)
	parents := make(map[Point]Point)

	priorityQueue := []Point{start}
	costs[start] = 0

	for len(priorityQueue) > 0 {
		current := extractMin(&priorityQueue, costs)

		if current.X == target.X && current.Y == target.Y {
			return reconstructPath(parents, current)
		}

		visited[current] = true

		for _, dir := range directions {
			next := Point{current.X + dir.X, current.Y + dir.Y, current.Cost + dir.Cost}

			if isValidPoint(next, rows, cols, grid) && !visited[next] {
				_, exists := costs[next]
				if !exists || next.Cost < costs[next] {
					costs[next] = next.Cost
					parents[next] = current
					priorityQueue = append(priorityQueue, next)
				}
			}
		}
	}

	return nil
}

func isValidPoint(p Point, rows, cols int, grid [][]int) bool {
	return p.X >= 0 && p.X < rows && p.Y >= 0 && p.Y < cols && grid[p.X][p.Y] != 0
}

const MaxInt = int(^uint(0) >> 1)

func extractMin(queue *[]Point, costs map[Point]int) Point {
	minCost := MaxInt
	var minPoint Point
	var minIndex int

	for i, p := range *queue {
		if costs[p] < minCost {
			minCost = costs[p]
			minPoint = p
			minIndex = i
		}
	}

	*queue = append((*queue)[:minIndex], (*queue)[minIndex+1:]...)

	return minPoint
}

func reconstructPath(parents map[Point]Point, current Point) []Point {
	path := []Point{current}

	for {
		parent, exists := parents[current]
		if !exists {
			break
		}
		path = append([]Point{parent}, path...)
		current = parent
	}

	return path
}
