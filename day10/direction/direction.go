package direction

const NORTH int = 0
const EAST int = 1
const SOUTH int = 2
const WEST int = 3

func IsNorth(d int) bool {
    return d == NORTH
}

func IsEast(d int) bool {
    return d == EAST
}

func IsSouth(d int) bool {
    return d == SOUTH
}

func IsWest(d int) bool {
    return d == WEST
}
