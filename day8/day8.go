package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func Day8() {
	f := getMatrixInput()
	nr, nc := getMatrixDimensions(f)
	fmt.Printf("No rows %v, No Cols %v \n", nr, nc)

	forest := Forest{Trees: make([]Tree, 0)}

	visibleTrees := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			// fmt.Printf("%v", string(f[r][c]))
			th, e := strconv.Atoi(string(f[r][c]))
			if e != nil {
				panic(e)
			}
			tree := Tree{Row: r, Col: c, TreeHeight: th, TreeRow: getForestRow(r, nc, f), TreeCol: getForestCol(c, nc, f)}
			tree.GetTopTrees()
			tree.GetBottomTrees()
			tree.GetLeftTrees()
			tree.GetRightTrees()
			tree.IsItVisible()
			tree.calcScenicScore()
			if tree.Visible {
				visibleTrees = visibleTrees + 1
			}
			forest.Trees = append(forest.Trees, tree)
		}
	}

	fmt.Println("Number of visible trees are", visibleTrees)

	sort.Sort(sort.Reverse(scenicSort(forest.Trees)))

	fmt.Println("Best scenic score is", forest.Trees[0].Score)
}

type scenicSort []Tree

func (t scenicSort) Len() int {
	return len(t)
}
func (t scenicSort) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t scenicSort) Less(i, j int) bool {
	return t[i].Score < t[j].Score
}

func getForestRow(row int, noCols int, f []string) []int {
	r := []int{}
	for c := 0; c < noCols; c++ {
		tstr, e := strconv.Atoi(string(f[row][c]))
		if e != nil {
			panic(e)
		}
		r = append(r, tstr)
	}
	return r
}

func getForestCol(col int, noRows int, f []string) []int {
	c := []int{}
	for r := 0; r < noRows; r++ {
		tstr, e := strconv.Atoi(string(f[r][col]))
		if e != nil {
			panic(e)
		}
		c = append(c, tstr)
	}
	return c
}

func getMatrixDimensions(f []string) (int, int) {
	return len(f), len(f[0])
}

func getMatrixInput() []string {
	b := getFile()
	c := string(b)
	ms := strings.Split(c, "\n")

	return ms
}

func getFile() []byte {
	b, e := file.ReadFile("./day8/day8.txt")

	if e != nil {
		panic(e)
	}
	return b
}

type Forest struct {
	Trees []Tree
}

type Tree struct {
	TreeHeight int
	TreeRow    []int
	TreeCol    []int

	Row int
	Col int

	TopTrees    []int
	BottomTrees []int
	LeftTrees   []int
	RightTrees  []int

	VisibleFromTop    bool
	VisibleFromBottom bool
	VisibleFromLeft   bool
	VisibleFromRight  bool
	Visible           bool

	topScore    int
	bottomScore int
	leftScore   int
	rightScore  int
	Score       int
}

func (t *Tree) GetTopTrees() {
	if t.Row == 0 {
		t.TopTrees = []int{}
		return
	}

	t.TopTrees = make([]int, 0)
	for r := 0; r < t.Row; r++ {
		t.TopTrees = append(t.TopTrees, t.TreeCol[r])
	}
}

func (t *Tree) GetBottomTrees() {
	if t.Row == len(t.TreeRow) {
		t.BottomTrees = []int{}
		return
	}

	t.BottomTrees = make([]int, 0)
	for r := t.Row + 1; r < len(t.TreeRow); r++ {
		t.BottomTrees = append(t.BottomTrees, t.TreeCol[r])
	}
}

func (t *Tree) GetLeftTrees() {
	if t.Col == 0 {
		t.LeftTrees = []int{}
		return
	}

	t.LeftTrees = make([]int, 0)
	for c := 0; c < t.Col; c++ {
		t.LeftTrees = append(t.LeftTrees, t.TreeRow[c])
	}
}

func (t *Tree) GetRightTrees() {
	if t.Col == len(t.TreeCol) {
		t.RightTrees = []int{}
		return
	}

	t.RightTrees = make([]int, 0)
	for c := t.Col + 1; c < len(t.TreeCol); c++ {
		t.RightTrees = append(t.RightTrees, t.TreeRow[c])
	}
}

func (t *Tree) IsItVisible() {
	// if the tree is on the edge, then it is always visible
	t.VisibleFromTop = len(t.TopTrees) == 0
	t.VisibleFromBottom = len(t.BottomTrees) == 0
	t.VisibleFromLeft = len(t.LeftTrees) == 0
	t.VisibleFromRight = len(t.RightTrees) == 0
	t.Visible = t.VisibleFromTop || t.VisibleFromBottom || t.VisibleFromLeft || t.VisibleFromRight

	if t.Visible {
		return
	}

	// Check from each side
	t.VisibleFromLeft = true
	for c := 0; c < len(t.LeftTrees); c++ {
		if t.LeftTrees[c] >= t.TreeHeight {
			t.VisibleFromLeft = false
			break
		}
	}

	t.VisibleFromRight = true
	for c := 0; c < len(t.RightTrees); c++ {
		if t.RightTrees[c] >= t.TreeHeight {
			t.VisibleFromRight = false
			break
		}
	}

	t.VisibleFromTop = true
	for r := 0; r < len(t.TopTrees); r++ {
		if t.TopTrees[r] >= t.TreeHeight {
			t.VisibleFromTop = false
			break
		}
	}

	t.VisibleFromBottom = true
	for r := 0; r < len(t.BottomTrees); r++ {
		if t.BottomTrees[r] >= t.TreeHeight {
			t.VisibleFromBottom = false
			break
		}
	}

	t.Visible = t.VisibleFromTop || t.VisibleFromBottom || t.VisibleFromLeft || t.VisibleFromRight
}

func (t *Tree) calcScenicScore() {

	if len(t.TopTrees) == 0 || len(t.BottomTrees) == 0 || len(t.LeftTrees) == 0 || len(t.RightTrees) == 0 {
		t.Score = 0
		return
	}

	leftScore := 0
	for c := len(t.LeftTrees) - 1; c >= 0; c-- {
		leftScore = leftScore + 1
		if t.LeftTrees[c] >= t.TreeHeight {
			break
		}
	}
	t.leftScore = leftScore

	rightScore := 0
	for c := 0; c < len(t.RightTrees); c++ {
		rightScore = rightScore + 1
		if t.RightTrees[c] >= t.TreeHeight {
			break
		}
	}
	t.rightScore = rightScore

	topScore := 0
	for r := len(t.TopTrees) - 1; r >= 0; r-- {
		topScore = topScore + 1
		if t.TopTrees[r] >= t.TreeHeight {
			break
		}
	}
	t.topScore = topScore

	bottomScore := 0
	for r := 0; r < len(t.BottomTrees); r++ {
		bottomScore = bottomScore + 1
		if t.BottomTrees[r] >= t.TreeHeight {
			break
		}
	}
	t.bottomScore = bottomScore

	t.Score = bottomScore * topScore * rightScore * leftScore
}
