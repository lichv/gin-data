import request from '../utils/request'

let getDataStateDistribution = {};

getDataStateDistribution.send = function(data) {
    return request({
        url: `/api/data/v1/data/getStateDistribution`,
        data,
        method: 'POST'
    })
}

export default getDataStateDistribution