// generated by stringer -type=op; DO NOT EDIT

package sesstype

import "fmt"

const _op_name = "NoOpNewChanOpSendOpRecvOpEndOp"

var _op_index = [...]uint8{0, 4, 13, 19, 25, 30}

func (i op) String() string {
	if i < 0 || i >= op(len(_op_index)-1) {
		return fmt.Sprintf("op(%d)", i)
	}
	return _op_name[_op_index[i]:_op_index[i+1]]
}
