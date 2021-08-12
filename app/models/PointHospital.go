package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type PointHospital struct {
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

var PointHospitalFields = []string{"id", "name", "address", "adcode", "province", "city", "area", "lat", "lng", "grade", "nature", "type", "telephone", "flag", "state", "status"}

func FindPointHospitalById(id string) (*PointHospital, error) {
	var pointHospital PointHospital
	err := db.Model(&PointHospital{}).Where("id = ? ", id).First(&pointHospital).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospital{}, err
	}
	return &pointHospital, nil
}

func GetPointHospitalOne(query map[string]interface{}, orderBy interface{}) (*PointHospital, error) {
	var pointHospital PointHospital
	model := db.Model(&PointHospital{})
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
	err := model.Order(orderBy).First(&pointHospital).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospital{}, nil
	}
	return &pointHospital, nil
}

func GetPointHospitalPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*PointHospital, error) {
	var pointHospitals []*PointHospital
	var errs []error
	var err error

	model := db.Model(&PointHospital{})
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
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&pointHospitals)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return pointHospitals, err
}

func GetPointHospitalList(query map[string]interface{}, orderBy interface{}, limit int) ([]*PointHospital, []error) {
	var PointHospitals []*PointHospital
	var errs []error
	model := db.Model(&PointHospital{})
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
	errs = model.Order(orderBy).Find(&PointHospitals).GetErrors()

	return PointHospitals, errs
}

func GetPointHospitalTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&PointHospital{})
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
func AddPointHospital(data map[string]interface{}) (string, error) {
	pointHospital := PointHospital{}
	_, ok := data["id"]
	if ok {
		pointHospital.Id = lichv.Strval(data["id"])
	}
	_, ok = data["name"]
	if ok {
		pointHospital.Name = lichv.Strval(data["name"])
	}
	_, ok = data["address"]
	if ok {
		pointHospital.Address = lichv.Strval(data["address"])
	}
	_, ok = data["adcode"]
	if ok {
		pointHospital.Adcode = lichv.Strval(data["adcode"])
	}
	_, ok = data["province"]
	if ok {
		pointHospital.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		pointHospital.City = lichv.Strval(data["city"])
	}
	_, ok = data["area"]
	if ok {
		pointHospital.Area = lichv.Strval(data["area"])
	}
	_, ok = data["lat"]
	if ok {
		pointHospital.Lat = lichv.Strval(data["lat"])
	}
	_, ok = data["lng"]
	if ok {
		pointHospital.Lng = lichv.Strval(data["lng"])
	}
	_, ok = data["grade"]
	if ok {
		pointHospital.Grade = lichv.Strval(data["grade"])
	}
	_, ok = data["natrue"]
	if ok {
		pointHospital.Natrue = lichv.Strval(data["natrue"])
	}
	_, ok = data["type"]
	if ok {
		pointHospital.Type = lichv.Strval(data["type"])
	}
	_, ok = data["telephone"]
	if ok {
		pointHospital.Telephone = lichv.Strval(data["telephone"])
	}
	_, ok = data["flag"]
	if ok {
		pointHospital.Flag = lichv.IntVal(data["flag"])
	} else {
		pointHospital.Flag = 0
	}
	_, ok = data["state"]
	if ok {
		pointHospital.State = lichv.IntVal(data["state"])
	} else {
		pointHospital.State = 0
	}
	_, ok = data["status"]
	if ok {
		pointHospital.Status = lichv.IntVal(data["status"])
	} else {
		pointHospital.Status = 0
	}
	fmt.Println(pointHospital)
	if pointHospital.Name == "" {
		return "", errors.New("参数为空")
	}

	if err := db.Create(&pointHospital).Error; err != nil {
		return "", err
	}
	return pointHospital.Id, nil
}

func ModifyPointHospitalById(id string, data map[string]interface{}) error {
	if err := db.Model(&PointHospital{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitalById(id string) error {
	if err := db.Where("id=?", id).Delete(PointHospital{}).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitals(maps map[string]interface{}) error {
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
	if err := model.Delete(&PointHospital{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllPointHospital() error {
	if err := db.Unscoped().Delete(&PointHospital{}).Error; err != nil {
		return err
	}
	return nil
}
