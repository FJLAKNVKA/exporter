<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-14'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>
        <el-main>
          <!-- 费用信息表格 -->
          <el-table :data="paginatedCostData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column prop="ExpType" label="费用类型" width="220%"></el-table-column>
            <el-table-column prop="Rates" label="收费标准" width="220%"></el-table-column>
            <el-table-column prop="UnitPrice" label="单价" width="220%"></el-table-column>
            <el-table-column prop="Number" label="数量" width="220%"></el-table-column>
            <el-table-column prop="Amount" label="金额" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="CreatedAt" label="创建时间" width="240%"></el-table-column>
            <el-table-column prop="UpdatedAt" label="更新时间" width="240%"></el-table-column>
            <el-table-column label="操作" fixed="right" width="200">
              <template #default="scope">
                <el-row type="flex" justify="space-between">
                  <el-button @click="handleView(scope.$index, scope.row)" type="text" size="small">查看</el-button>
                  <el-button @click="handleEdit(scope.$index, scope.row)" type="text" size="small">编辑</el-button>
                  <el-button @click="handleDelete(scope.$index, scope.row)" type="text" size="small">删除</el-button>
                </el-row>
              </template>
            </el-table-column>
          </el-table>

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="costData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加费用信息的对话框 -->
    <el-dialog v-model="showCostDialog" title="费用信息" width="80%" @close="resetCostForm">
      <el-form :model="costForm" label-width="150px" :rules="costRules" ref="costFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="费用类型" prop="ExpType">
              <el-select v-model="costForm.ExpType" placeholder="请选择费用类型">
                <el-option v-for="item in expTypeData" :key="item.ExpType" :label="item.ExpType"
                  :value="item.ExpType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收费标准" prop="Rates">
              <el-select v-model="costForm.Rates" placeholder="请选择收费标准">
                <el-option v-for="item in ratesData" :key="item.Rates" :label="item.Rates"
                  :value="item.Rates"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="costForm.UnitPrice" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数量" prop="Number">
              <el-input v-model="costForm.Number" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="costForm.Amount" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-select v-model="costForm.Currency" placeholder="请选择币种">
                <el-option v-for="item in currencyData" :key="item.Currency" :label="item.Currency"
                  :value="item.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitCostForm">保存</el-button>
            <el-button @click="showCostDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看费用信息的对话框 -->
    <el-dialog v-model="showViewCostDialog" title="费用信息" width="80%" @close="resetCostForm">
      <el-form :model="costForm" label-width="150px" :rules="costRules" ref="costFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="费用类型" prop="ExpType">
              <el-input v-model="costForm.ExpType" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="收费标准" prop="Rates">
              <el-input v-model="costForm.Rates" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="costForm.UnitPrice" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="数量" prop="Number">
              <el-input v-model="costForm.Number" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="costForm.Amount" type="number" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-input v-model="costForm.Currency" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="创建时间" prop="CreatedAt">
              <el-input v-model="costForm.CreatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="更新时间" prop="UpdatedAt">
              <el-input v-model="costForm.UpdatedAt" :readonly="true"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showViewCostDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';
import HeaderComponent from '@/components/HeaderComponent.vue';
const route = useRoute();
const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const costData = ref([]);
const showCostDialog = ref(false);
const showViewCostDialog = ref(false);
const costForm = ref({
  ExpType: '',
  Rates: '',
  UnitPrice: 0,
  Number: 0,
  Amount: 0,
  Currency: '',
  CostID: '',
});
const costFormRef = ref(null);

// 字典数据
const expTypeData = ref([]); // 费用类型
const ratesData = ref([]); // 收费标准
const currencyData = ref([]); // 币种

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedCostData = computed(() => {
  let filteredData = costData.value;

  if (searchQuery.value) {
    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        console.log("打印了");
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          item.ExpType.includes(searchQuery.value) ||
          item.Rates.includes(searchQuery.value) ||
          item.UnitPrice.toString().includes(searchQuery.value) ||
          item.Number.toString().includes(searchQuery.value) ||
          item.Amount.toString().includes(searchQuery.value) ||
          item.Currency.includes(searchQuery.value) ||
          item.CreatedAt.includes(searchQuery.value) ||
          item.UpdatedAt.includes(searchQuery.value)
        );
      } else {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value)
        );
      }
    } else {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString() === searchQuery.value ||
          item.ExpType === searchQuery.value ||
          item.Rates === searchQuery.value ||
          item.UnitPrice.toString() === searchQuery.value ||
          item.Number.toString() === searchQuery.value ||
          item.Amount.toString() === searchQuery.value ||
          item.Currency === searchQuery.value ||
          item.CreatedAt === searchQuery.value ||
          item.UpdatedAt === searchQuery.value
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

const isExactMatch = ref(true);
const onlyID = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
  console.log("isExactMatch", isExactMatch.value)
};

const toggleIDMode = () => {
  onlyID.value = !onlyID.value;
};
const handleAdd = () => {
  showCostDialog.value = true;
};

const handleEdit = (index, row) => {
  costForm.value = { ...row };
  showCostDialog.value = true;
};

const handleView = (index, row) => {
  costForm.value = { ...row };
  showViewCostDialog.value = true;
};

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定要删除该费用信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    costForm.value = { ...row };
    axios.post('/delete/costInfo', costForm.value)
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchCostData();
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

const resetCostForm = () => {
  costForm.value = {
    ExpType: '',
    Rates: '',
    UnitPrice: 0,
    Number: 0,
    Amount: 0,
    Currency: '',
    CostID: '',
  };
};

const submitCostForm = async () => {
  try {

    const isValid = await costFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    costForm.value.Amount = parseInt(costForm.value.Amount, 10);
    costForm.value.Number = parseInt(costForm.value.Number, 10);
    costForm.value.UnitPrice = parseInt(costForm.value.UnitPrice, 10);
    const response = await axios.post('/save/costInfo', costForm.value);
    if (response.status === 200) {
      ElMessage.success('费用信息保存成功');
      showCostDialog.value = fetchCostData();
      showCostDialog.value = false
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存费用信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 获取字典数据
const fetchDictionaryData = async () => {
  try {
    const [
      expTypeResponse,
      ratesResponse,
      currencyResponse,
    ] = await Promise.all([
      axios.get('/find/expType'), // 费用类型
      axios.get('/find/rates'), // 收费标准
      axios.get('/find/currency'), // 币种
    ]);

    // 提取字典数据
    expTypeData.value = expTypeResponse.data.ExpType || [];
    ratesData.value = ratesResponse.data.Rates || [];
    currencyData.value = currencyResponse.data.Currency || [];
  } catch (error) {
    console.error('获取字典数据失败:', error);
    ElMessage.error('获取字典数据失败，请稍后重试');
  }
};

// 获取费用数据
const fetchCostData = async () => {
  try {
    const response = await axios.get('/find/costInfo');
    costData.value = response.data.CostInfo || [];
    console.log(response.data.CostInfo)
  } catch (error) {
    console.error('获取费用信息失败:', error);
    ElMessage.error('获取费用信息失败，请稍后重试');
  }
};

onMounted(() => {
  fetchCostData();
  fetchDictionaryData();

  searchQuery.value = route.query.searchQuery || '';
});

const headerTitle = computed(() => '费用信息');
const addButtonText = computed(() => '添加费用信息');

// 自定义验证函数：检查非空
const validateNotEmpty = (rule, value, callback) => {
  if (value === '' || value === null || value === undefined) {
    callback(new Error(rule.message));  // 使用 rule.message 作为错误消息
  } else {
    callback();  // 验证通过
  }
};

// 自定义验证函数：检查大于 0
const validateGreaterThanZero = (rule, value, callback) => {
  if (value <= 0) {
    callback(new Error(rule.message || '该字段必须大于 0'));  // 提供字段的错误消息
  } else {
    callback();  // 验证通过
  }
};

const costRules = {
  ExpType: [{ required: true, message: '请选择费用类型', trigger: 'blur' }],
  Rates: [{ required: true, message: '请选择收费标准', trigger: 'blur' }],
  UnitPrice: [
    { required: true, validator: validateNotEmpty, message: '请输入单价', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '单价必须大于 0', trigger: 'blur' }  // 新增验证：单价大于 0
  ],
  Number: [
    { required: true, validator: validateNotEmpty, message: '请输入数量', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '数量必须大于 0', trigger: 'blur' }  // 新增验证：数量大于 0
  ],
  Amount: [
    { required: true, validator: validateNotEmpty, message: '请输入金额', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '金额必须大于 0', trigger: 'blur' }  // 新增验证：金额大于 0
  ],
  Currency: [{ required: true, message: '请选择币种', trigger: 'blur' }]
};
</script>

<style src="../assets/styles/Bottom.css"></style>
