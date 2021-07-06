import request from '../utils/request'

let getTempPage = {};

getTempPage.send = function(data) {
    return request({
        url: `/api/data/v1/temp/getPage`,
        data,
        method: 'POST'
    })
}

export default getTempPage