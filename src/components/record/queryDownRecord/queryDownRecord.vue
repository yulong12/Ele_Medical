<template>

<div>

<div align="center">
  <el-row>
      <el-input v-model="input" width="20dp" icon="search" @click="handelSearch" placeholder="请输入要查询患者的ID"></el-input>
    <el-button type="primary" round>查询</el-button>
  </el-row>
</div>
<!-- //DocterNo:医生职工编号
//HisNo：医院编号
// recordNo: 病历编号,
// recordName: 病历名称,
// recordPath: 病历存储路径,
// recordSize: 病历大小,
// recordHash: 病历的hash值 -->
    <div>
<el-table
    :data="tableData5"
    style="width: 100%">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
            <el-form-item label="病历编号">
            <span>{{ props.row.recordNo }}</span>
          </el-form-item>
          <el-form-item label="病历存储路径">
            <span>{{ props.row.recordPath }}</span>
          </el-form-item>
          <el-form-item label="所在医院编号">
            <span>{{ props.row.hisNo}}</span>
          </el-form-item>
          <el-form-item label="时间">
            <span>{{ props.row.recordTime }}</span>
          </el-form-item>
          <el-form-item label="医生职工编号">
            <span>{{ props.row.docterNo }}</span>
          </el-form-item>
          <el-form-item label="病历名称">
            <span>{{ props.row.recordName }}</span>
          </el-form-item>
          <el-form-item label="病历大小">
            <span>{{ props.row.recordSize }}</span>
          </el-form-item>
          <el-form-item label="病历的hash值">
            <span>{{ props.row.recordHash }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column
      label="时间"
      prop="recordTime">
    </el-table-column>
    <el-table-column
      label="病历名称"
      prop="recordName">
    </el-table-column>
    <el-table-column
      label="病历编号"
      prop="recordNo">
    </el-table-column>
  </el-table>
    </div>
    
      <div   align="center">
      <el-button slot="trigger" size="big" type="primary">下载</el-button>
      </div>
</div>


</template>

<script type="text/ecmascript-6">
  export default {
    data() {
      return {
        tableData5: [{
          recordNo:'2120170729',
          recordPath: 'public/images/2019-03-20',
          hisNo: '2211222',
          recordTime: '2019-03-20 21:05:16 +0800 CST',
          docterNo: '2211222012',
          recordName: 'x光片',
          recordSize: '789',
          recordHash: 'afdsaffasfdgfdjhfgjt'
        }, {
          recordNo:'2120170730',
          recordPath: 'public/images/2019-03-21',
          hisNo: '1111111',
          recordTime: '2019-03-20 21:05:16 +0800 CST',
          docterNo: '2211222012',
          recordName: 'CT',
          recordSize: '789',
          recordHash: 'xvczcvzxcvxczv'
        },{
           recordNo:'2120170731',
          recordPath: 'public/images/2019-03-22',
          hisNo: '2211222',
          recordTime: '2019-03-20 21:05:16 +0800 CST',
          docterNo: '2211222012',
          recordName: 'x光片',
          recordSize: '789',
          recordHash: 'qwrewqrqwerwrw'
        },{
           recordNo:'2120170732',
          recordPath: 'public/images/2019-03-23',
          hisNo: '2211222',
          recordTime: '2019-03-20 21:05:16 +0800 CST',
          docterNo: '2211222012',
          recordName: 'x光片',
          recordSize: '789',
          recordHash: 'khjkhkhfkghgfhk'
        }]
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