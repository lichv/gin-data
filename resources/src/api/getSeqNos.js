import request from '../utils/request'

let getSeqNos = {};

getSeqNos.send = function(data) {
    return request({
        url: `/api/monitor/monitor/getSeqNos`,
        data,
        method: 'POST'
    })
}

export default getSeqNos