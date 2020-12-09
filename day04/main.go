package main

import (
    "bufio"
    "fmt"
    "github.com/elliotchance/pie/pie"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"

    //"../reader"
)

type Passport struct {
    fields map[string]string
}

var validateKey = map[string]func (string) bool{
    "byr" : func(s string) bool{
        yr, _ := strconv.Atoi(s)
        return len(s) == 4 && yr >=1920 && yr <= 2002
    },
    "iyr" : func(s string) bool{
        yr, _ := strconv.Atoi(s)
        return len(s) == 4 && yr >=2010 && yr <= 2020
    },
    "eyr" : func(s string) bool{
        yr, _ := strconv.Atoi(s)
        return len(s) == 4 && yr >=2020 && yr <= 2030
    },
    "hgt" : func(s string) bool{
        re := regexp.MustCompile(`^((1([5-8][0-9]|9[0-3])cm)|((59|6[0-9]|7[0-6])in))$`)
        return re.MatchString(s)
    },
    "hcl" : func(s string) bool{
        re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
        return re.MatchString(s)
    },
    "ecl" : func(s string) bool{
        ecls := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
        return pie.Strings(ecls).Contains(s)
    },
    "pid": func(s string) bool{
        return len(s) == 9
    },
}

func Lines(path string) []string {
    lines := make([]string, 0)

    file, err := os.Open(path)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        str := strings.TrimSpace(scanner.Text())
        lines = append(lines, str)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    return lines
}

func parse(data string) map[string]string {
    str := strings.TrimSpace(data)
    pairs := strings.Split(str, " ")
    fields := make(map[string]string)

    for _, pair := range pairs {
        kv := strings.Split(pair, ":")
        k := kv[0]
        v := kv[1]
        fields[k] = v
    }

    return fields
}

func createPassports() []Passport {
    passports := make([]Passport, 0)
    lines := Lines("input.txt")
    fields := ""

    for _, line := range lines {
        if line == "" {
            passports = append(passports, Passport{fields: parse(fields)})
            fields = ""
        } else  {
            fields = fields + " " + line
        }
    }
    passports = append(passports, Passport{fields: parse(fields)})
    return passports
}

func (p Passport) optionalKeys() pie.Strings {
    keys := []string{"cid"}
    return pie.Strings(keys).Sort()
}

func (p Passport) requiredKeys() pie.Strings {
    keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
    return pie.Strings(keys).Sort()
}

func (p Passport) keys() pie.Strings {
    keys := make([]string, 0)
    for k, _ := range p.fields {
        if k != "cid" {
            keys = append(keys, k)
        }
    }
    return pie.Strings(keys).Sort()
}

func (p Passport) CheckKeys () bool {
    rkeys := p.requiredKeys()
    ikeys := p.keys().Intersect(rkeys).Sort()
    return rkeys.Equals(ikeys)
}

func (p Passport) validateKeys() bool {
    res := p.CheckKeys()

    if res {
        for _, k := range p.requiredKeys() {
            val := p.fields[k]
            res = validateKey[k](val)
            if !res {
                break
            }
        }
    }

    return res
}

func part1() int {
    count := 0
    for _, p := range createPassports() {
        fmt.Println(p)
        if p.CheckKeys() {
            count = count + 1
        }
    }
    return count
}

func part2() int {
    count := 0
    for _, p := range createPassports() {
        if p.validateKeys() {
            count = count + 1
        }
    }
    return count
}

func main() {
    fmt.Println("Day 4")
    fmt.Println("Part 1 -- ", part1())
    fmt.Println("Part 2 -- ", part2())
}
