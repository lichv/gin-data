import request from '../utils/request'

let getWechatConfig = {};

getWechatConfig.send = function(data) {
    return request({
        url: `https://auth.lichv.com/wechat/config`,
        data,
        method: 'POST'
    })
}

export default getWechatConfig