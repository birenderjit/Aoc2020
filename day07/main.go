package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {
    fmt.Println("Part 1 --", part1())
    //fmt.Println("Part 2 --", part2())
}

func part1() int {
    count := 0
    lines := Lines("in.txt")
    for _, line := range lines {
        fmt.Println(strings.Split(line, "contain"))
    }
    return count

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
