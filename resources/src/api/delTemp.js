import request from '../utils/request'

let delTemp = {};

delTemp.send = function(data) {
    return request({
        url: `/api/data/v1/temp/del`,
        data,
        method: 'POST'
    })
}

export default delTemp