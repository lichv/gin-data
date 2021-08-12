package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
	"time"
)

type Input struct {
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

var InputFields = []string{"id", "code", "name", "alias", "logo", "grade", "nature", "type", "regyear", "size", "province", "city", "town", "address", "location", "members", "beds", "emergency", "inpatients", "creditcode", "corporations", "typename", "introduction", "others", "source", "tianyancha", "yixue", "url", "phone", "email", "addtime", "flag", "state", "status"}

func FindInputById(id int) (*Input, error) {
	var input Input
	err := db.Model(&Input{}).Where("id = ? ", id).First(&input).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Input{}, err
	}
	return &input, nil
}

func GetInputOne(query map[string]interface{}, orderBy interface{}) (*Input, error) {
	var input Input
	model := db.Model(&Input{})
	for key, value := range query {
		if lichv.In(InputFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	err := model.First(&input).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Input{}, nil
	}
	return &input, nil
}

func GetInputPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*Input, error) {
	var inputs []*Input
	var errs []error
	var err error

	model := db.Model(&Input{})
	for key, value := range query {
		if lichv.In(InputFields,key) {
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
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&inputs)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return inputs, err
}

func GetInputList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Input, []error) {
	var Inputs []*Input
	var errs []error
	model := db.Model(&Input{})
	for key, value := range query {
		if lichv.In(InputFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Inputs).GetErrors()

	return Inputs, errs
}

func GetInputTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&Input{})
	for key, value := range query {
		if lichv.In(InputFields,key) {
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
func AddInput(data map[string]interface{}) (int, error) {
	input := Input{}
	_, ok := data["id"]
	if ok {
		input.Id = lichv.IntVal(data["id"])
	}
	_, ok = data["code"]
	if ok {
		input.Code = lichv.Strval(data["code"])
	}
	_, ok = data["name"]
	if ok {
		input.Name = lichv.Strval(data["name"])
	}
	_, ok = data["alias"]
	if ok {
		input.Alias = lichv.Strval(data["alias"])
	}
	_, ok = data["logo"]
	if ok {
		input.Logo = lichv.Strval(data["logo"])
	}
	_, ok = data["grade"]
	if ok {
		input.Grade = lichv.Strval(data["grade"])
	}
	_, ok = data["nature"]
	if ok {
		input.Nature = lichv.Strval(data["nature"])
	}
	_, ok = data["type"]
	if ok {
		input.Type = lichv.Strval(data["type"])
	}
	_, ok = data["regyear"]
	if ok {
		input.Regyear = lichv.Strval(data["regyear"])
	}
	_, ok = data["size"]
	if ok {
		input.Size = lichv.Strval(data["size"])
	}
	_, ok = data["province"]
	if ok {
		input.Province = lichv.Strval(data["province"])
	}
	_, ok = data["city"]
	if ok {
		input.City = lichv.Strval(data["city"])
	}
	_, ok = data["town"]
	if ok {
		input.Town = lichv.Strval(data["town"])
	}
	_, ok = data["address"]
	if ok {
		input.Address = lichv.Strval(data["address"])
	}
	_, ok = data["location"]
	if ok {
		input.Location = lichv.Strval(data["location"])
	}
	_, ok = data["members"]
	if ok {
		input.Members = lichv.Strval(data["members"])
	}
	_, ok = data["beds"]
	if ok {
		input.Beds = lichv.Strval(data["beds"])
	}
	_, ok = data["emergency"]
	if ok {
		input.Emergency = lichv.Strval(data["emergency"])
	}
	_, ok = data["inpatients"]
	if ok {
		input.Inpatients = lichv.Strval(data["inpatients"])
	}
	_, ok = data["creditcode"]
	if ok {
		input.Creditcode = lichv.Strval(data["creditcode"])
	}
	_, ok = data["corporations"]
	if ok {
		input.Corporations = lichv.Strval(data["corporations"])
	}
	_, ok = data["typename"]
	if ok {
		input.Typename = lichv.Strval(data["typename"])
	}
	_, ok = data["introduction"]
	if ok {
		input.Introduction = lichv.Strval(data["introduction"])
	}
	_, ok = data["others"]
	if ok {
		input.Others = lichv.Strval(data["others"])
	}
	_, ok = data["source"]
	if ok {
		input.Source = lichv.Strval(data["source"])
	}
	_, ok = data["url"]
	if ok {
		input.Url = lichv.Strval(data["url"])
	}
	_, ok = data["phone"]
	if ok {
		input.Phone = lichv.Strval(data["phone"])
	}
	_, ok = data["email"]
	if ok {
		input.Email = lichv.Strval(data["email"])
	}

	input.Addtime = time.Now().Unix()

	_, ok = data["flag"]
	if ok {
		input.Flag = lichv.IntVal(data["flag"])
	} else {
		input.Flag = 1
	}
	_, ok = data["state"]
	if ok {
		input.State = lichv.IntVal(data["state"])
	} else {
		input.State = 1
	}
	_, ok = data["status"]
	if ok {
		input.Status = lichv.IntVal(data["status"])
	} else {
		input.Status = 1
	}
	fmt.Println(input)
	if input.Name == "" {
		return 0, errors.New("参数为空")
	}

	if err := db.Create(&input).Error; err != nil {
		return 0, err
	}
	return input.Id, nil
}

func ModifyInputById(id int, data map[string]interface{}) error {
	if err := db.Model(&Input{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteInputById(id int) error {
	if err := db.Where("id=?", id).Delete(Input{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteInputs(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(InputFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if err := model.Delete(&Input{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllInput() error {
	if err := db.Unscoped().Delete(&Input{}).Error; err != nil {
		return err
	}
	return nil
}
