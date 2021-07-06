import request from '../utils/request'

let getMonitorPages = {};

getMonitorPages.send = function(data) {
    return request({
        url: `/api/monitor/monitor/getPages`,
        data,
        method: 'POST'
    })
}

export default getMonitorPages