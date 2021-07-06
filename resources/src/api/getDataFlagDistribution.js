import request from '../utils/request'

let getDataFlagDistribution = {};

getDataFlagDistribution.send = function(data) {
    return request({
        url: `/api/data/v1/data/getFlagDistribution`,
        data,
        method: 'POST'
    })
}

export default getDataFlagDistribution