package serializer

//  用户序列化器
type BindInfo struct {
	ErrCode  int    `json:"errcode"`
	IsBind   int    `json:"is_bind"`
	CorpCode string `json:"corp_code"`
}

//  序列化
func BuildBindInfo(errCode int, isBind int, corpCode string) BindInfo {
	return BindInfo{
		ErrCode:  errCode,
		IsBind:   isBind,
		CorpCode: corpCode,
	}
}

func BuildBindInfoResponse(errCode int, isBind int, corpCode string) Response {
	return Response{
		Data: BuildBindInfo(errCode, isBind, corpCode),
	}
}
