import request from '../utils/request'

let savePointGaode = {};

savePointGaode.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/gaode/save`,
        data,
        method: 'POST'
    })
}

export default savePointGaode