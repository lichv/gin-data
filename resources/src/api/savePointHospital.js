import request from '../utils/request'

let savePointHospital = {};

savePointHospital.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/save`,
        data,
        method: 'POST'
    })
}

export default savePointHospital