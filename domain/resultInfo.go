package domain

// ResultInfo 统一数据返回类型
type ResultInfo struct {
	Code int         `json:"code,1" thrift:"code"`
	Mes  string      `json:"mes,2" thrift:"mes"`
	Data interface{} `json:"data,3" thrift:"data"`
}

func NewResultInfo(code int, mes string, data interface{}) *ResultInfo {
	return &ResultInfo{Code: code, Mes: mes, Data: data}
}
