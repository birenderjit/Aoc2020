package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

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

func part1(lines []string, right int, down int) int{
    row := 0
    col := 0
    count := 0
    patWidth := len(lines[0])
    for row < len(lines) {
        if lines[row][col] == '#' {
            count++
        }
        row += down
        if col + right > patWidth - 1 {
            col = col + right - patWidth
        } else {
            col += right
        }
    }
    return count
}

func part2(lines []string)  int{
    result := 1
    slopes := [5][2]int{{1,1}, {3,1},{5,1},{7,1},{1,2}}
    for _, slope := range slopes {
        result *= part1(lines, slope[0], slope[1])
    }
    return result
}

func main() {
    input := Lines("input.txt")
    fmt.Println(part1(input, 3, 1))
    fmt.Println(part2(input))
}
