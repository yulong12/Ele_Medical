<template>
<div align="center">
  <span class="title">上传电子病历</span>
  
  <div>
<el-upload
  class="upload-demo"
  drag
  action="https://jsonplaceholder.typicode.com/posts/"
  multiple>
  <i class="el-icon-upload"></i>
  <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
  <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过500kb</div>
</el-upload>
  </div>
<el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="150px" class="demo-ruleForm" >
  <el-form-item label="病人编号" prop="patientNo">
    <el-input v-model.number="ruleForm2.patientNo"></el-input>
  </el-form-item>
</el-form>

  <div>
      <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="hashValue"
        label="文件hash值"
        align="center"
      >
      </el-table-column>
    </el-table>
  </div>

  <div align="center">
      <el-button slot="trigger" size="big" type="primary">哈希值上链</el-button>
  </div>


</div>
</template>

<script>
 export default {
    data() {
      var check = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('不能为空'));
        }
       
      };


      return {
        ruleForm2: {
          patientNo: ''
       
        },
        rules2: {
          patientNo: [
            { validator: check, trigger: 'blur' }
          ]
        }
      };
    },
 }
</script>

<style lang="stylus-loader" rel="stylesheet/stylus">
  .patientFeedback .title
    font-size:24px
    margin-left:300px
  .patientFeedback .form
     margin-top:80px
  .patientFeedback
    height:850px
    width:900px
    margin:auto
    border:1px solid black
  .patientFeedback .checkbox
    padding-bottom:50px
  .patientFeedback .question
    padding-right:100px

</style>