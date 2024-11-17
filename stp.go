package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
    "sync"
    "time"
)

func cpuIntensiveTask(complexity int) {
    result := 0
    for i := 0; i < complexity*complexity*10000; i++ {
        result += i
    }
}

func computeTask(complexity int) {
    cpuIntensiveTask(complexity)
}

func multiProcess(iterations, complexity int) {
    var wg sync.WaitGroup

    for i := 0; i < iterations; i++ {
        cmd := exec.Command(os.Args[0], "-m", "s", "-c", fmt.Sprint(complexity))
        if err := cmd.Start(); err != nil {
            fmt.Println("Error spawning subprocess:", err)
            continue
        }
        wg.Add(1)
        go func() {
            defer wg.Done()
            cmd.Wait()
        }()
    }

    wg.Wait()
}

func multiThread(iterations, complexity int) {
    var wg sync.WaitGroup

    for i := 0; i < iterations; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            computeTask(complexity)
        }()
    }

    wg.Wait()
}


func main() {
    mode := flag.String("m", "", "Select the mode of execution: singleprocess(s), multiprocess(p), or multithread(t)")
    complexity := flag.Int("c", 1, "Set the complexity level")
    iterations := flag.Int("i", 1, "Set the number of iterations")

    flag.Parse()

    if *mode == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

    start := time.Now()

    switch *mode {
    case "singleprocess", "s":
        for i := 0; i < *iterations; i++ {
            computeTask(*complexity)
        }
    case "multiprocess", "p":
        multiProcess(*iterations, *complexity)
    case "multithread", "t":
        multiThread(*iterations, *complexity)
    default:
        fmt.Println("Invalid mode. Please use singleprocess(s), multiprocess(p), or multithread(t).")
        flag.PrintDefaults()
        os.Exit(1)
    }

    elapsed := time.Since(start)
    fmt.Println("Execution time:", elapsed)
}

