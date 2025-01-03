<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-1'" />
        </div>
      </el-aside>

      <el-container>
        <el-header style="display: flex; justify-content: space-between; align-items: center;">
          <h2>{{ headerTitle }}</h2>
          <div>
            搜索：
            <el-input v-model="searchQuery" placeholder="输入要搜索的关键字" style="width: 200px;" />
            <el-button type="primary" @click="handleAdd">{{ addButtonText }}</el-button>
            <el-button @click="navigateToDictionaryManager">数据字典</el-button>
          </div>
        </el-header>
        <el-main>
          <!-- 会计实体信息表格 -->
          <el-table :data="paginatedAcctData" style="width: 100%" max-height="450">
            <el-table-column prop="AcctCode" label="会计实体编码" width="160%"></el-table-column>
            <el-table-column prop="AcctAbbr" label="会计实体缩写" width="160%"></el-table-column>
            <el-table-column prop="EtyAbbr" label="实体简称" width="160%"></el-table-column>
            <el-table-column prop="AcctName" label="会计实体名称" width="220%"></el-table-column>

            <el-table-column label="绑定的银行账户信息" width="220%">
              <template #default="scope">
                <span v-if="scope.row.AcctBanks && scope.row.AcctBanks.length > 0">
                  {{ scope.row.AcctBanks.map(bank => bank.AccNum).join(', ') }}
                </span>
                <span v-else>无</span>
              </template>
            </el-table-column>

            <el-table-column prop="AcctEngName" label="会计实体英文名称" width="220%"></el-table-column>
            <el-table-column prop="AcctAddr" label="会计实体地址" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家"></el-table-column>
            <el-table-column prop="TaxType" label="税号类型" width="220%"></el-table-column>
            <el-table-column prop="TaxCode" label="税号" width="220%"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="Website" label="网站" width="220%"></el-table-column>
            <el-table-column prop="RegDate" label="注册日期" width="450%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>

            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.AcctId)" type="text"
                    size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="acctData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加会计实体信息的对话框 -->
    <el-dialog v-model="showAcctDialog" title="会计实体信息" width="80%" @close="resetAcctForm">
      <el-form :model="acctForm" label-width="150px" :rules="acctRules" ref="acctFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体编码" prop="AcctCode" :required="true">
              <el-input v-model="acctForm.AcctCode" placeholder="国家代码+流水码"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体缩写" prop="AcctAbbr" :required="true">
              <el-input v-model="acctForm.AcctAbbr" placeholder="名称缩写，一般2-3个字母"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实体简称" prop="EtyAbbr" :required="true">
              <el-input v-model="acctForm.EtyAbbr" placeholder="客户名称简称，一般1-2个单词"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体名称" prop="AcctName" :required="true">
              <el-input v-model="acctForm.AcctName" placeholder="全称"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体英文名称" prop="AcctEngName">
              <el-input v-model="acctForm.AcctEngName" placeholder="英文全称（仅限注册在中国大陆的会计实体）"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体地址" prop="AcctAddr">
              <el-input v-model="acctForm.AcctAddr"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="acctForm.Nation"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-select v-model="acctForm.TaxType" placeholder="请选择">
                <el-option label="类型1" value="type1"></el-option>
                <el-option label="类型2" value="type2"></el-option>
                <!-- 添加更多选项 -->
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="acctForm.TaxCode"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="acctForm.PhoneNum"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="acctForm.Email"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="acctForm.Website"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册日期" prop="RegDate">
              <el-date-picker v-model="acctForm.RegDate" type="date" placeholder="选择日期"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="acctForm.Notes" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!acctForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(acctForm.FileId, acctForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="acctForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体银行账户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">

              <el-input v-model="acctForm.AcctBanksDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitAcctForm">保存</el-button>
            <el-button @click="showAcctDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>


    <!-- 查看会计实体信息的对话框 -->
    <el-dialog v-model="showshowAcctDialog" title="会计实体信息" width="80%" @close="resetAcctForm">
      <el-form :model="acctForm" label-width="150px" :rules="acctRules" ref="acctFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体编码" prop="AcctCode" :required="true">
              <el-input v-model="acctForm.AcctCode" placeholder="国家代码+流水码" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体缩写" prop="AcctAbbr" :required="true">
              <el-input v-model="acctForm.AcctAbbr" placeholder="名称缩写，一般2-3个字母" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="实体简称" prop="EtyAbbr" :required="true">
              <el-input v-model="acctForm.EtyAbbr" placeholder="客户名称简称，一般1-2个单词" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体名称" prop="AcctName" :required="true">
              <el-input v-model="acctForm.AcctName" placeholder="全称" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="会计实体英文名称" prop="AcctEngName">
              <el-input v-model="acctForm.AcctEngName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="会计实体地址" prop="AcctAddr">
              <el-input v-model="acctForm.AcctAddr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="acctForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>

          <el-col :span="12">
            <el-form-item label="税号类型" prop="TaxType">
              <el-input v-model="acctForm.TaxType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="税号" prop="TaxCode">
              <el-input v-model="acctForm.TaxCode" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="acctForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="acctForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="网站" prop="Website">
              <el-input v-model="acctForm.Website" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="注册日期" prop="RegDate">
              <el-date-picker v-model="acctForm.RegDate" type="date" placeholder="选择日期"
                :readonly="true"></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="acctForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!acctForm.FileId" ref="uploadRef" action="" :limit="1" :on-change="handleFileChange"
                :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(acctForm.FileId, acctForm.FileName)">下载文件</el-button>
            </el-form-item>

            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="acctForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 会计实体银行账户信息选择 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="关联银行账户信息">
              <el-input v-model="acctForm.AcctBanksDisplay" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showshowAcctDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>


    <!-- <header class="top-bar">
      <el-button @click="navigateToDictionaryManager">数据字典</el-button>
    </header> -->


  </div>
</template>

<script setup>

import { ref, onMounted, computed } from 'vue';

import SideMenu from '@/components/SideMenu.vue'; // 引入 SideMenu 组件

import { ElMessage, ElMessageBox } from 'element-plus'; // 确保 ElMessageBox 被引入
import axios from 'axios'; // 引入 axios
import { useRouter } from 'vue-router';

const router = useRouter();

// 方法：导航到 DictionaryManager 页面
const navigateToDictionaryManager = () => {
  router.push('/dictionaryManager');
};


const searchQuery = ref(''); // 添加搜索查询字段


const downloadFile = async (fileId, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileId: fileId }, // 提供 FileId 表单
      {
        responseType: 'blob', // 指定响应类型为二进制数据
      }
    );

    // 创建下载链接
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    // link.setAttribute('download', `file_${fileId}`); // 设置下载文件名
    link.setAttribute('download', `${fileName}` || `file_${fileId}`); // 使用 fileName 作为文件名，如果不存在则使用默认文件名
    document.body.appendChild(link);
    link.click(); // 触发下载
    document.body.removeChild(link); // 移除链接
    window.URL.revokeObjectURL(url); // 释放 URL 对象
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

// 获取 el-upload 组件的引用
const uploadRef = ref(null);
// 文件数据
const file = ref(null);

// 文件选择事件
const handleFileChange = (uploadFile) => {
  file.value = uploadFile.raw; // 保存选择的文件
};
// 查看按钮逻辑
const handleView = (index, row) => {
  // 填充会计实体信息表单
  acctForm.value = { ...row }; // 将当前行的数据赋值给 acctForm
  // 将 AcctBanks 转换为可显示的字符串
  acctForm.value.AcctBanksDisplay = row.AcctBanks && row.AcctBanks.length > 0
    ? row.AcctBanks.map(bank => bank.AccNum).join(', ')
    : '无';
  showshowAcctDialog.value = true; // 打开会计实体信息对话框
  // 检查是否已上传文件
  if (row.FileId) {
    acctForm.value.FileId = row.FileId; // 保存 FileId
    acctForm.value.FileName = row.FileName; // 保存 FileId
  }
};
// 重置会计实体信息表单
const resetAcctForm = () => {
  acctForm.value = {
    AcctCode: '',
    AcctAbbr: '',
    EtyAbbr: '',
    AcctName: '',
    AcctEngName: '',
    AcctAddr: '',
    Nation: '',
    TaxType: '',
    TaxCode: '',
    PhoneNum: '',
    Email: '',
    Website: '',
    RegDate: '',
    Notes: '',
    License: '',
    BankAccounts: [] // 关联的银行账户信息
  };
  file.value = null; // 重置文件数据
  if (uploadRef.value) {
    uploadRef.value.clearFiles(); // 清空文件列表
  }
};
// 会计实体信息表单提交逻辑
const submitAcctForm = async () => {
  try {

    // acctForm.value.AcctId = parseInt(acctForm.value.AcctId, 10);
    // const response = await axios.post('/save/acct', acctForm.value); // 调用保存会计实体信息接口
    const formData = new FormData(); // 创建 FormData 对象

    // 添加表单数据
    Object.keys(acctForm.value).forEach((key) => {
      formData.append(key, acctForm.value[key]);
    });

    // 添加文件
    if (file.value) {
      formData.append('file', file.value);
    }

    const response = await axios.post('/save/acct', formData, {
      headers: {
        'Content-Type': 'multipart/form-data', // 设置请求头
      },
    });
    if (response.status === 200) {
      ElMessage.success('会计实体信息保存成功');
      showAcctDialog.value = false; // 关闭对话框
      fetchAcctData(); // 重新获取会计实体信息数据
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 删除按钮逻辑
const handleDelete = (index, AcctId) => {
  // console.log('Delete button clicked', index, row); // 添加调试信息
  ElMessageBox.confirm('确定要删除该会计实体信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    console.log('Confirmed delete', AcctId); // 添加调试信息
    axios.post('/delete/acct', {
      "AcctId": AcctId,
      "AcctAbbr": "ss",
      "AcctCode": "ss",
      "AcctName": "ss",
      "EtyAbbr": "ss",
    })
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchAcctData(); // 重新获取会计实体信息数据
        } else {
          ElMessage.error(response.data.RetMessage || '删除失败');
        }
      })
      .catch(error => {
        ElMessage.error(error.response.data.RetMessage);
      });
  }).catch(() => {
    ElMessage.info('已取消删除');
  });
};


const handlePageChange = (page) => {
  currentPage.value = page;
};
const currentPage = ref(1); // 当前页码
const pageSize = 8; // 每页显示的行数
// 计算当前页显示的会计实体信息数据

const paginatedAcctData = computed(() => {
  let filteredData = acctData.value;
  if (searchQuery.value) {
    filteredData = filteredData.filter(item =>
      item.AcctCode.includes(searchQuery.value) ||
      item.AcctAbbr.includes(searchQuery.value) ||
      item.EtyAbbr.includes(searchQuery.value) ||
      item.AcctName.includes(searchQuery.value) ||
      item.AcctEngName.includes(searchQuery.value) ||
      item.AcctAddr.includes(searchQuery.value) ||
      item.Nation.includes(searchQuery.value) ||
      item.TaxType.includes(searchQuery.value) ||
      item.TaxCode.includes(searchQuery.value) ||
      item.PhoneNum.includes(searchQuery.value) ||
      item.Email.includes(searchQuery.value) ||
      item.Website.includes(searchQuery.value) ||
      item.RegDate.includes(searchQuery.value) ||
      item.Notes.includes(searchQuery.value) ||
      (item.AcctBanks && item.AcctBanks.some(bank => bank.AccNum.includes(searchQuery.value))) // 新增：检查 AcctBanks 中的 AccNum
    );
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});
onMounted(() => {
  fetchAcctData(); // 获取会计实体信息
});
// 定义接口请求函数
const fetchAcctData = async () => {
  try {
    const response = await axios.get('/find/acct'); // 调用会计实体信息接口
    acctData.value = response.data.Acct; // 假设返回的数据结构中有 Acct 字段
  } catch (error) {
    console.error('获取会计实体信息失败:', error);
    ElMessage.error('获取会计实体信息失败，请稍后重试');
  }
};


// 会计实体信息对话框显示状态
const showAcctDialog = ref(false);

// 会计实体信息对话框显示状态
const showshowAcctDialog = ref(false);

// 会计实体信息表单数据
var acctForm = ref({
  AcctCode: '',
  AcctAbbr: '',
  EtyAbbr: '',
  AcctName: '',
  AcctEngName: '',
  AcctAddr: '',
  Nation: '',
  TaxType: '',
  TaxCode: '',
  PhoneNum: '',
  Email: '',
  Website: '',
  RegDate: '',
  Notes: '',
  License: '',
  AcctId: '',
  FileName: '', // 添加 FileName 字段
  BankAccounts: [] // 关联的银行账户信息
});

// 会计实体信息表单验证规则
const acctRules = {
  AcctCode: [{ required: true, message: '请输入会计实体编码', trigger: 'blur' }],
  AcctAbbr: [{ required: true, message: '请输入会计实体缩写', trigger: 'blur' }],
  EtyAbbr: [{ required: true, message: '请输入实体简称', trigger: 'blur' }]
};


// 表格数据（初始值为空数组）
const acctData = ref([]); // 会计实体信息

const headerTitle = computed(() => {
  return '会计实体信息';
});

const addButtonText = computed(() => {
  return '添加会计实体信息';
});

// 添加按钮点击事件
const handleAdd = () => {
  showAcctDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  acctForm.value = { ...row }; // 将当前行的数据赋值给 acctForm
  acctForm.value.AcctBanksDisplay = row.AcctBanks && row.AcctBanks.length > 0
    ? row.AcctBanks.map(bank => bank.AccNum).join(', ')
    : '无';
  showAcctDialog.value = true; // 打开会计实体信息对话框
};
</script>

<style src="../assets/styles/Bottom.css"></style>


<style scoped>
.top-bar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  background-color: #f0f0f0;
  padding: 10px;
  display: flex;
  justify-content: center;
}
</style>
