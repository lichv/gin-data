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
}

export default api