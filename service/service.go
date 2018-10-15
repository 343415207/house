package service

import (
	"errors"

	"house/controller/form"
	"house/model"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//LoginOrCreateService ...
func LoginOrCreateService(form *form.Login) (*model.User, error) {
	u := &model.User{}
	if err := model.GetInfo(u, map[string]interface{}{"phone": form.User}); err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := model.SaveOrUpdate(u,
				map[string]interface{}{"phone": form.User, "pwd": form.Password},
				map[string]interface{}{"phone": form.User}); err != nil {
				return nil, err
			}
			return u, nil
		}
		return nil, err
	}

	if u.Pwd == form.Password {
		return u, nil
	} else {
		return nil, errors.New("密码错误")
	}

	return nil, errors.New("未知登录错误")

}

//UpPwdService ...
func UpPwdService(form *form.Login) (*model.User, error) {
	u := &model.User{}
	if err := model.GetInfo(u, map[string]interface{}{"phone": form.User}); err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("无此用户")
		}
		return nil, err
	}
	if err := model.SaveOrUpdate(u,
		map[string]interface{}{"phone": form.User, "pwd": form.Password},
		map[string]interface{}{"phone": form.User}); err != nil {
		return nil, err
	}
	return u, nil
}

//GetPListService ...
func GetPListService(fm *form.PListForm) ([]*model.Project, error) {
	entry := logrus.WithField("api", "GetPListService")
	entry.Info(fm)
	l := make([]*model.Project, 0)

	if err := model.DB().Find(l).Where(map[string]interface{}{
		"is_deleted": 0,
	}).Error; err != nil {
		entry.Infof("error for got projects %+v", fm)
		return nil, err
	}
	return l, nil
}

//SaveOrUpdateOrderService ...
func SaveOrUpdateOrderService(fm *form.OrderForm) error {
	info := &model.COrder{}
	info.CPhone = fm.CPhone
	info.CName = fm.CName
	info.CSex = fm.CSex
	info.Desc = fm.Desc
	info.EscorteType = fm.EscorteType
	//info.AccessTime = fm.AccessTime
	if err := model.GetInfo(info, map[string]interface{}{"c_phone": fm.CPhone}); err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.DB().Create(info).Error
		}
		return err
	}
	return model.DB().Save(info).Error
}
