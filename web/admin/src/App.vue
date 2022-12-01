<template>
    <!-- ant ui国际化（中文） -->
  <a-config-provider :locale="zh_CN">
    <div id="app">
        <router-view></router-view>
    </div>
  </a-config-provider>
</template>

<style scoped>


</style>

<script>
    import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN';
    import moment from 'moment';
    import 'moment/locale/zh-cn';
  
    moment.locale('zh-cn');
    export default {
      data() {
        return {
          zh_CN,
        };
      },
    };
  </script>
