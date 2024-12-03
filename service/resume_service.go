package service

import (
	"log"
	"sprout/repository"
	"sprout/repository/entity"
)

func QueryResumeBaseList() {
	var resumeBase entity.AtsResumeBase
	repository.DB.Table("ats_resume_base").Where("id = ?", 1).Find(&resumeBase)

	log.Println("查询到数据 ", resumeBase)
	var resumeBaseList []entity.AtsResumeBase
	repository.DB.Table("ats_resume_base").Find(&resumeBaseList)
}

// 插入简历基本信息
func InsertResumeBase(resumeBase entity.AtsResumeBase) error {
	tx := repository.DB.Table("ats_resume_base").Create(&resumeBase)
	if tx.Error != nil {
		log.Println("insert user in do fail")
		return tx.Error
	}
	return nil
}
