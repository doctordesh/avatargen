package main

import (
    "errors"
    "math/rand"
)

type Direction int

const (
    DirectionUp        Direction = 0
    DirectionUpRight   Direction = 1
    DirectionRight     Direction = 2
    DirectionDownRight Direction = 3
    DirectionDown      Direction = 4
    DirectionDownLeft  Direction = 5
    DirectionLeft      Direction = 6
    DirectionUpLeft    Direction = 7
)

func (d Direction) String() string {
    if d == DirectionUp {
        return "DirectionUp"
    }
    if d == DirectionUpRight {
        return "DirectionUpRight"
    }
    if d == DirectionRight {
        return "DirectionRight"
    }
    if d == DirectionDownRight {
        return "DirectionDownRight"
    }
    if d == DirectionDown {
        return "DirectionDown"
    }
    if d == DirectionDownLeft {
        return "DirectionDownLeft"
    }
    if d == DirectionLeft {
        return "DirectionLeft"
    }
    if d == DirectionUpLeft {
        return "DirectionUpLeft"
    }

    return "Unknown Direction"
}

func RandomDirection() Direction {
    return Direction(rand.Intn(8))
}

type SetOfDirections struct {
    list []Direction
}

func (s *SetOfDirections) Add(d Direction) {
    for i := range s.list {
        if s.list[i] == d {
            return
        }
    }

    s.list = append(s.list, d)
}

func (s *SetOfDirections) Contains(d Direction) bool {
    for i := range s.list {
        if s.list[i] == d {
            return true
        }
    }

    return false
}

func (s *SetOfDirections) Remove(d Direction) {
    for k, v := range s.list {
        if v == d {
            s.list = append(s.list[:k], s.list[k+1:]...)
        }
    }
}

func (s *SetOfDirections) Len() int {
    return len(s.list)
}

type Point struct {
    X int
    Y int
}

func MovePointInDirection(p Point, dir Direction, min_x, min_y, max_x, max_y int) (Point, error) {
    err := errors.New("Position out of bounds")

    if dir == DirectionUp {
        p.Y -= 1
    }
    if dir == DirectionDown {
        p.Y += 1
    }
    if dir == DirectionRight {
        p.X += 1
    }
    if dir == DirectionLeft {
        p.X -= 1
    }

    if dir == DirectionUpRight {
        p.X += 1
        p.Y -= 1
    }
    if dir == DirectionDownRight {
        p.X += 1
        p.Y += 1
    }
    if dir == DirectionDownLeft {
        p.X -= 1
        p.Y += 1
    }
    if dir == DirectionUpLeft {
        p.X -= 1
        p.Y -= 1
    }

    if p.X < min_x {
        return p, err
    }

    if p.X > max_x {
        return p, err
    }

    if p.Y < min_y {
        return p, err
    }

    if p.Y > max_y {
        return p, err
    }

    return p, nil
}
