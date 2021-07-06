<template>
  <div class="user-notice">
    <el-badge :value="12" class="item">
      <i class="el-icon-s-comment"></i>
    </el-badge>
    <el-badge :value="12" class="item">
      <i class="el-icon-message-solid"></i>
    </el-badge>
    <div class="avatar-block">
      <el-dropdown>
        <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
        <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>修改密码</el-dropdown-item>
          <el-dropdown-item divided>退出登录</el-dropdown-item>
        </el-dropdown-menu>
      </template>
      </el-dropdown>
    </div>
  </div>
</template>
<script>
  export default {
    name: 'ElUserNotice',
    data(){
      return{
        isCollapse: false,
      };
    },
    props: {
      separator: {
        type: String,
        default: '/'
      },
      separatorClass: {
        type: String,
        default: ''
      }
    },

    mounted() {

    }
  };
</script>
<style>
.user-notice{
  width: 128px;
}
.user-notice{
  display: flex;
  line-height: 50px;
}
.user-notice i{
  font-size: 20px;
  margin: 20px 12px ;
  color: #bebec0;
}
.user-notice .el-badge__content.is-fixed{
  top: 16px;
  right: 24px;
}
.user-notice .el-avatar{
  font-size: 20px;
  margin: 10px;
}
</style>
