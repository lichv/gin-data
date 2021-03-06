package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
)

type PointHospitalGaode struct {
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

var PointHospitalGaodeFields = []string{"id", "name", "address", "adcode", "province", "city", "area", "lat", "lng", "grade", "nature", "type", "telephone", "flag", "state", "status"}

func FindPointHospitalGaodeById(id string) (*PointHospitalGaode, error) {
	var pointHospitalBaidu PointHospitalGaode
	err := db.Model(&PointHospitalGaode{}).Where("id = ? ", id).First(&pointHospitalBaidu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospitalGaode{}, err
	}
	return &pointHospitalBaidu, nil
}

func GetPointHospitalGaodeOne(query map[string]interface{}, orderBy interface{}) (*PointHospitalGaode, error) {
	var pointHospitalBaidu PointHospitalGaode
	model := db.Model(&PointHospitalGaode{})
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
	err := model.Order(orderBy).First(&pointHospitalBaidu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &PointHospitalGaode{}, nil
	}
	return &pointHospitalBaidu, nil
}

func GetPointHospitalGaodePages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*PointHospitalGaode, error) {
	var pointHospitalBaidus []*PointHospitalGaode
	var errs []error
	var err error

	model := db.Model(&PointHospitalGaode{})
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
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&pointHospitalBaidus)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return pointHospitalBaidus, err
}

func GetPointHospitalGaodeList(query map[string]interface{}, orderBy interface{}, limit int) ([]*PointHospitalGaode, []error) {
	var PointHospitalGaodes []*PointHospitalGaode
	var errs []error
	model := db.Model(&PointHospitalGaode{})
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
	errs = model.Order(orderBy).Find(&PointHospitalGaodes).GetErrors()

	return PointHospitalGaodes, errs
}

func GetPointHospitalGaodeTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&PointHospitalGaode{})
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
func AddPointHospitalGaode(data map[string]interface{}) (string, error) {
	pointHospitalBaidu := PointHospitalGaode{}
	_, ok := data["id"]
	if ok {
		pointHospitalBaidu.Id = lichv.Strval(data["id"])
	}
	_, ok = data["name"]
	if ok {
		pointHospitalBaidu.Name = lichv.Strval(data["name"])
	}
	_, ok = data["address"]
	if ok {
		pointHospitalBaidu.Address = lichv.Strval(data["address"])
	}
	_, ok = data["adcode"]
	if ok {
		pointHospitalBaidu.Adcode = lichv.Strval(data["adcode"])
	}
	_, ok = data["province"]
	if ok {
		pointHospitalBaidu.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		pointHospitalBaidu.City = lichv.Strval(data["city"])
	}
	_, ok = data["area"]
	if ok {
		pointHospitalBaidu.Area = lichv.Strval(data["area"])
	}
	_, ok = data["lat"]
	if ok {
		pointHospitalBaidu.Lat = lichv.Strval(data["lat"])
	}
	_, ok = data["lng"]
	if ok {
		pointHospitalBaidu.Lng = lichv.Strval(data["lng"])
	}
	_, ok = data["grade"]
	if ok {
		pointHospitalBaidu.Grade = lichv.Strval(data["grade"])
	}
	_, ok = data["natrue"]
	if ok {
		pointHospitalBaidu.Natrue = lichv.Strval(data["natrue"])
	}
	_, ok = data["type"]
	if ok {
		pointHospitalBaidu.Type = lichv.Strval(data["type"])
	}
	_, ok = data["telephone"]
	if ok {
		pointHospitalBaidu.Telephone = lichv.Strval(data["telephone"])
	}
	_, ok = data["flag"]
	if ok {
		pointHospitalBaidu.Flag = lichv.IntVal(data["flag"])
	} else {
		pointHospitalBaidu.Flag = 0
	}
	_, ok = data["state"]
	if ok {
		pointHospitalBaidu.State = lichv.IntVal(data["state"])
	} else {
		pointHospitalBaidu.State = 0
	}
	_, ok = data["status"]
	if ok {
		pointHospitalBaidu.Status = lichv.IntVal(data["status"])
	} else {
		pointHospitalBaidu.Status = 0
	}
	fmt.Println(pointHospitalBaidu)
	if pointHospitalBaidu.Name == "" {
		return "", errors.New("????????????")
	}

	if err := db.Create(&pointHospitalBaidu).Error; err != nil {
		return "", err
	}
	return pointHospitalBaidu.Id, nil
}

func ModifyPointHospitalGaodeById(id string, data map[string]interface{}) error {
	if err := db.Model(&PointHospitalGaode{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitalGaodeById(id string) error {
	if err := db.Where("id=?", id).Delete(PointHospitalGaode{}).Error; err != nil {
		return err
	}
	return nil
}

func DeletePointHospitalGaodes(maps map[string]interface{}) error {
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
	if err := model.Delete(&PointHospitalGaode{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllPointHospitalGaode() error {
	if err := db.Unscoped().Delete(&PointHospitalGaode{}).Error; err != nil {
		return err
	}
	return nil
}
