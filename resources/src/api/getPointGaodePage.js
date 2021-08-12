import request from '../utils/request'

let getPointGaodePage = {};

getPointGaodePage.send = function(data) {
    return request({
        url: `/api/data/v1/point/hospital/gaode/getPage`,
        data,
        method: 'POST'
    })
}

export default getPointGaodePage