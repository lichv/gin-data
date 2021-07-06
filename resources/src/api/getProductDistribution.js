import request from '../utils/request'

let getProductDistribution = {};

getProductDistribution.send = function(data) {
    return request({
        url: `/api/monitor/monitor/productDistribution`,
        data,
        method: 'POST'
    })
}

export default getProductDistribution