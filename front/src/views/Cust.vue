<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-4'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>
        <el-main>
          <!-- 联系人信息表格 -->
          <el-table :data="paginatedCustData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="Name" label="联系人姓名" width="220%"></el-table-column>
            <el-table-column prop="Gender" label="性别" width="220%"></el-table-column>
            <el-table-column prop="Nation" label="国家" width="220%"></el-table-column>
            <el-table-column prop="Addr" label="地址" width="220%"></el-table-column>
            <el-table-column prop="Email" label="邮箱" width="220%"></el-table-column>
            <el-table-column prop="PhoneNum" label="联系电话" width="220%"></el-table-column>
            <el-table-column prop="QQ" label="QQ" width="220%"></el-table-column>
            <el-table-column prop="Wechat" label="微信" width="220%"></el-table-column>
            <el-table-column label="所属商户" width="320%">
              <template #default="scope">
                <span v-if="scope.row.Merchant.Merc">{{ scope.row.Merchant.Merc }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Post" label="职位" width="220%"></el-table-column>
            <el-table-column prop="Notes" label="备注" width="220%"></el-table-column>
            <el-table-column prop="FileName" label="文件名" width="220%"></el-table-column>

            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row.ID)" type="text" size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="custData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加联系人信息的对话框 -->
    <el-dialog v-model="showCustDialog" title="联系人信息" width="80%" @close="resetCustForm">
      <el-form :model="custForm" label-width="150px" :rules="custRules" ref="custFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="联系人姓名" prop="Name" :required="true">
              <el-input v-model="custForm.Name" placeholder="请输入联系人姓名"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender" :required="true">
              <el-select v-model="custForm.Gender" placeholder="请选择性别">
                <el-option label="男" value="男"></el-option>
                <el-option label="女" value="女"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">

            <el-form-item label="国家" prop="Nation">
              <el-select v-model="custForm.Nation" placeholder="请选择国家">
                <el-option v-for="nation in nationData" :key="nation.NationID" :label="nation.Nation"
                  :value="nation.Nation"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Addr">
              <el-input v-model="custForm.Addr" placeholder="请输入地址"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="custForm.Email" placeholder="请输入邮箱"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="custForm.PhoneNum" placeholder="请输入联系电话"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="QQ" prop="QQ">
              <el-input v-model="custForm.QQ" placeholder="请输入QQ"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="微信" prop="Wechat">
              <el-input v-model="custForm.Wechat" placeholder="请输入微信"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="所属商户" prop="MerchantID">
              <el-select v-model="custForm.MerchantID" @change="onMercChange" placeholder="请选择商户信息">
                <el-option v-for="merc in merchantData" :key="merc.ID" :label="`${merc.Merc} (${merc.MercCode})`"
                  :value="merc.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="职位" prop="Post">
              <el-input v-model="custForm.Post" placeholder="请输入职位"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="custForm.Notes" type="textarea" placeholder="请输入备注"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!custForm.FileID" ref="custUploadRef" action="" :limit="1"
                :on-change="handleCustFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(custForm.FileID, custForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="custForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitCustForm">保存</el-button>
            <el-button @click="showCustDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看联系人信息的对话框 -->
    <el-dialog v-model="showViewCustDialog" title="联系人信息" width="80%" @close="resetCustForm">
      <el-form :model="custForm" label-width="150px" :rules="custRules" ref="custFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="联系人姓名" prop="Name">
              <el-input v-model="custForm.Name" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="性别" prop="Gender">
              <el-input v-model="custForm.Gender" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="国家" prop="Nation">
              <el-input v-model="custForm.Nation" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址" prop="Addr">
              <el-input v-model="custForm.Addr" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="邮箱" prop="Email">
              <el-input v-model="custForm.Email" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系电话" prop="PhoneNum">
              <el-input v-model="custForm.PhoneNum" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="QQ" prop="QQ">
              <el-input v-model="custForm.QQ" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="微信" prop="Wechat">
              <el-input v-model="custForm.Wechat" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">

          <el-col :span="12">
            <el-form-item label="所属商户" prop="MerchantID">
              <el-select v-model="custForm.MerchantID" :disabled="true" @change="onMercChange" placeholder="请选择商户信息">
                <el-option v-for="merc in merchantData" :key="merc.ID" :label="`${merc.Merc} (${merc.MercCode})`"
                  :value="merc.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
          </el-col>
          <el-col :span="12">
            <el-form-item label="职位" prop="Post">
              <el-input v-model="custForm.Post" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="备注" prop="Notes">
              <el-input v-model="custForm.Notes" type="textarea" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <!-- 上传文件或下载文件 -->
        <el-row :gutter="20">
          <el-col :span="24">
            <el-form-item label="文件">
              <el-upload v-if="!custForm.FileID" ref="custUploadRef" action="" :limit="1"
                :on-change="handleCustFileChange" :auto-upload="false" :show-file-list="true">
                <el-button type="primary">选择文件</el-button>
              </el-upload>
              <el-button v-else type="success"
                @click="downloadFile(custForm.FileID, custForm.FileName)">下载文件</el-button>
            </el-form-item>
            <el-form-item label="文件名" prop="FileName">
              <el-input v-model="custForm.FileName" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewCustDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';
import HeaderComponent from '@/components/HeaderComponent.vue';

const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const showCustDialog = ref(false);
const showViewCustDialog = ref(false);
const custForm = ref({
  Name: '',
  Gender: '',
  Nation: '',
  Addr: '',
  Email: '',
  PhoneNum: '',
  QQ: '',
  Wechat: '',
  MerchantID: '',
  Merc: '',
  Post: '',
  Notes: '',
  FileName: '',
  FileID: ''
});

const custFormRef = ref(null);

// 自定义验证函数
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 客户表单验证规则
const custRules = {
  Name: [
    { required: true, validator: validateNotEmpty, message: '请输入联系人姓名', trigger: 'blur' }
  ],
  Gender: [
    { required: true, validator: validateNotEmpty, message: '请选择性别', trigger: 'blur' }
  ],
  Nation: [
    { required: true, validator: validateNotEmpty, message: '请输入国家', trigger: 'blur' }
  ]
};

const custData = ref([]);
const merchantData = ref([]);
const custUploadRef = ref(null);
const custFile = ref(null);

// 分页计算
const paginatedCustData = computed(() => {
  let filteredData = custData.value;
  if (searchQuery.value) {
    console.log(custData.value);
    //console.log(onlyID.value);
    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.Name.includes(searchQuery.value) ||
          item.Gender.includes(searchQuery.value) ||
          item.Nation.includes(searchQuery.value) ||
          item.Addr.includes(searchQuery.value) ||
          item.Email.includes(searchQuery.value) ||
          item.PhoneNum.includes(searchQuery.value) ||
          item.QQ.includes(searchQuery.value) ||
          item.Wechat.includes(searchQuery.value) ||
          item.Merc.includes(searchQuery.value) ||
          item.Post.includes(searchQuery.value) ||
          item.Notes.includes(searchQuery.value)
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.Name === searchQuery.value ||
          item.Gender === searchQuery.value ||
          item.Nation === searchQuery.value ||
          item.Addr === searchQuery.value ||
          item.Email === searchQuery.value ||
          item.PhoneNum === searchQuery.value ||
          item.QQ === searchQuery.value ||
          item.Wechat === searchQuery.value ||
          item.Merc === searchQuery.value ||
          item.Post === searchQuery.value ||
          item.Notes === searchQuery.value
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value
        );
      }
    }
  }
  const start = (currentPage.value - 1) * pageSize;
  const end = start + pageSize;
  return filteredData.slice(start, end);
});

const toggleMatchMode = () => {
  console.log("check onlyID", isExactMatch.value)
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {

console.log("check match", onlyID.value)
onlyID.value = !onlyID.value;
};



// 文件选择事件
const handleCustFileChange = (uploadFile) => {
  custFile.value = uploadFile.raw;
};

// 下载文件
const downloadFile = async (fileID, fileName) => {
  try {
    const response = await axios.post(
      '/file',
      { FileID: fileID },
      { responseType: 'blob' }
    );
    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement('a');
    link.href = url;
    link.setAttribute('download', fileName || `file_${fileID}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('文件下载失败:', error);
    ElMessage.error('文件下载失败，请稍后重试');
  }
};

// 查看按钮逻辑
const handleView = (index, row) => {
  custForm.value = { ...row };
  showViewCustDialog.value = true;
};

// 重置表单
const resetCustForm = () => {
  custForm.value = {
    Name: '',
    Gender: '',
    Nation: '',
    Addr: '',
    Email: '',
    PhoneNum: '',
    QQ: '',
    Wechat: '',
    MerchantID: '',
    Merc: '',
    Post: '',
    Notes: '',
    FileName: '',
    FileID: '',
    ID: ''
  };
  custFile.value = null;
  if (custUploadRef.value) {
    custUploadRef.value.clearFiles();
  }
};
const isExactMatch = ref(true);
const onlyID = ref(true);

// 提交表单
const submitCustForm = async () => {
  try {

    const isValid = await custFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    const formData = new FormData();
    Object.keys(custForm.value).forEach((key) => {

      if (key != 'Merchant') {
        formData.append(key, custForm.value[key]);
      }
    });
    if (custFile.value) {
      formData.append('file', custFile.value);
    }
    formData.append("MercCode", "ss")
    formData.append("ShortMerc", "wss")
    const response = await axios.post('/save/cust', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    if (response.status === 200) {
      ElMessage.success('联系人信息保存成功');
      showCustDialog.value = false;
      fetchCustData();
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 删除按钮逻辑
const handleDelete = (index, CustID) => {
  ElMessageBox.confirm('确定要删除该联系人信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    const params = new URLSearchParams();
    params.append('ID', CustID); // 添加表单字段
    axios.post('/delete/cust', params, {

      headers: {
        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    })
      .then(response => {
        if (response.status === 200) {

          ElMessage.success('删除成功');
          fetchCustData(); // 重新获取联系人信息数据
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

// 分页切换
const handlePageChange = (page) => {
  currentPage.value = page;
};

// 添加按钮点击事件
const handleAdd = () => {
  showCustDialog.value = true;
};

// 编辑按钮逻辑
const handleEdit = (index, row) => {
  custForm.value = { ...row };
  showCustDialog.value = true;
};

// 获取联系人数据
const fetchCustData = async () => {
  try {
    const response = await axios.get('/find/cust');
    custData.value = response.data.Cust;
    console.log(custData.value)
  } catch (error) {
    console.error('获取联系人信息失败:', error);
    ElMessage.error('获取联系人信息失败，请稍后重试');
  }
};

// 获取商户数据
const fetchMerchantData = async () => {
  try {
    const response = await axios.get('/find/merchant');
    merchantData.value = response.data.Merchant;
  } catch (error) {
    console.error('获取商户信息失败:', error);
    ElMessage.error('获取商户信息失败，请稍后重试');
  }
};

// 监听商户选择事件
const onMercChange = (value) => {
  const selectedMerc = merchantData.value.find(merc => merc.MerchantID === value);
  if (selectedMerc) {
    custForm.value.Merc = selectedMerc.Merc;
  }
};

const nationData = ref([]);

const fetchNationData = async () => {
  try {
    const response = await axios.get('/find/nation');
    nationData.value = response.data.Nation;
  } catch (error) {
    console.error('获取国家信息失败:', error);
    ElMessage.error('获取国家信息失败，请稍后重试');
  }
};
onMounted(() => {
  fetchNationData();
  fetchCustData();
  fetchMerchantData();
});

const headerTitle = computed(() => {
  return '联系人信息';
});

const addButtonText = computed(() => {
  return '添加联系人信息';
});
</script>

<style src="../assets/styles/Bottom.css"></style>
