import request from '../utils/request'

let getHospitalPage = {};

getHospitalPage.send = function(data) {
    return request({
        url: `/api/data/v1/hospital/del`,
        data,
        method: 'POST'
    })
}

export default getHospitalPage