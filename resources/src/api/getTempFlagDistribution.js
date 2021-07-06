import request from '../utils/request'

let getTempFlagDistribution = {};

getTempFlagDistribution.send = function(data) {
    return request({
        url: `/api/data/v1/temp/getFlagDistribution`,
        data,
        method: 'POST'
    })
}

export default getTempFlagDistribution