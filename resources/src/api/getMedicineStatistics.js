import request from '../utils/request'

let getMedicineStatistics = {};

getMedicineStatistics.send = function(data) {
    return request({
        url: `/api/monitor/medicine/statistics`,
        data,
        method: 'POST'
    })
}

export default getMedicineStatistics