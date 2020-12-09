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
)

type Ticket struct {
   code string
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

func part1() int {
    return sortedIdList().Last()
}

func part2() int {
    prev, list := sortedIdList().Shift()

    for _, id := range list {
        if id - prev == 2 {
            break
        } else {
            prev = id
        }
    }
    return prev + 1
}

func makeTickets() []Ticket {
    codes := Lines("input.txt")
    tickets := make([]Ticket, len(codes))
    for i, code := range codes {
        tickets[i] = Ticket{code: code}
    }
    return tickets
}

func (t Ticket) getId() int {
    re0 := regexp.MustCompile(`[F|L]`)
    re1 := regexp.MustCompile(`[B|R]`)

    code := t.code
    code = re0.ReplaceAllString(code, "0")
    code = re1.ReplaceAllString(code, "1")

    num, _ := strconv.ParseInt(code, 2, 32)
    return int(num)
}

func sortedIdList() pie.Ints {
     tickets := makeTickets()
     ids := pie.Ints(make([]int, len(tickets)))
     for i, t := range tickets {
         ids[i] = t.getId()
     }
     return ids.Sort()
}

func main() {
    fmt.Println("Day 5")
    fmt.Println("Part 1 -- ", part1())
    fmt.Println("Part 2 -- ", part2())

}
