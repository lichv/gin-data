import request from '../utils/request'

let savePointBaidu = {};

savePointBaidu.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/baidu/save`,
        data,
        method: 'POST'
    })
}

export default savePointBaidu