import request from '../utils/request'

let saveHospital = {};

saveHospital.send = function(data) {
    return request({
        url: `/api/data/v1/hospital/save`,
        data,
        method: 'POST'
    })
}

export default saveHospital