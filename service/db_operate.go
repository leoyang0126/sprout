package service

import (
	"context"
	"sprout/app/system"
)

type IDbOperate interface {
	Add(ctx context.Context, req *system.DictDataAddReq, userId uint64) (err error)
	Get(ctx context.Context, dictCode uint) (res *system.DictDataGetRes, err error)
	Edit(ctx context.Context, req *system.DictDataEditReq, userId uint64) (err error)
	Delete(ctx context.Context, ids []int) (err error)
}

var iDbOperate IDbOperate

func GetIDbOperate() IDbOperate {
	if iDbOperate == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return iDbOperate
}

func RegisterIDbOperate(i IDbOperate) {
	iDbOperate = i
}
