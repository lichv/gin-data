package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
	"strings"
)

type Temp struct {
	Id    int    `json:"id" form:"id" gorm:"id"`
	Code  string `json:"code" form:"code" gorm:"code"`
	Name  string `json:"name" form:"name" gorm:"name"`
	Data  string `json:"alias" form:"alias" gorm:"alias"`
	Flag  int    `json:"flag" form:"flag" gorm:"flag"`
	State int    `json:"state" form:"state" gorm:"state"`
}

type Flag struct {
	Flag int
}

var TempFields = []string{"id", "code", "name", "data", "flag", "state"}

func FindTempById(id int) (*Temp, error) {
	var temp Temp
	err := db.Model(&Temp{}).Where("id = ? ", id).First(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Temp{}, err
	}
	return &temp, nil
}

func GetTempOne(query map[string]interface{}, orderBy interface{}) (*Temp, error) {
	var temp Temp
	model := db.Model(&Temp{})
	for key, value := range query {
		if lichv.In(TempFields,key){
			model = model.Where(key+"= ?", value)
		}
	}
	err := model.Order(orderBy).First(&temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Temp{}, nil
	}
	return &temp, nil
}

func GetTempPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*Temp, error) {
	var temps []*Temp
	var errs []error
	var err error

	model := db.Model(&Temp{})
	for key, value := range query {
		if lichv.In(TempFields,key){
			name := lichv.Strval(value)
			if key == "name" {
				if name != ""{
					model = model.Where("name like ?", "%"+strings.Join(strings.Split(name, ""), "%")+"%")
				}
			} else {
				if lichv.BoolVal(value) {
					model = model.Where(key+"= ?", value)
				}
			}
		}

	}
	beginNum := (pageNum - 1) * pageSize
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&temps)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return temps, err
}

func GetTempList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Temp, []error) {
	var Temps []*Temp
	var errs []error
	model := db.Model(&Temp{})
	for key, value := range query {
		if lichv.In(TempFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Temps).GetErrors()

	return Temps, errs
}

func GetTempTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&Temp{})
	fmt.Println(query)
	for key, value := range query {
		if lichv.In(TempFields,key) {
			name := lichv.Strval(value)
			if key == "name" {
				if name != ""{
					model = model.Where("name like ?", "%"+strings.Join(strings.Split(name, ""), "%")+"%")
				}
			} else {
				model = model.Where(key+"= ?", value)
			}
		}
	}
	err = model.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}


func AddTemp(data map[string]interface{}) (int, error) {
	temp := Temp{}
	_, ok := data["id"]
	if ok {
		temp.Id = lichv.IntVal(data["id"])
	}
	_, ok = data["code"]
	if ok {
		temp.Code = lichv.Strval(data["code"])
	}
	_, ok = data["name"]
	if ok {
		temp.Name = lichv.Strval(data["name"])
	}
	_, ok = data["data"]
	if ok {
		temp.Data = lichv.Strval(data["data"])
	}

	_, ok = data["flag"]
	if ok {
		temp.Flag = lichv.IntVal(data["flag"])
	} else {
		temp.Flag = 1
	}
	_, ok = data["state"]
	if ok {
		temp.State = lichv.IntVal(data["state"])
	} else {
		temp.State = 1
	}

	fmt.Println(temp)
	if temp.Name == "" {
		return 0, errors.New("参数为空")
	}

	if err := db.Create(&temp).Error; err != nil {
		return 0, err
	}
	return temp.Id, nil
}

func ModifyTempById(id int, data map[string]interface{}) error {
	if err := db.Model(&Temp{}).Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTempById(id int) error {
	if err := db.Where("id=?", id).Delete(Temp{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTemps(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(TempFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if err := model.Delete(&Temp{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllTemp() error {
	if err := db.Unscoped().Delete(&Temp{}).Error; err != nil {
		return err
	}
	return nil
}

func GetTempFlagDistribution(maps map[string]interface{}) ([]DistributionResult, error) {
	var result []DistributionResult

	var flagMap = map[string]string{"0": "初始状态", "1": "进行中", "2": "初步处理", "3": "已完成"}
	model := db.Model(&Temp{}).Select("flag as name,count(1) as count")
	for key, value := range maps {
		if lichv.In(TempFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}

	rows, err := model.Group("flag").Order("name asc,count desc").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var name = ""
		var count = 0
		rows.Scan(&name, &count)
		result = append(result, DistributionResult{Name: flagMap[name], Count: count})

	}

	return result, nil
}

func GetTempStateDistribution(maps map[string]interface{}) ([]DistributionResult, error) {
	var result []DistributionResult
	var stateMap = map[string]string{"0": "初始状态", "1": "进行中", "2": "初步处理", "3": "已完成"}
	model := db.Model(&Temp{}).Select("state as name,count(1) as count")
	for key, value := range maps {
		if  lichv.In(TempFields,key){
			model = model.Where(key+"= ?", value)
		}
	}

	rows, err := model.Group("state").Order("name asc,count desc").Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		var name = ""
		var count = 0
		rows.Scan(&name, &count)
		result = append(result, DistributionResult{Name: stateMap[name], Count: count})

	}

	return result, nil
}