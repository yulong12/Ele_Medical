<template>

<div>

<div align="center">
  <el-row>
      <el-input v-model="input" width="20dp" icon="search" @click="handelSearch" placeholder="请输入要删除的电子病历模版编号"></el-input>
    <el-button type="primary" round>删 除</el-button>
  </el-row>
</div>
    

</div>


</template>

<script type="text/ecmascript-6">
  export default {
    data() {
      return {
     
      }
    }
  }
</script>

<style lang="stylus-loader" rel="stylesheet/stylus">
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>