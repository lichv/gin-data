package services

import (
	"gin-data/app/models"
)

type Temp struct {
	Id    int    `json:"id" form:"id" gorm:"id"`
	Code  string `json:"code" form:"code" gorm:"code"`
	Name  string `json:"name" form:"name" gorm:"name"`
	Data  string `json:"alias" form:"alias" gorm:"alias"`
	Flag  int    `json:"flag" form:"flag" gorm:"flag"`
	State int    `json:"state" form:"state" gorm:"state"`
}

func GetTempTotal(maps map[string]interface{}) (count int, err error) {
	count, err = models.GetTempTotal(maps)
	return count, err
}

func GetTempOne(query map[string]interface{}, orderBy interface{}) (*Temp, error) {
	var nu *models.Temp
	nu, err := models.GetTempOne(query, orderBy)
	if err != nil {
		return nil, err
	}
	return TransferTempModel(nu), nil
}

func GetTempPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) (temps []*Temp, total int, err error) {
	total, err = models.GetTempTotal(query)
	if err != nil {
		return nil, 0, err
	}
	us, err := models.GetTempPages(query, orderBy, pageNum, pageSize)
	temps = TransferTemps(us)
	return temps, total, nil
}
func GetTempList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Temp, []error) {
	users, errors := models.GetTempList(query, orderBy, limit)
	temps := TransferTemps(users)
	return temps, errors
}

func AddTemp(data map[string]interface{}) (int, error) {
	return models.AddTemp(data)
}

func ModifyTemp(id int, data map[string]interface{}) (err error) {
	err = models.ModifyTempById(id, data)
	return err
}

func DeleteTemp(maps map[string]interface{}) (err error) {
	err = models.DeleteTemps(maps)
	return nil
}

func ClearAllTemp() (err error) {
	err = models.ClearAllTemp()
	return err
}

func TransferTempModel(u *models.Temp) *Temp {
	temp := &Temp{
		Id:    u.Id,
		Code:  u.Code,
		Name:  u.Name,
		Data:  u.Data,
		Flag:  u.Flag,
		State: u.State,
	}
	return temp
}
func TransferTemps(us []*models.Temp) []*Temp {
	var temps []*Temp
	for _, value := range us {
		temp := TransferTempModel(value)
		temps = append(temps, temp)
	}
	return temps
}

func GetTempFlagDistribution() ([]models.DistributionResult,error){
	return models.GetTempFlagDistribution(map[string]interface{}{})
}

func GetTempStateDistribution() ([]models.DistributionResult,error){
	return models.GetTempStateDistribution(map[string]interface{}{})
}

func MoveTemp2Hospital(query map[string]interface{}) (bool,error) {
	nu, err := models.GetTempOne(query, "id asc")
	if err != nil{
		return false,err
	}
	data := map[string]interface{}{"name":nu.Name}
	hospital, err := models.AddHospital(data)
	if err != nil{
		return false,err
	}

	models.DeleteTempById(nu.Id)
	if hospital > 0 {
		return true,nil
	}
	return false,nil
}