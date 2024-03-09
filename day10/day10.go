package main

import (
	"bufio"
	"fmt"
	"os"

	pipefield "github.com/Floriszenz/advent-of-code-2023/day10/pipeField"
)

func main()  {
    if len(os.Args) != 2 {
        fmt.Println("You have to provide the path of the input file as argument.")
        return
    }

    pipeFieldSketch, err := os.Open(os.Args[1])

    if err != nil {
        fmt.Println("Could not open pipe field file.", err)
        return
    }
    defer pipeFieldSketch.Close()

    scanner := bufio.NewScanner(pipeFieldSketch)

    field := pipefield.InitializePipeField(scanner)
    fieldMap := pipefield.InitializeFieldMap(len(field), len(field[0]))

    err = pipefield.FollowPipe(field, fieldMap)

    if err != nil {
        fmt.Println(err)
        return
    }
        
    stepsForFarthestPoint := pipefield.GetStepsForFarthestPoint(fieldMap)
    enclosedTileCount := pipefield.GetNumberOfEnclosedTiles(field, fieldMap)

    fmt.Printf("Number of steps required to reach farthest point from pipe start: %v\n", stepsForFarthestPoint)
    fmt.Printf("Number of tiles enclosed by the loop: %v\n", enclosedTileCount)
}
