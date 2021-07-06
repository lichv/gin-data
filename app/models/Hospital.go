package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
	"time"
)

type Hospital struct {
	Id           int    `json:"id" form:"id" gorm:"id"`
	Code         string `json:"code" form:"code" gorm:"code"`
	Name         string `json:"name" form:"name" gorm:"name"`
	Alias        string `json:"alias" form:"alias" gorm:"alias"`
	Logo         string `json:"logo" form:"logo" gorm:"logo"`
	Grade        string `json:"grade" form:"grade" gorm:"grade"`
	Nature       string `json:"nature" form:"nature" gorm:"nature"`
	Type         string `json:"type" form:"type" gorm:"type"`
	Regyear      string `json:"regyear" form:"regyear" gorm:"regyear"`
	Size         string `json:"size" form:"size" gorm:"size"`
	Province     string `json:"province" form:"province" gorm:"province"`
	City         string `json:"city" form:"city" gorm:"city"`
	Town         string `json:"town" form:"town" gorm:"town"`
	Address      string `json:"address" form:"address" gorm:"address"`
	Location     string `json:"location" form:"location" gorm:"location"`
	Members      string `json:"members" form:"members" gorm:"members"`
	Beds         string `json:"beds" form:"beds" gorm:"beds"`
	Emergency    string `json:"emergency" form:"emergency" gorm:"emergency"`
	Inpatients   string `json:"inpatients" form:"inpatients" gorm:"inpatients"`
	Creditcode   string `json:"creditcode" form:"creditcode" gorm:"creditcode"`
	Corporations string `json:"corporations" form:"corporations" gorm:"corporations"`
	Typename     string `json:"typename" form:"typename" gorm:"typename"`
	Introduction string `json:"introduction" form:"introduction" gorm:"introduction"`
	Others       string `json:"others" form:"others" gorm:"others"`
	Source       string `json:"source" form:"source" gorm:"source"`
	Tianyancha   string `json:"tianyancha" form:"tianyancha" gorm:"tianyancha"`
	Yixue        string `json:"yixue" form:"yixue" gorm:"yixue"`
	Url          string `json:"url" form:"url" gorm:"url"`
	Phone        string `json:"phone" form:"phone" gorm:"phone"`
	Email        string `json:"email" form:"email" gorm:"email"`
	Addtime      int64  `json:"addtime" form:"addtime" gorm:"addtime"`
	Flag         int    `json:"flag" form:"flag" gorm:"flag"`
	State        int    `json:"state" form:"state" gorm:"state"`
	Status       int    `json:"status" form:"status" gorm:"status"`
}

var HospitalFields = []string{"id", "code", "name", "alias", "logo", "grade", "nature", "type", "regyear", "size", "province", "city", "town", "address", "location", "members", "beds", "emergency", "inpatients", "creditcode", "corporations", "typename", "introduction", "others", "source", "tianyancha", "yixue", "url", "phone", "email", "addtime", "flag", "state", "status"}

func FindHospitalById(id int) (*Hospital, error) {
	var hospital Hospital
	err := db.Model(&Hospital{}).Where("id = ? ", id).First(&hospital).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Hospital{}, err
	}
	return &hospital, nil
}

func GetHospitalOne(query map[string]interface{}, orderBy interface{}) (*Hospital, error) {
	var hospital Hospital
	model := db.Model(&Hospital{})
	for key, value := range query {
		if lichv.In(HospitalFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	err := model.First(&hospital).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Hospital{}, nil
	}
	return &hospital, nil
}

func GetHospitalPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*Hospital, error) {
	var hospitals []*Hospital
	var errs []error
	var err error

	model := db.Model(&Hospital{})
	for key, value := range query {
		if lichv.In(HospitalFields,key) {
			name := lichv.Strval(value)
			if key == "name" {
				if name != ""{
					model = model.Where("name like ?", "%"+name+"%")
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}

	}
	beginNum := (pageNum - 1) * pageSize
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&hospitals)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return hospitals, err
}

func GetHospitalList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Hospital, []error) {
	var Hospitals []*Hospital
	var errs []error
	model := db.Model(&Hospital{})
	for key, value := range query {
		if lichv.In(HospitalFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Hospitals).GetErrors()

	return Hospitals, errs
}

func GetHospitalTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&Hospital{})
	for key, value := range query {
		if lichv.In(HospitalFields,key) {
			name := lichv.Strval(value)
			if key == "name" {
				if name != ""{
					model = model.Where("name like ?", "%"+name+"%")
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
func AddHospital(data map[string]interface{}) (int, error) {
	hospital := Hospital{}
	_, ok := data["id"]
	if ok {
		hospital.Id = lichv.IntVal(data["id"])
	}
	_, ok = data["code"]
	if ok {
		hospital.Code = lichv.Strval(data["code"])
	}
	_, ok = data["name"]
	if ok {
		hospital.Name = lichv.Strval(data["name"])
	}
	_, ok = data["alias"]
	if ok {
		hospital.Alias = lichv.Strval(data["alias"])
	}
	_, ok = data["logo"]
	if ok {
		hospital.Logo = lichv.Strval(data["logo"])
	}
	_, ok = data["grade"]
	if ok {
		hospital.Grade = lichv.Strval(data["grade"])
	}
	_, ok = data["nature"]
	if ok {
		hospital.Nature = lichv.Strval(data["nature"])
	}
	_, ok = data["type"]
	if ok {
		hospital.Type = lichv.Strval(data["type"])
	}
	_, ok = data["regyear"]
	if ok {
		hospital.Regyear = lichv.Strval(data["regyear"])
	}
	_, ok = data["size"]
	if ok {
		hospital.Size = lichv.Strval(data["size"])
	}
	_, ok = data["province"]
	if ok {
		hospital.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		hospital.City = lichv.Strval(data["city"])
	}
	_, ok = data["town"]
	if ok {
		hospital.Town = lichv.Strval(data["town"])
	}
	_, ok = data["address"]
	if ok {
		hospital.Address = lichv.Strval(data["address"])
	}
	_, ok = data["location"]
	if ok {
		hospital.Location = lichv.Strval(data["location"])
	}
	_, ok = data["members"]
	if ok {
		hospital.Members = lichv.Strval(data["members"])
	}
	_, ok = data["beds"]
	if ok {
		hospital.Beds = lichv.Strval(data["beds"])
	}
	_, ok = data["emergency"]
	if ok {
		hospital.Emergency = lichv.Strval(data["emergency"])
	}
	_, ok = data["inpatients"]
	if ok {
		hospital.Inpatients = lichv.Strval(data["inpatients"])
	}
	_, ok = data["creditcode"]
	if ok {
		hospital.Creditcode = lichv.Strval(data["creditcode"])
	}
	_, ok = data["corporations"]
	if ok {
		hospital.Corporations = lichv.Strval(data["corporations"])
	}
	_, ok = data["typename"]
	if ok {
		hospital.Typename = lichv.Strval(data["typename"])
	}
	_, ok = data["introduction"]
	if ok {
		hospital.Introduction = lichv.Strval(data["introduction"])
	}
	_, ok = data["others"]
	if ok {
		hospital.Others = lichv.Strval(data["others"])
	}
	_, ok = data["source"]
	if ok {
		hospital.Source = lichv.Strval(data["source"])
	}
	_, ok = data["url"]
	if ok {
		hospital.Url = lichv.Strval(data["url"])
	}
	_, ok = data["phone"]
	if ok {
		hospital.Phone = lichv.Strval(data["phone"])
	}
	_, ok = data["email"]
	if ok {
		hospital.Email = lichv.Strval(data["email"])
	}

	hospital.Addtime = time.Now().Unix()

	_, ok = data["flag"]
	if ok {
		hospital.Flag = lichv.IntVal(data["flag"])
	} else {
		hospital.Flag = 1
	}
	_, ok = data["state"]
	if ok {
		hospital.State = lichv.IntVal(data["state"])
	} else {
		hospital.State = 1
	}
	_, ok = data["status"]
	if ok {
		hospital.Status = lichv.IntVal(data["status"])
	} else {
		hospital.Status = 1
	}
	fmt.Println(hospital)
	if hospital.Name == "" {
		return 0, errors.New("参数为空")
	}

	if err := db.Create(&hospital).Error; err != nil {
		return 0, err
	}
	return hospital.Id, nil
}

func ModifyHospitalById(id int, data map[string]interface{}) error {
	if err := db.Model(&Hospital{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHospitalById(id int) error {
	if err := db.Where("id=?", id).Delete(Hospital{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteHospitals(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(HospitalFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if err := model.Delete(&Hospital{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllHospital() error {
	if err := db.Unscoped().Delete(&Hospital{}).Error; err != nil {
		return err
	}
	return nil
}
