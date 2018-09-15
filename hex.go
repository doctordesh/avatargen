package main

import (
    "fmt"
    "image/color"
    "strconv"
)

func HexStringToRGB(s string) (color.RGBA, error) {
    var c color.RGBA
    var err error
    var r, g, b int64

    if len(s) != 6 {
        return c, fmt.Errorf("Invalid HEX string")
    }

    r, err = hexPairToInt(string(s[0]), string(s[1]))
    if err != nil {
        return c, err
    }

    g, err = hexPairToInt(string(s[2]), string(s[3]))
    if err != nil {
        return c, err
    }

    b, err = hexPairToInt(string(s[4]), string(s[5]))
    if err != nil {
        return c, err
    }

    c.R = uint8(r)
    c.G = uint8(g)
    c.B = uint8(b)
    c.A = 255

    return c, nil
}

func hexPairToInt(s1, s2 string) (int64, error) {
    color_part, err := strconv.ParseInt(s1, 16, 8)
    if err != nil {
        return 0, err
    }
    i := color_part << 4
    color_part, err = strconv.ParseInt(s2, 16, 8)
    if err != nil {
        return 0, err
    }
    i = i | color_part

    return i, nil
}
