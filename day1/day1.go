package day1

import "os"

func readFile() ([]byte, error) {
	return os.ReadFile("./day1/day1a.txt")
}
