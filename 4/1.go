package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const FN = "input.txt"

var NEEDED_FIELDS map[string]bool = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
}

func main() {
	f, _ := os.Open(FN)

	scanner := bufio.NewScanner(f)

	var allPp [][]string

	var ppData []string
	for {
		s := scanner.Scan()
		if !s {
			fmt.Println("passport:", strings.Join(ppData, " | "))
			allPp = append(allPp, ppData)
			break
		}
		r := scanner.Text()
		if len(r) == 0 {
			fmt.Println("passport:", strings.Join(ppData, " | "))
			allPp = append(allPp, ppData)
			ppData = make([]string, 0)
			continue
		}
		vals := strings.Fields(r)
		ppData = append(ppData, vals...)
	}

	valids := 0
	fmt.Printf("total of %d passports\n", len(allPp))
	for i, pp := range allPp {
		var target map[string]bool = map[string]bool{}
		for k, v := range NEEDED_FIELDS {
			target[k] = v
		}
		for _, kv := range pp {
			fieldS := strings.Split(kv, ":")

			field := fieldS[0]
			value := fieldS[1]

					delete(target, field)
			switch field {
			case "byr":
				tmp, _ := strconv.Atoi(value)
				if tmp >= 1920 && tmp <= 2002 {
					delete(target, field)
				}
			case "iyr":
				tmp, _ := strconv.Atoi(value)
				if tmp >= 2010 && tmp <= 2020 {
					delete(target, field)
				}
			case "eyr":
				tmp, _ := strconv.Atoi(value)
				if tmp >= 2020 && tmp <= 2030 {
					delete(target, field)
				}
			case "hgt":
				var num int
				var unit string
				fmt.Sscanf(value, "%d%s", &num, &unit)
				if unit == "cm" {
					if num >= 150 && num <= 193 {
						delete(target, field)
					}
				}

				if unit == "in" {
					if num >= 59 && num <= 76 {
						delete(target, field)
					}
				}
			case "hcl":
				re, err := regexp.Compile(`^#[0-9a-f]{6}`)
				check(err)
				if re.MatchString(value) {
					delete(target, field)
				}
			case "ecl":

				for _, i := range strings.Fields("amb blu brn gry grn hzl oth") {
					if i == value {
						delete(target, field)
						continue
					}
				}
			case "pid":
				_, err := strconv.Atoi(value)
				if err == nil && len(value) == 9 {
					delete(target, field)
				}
			}
		}
		if len(target) == 0 {
			fmt.Printf("%d is %v\n", i, "valid")
			valids++
		}
	}
	fmt.Println("valids:", valids)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
