// Code generated by "stringer -type=ConstValueType"; DO NOT EDIT.

package meta

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Undefined-0]
	_ = x[Integer-1]
	_ = x[Float-2]
	_ = x[String-3]
	_ = x[Bool-4]
	_ = x[Array-5]
}

const _ConstValueType_name = "UndefinedIntegerFloatStringBoolArray"

var _ConstValueType_index = [...]uint8{0, 9, 16, 21, 27, 31, 36}

func (i ConstValueType) String() string {
	if i >= ConstValueType(len(_ConstValueType_index)-1) {
		return "ConstValueType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ConstValueType_name[_ConstValueType_index[i]:_ConstValueType_index[i+1]]
}
