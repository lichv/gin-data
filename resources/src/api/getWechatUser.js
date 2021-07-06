import request from '../utils/request'

let getWechatUser = {};

getWechatUser.send = function(data) {
    return request({
        url: `https://auth.lichv.com/api/v1/user/self`,
        data,
        method: 'POST'
    })
}

export default getWechatUser