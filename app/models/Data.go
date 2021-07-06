package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	lichv "github.com/lichv/go"
	"strings"
)

type Data struct {
	Id    int    `json:"id" form:"id" gorm:"id"`
	Code  string `json:"code" form:"code" gorm:"code"`
	Name  string `json:"name" form:"name" gorm:"name"`
	Data  string `json:"alias" form:"alias" gorm:"alias"`
	Flag  int    `json:"flag" form:"flag" gorm:"flag"`
	State int    `json:"state" form:"state" gorm:"state"`
}

var DataFields = []string{"id", "code", "name", "data", "flag", "state"}

func FindDataById(id int) (*Data, error) {
	var data Data
	err := db.Model(&Data{}).Where("id = ? ", id).First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Data{}, err
	}
	return &data, nil
}

func GetDataOne(query map[string]interface{}, orderBy interface{}) (*Data, error) {
	var data Data
	model := db.Model(&Data{})
	for key, value := range query {
		if lichv.In(DataFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	err := model.First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &Data{}, nil
	}
	return &data, nil
}

func GetDataPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) ([]*Data, error) {
	var datas []*Data
	var errs []error
	var err error

	model := db.Model(&Data{})
	for key, value := range query {
		if lichv.In(DataFields,key){
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
	model = model.Offset(beginNum).Limit(pageSize).Order(orderBy).Find(&datas)
	errs = model.GetErrors()
	for _, v := range errs {
		if v != nil {
			err = v
		}
	}

	return datas, err
}

func GetDataList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Data, []error) {
	var Datas []*Data
	var errs []error
	model := db.Model(&Data{})
	for key, value := range query {
		if lichv.In(DataFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if limit > 0 {
		model = model.Limit(limit)
	}
	errs = model.Order(orderBy).Find(&Datas).GetErrors()

	return Datas, errs
}

func GetDataTotal(query map[string]interface{}) (count int, err error) {
	model := db.Model(&Data{})
	for key, value := range query {
		if lichv.In(DataFields,key) {
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

func GetDataFlag() ([]int, error) {
	var flags []CountResult
	var result []int
	err := db.Table("data").Select("distinct flag").Scan(&flags).Error
	if err != nil {
		return result, err
	}
	for _, flag := range flags {
		result = append(result, flag.Count)
	}
	return result, nil
}
func GetDataState() ([]int, error) {
	var flags []CountResult
	var result []int
	err := db.Table("data").Select("distinct state").Scan(&flags).Error
	if err != nil {
		return result, err
	}
	for _, flag := range flags {
		result = append(result, flag.Count)
	}
	return result, nil
}
func AddData(post map[string]interface{}) (int, error) {
	data := Data{}
	_, ok := post["id"]
	if ok {
		data.Id = lichv.IntVal(post["id"])
	}
	_, ok = post["code"]
	if ok {
		data.Code = lichv.Strval(post["code"])
	}
	_, ok = post["name"]
	if ok {
		data.Name = lichv.Strval(post["name"])
	}
	_, ok = post["data"]
	if ok {
		data.Data = lichv.Strval(post["data"])
	}

	_, ok = post["flag"]
	if ok {
		data.Flag = lichv.IntVal(post["flag"])
	} else {
		data.Flag = 1
	}
	_, ok = post["state"]
	if ok {
		data.State = lichv.IntVal(post["state"])
	} else {
		data.State = 1
	}

	fmt.Println(data)
	if data.Name == "" {
		return 0, errors.New("参数为空")
	}

	if err := db.Create(&data).Error; err != nil {
		return 0, err
	}
	return data.Id, nil
}

func ModifyDataById(id int, post map[string]interface{}) error {
	if err := db.Model(&Data{}).Where("id=?", id).Updates(post).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDataById(id int) error {
	if err := db.Where("id=?", id).Delete(Data{}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDatas(maps map[string]interface{}) error {
	model := db
	for key, value := range maps {
		if lichv.In(DataFields,key) {
			model = model.Where(key+"= ?", value)
		}
	}
	if err := model.Delete(&Data{}).Error; err != nil {
		return err
	}
	return nil
}

func ClearAllData() error {
	if err := db.Unscoped().Delete(&Data{}).Error; err != nil {
		return err
	}
	return nil
}

func GetDataFlagDistribution(maps map[string]interface{}) ([]DistributionResult, error) {
	var result []DistributionResult

	var flagMap = map[string]string{"0": "初始状态", "1": "进行中", "2": "初步处理", "3": "已完成"}
	model := db.Model(&Data{}).Select("flag as name,count(1) as count")
	for key, value := range maps {
		if lichv.In(DataFields,key) {
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

func GetDataStateDistribution(maps map[string]interface{}) ([]DistributionResult, error) {
	var result []DistributionResult
	var stateMap = map[string]string{"0": "初始状态", "1": "进行中", "2": "初步处理", "3": "已完成"}
	model := db.Model(&Data{}).Select("state as name,count(1) as count")
	for key, value := range maps {
		if lichv.In(DataFields,key){
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