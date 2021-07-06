import request from '../utils/request'

let saveTemp = {};

saveTemp.send = function(data) {
    return request({
        url: `/api/data/v1/temp/save`,
        data,
        method: 'POST'
    })
}

export default saveTemp