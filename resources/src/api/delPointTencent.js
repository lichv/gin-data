import request from '../utils/request'

let delPointTencent = {};

delPointTencent.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/tencent/del`,
        data,
        method: 'POST'
    })
}

export default delPointTencent