package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nttu-ysc/sudoku/pkg/sudoku"

	"github.com/pkg/term"
)

const (
	up     byte = 65
	down   byte = 66
	right  byte = 67
	left   byte = 68
	escape byte = 27
	enter  byte = 13
	ctrlC  byte = 3
)

var keys = map[byte]bool{
	up:    true,
	down:  true,
	right: true,
	left:  true,
}

type Cli struct {
	Sudoku    [9][9]int
	CursorPos [2]int
}

func NewCli() *Cli {
	return &Cli{
		Sudoku:    [9][9]int{},
		CursorPos: [2]int{0, 0},
	}
}

func (c *Cli) Display() {
	// turn the terminal cursor off
	fmt.Printf("\033[?25l")
	// show the terminal cursor
	defer fmt.Printf("\033[?25h")
	fmt.Printf(`
                      _____             __        __         
                     / ___/ __  __ ____/ /____   / /__ __  __
                     \__ \ / / / // __  // __ \ / //_// / / /
                    ___/ // /_/ // /_/ // /_/ // ,<  / /_/ / 
                   /____/ \__,_/ \__,_/ \____//_/|_| \__,_/  
                                                             
                       __        _                  __                 __
         _____ ____   / /_   __ (_)____   ____ _   / /_ ____   ____   / /
        / ___// __ \ / /| | / // // __ \ / __  /  / __// __ \ / __ \ / / 
       (__  )/ /_/ // / | |/ // // / / // /_/ /  / /_ / /_/ // /_/ // /  
      /____/ \____//_/  |___//_//_/ /_/ \__, /   \__/ \____/ \____//_/   
                                       /____/                            

Author: Shun
GitHub: https://github.com/nttu-ysc/suduku

This is a Sudoku solving tool.
Here are some instructions:

0 represents an empty cell.
You can use the numbers 1 to 9 to fill in the grid.
The arrow keys can be used to move around.
ESC: Exit the tool.
Enter: Start solving.
R: Reset the grid.
`)

	var row int
	row += printGrid(c.Sudoku, c.CursorPos)

	var canInteractive bool = true
	for canInteractive {
		// fmt.Print(row)
		keyCode := getInput()
		for i := 0; i < row; i++ {
			fmt.Print("\033[K\033[A\033[K")
		}
		row = 0
		switch keyCode {
		case escape:
			fallthrough
		case ctrlC:
			canInteractive = false
		case up:
			x := c.CursorPos[0] - 1
			if x >= 0 {
				c.CursorPos[0] = x
			}
		case down:
			x := c.CursorPos[0] + 1
			if x < 9 {
				c.CursorPos[0] = x
			}
		case left:
			y := c.CursorPos[1] - 1
			if y >= 0 {
				c.CursorPos[1] = y
			}
		case right:
			y := c.CursorPos[1] + 1
			if y < 9 {
				c.CursorPos[1] = y
			}
		case '0':
			fallthrough
		case '1':
			fallthrough
		case '2':
			fallthrough
		case '3':
			fallthrough
		case '4':
			fallthrough
		case '5':
			fallthrough
		case '6':
			fallthrough
		case '7':
			fallthrough
		case '8':
			fallthrough
		case '9':
			c.Sudoku[c.CursorPos[0]][c.CursorPos[1]] = int(keyCode - '0')
		case enter:
			if sudoku.IsValidSudoku(c.Sudoku) {
				sudoku.SolveSudoku(&(c.Sudoku))
				canInteractive = false
			} else {
				fmt.Println("\033[41mInvalid Sudoku.\033[0m")
				row++
			}
		case 'R':
			c.Sudoku = [9][9]int{}
		}
		row += printGrid(c.Sudoku, c.CursorPos)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Sudoku has completed.")
}

func getInput() byte {
	t, err := term.Open("/dev/tty")
	if err != nil {
		log.Fatalln(err)
	}

	if err := term.RawMode(t); err != nil {
		log.Fatal(err)
	}

	var read int
	readBytes := make([]byte, 3)
	read, err = t.Read(readBytes)
	if err != nil {
		log.Fatalln(err)
	}

	t.Restore()
	t.Close()

	if read == 3 {
		if _, ok := keys[readBytes[2]]; ok {
			return readBytes[2]
		}
	} else {
		return readBytes[0]
	}
	return 0
}

func printGrid(grid [9][9]int, curPos [2]int) int {
	var row int
	for i := range grid {
		for k := 0; k < 9; k++ {
			if i%3 == 0 {
				fmt.Printf("\033[31m%s\033[0m", "+---")
			} else {
				if k == 0 {
					fmt.Printf("\033[31m+\033[0m---")
				} else {
					fmt.Printf("+---")
				}
			}
		}
		fmt.Println("\033[31m+\033[0m")

		for j := range grid[i] {
			if j%3 == 0 {
				fmt.Printf("\033[31m|\033[0m ")
			} else {
				fmt.Printf("| ")
			}

			if i == curPos[0] && j == curPos[1] {
				fmt.Printf("\033[45m%d\033[0m ", grid[i][j])
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println("\033[31m|\033[0m")
		row += 2
	}
	for i := 0; i < 9; i++ {
		fmt.Printf("\033[31m%s\033[0m", "+---")
	}
	fmt.Println("\033[31m+\033[0m")
	row++
	return row
}
