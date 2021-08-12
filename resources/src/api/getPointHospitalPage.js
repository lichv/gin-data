import request from '../utils/request'

let getPointHospitalPage = {};

getPointHospitalPage.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/getPage`,
        data,
        method: 'POST'
    })
}

export default getPointHospitalPage