package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func loadDataFromFile(name string) []string {
    file, err := os.Open(name)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var data []string

    for scanner.Scan() {
        data = append(data, scanner.Text())
    }
    return data
}

func part1(list []string) int {
    count := 0
    for _, line := range list {
        var splits = strings.Split(line, " ")
        var c = strings.Split(splits[0], "-")
        var charToCheck = string(splits[1][0])
        //fmt.Println(c[0], c[1], charToCheck)
        cc := strings.Count(splits[2], charToCheck)
        low, _ := strconv.Atoi(c[0])
        high, _ := strconv.Atoi(c[1])
        if (cc >= low && cc <= high) {
            count++
        }
    }
    return count
}

func part2(list []string) int {
    count := 0
    for _, line := range list {
        var splits = strings.Split(line, " ")
        var c = strings.Split(splits[0], "-")
        low, _ := strconv.Atoi(c[0])
        high, _ := strconv.Atoi(c[1])
        if (splits[2][low-1] == splits[1][0] &&
            splits[2][high-1] != splits[1][0]) ||
            (splits[2][low-1] != splits[1][0] &&
                splits[2][high-1] == splits[1][0]){
            count++
        }
    }
    return count
}

func main()  {
    data := loadDataFromFile("input.txt")

    fmt.Println("Part 1 -- ", part1(data))
    fmt.Println("Part 2 -- ", part2(data))
}
