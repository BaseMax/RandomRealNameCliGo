package main

import (
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
)

var MiddleChar []string = []string{".", "_"}

type Gender int

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

	rand.Seed(time.Now().Unix())
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
	var want_limit = 1
	var want_gender = Both

	if len(os.Args) <= 1 {
		fmt.Printf("%s -limit n -gender both", os.Args[0])
	} else {
		var count = len(os.Args)
		var isValid = true
		for i := 1; i < count; i++ {
			if os.Args[i] == "-limit" {
				if i+1 <= count {
					val, err := strconv.Atoi(os.Args[i+1])
					if err != nil {
						panic(err)
					}
					want_limit = val
				} else {
					isValid = false
					fmt.Printf("Error: limit value missed!")
					break
				}
				i++
			} else if os.Args[i] == "-gender" {
				if i+1 <= count {
					val := strings.ToLower(os.Args[i+1])
					if val == "male" {
						want_gender = Male
					} else if val == "female" {
						want_gender = Female
					} else if val == "both" {
						want_gender = Both
					} else {
						isValid = false
						fmt.Printf("Error: %s mode is not supported as gender value!", os.Args[i+1])
						break
					}
					i++
				} else {
					isValid = false
					fmt.Printf("Error: gender value missed!")
					break
				}
			} else {
				isValid = false
				fmt.Printf("Error: %s not supporting as an argument!", os.Args[i])
				break
			}
		}

		if isValid {
			names := GenerateRandomNames(want_limit, want_gender)

			fmt.Printf("List of %d name(s):\n", want_limit)
			for _, name := range names {
				fmt.Printf(" - %s\n", name)
			}
		}
	}
}
