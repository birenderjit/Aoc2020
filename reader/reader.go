package reader

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
        fmt.Println(str)
        lines = append(lines, str)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

    return lines
}
