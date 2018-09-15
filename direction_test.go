package main

import "testing"

func checkError(t *testing.T, err error) {
    if err != nil {
        t.Error(err.Error())
    }
}

func TestMoveInDirection(t *testing.T) {
    var err error
    p0 := Point{X: 5, Y: 5}
    p1 := Point{}

    p1, err = MovePointInDirection(p0, DirectionUp, 0, 0, 10, 10)
    t.Log(p1)
    checkError(t, err)

    if p1.Y != p0.Y-1 {
        t.Fail()
    }
}
