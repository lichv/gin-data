import api from '../api';
import Cookies from 'js-cookie'

let actions = {
	getLoginUser(context){
		var user_token = localStorage.getItem('user_token')
		if (!user_token && typeof(user_token)!="undefined" && user_token!=0) {
			user_token = Cookies.get('user_token')
		}
		console.log('user_token',user_token)
		if (user_token) {
			api.getWechatUser.send().
			then(result=>{
				console.log('getWechatUser',result)
				if (result.state==2000) {
					context.state.userinfo = result.data
					Cookies.set('nickname',result.data.nickname)
				}else{
					Cookies.remove('user_token')
					Cookies.remove('nickname')
					localStorage.removeItem('user_token')
					localStorage.removeItem('nickname')
				}

			})
			.catch(e => {
				Cookies.remove('user_token')
				Cookies.remove('nickname')
				console.log('getLoginUser failed')
				localStorage.removeItem('user_token')
				localStorage.removeItem('nickname')
				console.log(e)
			})
		}
	}
}

export default actions