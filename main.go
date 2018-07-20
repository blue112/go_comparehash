package main

import (
    "os"
    "fmt"
    "bufio"
    "time"
)

func main () {
    if len(os.Args) != 2 {
        println("it must take args")
        os.Exit(1)
    }

    filename := os.Args[1]

    f, _ := os.Open(filename)
    r := bufio.NewReader(f)

    var err error
    var d []byte

    var hashes [767690][40]byte

    start := time.Now().UnixNano()
    n := 0
    for err == nil {
        d, err = r.ReadBytes('\n')
        if len(d) > 0 {
            d = d[:len(d) - 1]
            copy(hashes[n][:], d)
            n++
        }
    }
    f.Close()

    delta := float64(time.Now().UnixNano() - start) / (1000 * 1000 * 1000)
    fmt.Printf("Loaded %d hashes in %f s\n", len(hashes), delta)

    min_dist_all := 40
    var min_dist_hash_all []byte

    for i1 := 0; i1 < len(hashes); i1++ {
        min_dist := 40
        var min_dist_hash []byte

        start_t := time.Now().UnixNano()
        for i2 := i1 + 1; i2 < len(hashes); i2++ {
            dist := getDist(hashes[i1][:], hashes[i2][:])
            if dist < min_dist {
                min_dist = dist
                min_dist_hash = hashes[i2][:]
            }
        }

        if min_dist < min_dist_all {
            min_dist_all = min_dist
            min_dist_hash_all = min_dist_hash
            fmt.Printf("%s ~= %s (%d)\n", hashes[i1], min_dist_hash_all, 40 - min_dist)
        }

        end_t := time.Now().UnixNano()
        delta = float64(end_t - start_t) / (1000 * 1000)
        fmt.Printf("One hash = %fms\n", delta)
        break
    }
}

func getDist(a, b []byte) int {
    d := 0
    l := len(a)
    for i := 0; i < l; i++ {
        if a[i] != b[i] {
            d++
        }
    }

    return d
}
