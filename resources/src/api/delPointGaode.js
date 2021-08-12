import request from '../utils/request'

let delPointGaode = {};

delPointGaode.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/gaode/del`,
        data,
        method: 'POST'
    })
}

export default delPointGaode