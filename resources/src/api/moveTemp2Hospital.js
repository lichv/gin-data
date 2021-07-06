import request from '../utils/request'

let moveTemp2Hospital = {};

moveTemp2Hospital.send = function(data) {
    return request({
        url: `/api/data/v1/temp/move2Hospital`,
        data,
        method: 'POST'
    })
}

export default moveTemp2Hospital