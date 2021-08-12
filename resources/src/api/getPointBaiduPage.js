import request from '../utils/request'

let getPointBaiduPage = {};

getPointBaiduPage.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/baidu/getPage`,
        data,
        method: 'POST'
    })
}

export default getPointBaiduPage