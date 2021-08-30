package singleton

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	// fmt.Println(db.capitals)
	return db.capitals[name]
}

// both init and sync.Once are thread-safe
// but only sync.Once is lazy

var once sync.Once

var instance *singletonDatabase

// func GetSingletonDatabase() *singletonDatabase {
//  once.Do(func() {
//    if instance == nil {
//      instance = &singletonDatabase{}
//    }
//  })
//  return instance
// }

// init() — we could, but it's not lazy

func readData(path string) (map[string]int, error) {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	// fmt.Println(exPath, path)

	file, err := os.Open("singleton/" + path)
	// fmt.Println(file, err)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("capitals.txt")
		db := singletonDatabase{}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)

}
