<template>

<div>
<div align="center">
  <el-row>
      <el-input v-model="input" width="20dp" icon="search" @click="handelSearch" placeholder="请输入要查询电子病历的编号"></el-input>
    <el-button type="primary" round>查询</el-button>
  </el-row>
</div>

<div>

</div>
      <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="recordInfo"
        label="电子病历信息"
        align="center"
      >
      </el-table-column>
    </el-table>
      <div      align="center">
      <el-button slot="trigger" size="big" type="primary">下载</el-button>
  </div>
</div>


</template>

<script type="text/ecmascript-6">
  
</script>

<style lang="stylus-loader" rel="stylesheet/stylus">
  .hisIntroduction .span_font
    font-size:24px
    color:blue
  .hisIntroduction .introduction
    width:700px
    height:550px
    float:left
    position:relative
    margin-top:50px
  .hisIntroduction .picture-display
    float:right
    width:600px
    height:500px
  .hisIntroduction .picture-display ul li
    float:left
    margin-top:50px
  .hisIntroduction .card
    float:left
    margin-top:50px
    padding-left:20px
  .hisIntroduction .text
    font-size:14px
  .hisIntroduction .item
    padding:18px 0
  .hisIntroduction .clearfix:before, .clearfix:after
    display: table
    content: ""
  .hisIntroduction .clearfix:after
    clear: both
  .hisIntroduction .box-card
    width: 420px
    background:green
  .hisIntroduction .culture
    margin:auto
    margin-top:50px
    width:500px 
    padding-left:150px
    float:left
</style>