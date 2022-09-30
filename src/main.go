package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

var (
	MaleName   []string = ReadFromResource("../male-first-names.txt")
	FemaleName []string = ReadFromResource("../female-first-names.txt")
	FamilyName []string = ReadFromResource("../last-names.txt")
	MiddleChar []string = []string{".", "_"}
)

type Gender uint8

const (
	Male Gender = iota
	Female
	Both
)

func GenerateRandomNames(limit int, gender Gender) []string {
	var nameList []string
	switch gender {
	case Male:
		nameList = MaleName
	case Female:
		nameList = FemaleName
	case Both:
		nameList = append(MaleName, FemaleName...)
	}

	var result []string

	for len(result) < limit {
		newName := GenerateName(nameList)
		if slices.Contains(result, newName) {
			continue
		}
		result = append(result, strings.Trim(newName, "\r"))
	}
	return result
}

func GenerateName(nameList []string) string {
	var name string
	name += nameList[rand.Intn(len(nameList))]
	name += GetRandomChars()
	name += FamilyName[rand.Intn(len(FamilyName))]
	name += GenerateRandomNumber()
	return name
}

func GenerateRandomNumber() string {
	if !GenerateRandomBool() {
		return ""
	}
	return strconv.Itoa(rand.Intn(999))
}

func GetRandomChars() string {
	if !GenerateRandomBool() {
		return ""
	}
	return MiddleChar[rand.Intn(len(MiddleChar))]
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}

func ReadFromResource(fileName string) []string {
	dat, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(dat), "\n")
}

func main() {
	rand.Seed(time.Now().Unix())
	var (
		limit  int
		gender Gender
	)

	flag.IntVar(&limit, "limit", 1, "limit of generated names")
	flag.Func("gender", "gender of names", func(s string) error {
		switch strings.ToLower(s) {
		case "male":
			gender = Male
		case "female":
			gender = Female
		case "both":
			gender = Both
		default:
			gender = Both
		}

		return nil
	})
	flag.Parse()

	names := GenerateRandomNames(limit, gender)

	fmt.Printf("List of %d name(s):\n", limit)
	for _, name := range names {
		fmt.Printf(" - %s\n", name)
	}
}
