import request from '../utils/request'

let saveTest = {};

saveTest.send = function(data) {
    return request({
        url: `/api/test/save`,
        data,
        method: 'POST'
    })
}

export default saveTest