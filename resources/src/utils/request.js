import axios from 'axios'
import Cookies from 'js-cookie'

const service = axios.create({
    baseURL: '/',
    timeout: 60000, // 请求超时时间
})

// 请求拦截
service.interceptors.request.use(
    (config) => {
        // 请求拦截
        // 例如判断用户是否登录，向请求头中加token
        var user_token = localStorage.getItem('user_token')
        if (!user_token && typeof(user_token)!="undefined" && user_token!=0) {
            user_token = Cookies.get('user_token')
        }
        if (user_token) {
            config.headers.common['X-TOKEN'] = user_token
        }
        return config
    },
    (error) => {
        // 发送失败
        return Promise.reject(error);
    }
    )

// 响应拦截
service.interceptors.response.use(
    (response) => {
        console.log('service.interceptors.response.success',response)
        const { headers, data, status } = response;
        if(status == '200' || data.success || data.result == 'success') {
            console.log('headers',headers)
            return data
        } else {
            // Cookies.remove('user_token')
            // localStorage.removeItem('user_token')
            return Promise.reject();
        }
    },
    (error) => {
        console.log('service.interceptors.response.error',error.code)

        if (error.code==401) {
            console.log('登录失败')
            Cookies.remove('user_token')
            Cookies.remove('nickname')
            console.log('getLoginUser failed')
            localStorage.removeItem('user_token')
            localStorage.removeItem('nickname')
        }
        
        return Promise.reject(error);
    }
    )

export default service