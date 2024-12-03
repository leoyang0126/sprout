package system

// DictDataAddReq 分页请求参数
type DictDataAddReq struct {
	Name string
	Age  int
}

// DictDataGetRes 分页请求参数
type DictDataGetRes struct {
	Name string
	Age  int
}

// DictDataAddReq 分页请求参数
type DictDataEditReq struct {
	Id   int64
	Name string
	Age  int
}
