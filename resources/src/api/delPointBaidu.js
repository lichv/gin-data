import request from '../utils/request'

let delPointBaidu = {};

delPointBaidu.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/baidu/del`,
        data,
        method: 'POST'
    })
}

export default delPointBaidu