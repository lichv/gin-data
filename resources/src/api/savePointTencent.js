import request from '../utils/request'

let savePointTencent = {};

savePointTencent.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/baidu/save`,
        data,
        method: 'POST'
    })
}

export default savePointTencent