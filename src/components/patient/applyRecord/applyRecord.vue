<template>
<div  class= "applyRecord">
  <div align="center">
    <h >远程请求电子病历</h>
    </div>
    <div align="center" >
<el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="150px" class="demo-ruleForm" >
  <el-form-item label="请求医院编号" prop="targetNo">
    <el-input v-model.number="ruleForm2.targetNo"></el-input>
  </el-form-item>
  <el-form-item label="请求病历编号" prop="redicalNO">
    <el-input v-model.number="ruleForm2.redicalNO"></el-input>
  </el-form-item>


  <el-form-item label="请求人" prop="applier">
    <el-input v-model.number="ruleForm2.applier"></el-input>
  </el-form-item>


  <el-form-item label="请求人编号" prop="applierNo">
    <el-input v-model.number="ruleForm2.applierNo"></el-input>
  </el-form-item>

  <el-form-item label="请求人所在医院编号" prop="applierHisNo">
    <el-input v-model.number="ruleForm2.applierHisNo"></el-input>
  </el-form-item>

  <el-form-item>
    <el-button type="primary" @click="submitForm('ruleForm2')">发出请求</el-button>
    <el-button @click="resetForm('ruleForm2')">重置</el-button>
  </el-form-item>
</el-form>
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
          targetNo: '',
          redicalNO: '',
          applier: '',
          applierNo: '',
          applierHisNo:''
        },
        rules2: {
          targetNo: [
            { validator: check, trigger: 'blur' }
          ],
          redicalNO: [
            { validator: check, trigger: 'blur' }
          ],
          applier: [
            { validator: check, trigger: 'blur' }
          ],
          applierNo: [
            { validator: check, trigger: 'blur' }
          ],
          applierHisNo: [
            { validator: check, trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
     
        this.$refs[formName].validate((valid) => {
        
          if (valid) {
                 alert("请患者授权");
            alert('submit!');
          } else {
             alert('error submit!!');
        
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
</script>
<style lang="stylus-loader" rel="stylesheet/stylus">   
  .applyRecord
    padding:10px
 

</style>
