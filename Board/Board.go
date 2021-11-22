package Board

import "fmt"

type Board struct {
	Board [3][3]string
}

func (b *Board) GetBoardString() {
	currentBoard := fmt.Sprintf(`
 %1s | %1s | %1s 
 %1s | %1s | %1s 
 %1s | %1s | %1s `, b.Board[0][0], b.Board[0][1], b.Board[0][2],
					b.Board[1][0], b.Board[1][1], b.Board[1][2],
					b.Board[2][0], b.Board[2][1], b.Board[2][2],)

	fmt.Println(currentBoard)
}

func (b *Board) GetBoardArray() [3][3]string {
	return b.Board
}

func (b *Board) updateBoardRecord(x, y int, val string) {
	b.Board[x][y] = val
}

func (b *Board) UpdateBoard(x, y int, val string) error{
	if b.Board[x][y] != "" {
		err := fmt.Errorf("This cell [%d,%d] is occupied", x+1, y+1)
		return err
	} else if val != "O" && val != "X" {
		err := fmt.Errorf("Bad sign %s", val)
		return err
	} else {
		b.updateBoardRecord(x, y, val)
		return nil
	}
}