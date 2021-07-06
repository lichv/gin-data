<template>
  <div :id="id" :style="{ width: '100%', height: '100%',minWidth:'320px',minHeight:'180px' }"></div>
</template>
<script>
  import api from '../api';
  export default {
    name: 'ElChart',
    data(){
      return{
        id:Math.random().toString(36).substr(2),
        data:[],
        chartPie: {},
      };
    },
    props: {
      
    },
    mounted() {
      this.getData()
      // this.initChartBar(this.id,{title:"总用户量",x:["12-3", "12-4", "12-5", "12-6", "12-7", "12-8"],y:[5, 20, 36, 10, 10, 20]})
      // this.initChartPie(this.id,{name:"Sanger分布",data:[{"value":100,"name":"异常周期样本"},{"value":80,"name":"生成验证报告"},{"value":50,"name":"样本开始检测"}]})
    },
    methods:{
      initChartBar(domId, echartData) {
      var option = {
        // //dataZoom-inside 内置型数据区域缩放组件 所谓内置 1平移：在坐标系中滑动拖拽进行数据区域平移。2缩放：PC端：鼠标在坐标系范围内滚轮滚动（MAC触控板类同;移动端：在移动端触屏上，支持两指滑动缩放。
        dataZoom: [
          {
            type: "inside", //1平移 缩放
            throttle: "50", //设置触发视图刷新的频率。单位为毫秒（ms）。
            minValueSpan: 10, //用于限制窗口大小的最小值,在类目轴上可以设置为 5 表示 5 个类目
            start: 100, //数据窗口范围的起始百分比 范围是：0 ~ 100。表示 0% ~ 100%。
            end: 100, //数据窗口范围的结束百分比。范围是：0 ~ 100。
            zoomLock: true, //如果设置为 true 则锁定选择区域的大小，也就是说，只能平移，不能缩放。
          },
        ],
        title: {
          text: echartData.title,
        },
        // 提示框组件
        tooltip: {
          trigger: "axis", //坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
          // backgroundColor: "#377CFF", //提示框浮层的背景颜色。
          axisPointer: {
            //去掉移动的指示线
            type: "none",
          },
        },
        xAxis: {
          type: "category",
          data: echartData.x,
          axisTick: {
            show: false, //x轴刻度线设置
          },
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            type: "bar",
            data: echartData.y,
          },
        ],
      };

      let getchart = this.$echarts.init(document.getElementById(domId));

      getchart.setOption(option);

      //随着屏幕大小调节图表
      window.addEventListener("resize", () => {
        getchart.resize();
      });
    },
    initChartPie(domId, pData) {
      if (!this.chartPie[domId]) {
        let getchart = this.$echarts.init(document.getElementById(domId));

        var option = {
          // 提示框组件
          tooltip: {
            trigger: "item",
          },
          series: [
            {
              name: pData.name,
              type: "pie",
              radius: "50%",
              data: pData.data,
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: "rgba(0, 0, 0, 0.5)",
                },
              },
            },
          ],
        };
        getchart.setOption(option);

        //随着屏幕大小调节图表
        window.addEventListener("resize", () => {
          getchart.resize();
        });

        this.chartPie[domId] = getchart;
      } else {
        var option = this.chartPie[domId].getOption();
        // 先清空
        option.series[0].data = [];
        this.chartPie[domId].setOption(option);
        // 再设置
        option.series[0].data = pData.data;
        this.chartPie[domId].setOption(option);
      }
    },
    initChartLine(xData, yData) {
      if (!this.chartLine) {
        let getchart = this.$echarts.init(
          document.getElementById("echart-line")
        );

        var option = {
          //dataZoom-inside 内置型数据区域缩放组件 所谓内置 1平移：在坐标系中滑动拖拽进行数据区域平移。2缩放：PC端：鼠标在坐标系范围内滚轮滚动（MAC触控板类同;移动端：在移动端触屏上，支持两指滑动缩放。
          dataZoom: [
            {
              type: "inside", //1平移 缩放
              throttle: "50", //设置触发视图刷新的频率。单位为毫秒（ms）。
              minValueSpan: 5, //用于限制窗口大小的最小值,在类目轴上可以设置为 5 表示 5 个类目
              start: 100, //数据窗口范围的起始百分比 范围是：0 ~ 100。表示 0% ~ 100%。
              end: 100, //数据窗口范围的结束百分比。范围是：0 ~ 100。
              zoomLock: true, //如果设置为 true 则锁定选择区域的大小，也就是说，只能平移，不能缩放。
            },
          ],
          // 提示框组件
          tooltip: {
            trigger: "axis", //坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
            // backgroundColor: "#377CFF", //提示框浮层的背景颜色。
            axisPointer: {
              //去掉移动的指示线
              // type: "none",
            },
          },
          xAxis: {
            type: "category",
            data: xData,
            axisTick: {
              show: false, //x轴刻度线设置
            },
          },
          yAxis: {
            type: "value",
          },
          series: [
            {
              type: "line",
              data: yData,
            },
          ],
        };
        getchart.setOption(option);

        //随着屏幕大小调节图表
        window.addEventListener("resize", () => {
          getchart.resize();
        });

        this.chartLine = getchart;
      } else {
        var option = this.chartLine.getOption();
        // 先清空
        option.xAxis[0].data = [];
        option.series[0].data = [];
        this.chartLine.setOption(option);
        // 再设置
        option.xAxis[0].data = xData;
        option.series[0].data = yData;
        option.dataZoom[0].start = 100;
        option.dataZoom[0].end = 100;
        this.chartLine.setOption(option);
      }
    },
    getData:function(){
        let _this = this;
        api.getProductDistribution.send()
        .then(result => {
          console.log('result',result)
          if (result.state==2000) {
            _this.data = result.data
            _this.initChartPie(_this.id,{name:"产品分布",data:_this.data})
          }
        })
        .catch(e => {
          console.log(e)
        })
      },
    },
  };
</script>
<style>

</style>
