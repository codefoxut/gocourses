package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	main2()
	var filename string
	fmt.Print("Please enter a filename: ")
	fmt.Scan(&filename)

	// 1. open the file
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// 2. create slice of 20 persons as I know there are 20 names.
	person_list := make([]Person, 0, 20)
	var seek_pos int = 0
	for {
		// 3. create a byte array for 20 characters.
		barr := make([]byte, 20)
		// 4. seek is needed to back-track extra characters before reading next name.
		_, _ = f.Seek(int64(seek_pos), 1)
		n1, err := f.Read(barr)
		if err != nil {
			break
		}
		// 5. convert the byte to unicode.
		names_str := string(barr[:n1])
		// fmt.Printf("%d %d bytes: %s\n", n1, seek_pos, names_str)
		// 6. split the string in 2 pieces about "\n"
		names_temp := strings.SplitN(names_str, "\n", 2)

		// 7. create p1 object with first name and last name adn append it to slice.
		names := strings.Split(names_temp[0], " ")
		// fmt.Printf("%s %s\n", names, names_temp)
		p1 := Person{fname: names[0], lname: names[1]}
		person_list = append(person_list, p1)

		// calculate next position to read the full name as we must have read more characters.
		if len(names_temp) > 1 {
			seek_pos = -len(names_temp[1])
		} else {
			break
		}
	}
	// fmt.Print(person_list)
	for i, p2 := range person_list {
		fmt.Println(i+1, p2.fname, p2.lname)
	}

}

type FullName struct {
	fname, lname string
}

func trim_str_to_length(str string, length int) string {
	if len(str) < length {
		return str
	} else {
		return str[0:length]
	}
}

func main2() {
	myslice := make([]FullName, 0, 20)
	var input_file_path string
	fmt.Print("Enter file name of the text file: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input_file_path = scanner.Text()
	}
	data, err := ioutil.ReadFile(input_file_path)
	if err != nil {
		fmt.Print("Read file errored", err)
		return
	}
	text := strings.Split(string(data), "\n")
	for i := 0; i < len(text); i++ {
		if len(text[i]) < 3 {
			continue
		}
		splitted := strings.Fields(text[i])
		fname := trim_str_to_length(splitted[0], 20)
		lname := trim_str_to_length(splitted[1], 20)
		full_name := FullName{fname: fname, lname: lname}
		myslice = append(myslice, full_name)
	}

	for _, full_name := range myslice {
		fmt.Printf("%s %s\n", full_name.fname, full_name.lname)

	}
}
