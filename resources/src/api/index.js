import getMonitorPages from './getMonitorPages'
import getProductDistribution from './getProductDistribution'
import getStatusDistribution from './getStatusDistribution'
import getMedicineStatistics from './getMedicineStatistics'
import getSeqNos from './getSeqNos'
import getWechatConfig from './getWechatConfig'
import getWechatUser from './getWechatUser'
import getHospitalPage from './getHospitalPage'
import saveHospital from './saveHospital'
import delHospital from './delHospital'

import getTempPage from './getTempPage'
import saveTemp from './saveTemp'
import delTemp from './delTemp'
import moveTemp2Hospital from './moveTemp2Hospital'

import getTempFlagDistribution from './getTempFlagDistribution'
import getTempStateDistribution from './getTempStateDistribution'

import getDataFlagDistribution from './getDataFlagDistribution'
import getDataStateDistribution from './getDataStateDistribution'



import getPointBaiduPage from './getPointBaiduPage'
import getPointGaodePage from './getPointGaodePage'
import getPointTencentPage from './getPointTencentPage'
import getPointHospitalPage from './getPointHospitalPage'

import delPointBaidu from './delPointBaidu'
import delPointGaode from './delPointGaode'
import delPointTencent from './delPointTencent'
import delPointHospital from './delPointHospital'

import savePointBaidu from './savePointBaidu'
import savePointGaode from './savePointGaode'
import savePointTencent from './savePointTencent'
import savePointHospital from './savePointHospital'


import getBaiduPages from './getBaiduPages'
import getGooglePages from './getGooglePages'
import getInputList from './getInputList'
import saveTest from './saveTest'
import delTest from './delTest'
import delInput from './delInput'

let api = {
    getMonitorPages,
    getProductDistribution,
    getStatusDistribution,
    getMedicineStatistics,
    getSeqNos,
    getWechatConfig,
    getWechatUser,
    getHospitalPage,
    saveHospital,
    delHospital,
    getTempPage,
    saveTemp,
    delTemp,
    moveTemp2Hospital,
    getTempFlagDistribution,
    getTempStateDistribution,
    getDataFlagDistribution,
    getDataStateDistribution,

    getPointBaiduPage,
    getPointGaodePage,
    getPointTencentPage,
    getPointHospitalPage,
    delPointBaidu,
    delPointGaode,
    delPointTencent,
    delPointHospital,
    savePointBaidu,
    savePointGaode,
    savePointTencent,
    savePointHospital,
    
    getBaiduPages,
    getGooglePages,
    getInputList,
    delInput,
    saveTest,
    delTest,
}

export default api