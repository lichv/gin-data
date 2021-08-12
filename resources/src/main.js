import { createApp } from 'vue'
import App from './App.vue'
import './base.css'
import './index.css'
import router from './router'
import store from './store'
import ElementPlus from 'element-plus';
import 'dayjs/locale/zh-cn'
import locale from 'element-plus/lib/locale/lang/zh-cn'
import 'element-plus/lib/theme-chalk/index.css';

import * as echarts from 'echarts'

const app = createApp(App)

app.use(ElementPlus, { size: 'small', zIndex: 3000,locale:locale })

import ElAsiderMenu from './components/AsiderMenu.vue'
import ElChart from './components/Chart.vue'
import ElUserNotice from './components/UserNotice.vue';
import ElWechatLogin from './components/WechatLogin.vue';
import ElMonitorProductDistribution from './components/MonitorProductDistribution.vue'
import ElMonitorStatusDistribution from './components/MonitorStatusDistribution.vue'
import ElMonitorMedicineStatistics from './components/MonitorMedicineStatistics.vue'

import ElHospitalForm from './components/HospitalForm.vue'
import ElHospitalTable from './components/HospitalTable.vue'
import ElTempTable from './components/TempTable.vue'

import ElTempFlagDistribution from './components/TempFlagDistribution.vue'
import ElTempStateDistribution from './components/TempStateDistribution.vue'

import ElDataFlagDistribution from './components/DataFlagDistribution.vue'
import ElDataStateDistribution from './components/DataStateDistribution.vue'

import ElAreaSelect from './components/AreaSelect.vue'

import ElPointBaiduTable from './components/PointBaiduTable.vue'
import ElPointGaodeTable from './components/PointGaodeTable.vue'
import ElPointTencentTable from './components/PointTencentTable.vue'
import ElPointHospitalTable from './components/PointHospitalTable.vue'

import ElInputTable from './components/InputTable.vue'
import ElBaiduTable from './components/BaiduTable.vue'
import ElGoogleTable from './components/GoogleTable.vue'

const components = [
ElAsiderMenu,
ElChart,
ElUserNotice,
ElWechatLogin,
ElMonitorProductDistribution,
ElMonitorStatusDistribution,
ElMonitorMedicineStatistics,
ElHospitalForm,
ElHospitalTable,
ElTempTable,
ElTempFlagDistribution,
ElTempStateDistribution,
ElDataFlagDistribution,
ElDataStateDistribution,
ElAreaSelect,

ElPointBaiduTable,
ElPointGaodeTable,
ElPointTencentTable,
ElPointHospitalTable,

	ElInputTable,
	ElBaiduTable,
	ElGoogleTable,
]
components.forEach(component => {
	app.component(component.name, component)
})

app.use(router)
app.use(store)

store.dispatch('getLoginUser')
app.config.globalProperties.$echarts = echarts
app.config.globalProperties.$filters = {
	findStatusByCode(code,list) {
		var label = ''
		let index;
		for (index in list){
			if (list[index].value==code) {
				label = list[index].label
			}
		}
		return label
	}
}

app.mount('#app')
