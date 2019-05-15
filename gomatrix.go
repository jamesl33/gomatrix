/* This file is part of gomatrix.

Copyright (C) 2019, James Lee <jamesl33info@gmail.com>.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main

import (
    "fmt"
    "math/rand"
    "time"
    "bytes"
)

type MatrixRain struct {
    width int
    height int
    view [][] string
}

func NewMatrixRain(width int, height int) *MatrixRain {
    view := make([][] string, height)

    for slice := range view {
        view[slice] = make([] string, width)

        for letter := range view[slice] {
            rand.Seed(time.Now().UnixNano())

            if (rand.Intn(2) == 0) {
                view[slice][letter] = string('a' + rand.Intn(26))
            } else {
                view[slice][letter] = string('A' + rand.Intn(26))
            }
        }
    }

    return &MatrixRain{width: width, height: height, view: view}
}

func (matrixrain *MatrixRain) String() string {
    var buf bytes.Buffer

    for y := 0; y < matrixrain.height; y++ {
        for x := 0; x < matrixrain.width; x++ {
            buf.WriteByte(byte([] rune(matrixrain.view[y][x])[0]))
        }

        buf.WriteByte('\n')
    }

    return buf.String()
}

func (matrixrain *MatrixRain) Step() {
    for index := matrixrain.height - 1; index > 0; index-- {
        matrixrain.view[index], matrixrain.view[index - 1] = matrixrain.view[index - 1], matrixrain.view[index]
    }

    for letter := range matrixrain.view[0] {
        rand.Seed(time.Now().UnixNano())

        if (rand.Intn(2) == 0) {
            matrixrain.view[0][letter] = string('a' + rand.Intn(26))
        } else {
            matrixrain.view[0][letter] = string('A' + rand.Intn(26))
        }
    }
}

func main() {
    view := NewMatrixRain(80, 24)

    for {
        fmt.Print("\x0c", view)
        view.Step()
        time.Sleep(time.Second / 30)
    }
}
