package tile

func IsStartingPosition(t string) bool {
    return t == "S"
}

func IsGround(t string) bool {
    return t == "."
}

func IsVerticalPipe(t string) bool {
    return t == "|"
}

func IsHorizontalPipe(t string) bool {
    return t == "-"
}

func IsNorthEastBend(t string) bool {
    return t == "L"
}

func IsNorthWestBend(t string) bool {
    return t == "J"
}

func IsSouthWestBend(t string) bool {
    return t == "7"
}

func IsSouthEastBend(t string) bool {
    return t == "F"
}

func IsNorthOpen(t string) bool {
    return IsVerticalPipe(t) || IsNorthEastBend(t) || IsNorthWestBend(t)
}

func IsEastOpen(t string) bool {
    return IsHorizontalPipe(t) || IsNorthEastBend(t) || IsSouthEastBend(t)
}

func IsSouthOpen(t string) bool {
    return IsVerticalPipe(t) || IsSouthEastBend(t)|| IsSouthWestBend(t)
}

func IsWestOpen(t string) bool {
    return IsHorizontalPipe(t) || IsSouthWestBend(t) || IsNorthWestBend(t)
}
