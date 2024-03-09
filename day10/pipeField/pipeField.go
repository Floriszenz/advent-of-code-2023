package pipefield

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Floriszenz/advent-of-code-2023/day10/direction"
	"github.com/Floriszenz/advent-of-code-2023/day10/tile"
)

func InitializePipeField(s *bufio.Scanner) [][]string {
    var pipeField [][]string

    for s.Scan() {
        row := s.Text()
        tile := strings.Split(row, "")
        pipeField  = append(pipeField , tile)
    }

    return pipeField
}

func InitializeFieldMap(rows, cols int) [][]string {
    fieldMap := make([][]string, rows)

    for y := 0; y < rows; y++ {
        fieldMap[y] = make([]string, cols)

        for x := 0; x < cols; x++ {
            fieldMap[y][x] = "."
        }
    }

    return fieldMap
}

func findStartingPosition(pipeField *[][]string) (int, int, error) {
    for y, row := range *pipeField {
        for x, t := range row {
            if tile.IsStartingPosition(t) {
                return x, y, nil
            }
        }
    }

    return 0, 0, errors.New("Couldn't find starting tile in given pipe field.")
}

func getInitialDirection(f [][]string, x, y int) (int, error) {
    if y > 0 && tile.IsSouthOpen(f[y - 1][x]) {
        return direction.NORTH, nil
    } else if x < len(f[0]) && tile.IsWestOpen(f[y][x + 1]) {
        return direction.EAST, nil
    } else if y < len(f) && tile.IsNorthOpen(f[y + 1][x]) {
        return direction.SOUTH, nil
    } else if x > 0 && tile.IsEastOpen(f[y][x - 1]) {
        return direction.WEST, nil
    }

    return 0, errors.New("Couldn't determine initial direction.")
}

func getTypeOfStartingTile(f [][]string, x, y int) (string, error) {
    firstDir, secondDir := -1, -1

    if y > 0 && tile.IsSouthOpen(f[y - 1][x]) {
        firstDir = direction.NORTH
    }
    if x < len(f[0]) && tile.IsWestOpen(f[y][x + 1]) {
        if firstDir == -1 {
            firstDir = direction.EAST
        } else {
            secondDir = direction.EAST
        }
    }
    if y < len(f) && tile.IsNorthOpen(f[y + 1][x]) {
        if firstDir == -1 {
            firstDir = direction.SOUTH
        } else {
            secondDir = direction.SOUTH
        }
    }
    if x > 0 && tile.IsEastOpen(f[y][x - 1]) {
        secondDir = direction.WEST
    }

    switch {
    case firstDir == direction.NORTH && secondDir == direction.EAST:
        return "L", nil
    case firstDir == direction.NORTH && secondDir == direction.SOUTH:
        return "|", nil
    case firstDir == direction.NORTH && secondDir == direction.WEST:
        return "J", nil
    case firstDir == direction.EAST && secondDir == direction.SOUTH:
        return "F", nil
    case firstDir == direction.EAST && secondDir == direction.WEST:
        return "-", nil
    case firstDir == direction.SOUTH && secondDir == direction.WEST:
        return "7", nil
    }

    return "", errors.New("Couldn't determine type of starting tile")
}

func updateDirection(t string, d int) int {
    switch d {
    case direction.NORTH:
        switch {
        case tile.IsSouthEastBend(t):
            return direction.EAST
        case tile.IsSouthWestBend(t):
            return direction.WEST
        }
    case direction.EAST:
        switch {
        case tile.IsNorthWestBend(t):
            return direction.NORTH
        case tile.IsSouthWestBend(t):
            return direction.SOUTH
        }
    case direction.SOUTH:
        switch {
        case tile.IsNorthEastBend(t):
            return direction.EAST
        case tile.IsNorthWestBend(t):
            return direction.WEST
        }
    case direction.WEST:
        switch {
        case tile.IsNorthEastBend(t):
            return direction.NORTH
        case tile.IsSouthEastBend(t):
            return direction.SOUTH
        }
    }

    return d
}


func FollowPipe(f, m [][]string) error {
    x, y, err := findStartingPosition(&f)

    if err != nil {
        return err
    }

    dir, err := getInitialDirection(f, x, y)

    if err != nil {
        return err
    }

    for {
        // Make next step
        switch dir {
        case direction.NORTH:
            y -= 1
        case direction.EAST:
            x += 1
        case direction.SOUTH:
            y += 1
        case direction.WEST:
            x -= 1
        }

        // Update direction based on tile
        dir = updateDirection(f[y][x], dir)

        m[y][x] = "x"

        if tile.IsStartingPosition(f[y][x]) {
            break
        }
    }

    // Note: This does modify the original data structure, making this function not idempotent
    startingType, err := getTypeOfStartingTile(f, x, y)

    if err != nil {
        return err
    }

    f[y][x] = startingType

    return nil
}

func GetStepsForFarthestPoint(m [][]string) int {
    steps := 0

    for _, row := range m {
        for _, t := range row {
            if t == "x" {
                steps++
            }
        }
    }

    return steps / 2
    
}

func GetNumberOfEnclosedTiles(f, m [][]string) int {
    for y, row := range m {
        boundaryCrosses := 0

        for x := 0; x < len(row); x++ {
            if m[y][x] == "." {
                m[y][x] = fmt.Sprint(boundaryCrosses)
            } else if tile.IsVerticalPipe(f[y][x]) {
                boundaryCrosses++
            } else if tile.IsHorizontalPipe(f[y][x]) {
                continue
            } else {
                var nextTile string
                var i int

                for i = x + 1; i < len(row); i++ {
                    if !tile.IsHorizontalPipe(f[y][i]) {
                        nextTile = f[y][i]
                        break
                    }
                }

                if tile.IsSouthOpen(f[y][x]) && tile.IsNorthOpen(nextTile) || tile.IsNorthOpen(f[y][x]) && tile.IsSouthOpen(nextTile) {
                    boundaryCrosses++
                }

                x = i
            }
        }
    }

    enclosedTileCount := 0

    for _, row := range m {
        for _, t := range row {
            if v, err := strconv.ParseInt(t, 10, 64); err == nil && v % 2 == 1 {
                enclosedTileCount++
            }
        }
    }

    return enclosedTileCount
}
