import request from '../utils/request'

let getTempStateDistribution = {};

getTempStateDistribution.send = function(data) {
    return request({
        url: `/api/data/v1/temp/getStateDistribution`,
        data,
        method: 'POST'
    })
}

export default getTempStateDistribution