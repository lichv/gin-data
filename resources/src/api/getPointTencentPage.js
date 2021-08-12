import request from '../utils/request'

let getPointTencentPage = {};

getPointTencentPage.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/tencent/getPage`,
        data,
        method: 'POST'
    })
}

export default getPointTencentPage