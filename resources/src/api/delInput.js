import request from '../utils/request'

let delInput = {};

delInput.send = function(data) {
    return request({
        url: `/api/input/del`,
        data,
        method: 'POST'
    })
}

export default delInput