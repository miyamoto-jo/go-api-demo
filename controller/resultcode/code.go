package resultcode

// api独自の結果コード
type ResultCode int

const (
	// CodeSuccess API成功
	CodeSuccess = ResultCode(100)

	// CodeErrorFailure システムエラー
	CodeErrorFailure = ResultCode(900)
	// resultcode.CodeRequestParseFailure リクエストに問題があるエラー
	CodeRequestParseFailure = ResultCode(910)
	// CodeErrorPanicFatal panicエラー
	CodeErrorPanicFatal = ResultCode(999)
)
