import request from '../utils/request'

let delPointHospital = {};

delPointHospital.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/del`,
        data,
        method: 'POST'
    })
}

export default delPointHospital