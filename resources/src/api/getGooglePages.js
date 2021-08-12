import request from '../utils/request'

let getBaiduPages = {};

getBaiduPages.send = function(data) {
    return request({
        url: `/api/google/search`,
        data,
        method: 'POST'
    })
}

export default getBaiduPages