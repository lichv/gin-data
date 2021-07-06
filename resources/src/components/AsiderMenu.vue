<template>
    <el-menu :uniqueOpened="true" default-active="2" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" background-color="#001529" text-color="#aaa" active-text-color="#fff" :collapse="isCollapse" router @select="handleSelect">
      <template v-for="menu in treeData">
        <el-submenu v-if="'children' in menu && Object.keys(menu.children).length > 0" :index="menu.link+''">
          <template #title>
            <i :class="menu.icon" v-if="menu.icon != ''"></i>
            <span slot="title">{{menu.title}}</span>
          </template>
          <template v-for="submenu in menu.children">
            <el-submenu v-if="'children' in submenu && Object.keys(submenu.children).length > 0"  :index="submenu.link+''" class="grandmenu">
              <template #title>
                <i :class="submenu.icon" v-if="submenu.icon != ''"></i>
                <span slot="title">{{submenu.title}}</span>
              </template>
              <template v-for="grandmenu in submenu.children">
                <el-menu-item :index="grandmenu.id+''">
                  <template #title>
                    <i :class="grandmenu.icon" v-if="grandmenu.icon != ''"></i>
                    <span slot="title">{{grandmenu.title}}</span>
                  </template>
                </el-menu-item>
              </template>
            </el-submenu>
            <el-menu-item v-else :index="submenu.link+''">
              <template #title>
                <i :class="submenu.icon" v-if="submenu.icon != ''"></i>
                <span slot="title">{{submenu.title}}</span>
              </template>
            </el-menu-item>
          </template>
        </el-submenu>
        <el-menu-item v-else  :index="menu.link+''">
          <i :class="menu.icon" v-if="menu.icon != ''"></i>
          <template #title>
            <span >{{menu.title}}</span>
          </template>
        </el-menu-item>
      </template>
    </el-menu>
</template>
<script>
  export default {
    name: 'ElAsiderMenu',
    data(){
      return{
        showCollapse: false,
        data:this.tree,
        treeData:[
        {
          'link':'/',
          'title':'面板',
          'icon':'el-icon-odometer',
        },
        {
          'link':'/hospital/',
          'title':'医院管理',
          'icon':'el-icon-video-camera',
          'children':[
          {
            'link':'/hospital/',
            'title':'医院管理',
            'icon':'el-icon-video-camera',
          },
          ]
        },
        ]
      };
    },
    props: {
      tree:{
        type: Array,
        default: [],
      },
      isCollapse:{
        type: Boolean,
        default: false,
      }
    },
    
    mounted() {
      console.log('this.data',this.data)
      this.showCollapse = this.isCollapse
    },
    methods:{
      showAside(){
        this.isCollapse = !this.isCollapse;
      },
      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClick(){
        console.log('isCollapse',this.isCollapse)
        this.isCollapse = !this.isCollapse;
      },
      handleSelect(val){
        console.log('val',val)
        this.$emit('select',val)
      }
    },
  };
</script>
<style>

</style>