<template>
  <div class="app-container">

    <div class="filter-container">

      <el-radio v-model="input.typ" class="filter-item" label="" @change="search">{{ $t('全部') }}</el-radio>
      <el-radio v-model="input.typ" class="filter-item" label="mysql" @change="search">{{ $t('mysql') }}</el-radio>
      <!--<el-radio @change="search" v-model="input.typ" class="filter-item" label="sqlserver">sqlserver</el-radio>-->
      <el-radio v-model="input.typ" class="filter-item" label="clickhouse" @change="search">{{ $t('clickhouse') }}
      </el-radio>
      <!--<el-radio @change="search" v-model="input.typ" class="filter-item" label="mongodb">mongodb</el-radio>-->
      <el-input
        v-model="input.remark"
        style="width: 300px"
        class="filter-item"
        clearable
        :placeholder="$t('备注')"
      />
      <el-button
        size="mini"
        type="success"
        class="filter-item"
        @click="search"
      >{{ $t('查询') }}
      </el-button>
      <el-button
        size="mini"
        type="primary"
        class="filter-item"
        @click.native="open = true"
      >{{ $t('新增') }}
      </el-button>
    </div>

    <el-table v-loading="tableLoading" :data="tableData">
      <el-table-column
        width="80"
        align="center"
        prop="id"
        label="id"
      />
      <el-table-column width="100" align="center" :label="$t('类型')">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.typ == 'clickhouse'" type="primary">
            {{ scope.row.typ }}
          </el-tag>
          <el-tag v-if="scope.row.typ == 'mysql'" type="success">
            {{ scope.row.typ }}
          </el-tag>
          <!--<el-tag v-if="scope.row.typ == 'mongodb'" type="warning">
            {{ scope.row.typ }}
          </el-tag>-->
        </template>
      </el-table-column>

      <el-table-column
        align="center"
        prop="ip"
        label="ip"
        width="180"
      />
      <el-table-column
        align="center"
        prop="port"
        :label="$t('端口')"
        width="80"
      />

      <el-table-column
        width="100"
        align="center"
        prop="username"
        :label="$t('用户名')"
      />

      <el-table-column
        width="100"
        prop="db_name"
        :label="$t('数据库名')"
      />
      <el-table-column
        align="center"
        width="100"
        prop="remark"
        :label="$t('备注')"
      />
      <el-table-column
        width="200"
        align="center"
        prop="created"
        :label="$t('创建时间')"
      />
      <el-table-column
        width="200"
        align="center"
        prop="updated"
        :label="$t('修改时间')"
      />
      <el-table-column align="center" :label="$t('操作')" fixed="right" width="300">
        <template slot-scope="scope">
          <el-button

            v-loading="loadingList[scope.$index].loading"
            icon="el-icon-link"
            type="success"
            size="mini"
            :disabled="loadingList[scope.$index].loading"
            @click="link({typ:scope.row.typ,db_name:scope.row.db_name,ip:scope.row.ip,
                          port:scope.row.port,username:scope.row.username,pwd:scope.row.pwd,
            },scope.$index)"
          >{{ $t('测试连接') }}
          </el-button>
          <el-button

            icon="el-icon-delete"
            type="danger"
            size="mini"
            @click="deleteById(scope.row.id)"
          >{{ $t('删除') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination
        v-if="pageshow"
        background
        :current-page="input.page"
        :page-size="input.limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="count"
        @current-change="search"
        @size-change="handleSizeChange"
      />
    </div>

    <el-dialog :close-on-click-modal="false" :visible.sync="open" :title="$t('新增数据源')" @close="closeDialog">
      <el-card class="box-card">
        <el-form label-width="300px" label-position="left">
          <el-form-item label="IP">
            <el-input v-model="form.ip" clearable placeholder="127.0.0.1" />
          </el-form-item>
          <el-form-item :label="$t('端口')">
            <el-input v-model="form.port" placeholder="3306" />
          </el-form-item>
          <el-form-item :label="$t('用户名')">
            <el-input v-model="form.username" clearable placeholder="mysql" />
          </el-form-item>
          <el-form-item :label="$t('密码')">
            <el-input v-model="form.pwd" clearable placeholder="root" />
          </el-form-item>
          <el-form-item :label="$t('数据库名')">
            <el-input v-model="form.db_name" clearable placeholder="mysql" />
          </el-form-item>
          <el-form-item :label="$t('备注')">
            <el-input v-model="form.remark" clearable placeholder="测试" />
          </el-form-item>
          <el-form-item :label="$t('数据源类型')">
            <el-radio v-model="form.typ" label="mysql">mysql</el-radio>
            <!--<el-radio v-model="form.typ" label="sqlserver">sqlserver</el-radio>-->
            <el-radio v-model="form.typ" label="clickhouse">clickhouse</el-radio>
            <!--<el-radio v-model="form.typ" label="mongodb">mongodb</el-radio>-->
          </el-form-item>
        </el-form>
        <div style="text-align:right;">
          <el-button
            size="mini"
            type="danger"
            icon="el-icon-close"
            @click="closeDialog"
          >{{ $t('取消') }}
          </el-button>
          <el-button
            size="mini"
            type="primary"
            icon="el-icon-check"
            @click="add"
          >{{ $t('确认') }}
          </el-button>
          <el-button
            v-loading="formLinkLoading"
            size="mini"
            :disabled="formLinkLoading"
            type="success"
            icon="el-icon-link"
            @click="link({typ:form.typ,db_name:form.db_name,ip:form.ip,
                          port:form.port,username:form.username,pwd:form.pwd,
            })"
          >{{ $t('测试连接') }}
          </el-button>
        </div>
      </el-card>
    </el-dialog>

  </div>
</template>

<script>

import { DelLinkById, InsertLink, LinkInfoList, TestLink } from '@/api/datax'

const defaultForm = {
  ip: '127.0.0.1',
  port: 3306,
  db_name: 'test',
  username: 'root',
  pwd: '',
  remark: 'test',
  typ: 'mysql'
}

export default {
  name: 'Link',
  data() {
    return {
      formLinkLoading: false,
      count: 0,
      pageshow: true,
      tableData: [],
      form: Object.assign({}, defaultForm),
      input: {
        remark: '',
        typ: '',
        page: 1,
        limit: 10
      },
      loadingList: [],
      tableLoading: false,
      open: false
    }
  },
  created() {
    this.search()
  },
  methods: {

    async link(form, index) {
      if (form.port == '') form.port = 0
      form.port = Number(form.port)
      if (index != undefined) {
        this.loadingList[index]['loading'] = true
      } else {
        this.formLinkLoading = true
      }

      const res = await TestLink(form)
      if (index != undefined) {
        this.loadingList[index].loading = false
      } else {
        this.formLinkLoading = false
      }
      if (res.code != 0) {
        this.$message({
          type: 'error',
          message: res.msg
        })
        return
      }

      this.$message({
        type: 'success',
        message: res.msg
      })
    },
    async deleteById(id) {
      const res = await DelLinkById({ 'id': id })
      if (res.code != 0) {
        this.$message({
          type: 'error',
          message: res.msg
        })
        return
      }
      this.$message({
        type: 'success',
        message: res.msg
      })
      this.search()
    },
    refreshPage() {
      this.pageshow = false
      this.count = this.tableData.length
      this.$nextTick(() => {
        this.pageshow = true
      })
    },
    handleSizeChange(v) {
      this.input.limit = v
      this.refreshPage()
    },
    async add() {
      if (this.form.port == '') this.form.port = 0
      this.form.port = Number(this.form.port)
      const res = await InsertLink(this.form)
      if (res.code != 0) {
        this.$message({
          type: 'error',
          message: res.msg
        })
        return
      }
      this.$message({
        type: 'success',
        message: res.msg
      })
      this.open = false
      this.search()
    },
    closeDialog() {
      this.open = false
      this.form = Object.assign({}, defaultForm)
      this.formLinkLoading = false
    },
    async search() {
      this.tableLoading = true
      const res = await LinkInfoList(this.input)
      if (res.code != 0) {
        this.$message({
          type: 'error',
          message: res.msg
        })
        return
      }
      this.tableData = res.data.data
      this.count = res.data.count
      if (this.tableData == null) {
        this.tableData = []
      }
      this.loadingList = []
      for (const k in this.tableData) {
        this.loadingList.push({
          loading: false
        })
      }

      this.$message({
        type: 'success',
        message: res.msg
      })
      this.tableLoading = false
    },
    handleClick() {

    }
  }
}
</script>

<style scoped>

</style>
