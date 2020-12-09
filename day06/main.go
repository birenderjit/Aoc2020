package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

type Group struct {
    answers []string
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

func createGroups() []Group {
    lines := Lines("input.txt")
    groups := make([]Group, 0)
    answers := make([]string, 0)

    for _, line := range lines {
        if line == "" {
            groups = append(groups, Group{answers: answers})
            answers = make([]string, 0)
        } else  {
            answers = append(answers,line)
        }
    }
    groups = append(groups, Group{answers: answers})
    return groups
}

func (g Group) getAnswerSet() map[string]int {
    ansSet := map[string]int{}
    ans := g.answers
    str := strings.Join(ans, "")
    chars := strings.Split(str, "")

    for _, ch := range chars {
        ansSet[ch]++
    }

    return ansSet
}


func (g Group) getUniqueAnswers() string {
    ansSet := g.getAnswerSet()
    ansStr := ""

    for k, _ := range ansSet {
        ansStr += k
    }

    return ansStr
}

func (g Group) getUnanimousAnswers() string {
    ansSet := g.getAnswerSet()
    ansStr := ""
    grpSize := len(g.answers)

    for k, v := range ansSet {
        if v == grpSize {
            ansStr += k
        }
    }

    return ansStr
}

func part1() int {
    total := 0
    for _, group := range createGroups() {
        uniqueAns := group.getUniqueAnswers()
        total += len(uniqueAns)
    }

    return total
}

func part2() int {
    total := 0
    for _, group := range createGroups() {
        unanimousAns := group.getUnanimousAnswers()
        total += len(unanimousAns)
    }

    return total
}

func main() {
    fmt.Println("Day 6")
    fmt.Println("Part 1 -- ", part1())
    fmt.Println("Part 2 -- ", part2())
}
