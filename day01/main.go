package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

func loadDataFromFile(name string) []int {
    file, err := os.Open(name)
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var data []int

    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatalln(err)
        }
        data = append(data, num)
    }
    sort.Ints(data)
    return data
}

func bruteForceSum2(data []int) int {
    res := 0
    for i := 0; i < len(data); i++ {
        for j := i+1; j < len(data); j++ {
            if data[i] + data[j] == 2020 {
                res = data[i] * data[j]
            }
        }
    }
    return res
}

func sum2(data []int) int {
    res := 0
    l := 0
    r := len(data) - 1

    for l < r {
        if (data[l] + data[r]) == 2020 {
            res = data[l] * data[r]
            break
        } else if (data[l] + data[r]) < 2020 {
            l++
        } else {
            r--
        }
    }

    return res
}

func bruteForceSum3(data []int) int {
    res := 0
    for i := 0; i < len(data)-3; i++ {
        for j := i+1; j < len(data)-2; j++ {
            for k := j+1; k < len(data)-1; k++ {
                //fmt.Println(i, j, k)
                if data[i] + data[j] + data[k] == 2020 {
                    res = data[i] * data[j] * data[k]
                }
            }
        }
    }
    return res
}

func main()  {
    data := loadDataFromFile("input.txt")

    fmt.Println(bruteForceSum2(data))
    fmt.Println(sum2(data))
    fmt.Println(bruteForceSum3(data))

}
