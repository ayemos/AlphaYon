// Code generated by "stringer -type GameStatus"; DO NOT EDIT

package alphaYon

import "fmt"

const _GameStatus_name = "WHITE_WONBLACK_WONRUNNINGDRAW"

var _GameStatus_index = [...]uint8{0, 9, 18, 25, 29}

func (i GameStatus) String() string {
	if i < 0 || i >= GameStatus(len(_GameStatus_index)-1) {
		return fmt.Sprintf("GameStatus(%d)", i)
	}
	return _GameStatus_name[_GameStatus_index[i]:_GameStatus_index[i+1]]
}
