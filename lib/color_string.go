// Code generated by "stringer -type Color"; DO NOT EDIT

package alphaYon

import "fmt"

const _Color_name = "EMPTYWHITEBLACK"

var _Color_index = [...]uint8{0, 5, 10, 15}

func (i Color) String() string {
	if i < 0 || i >= Color(len(_Color_index)-1) {
		return fmt.Sprintf("Color(%d)", i)
	}
	return _Color_name[_Color_index[i]:_Color_index[i+1]]
}
