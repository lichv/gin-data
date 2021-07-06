import request from '../utils/request'

let getStatusDistribution = {};

getStatusDistribution.send = function(data) {
    return request({
        url: `/api/monitor/monitor/statusDistribution`,
        data,
        method: 'POST'
    })
}

export default getStatusDistribution