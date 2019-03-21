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
 <!-- <el-form-item label="文件" label-width="120px" margin="20px">
        <el-upload class="upload-demo" action="" :auto-upload="false" :limit="1" ref="upload" :http-request="upload" multiple>
          <el-button size="small" type="primary">模拟上传2</el-button>
        </el-upload>
      </el-form-item>
 <el-button class="btn" size="small" type="primary" @click="upload">确定上传2</el-button> -->

  </div>
  
<el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="150px" class="demo-ruleForm" >
  <el-form-item label="病人ID" prop="patientNo">
    <el-input v-model.number="ruleForm2.patientNo"></el-input>
  </el-form-item>
    <el-form-item label="病历编号" prop="recordNO">
    <el-input v-model.number="ruleForm2.recordNO"></el-input>
  </el-form-item>
    <el-form-item label="医生编号" prop="doctorNo">
    <el-input v-model.number="ruleForm2.doctorNo"></el-input>
  </el-form-item>
    <el-form-item label="医院编号" prop="hisNo">
    <el-input v-model.number="ruleForm2.hisNo"></el-input>
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
      <el-button slot="trigger" size="big" type="primary">上传</el-button>
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
            
          ],
          recordNO: [
            { validator: check, trigger: 'blur' }
            
          ],
          doctorNo: [
            { validator: check, trigger: 'blur' }
            
          ],
          hisNo: [
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