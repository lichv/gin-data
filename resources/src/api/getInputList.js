import request from '../utils/request'

let getInputList = {};

getInputList.send = function(data) {
    return request({
        url: `/api/input/getList`,
        data,
        method: 'POST'
    })
}

export default getInputList