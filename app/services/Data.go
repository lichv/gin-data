package services

import (
	"gin-data/app/models"
)

type Data struct {
	Id    int    `json:"id" form:"id" gorm:"id"`
	Code  string `json:"code" form:"code" gorm:"code"`
	Name  string `json:"name" form:"name" gorm:"name"`
	Data  string `json:"alias" form:"alias" gorm:"alias"`
	Flag  int    `json:"flag" form:"flag" gorm:"flag"`
	State int    `json:"state" form:"state" gorm:"state"`
}

func GetDataTotal(maps map[string]interface{}) (count int, err error) {
	count, err = models.GetDataTotal(maps)
	return count, err
}

func GetDataOne(query map[string]interface{}, orderBy interface{}) (*Data, error) {
	var nu *models.Data
	nu, err := models.GetDataOne(query, orderBy)
	if err != nil {
		return nil, err
	}
	return TransferDataModel(nu), nil
}

func GetDataPages(query map[string]interface{}, orderBy interface{}, pageNum int, pageSize int) (datas []*Data, total int, err error) {
	total, err = models.GetDataTotal(query)
	if err != nil {
		return nil, 0, err
	}
	us, err := models.GetDataPages(query, orderBy, pageNum, pageSize)
	datas = TransferDatas(us)
	return datas, total, nil
}
func GetDataList(query map[string]interface{}, orderBy interface{}, limit int) ([]*Data, []error) {
	users, errors := models.GetDataList(query, orderBy, limit)
	datas := TransferDatas(users)
	return datas, errors
}

func GetDataFlag() ([]int,error){
	return models.GetDataFlag()
}

func GetDataState() ([]int,error){
	return models.GetDataState()
}

func AddData(post map[string]interface{}) (int, error) {
	return models.AddData(post)
}

func ModifyData(id int, post map[string]interface{}) (err error) {
	err = models.ModifyDataById(id, post)
	return err
}

func DeleteData(maps map[string]interface{}) (err error) {
	err = models.DeleteDatas(maps)
	return nil
}

func ClearAllData() (err error) {
	err = models.ClearAllData()
	return err
}

func TransferDataModel(u *models.Data) *Data {
	data := &Data{
		Id:    u.Id,
		Code:  u.Code,
		Name:  u.Name,
		Data:  u.Data,
		Flag:  u.Flag,
		State: u.State,
	}
	return data
}
func TransferDatas(us []*models.Data) []*Data {
	var datas []*Data
	for _, value := range us {
		data := TransferDataModel(value)
		datas = append(datas, data)
	}
	return datas
}

func GetDataFlagDistribution() ([]models.DistributionResult,error){
	return models.GetDataFlagDistribution(map[string]interface{}{})
}

func GetDataStateDistribution() ([]models.DistributionResult,error){
	return models.GetDataStateDistribution(map[string]interface{}{})
}