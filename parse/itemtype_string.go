// generated by stringer -type=itemType; DO NOT EDIT

package parse

import "fmt"

const _itemType_name = "itemErroritemEOFitemNameitemLeftCurlyitemRightCurlyitemLeftParenitemRightParenitemNumberitemColonitemCommaitemStringitemDotitemNilitemKeyworditemQueryitemMutationitemFragmentitemOnitemEllipsesitemTrueitemFalseitemComment"

var _itemType_index = [...]uint8{0, 9, 16, 24, 37, 51, 64, 78, 88, 97, 106, 116, 123, 130, 141, 150, 162, 174, 180, 192, 200, 209, 220}

func (i itemType) String() string {
	if i < 0 || i+1 >= itemType(len(_itemType_index)) {
		return fmt.Sprintf("itemType(%d)", i)
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}
