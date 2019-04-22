<template>
    <div class="deleteDepartment">
      <span class="tittle"></span>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px" class="addForm">
        <el-form-item label="科室名称" prop="name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>

        <el-form-item label="科室编号" prop="departNO">
          <el-input v-model="form.departNO"></el-input>
        </el-form-item>

        <el-form-item class="buttonAdd">
          <el-button type="primary" @click="submitForm('form')">删 除</el-button>
          <el-button type="danger" @click="resetForm ('form')">重 置</el-button>
        </el-form-item>
      </el-form>
    </div>
</template>

<script type="text/ecmascript-6">
    export default {
      data () {
        return {
          form: {
            name: '',
            departNO: '',
            attention: ''
          },
          // 验证规则
          rules: {
            name: [
              { required: true, message: '请输入科室名称', trigger: 'blur' },
              { min: 2, max: 5, message: '长度在 3 到 5 个字符', trigger: 'blur' }
            ],
            departNO: [
              { required: true, message: '请输入科室编号', trigger: 'change' }
            ]
          }
        };
      },
      methods: {
        // 增加的方法
        submitForm (form) {
        // 规则验证
          this.$refs[form].validate((valid) => {
            if (valid) {
        // 提示信息
              this.$message({
                showClose: true,
                message: '添加成功！',
                type: 'success'
              });
              var json = JSON.stringify(this.form);
              console.log('您修改后的参数为1：', JSON.stringify(json));
            } else {
              this.$message({
                showClose: true,
                message: '添加失败！',
                type: 'error'
              });
              return false;
            }
          });
        },
        // 重置方法
        resetForm (form) {
          this.$refs[form].resetFields();
        }
      }
};

</script>

<style lang="stylus-loader" rel="stylesheet/stylus">
.deleteDepartment .el-input
  width:384px
.deleteDepartment .tittle
  display:block
  text-align: center
  font-size:24px
  font-weight:600
  line-height:80px
.deleteDepartment .addForm
  width:470px
  margin:auto
.deleteDepartment .addForm .buttonAdd
  text-align: center;
</style>