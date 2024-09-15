# Simple algorithm visualizer in GO
Currently **WIP**

## Compile:
To build the program run:
```
go build .
```

# Usage:
The defaults are:
- alg = bubblesort (only bubblesort is implemented for now)
- ms = 50
- size = 10
- status = true

## Examples:
Default behaviour
```
./visualizer
```
--- 
Custom behaviour
```
./visualizer -ms 75 -size 15 -status=false
```
---
To see all the options:
```
./visualizer -h
```

# Known problems
The output will become disorganized if the array is printed in an area larger than the screen size.

Also horizontal visualization may work better.
