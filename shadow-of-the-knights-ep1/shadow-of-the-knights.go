package main

import ( 
    "fmt"
    "strings"
) 

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: width of the building.
    // H: height of the building.
    var W, H int
    fmt.Scan(&W, &H)
    
    // N: maximum number of turns before game over.
    var N int
    fmt.Scan(&N)
    
    var X0, Y0 int
    fmt.Scan(&X0, &Y0)

    var minx, miny  = 0, 0
    var maxx = W - 1 
    var maxy = H - 1

    for {
        
        // bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        var bombDir string
        fmt.Scan(&bombDir)

        if(strings.Contains(bombDir, "U")) {
            maxy = Y0 - 1
        } else if(strings.Contains(bombDir, "D")) {
            miny = Y0 + 1
        }
        
        if (strings.Contains(bombDir, "L")) {
            maxx = X0 - 1
        } else if (strings.Contains(bombDir, "R")) {
            minx = X0 + 1
        }

        X0 = minx + ((maxx  - minx)/2)
        Y0 = miny + ((maxy  - miny)/2)
        
        // the location of the next window Batman should jump to.
        fmt.Println(X0,Y0)
    }
}
