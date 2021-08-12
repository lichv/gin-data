package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type PointHospitalTencent struct {
	Id        string    `json:"id" form:"id" gorm:"id"`
	Name      string `json:"name" form:"name" gorm:"name"`
	Address   string `json:"address" form:"address" gorm:"address"`
	Adcode    string `json:"adcode" form:"adcode" gorm:"adcode"`
	Province  string `json:"province" form:"province" gorm:"province"`
	City      string `json:"city" form:"city" gorm:"city"`
	Area      string `json:"area" form:"area" gorm:"area"`
	Lat       string `json:"lat" form:"lat" gorm:"lat"`
	Lng       string `json:"lng" form:"lng" gorm:"lng"`
	Grade     string `json:"grade" form:"grade" gorm:"grade"`
	Natrue    string `json:"natrue" form:"nature" gorm:"natrue"`
	Type      string `json:"type" form:"type" gorm:"type"`
	Telephone string `json:"telephone" form:"telephone" gorm:"telephone"`
	Flag      int    `json:"flag" form:"flag" gorm:"flag"`
	State     int    `json:"state" form:"state" gorm:"state"`
	Status    int    `json:"status" form:"status" gorm:"status"`
}

var PointHospitalTencentFields = []string{"id", "name", "address", "adcode", "province", "city", "area", "lat", "lng", "grade", "nature", "type", "telephone", "flag", "state", "status"}

func FindPointHospitalTencentById(id string) (*PointHospitalTencent, error) {
	var pointHospitalTencent PointHospitalTencent
	err := db.Model(&PointHospitalTencent{}).Where("id = ? ", id).First(&pointHospitalTencent).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospitalTencent{}, err
	}
	return &pointHospitalTencent, nil
}

func GetPointHospitalTencentOne(query map[string]interface{}, orderBy interface{}) (*PointHospitalTencent, error) {
	var pointHospitalTencent PointHospitalTencent
	model := db.Model(&PointHospitalTencent{})
	for key, value := range query {
		if lichv.In(PointHospitalBaiduFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  key == "adcode" {
				model = model.Where(key +" like ?",lichv.Strval(value)+"%")
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	err := model.Order(orderBy).First(&pointHospitalTencent).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospitalTencent{}, nil
	}
	return &pointHospitalTencent, nil
}

func GetPointHospitalTencentPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*PointHospitalTencent, error) {
	var pointHospitalTencents []*PointHospitalTencent
	var errs []error
	var err error

	model := db.Model(&PointHospitalTencent{})
	for key, value := range query {
		if lichv.In(PointHospitalBaiduFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  key == "adcode" {
				model = model.Where(key +" like ?",lichv.Strval(value)+"%")
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}

	}
	beginNum := (pageNum - 1) * pageSize
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&pointHospitalTencents)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return pointHospitalTencents, err
}

func GetPointHospitalTencentList(query map[string]interface{}, orderBy interface{}, limit int) ([]*PointHospitalTencent, []error) {
	var PointHospitalTencents []*PointHospitalTencent
	var errs []error
	model := db.Model(&PointHospitalTencent{})
	for key, value := range query {
		if lichv.In(PointHospitalBaiduFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  key == "adcode" {
				model = model.Where(key +" like ?",lichv.Strval(value)+"%")
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&PointHospitalTencents).GetErrors()

	return PointHospitalTencents, errs
}

func GetPointHospitalTencentTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&PointHospitalTencent{})
	for key, value := range query {
		if lichv.In(PointHospitalBaiduFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  key == "adcode" {
				model = model.Where(key +" like ?",lichv.Strval(value)+"%")
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	err = model.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}
func AddPointHospitalTencent(data map[string]interface{}) (string, error) {
	pointHospitalTencent := PointHospitalTencent{}
	_, ok := data["id"]
	if ok {
		pointHospitalTencent.Id = lichv.Strval(data["id"])
	}
	_, ok = data["name"]
	if ok {
		pointHospitalTencent.Name = lichv.Strval(data["name"])
	}
	_, ok = data["address"]
	if ok {
		pointHospitalTencent.Address = lichv.Strval(data["address"])
	}
	_, ok = data["adcode"]
	if ok {
		pointHospitalTencent.Adcode = lichv.Strval(data["adcode"])
	}
	_, ok = data["province"]
	if ok {
		pointHospitalTencent.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		pointHospitalTencent.City = lichv.Strval(data["city"])
	}
	_, ok = data["area"]
	if ok {
		pointHospitalTencent.Area = lichv.Strval(data["area"])
	}
	_, ok = data["lat"]
	if ok {
		pointHospitalTencent.Lat = lichv.Strval(data["lat"])
	}
	_, ok = data["lng"]
	if ok {
		pointHospitalTencent.Lng = lichv.Strval(data["lng"])
	}
	_, ok = data["grade"]
	if ok {
		pointHospitalTencent.Grade = lichv.Strval(data["grade"])
	}
	_, ok = data["natrue"]
	if ok {
		pointHospitalTencent.Natrue = lichv.Strval(data["natrue"])
	}
	_, ok = data["type"]
	if ok {
		pointHospitalTencent.Type = lichv.Strval(data["type"])
	}
	_, ok = data["telephone"]
	if ok {
		pointHospitalTencent.Telephone = lichv.Strval(data["telephone"])
	}
	_, ok = data["flag"]
	if ok {
		pointHospitalTencent.Flag = lichv.IntVal(data["flag"])
	} else {
		pointHospitalTencent.Flag = 0
	}
	_, ok = data["state"]
	if ok {
		pointHospitalTencent.State = lichv.IntVal(data["state"])
	} else {
		pointHospitalTencent.State = 0
	}
	_, ok = data["status"]
	if ok {
		pointHospitalTencent.Status = lichv.IntVal(data["status"])
	} else {
		pointHospitalTencent.Status = 0
	}
	fmt.Println(pointHospitalTencent)
	if pointHospitalTencent.Name == "" {
		return "", errors.New("参数为空")
	}

	if err := db.Create(&pointHospitalTencent).Error; err != nil {
		return "", err
	}
	return pointHospitalTencent.Id, nil
}

func ModifyPointHospitalTencentById(id string, data map[string]interface{}) error {
	if err := db.Model(&PointHospitalTencent{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitalTencentById(id string) error {
	if err := db.Where("id=?", id).Delete(PointHospitalTencent{}).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitalTencents(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(PointHospitalBaiduFields, key) {
			if key == "name" {
				val := lichv.Strval(value)
				if val != "" {
					model = model.Where("name like ?", "%"+val+"%")
				}
			}else if  lichv.In([]string{"flag","state","status"},key) {
				if value == 0{
					model = model.Where(key +"=?",0)
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}
	}
	if err := model.Delete(&PointHospitalTencent{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllPointHospitalTencent() error {
	if err := db.Unscoped().Delete(&PointHospitalTencent{}).Error; err != nil {
		return err
	}
	return nil
}
