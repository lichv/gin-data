import request from '../utils/request'

let delTest = {};

delTest.send = function(data) {
    return request({
        url: `/api/test/del`,
        data,
        method: 'POST'
    })
}

export default delTest