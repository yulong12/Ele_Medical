<template>
<div>
  <span class="title">上传电子病历</span>
  <div>
    <el-upload
  class="upload-demo"
  ref="upload"
  action="https://jsonplaceholder.typicode.com/posts/"
  :on-preview="handlePreview"
  :on-remove="handleRemove"
  :file-list="fileList"
  :auto-upload="false">
  <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
  <el-button style="margin-left: 10px;" size="small" type="success" @click="submitUpload">上传到服务器</el-button>
   <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
   </el-upload>
  </div>
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

  <div      align="center">
      <el-button slot="trigger" size="big" type="primary">哈希值上链</el-button>
  </div>


</div>
</template>

<script>
  export default {
    data() {
      return {
        fileList: [{name: 'food.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}, {name: 'food2.jpeg', url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'}]
      };
    },
    methods: {
      submitUpload() {
        this.$refs.upload.submit();
      },
      handleRemove(file, fileList) {
        console.log(file, fileList);
      },
      handlePreview(file) {
        console.log(file);
      }
    }
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