import request from '../utils/request'

let getBaiduPages = {};

getBaiduPages.send = function(data) {
    return request({
        url: `/api/baidu/search`,
        data,
        method: 'POST'
    })
}

export default getBaiduPages