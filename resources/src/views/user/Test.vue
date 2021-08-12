<template>
  <div class="user-page">
    <el-card class="user-card card-board">
      <template #header>
        <div class="card-header">
          <div class="card-header-left">
            用户信息
          </div>
          <div class="card-header-right">
            <el-button size="mini" round  icon="el-icon-edit"></el-button>
          </div>
        </div>
      </template>
      <el-descriptions>
        <el-descriptions-item label="用户编号：">AP32134</el-descriptions-item>
        <el-descriptions-item label="家系编号：">A32134</el-descriptions-item>
        <el-descriptions-item label="参照用户：">张毅</el-descriptions-item>
        <el-descriptions-item label="参照关系：">爸爸</el-descriptions-item>
        <el-descriptions-item label="用户名称：">张佳豪</el-descriptions-item>
        <el-descriptions-item label="性别：">男</el-descriptions-item>
        <el-descriptions-item label="出生日期：">1999-09-09</el-descriptions-item>
        <el-descriptions-item label="年龄：">22</el-descriptions-item>
        <el-descriptions-item label="手机号码：">13012345678</el-descriptions-item>
        <el-descriptions-item label="地址：">广东省 深圳市 龙华区 宝能科技园</el-descriptions-item>
        <el-descriptions-item label="销售负责人：">赵明明</el-descriptions-item>
        <el-descriptions-item label="备注：">
          <el-tag size="small">学校</el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
    <el-card class="user-card card-board">
      <template #header>
        <div class="card-header">
          <div class="card-header-left">
            <span>用户：kooriookami</span>
          </div>
          <div class="card-header-right">
            <el-button size="mini" round  icon="el-icon-edit"></el-button>
          </div>
        </div>
      </template>
      <el-descriptions>
        <el-descriptions-item label="用户编号：">AP32134</el-descriptions-item>
        <el-descriptions-item label="家系编号：">A32134</el-descriptions-item>
        <el-descriptions-item label="参照用户：">张毅</el-descriptions-item>
        <el-descriptions-item label="参照关系：">爸爸</el-descriptions-item>
        <el-descriptions-item label="用户名称：">张佳豪</el-descriptions-item>
        <el-descriptions-item label="性别：">男</el-descriptions-item>
        <el-descriptions-item label="出生日期：">1999-09-09</el-descriptions-item>
        <el-descriptions-item label="年龄：">22</el-descriptions-item>
        <el-descriptions-item label="手机号码：">13012345678</el-descriptions-item>
        <el-descriptions-item label="地址：">广东省 深圳市 龙华区 宝能科技园</el-descriptions-item>
        <el-descriptions-item label="销售负责人：">赵明明</el-descriptions-item>
        <el-descriptions-item label="备注：">
          <el-tag size="small">学校</el-tag>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
    <el-card class="user-form card-board">
      <el-form ref="form" :model="form" class="user-form">
        <el-row>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="用户编号：" label-width="84px">

            </el-form-item>
          </el-col>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="家系编号：" label-width="84px">
              <el-input v-model="form.family_no"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item label="参照用户" label-width="96px">
              <el-select v-model="form.referent" placeholder="请选择样本类型">
                <el-option label="AP23423-张二" value="AP23423"></el-option>
                <el-option label="AP23424-李兰" value="AP23424"></el-option>
                <el-option label="AP23425-张易" value="AP23425"></el-option>
                <el-option label="AP23426-王澜" value="AP23426"></el-option>
                <el-option label="AP23427-张二狗" value="AP23427"></el-option>
              </el-select>

            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item label="参照关系" label-width="116px">
              <el-select v-model="form.relation" placeholder="请选择样本类型">
                <el-option label="爸爸" value="爸爸"></el-option>
                <el-option label="妈妈" value="妈妈"></el-option>
                <el-option label="爷爷" value="爷爷"></el-option>
                <el-option label="奶奶" value="奶奶"></el-option>
                <el-option label="叔叔" value="叔叔"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="4">
            <el-form-item >
              <el-button type="danger" @click="" icon="el-icon-delete">删除</el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="客户名称" label-width="84px">
              <el-input v-model="form.name"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="性别" label-width="84px">
              <el-radio-group v-model="form.gender" size="mini">
                <el-radio-button label="男" value="1"></el-radio-button>
                <el-radio-button label="女" value="2"></el-radio-button>
                <el-radio-button label="未知" value="0"></el-radio-button>
              </el-radio-group>
            </el-form-item>
          </el-col>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="出生日期" label-width="84px" class="form-address-detail">
              <el-input v-model="form.birth"></el-input>
              <el-input v-model="form.age" ><template #append>岁</template></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="手机号码" label-width="84px">
              <el-input v-model="form.mobile"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>

          <el-col :span="12" :xs="24" :lg="12">
            <el-form-item label="地址" label-width="84px" class="form-address-detail">
              <el-cascader :options="options" :props="{label:'name',value:'code'}" clearable class="width-160" ></el-cascader>
              <el-input v-model="form.address"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="6" :xs="12" :sm="10" :md="8" :lg="6" :xl="6">
            <el-form-item label="销售负责人" label-width="84px">
              <el-input v-model="form.saler" disabled></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item >
          <el-button type="primary" >保存</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: 'App',
    data(){
      return{
        value: [],
        options:[],
        form: {
          'ap':'AP23350',
          'af':'AF23330',
          'name': '*佳豪',
          'gender': '男',
          'birth': '1995/12/12',
          'mobile': '13012345678',
          'province': '广东',
          'city': '深圳',
          'county': '南山区',
          'address': '大冲城市花园',
          'referent': 'AP23333',
          'relation': '爸爸',
          'saler':'12312',
          'regtime': '2020/08/02 14:23:12',
          'is_contact': true,
          'remark': '',
        },

      };
    },
  }
</script>
<style>
.user-page{
  background-color: #f0f9eb;
}
.card-header{
  display: flex;
}
.card-header-left{
  width: 200px;
}
.card-board{
  margin-top: 6px;
}
</style>
