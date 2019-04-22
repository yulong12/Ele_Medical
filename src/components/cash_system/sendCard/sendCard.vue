<template>
  <div class="sendCard">
    <p class="sendCard-title">办理健康卡</p>
   <el-form :model="sendCardForm" status-icon :rules="sendCardrules" ref="sendCardForm" label-width="100px" class="sendCard-ruleForm">

  <el-form-item label="身份证号" prop="idCard">
    <el-input v-model="sendCardForm.idCard" ></el-input>
  </el-form-item>

  <el-form-item label="性别" prop="sex">
    <el-input  v-model="sendCardForm.sex" ></el-input>
  </el-form-item>

  <el-form-item label="年龄" prop="age">
    <el-input v-model.number="sendCardForm.age"></el-input>
  </el-form-item>

 <el-form-item label="家庭住址" prop="address">
    <el-input v-model.number="sendCardForm.address"></el-input>
  </el-form-item>

  <el-form-item>
    <el-button type="primary" @click="submitForm('sendCardForm')">办卡</el-button>
    <el-button @click="resetForm('sendCardForm')">重置</el-button>
  </el-form-item>
</el-form>
  </div>
</template>

<script type="text/ecmascript-6">
  export default {
    data() {
      // 验证输入身份证号格式是否正确的方法
      var  checkCertificate = (rule, value, callback) => {
        if (/^\d+$/.test(value) !== false) {
          if (value.length !== 18) {
            callback(new Error('身份证号应为18位数！'));
          } else {
            callback();
          }
        } else {
          callback(new Error('请输入数字值！'));
        }
      };
   var checkAge = (rule, value, callback) => {
        if (!value) {
          return callback(new Error('年龄不能为空'));
        }
        setTimeout(() => {
          if (!Number.isInteger(value)) {
            callback(new Error('请输入数字值'));
          } else {
            callback();
            }
  
        }, 1000);
      };
      return {
        sendCardForm: {
          idCard: '',
          sex: '',
          age: '',
          address:''
        },
        sendCardrules: {
          idCard: [
            { required:true, validator: checkCertificate, trigger: 'blur' }
          ],
          sex: [
            { required:true,message:"请输入性别" ,trigger: 'blur' }
          ],
          age: [
            {required:true, validator:checkAge, trigger: 'blur' }
          ],
         address: [
            { required:true,message:"请输入住址", trigger: 'blur' }
          ]
        }
      };
    },
    methods: {
      submitForm (formName) {
        this.$refs.sendCardForm.validate((valid) => {
          if (valid) {
            this.$notify({
              message: '办卡成功',
              type: 'success'
            });
          } else {
            this.$notify.error({
              message: '提交失败'
            });
            return false;
          }
        });
      },
      resetForm (formName) {
        this.$refs.sendCardForm.resetFields();
        this.departmentvalue = '';
      }
    }
  }
</script>

<style lang="stylus-loader" rel="stylesheet/stylus">

.sendCard 
    .sendCard-title
      font-weight: bold
      font-size: 28px
      width: 50%
      margin: 0 auto
      text-align: center
      padding-bottom: 40px
    .sendCard-ruleForm
      width: 30%
      margin: 0 auto
      padding-left: 20px
      .register-name
        width: 315px
      .submitBtn
        margin-left: 30px   
</style>
