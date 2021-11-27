// Code generated by "stringer -type Code -linecomment ./pkg/errcode"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ParamErr-1000]
	_ = x[SourceNotFind-1001]
	_ = x[SystemErr-1002]
	_ = x[CopierErr-1003]
	_ = x[WechatErr-1004]
	_ = x[JwtErr-1005]
	_ = x[MarshalErr-1006]
	_ = x[SQLErr-1007]
	_ = x[RedisErr-1008]
	_ = x[AuthErr-1009]
	_ = x[YunPianErr-1010]
}

const _Code_name = "参数错误资源不存在系统错误复制错误微信错误Token错误数据格式化错误数据库错误Redis错误请登录云片网错误"

var _Code_index = [...]uint8{0, 12, 27, 39, 51, 63, 74, 95, 110, 121, 130, 145}

func (i Code) String() string {
	i -= 1000
	if i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i+1000), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}
