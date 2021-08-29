package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// remove the entry at index int.
}

// separation of concerns

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// load the jornals.
}

func (j *Journal) LoadFromWeb(Url *url.URL) {

}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I went for shopping.")
	j.AddEntry("I am moving to some other city.")

	fmt.Println(j.String())

	SaveToFile(&j, "journal.txt")

	p := Persistence{"\r\n\r\n"}
	p.SaveToFile(&j, "journal2.txt")
}
