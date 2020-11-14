package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Int int
type String string

func (i Int) toString() string {
	return strconv.Itoa(int(i))
}

func (s String) toInt() (n int) {
	n, _ = strconv.Atoi(string(s))
	return
}

func (s *String) set(ns String) {
	*s = ns
}

func (s String) get() String {
	return s
}

func main() {

	var i Int = 12
	fmt.Printf("%q\n", i.toString())

	var s String = "12"
	fmt.Printf("%d\n", s.toInt())

	s.set("yyy")
	fmt.Println(s.get())

	myPrint(4)

	// var cred credential
	// err := json.Unmarshal([]byte(json))

	// fmt.Printf(" %#v ", trippleChar("abcdef"))
	// getdata()

	// var p *int
	// n := 5
	// p = &n

	// set(p)
	// fmt.Println(n)

}

func myPrint(n int) {
	defer fmt.Println(n)
	n = n + n

	defer func() {
		fmt.Println(n)
	}()

	fmt.Println(n)
}

func catchMe() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println(r)
		}

	}()

	// log.Fatal("test")
	log.Panic("test")
}

func set(n *int) {
	fmt.Println(*n)
	*n = *n + *n
}

func getdata() {
	dat, err := ioutil.ReadFile("./focusive.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bytes.NewReader(dat))
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	nameAndCount := map[string]int{}
	for _, record := range records {

		nameAndCount[record[3]]++
		// fmt.Printf("%#v\n", record[3])

	}

	for name, count := range nameAndCount {
		if count > 1 {
			fmt.Println(name, count)
		}
	}
}

func trippleChar(s string) []string {

	count := 3 - (len(s) % 3)
	s += strings.Repeat("*", count)

	round := len(s) / 3

	result := []string{}
	for i := 0; i < round; i++ {
		p := s[:3]
		s = s[3:]

		result = append(result, p)

	}

	return result

}
