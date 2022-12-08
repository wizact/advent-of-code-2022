package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func getFile() []byte {
	b, e := file.ReadFile("./day7/day7.txt")

	if e != nil {
		panic(e)
	}
	return b
}

func Day7() {
	i := getInstructions()
	p := &Folder{}
	runDownTheTree(i, -1, p)
	// Part a
	fmt.Printf("root (%v) size is (%v) and the sum of directories within threshold is (%v)\n", p.name, p.totalSize, calcSum(p, 100000, 0))

	// Part b
	ff := []Folder{}
	d := flattenFolders(p, ff)
	sort.Sort(folderStructure(d))

	const maxSpace int64 = 70000000
	var usedSpace int64 = p.totalSize
	const requiredSpace = 30000000
	var availableSpace = maxSpace - usedSpace
	var spaceToFreeUp = requiredSpace - availableSpace

	fmt.Printf("In order to update the system, we still need (%v) space\n", spaceToFreeUp)
	for i := 0; i < len(d); i++ {
		if spaceToFreeUp <= d[i].totalSize {
			fmt.Printf("Directory to delete is (%v) with total space of (%v)\n", d[i].name, d[i].totalSize)
			break
		}
	}

}

type folderStructure []Folder

func (s folderStructure) Len() int {
	return len(s)
}
func (s folderStructure) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s folderStructure) Less(i, j int) bool {
	return s[i].totalSize < s[j].totalSize
}

func flattenFolders(p *Folder, pl []Folder) []Folder {
	for _, v := range p.Folders {
		pl = append(pl, *v)
		pl = flattenFolders(v, pl)
	}
	return pl
}

func calcSum(p *Folder, threshold int64, ts int64) int64 {
	for _, v := range p.Folders {
		if v.totalSize <= threshold {
			ts = ts + v.totalSize
		}
		ts = calcSum(v, threshold, ts)
	}
	return ts
}

func runDownTheTree(t []string, index int, currentFolder *Folder) {
	index = index + 1
	newParent := currentFolder
	switch t[index] {
	case "$ cd /":
		// Parent directory
		newParent.name = "/"
	case "$ ls":
		// Set parent to the current parent
	case "$ cd ..":
		// go one level up
		runDownTheTree(t, index, currentFolder.parent)
		return
	default:
		if strings.Index(t[index], "dir") == 0 {
			folderName := strings.Split(t[index], " ")[1]
			currentFolder.Folders = append(currentFolder.Folders, &Folder{parent: currentFolder, name: folderName})
		} else if strings.Index(t[index], "$ cd") == 0 {
			folderName := strings.Split(t[index], " ")[2]
			// Find the array item
			var node *Folder
			for _, v := range currentFolder.Folders {
				if v.name == folderName {
					node = v
				}
			}
			runDownTheTree(t, index, node)
			currentFolder.totalSize = currentFolder.totalSize + node.totalSize
			return
		} else {
			fileName := strings.Split(t[index], " ")[1]
			fileSize := strings.Split(t[index], " ")[0]
			fs, e := strconv.ParseInt(fileSize, 10, 64)
			if e != nil {
				panic(e)
			}
			currentFolder.Files = append(currentFolder.Files, File{name: fileName, size: fs})
			currentFolder.totalSize = currentFolder.totalSize + fs
		}
	}

	if index < len(t)-1 {
		runDownTheTree(t, index, newParent)
	} else {
		return
	}
}

func getInstructions() []string {
	b := getFile()
	c := string(b)
	ms := strings.Split(c, "\n")

	return ms
}

type Folder struct {
	parent    *Folder
	name      string
	totalSize int64

	Folders []*Folder
	Files   []File
}

type File struct {
	name string
	size int64
}
