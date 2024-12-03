package service

import (
	"context"
	"sprout/app/system"
)

type SystemDictInfo struct {
}

func aa() {
	systemDict := new(SystemDictInfo)
	RegisterIDbOperate(systemDict)
}

func (*SystemDictInfo) Add(ctx context.Context, req *system.DictDataAddReq, userId uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (*SystemDictInfo) Get(ctx context.Context, dictCode uint) (res *system.DictDataGetRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (*SystemDictInfo) Edit(ctx context.Context, req *system.DictDataEditReq, userId uint64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (*SystemDictInfo) Delete(ctx context.Context, ids []int) (err error) {
	//TODO implement me
	panic("implement me")
}
