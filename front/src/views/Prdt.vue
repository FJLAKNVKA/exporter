<template>
  <div id="app">
    <el-container>
      <!-- 侧边栏 -->
      <el-aside width="200px">
        <div style="height: 100vh; overflow-y: auto;">
          <SideMenu :default-active="'1-12'" />
        </div>
      </el-aside>

      <!-- 主体内容 -->
      <el-container>
        <!-- 使用 HeaderComponent -->

        <HeaderComponent :header-title="headerTitle" :add-button-text="addButtonText" v-model:search-query="searchQuery"
          @toggle-match-mode="toggleMatchMode" @toggle-id-mode="toggleIDMode" @add="handleAdd" />
        <el-header height="1px">
        </el-header>

        <el-main>
          <!-- 产品信息表格 -->
          <el-table :data="paginatedPrdtData" style="width: 100%" max-height="450">
            <el-table-column prop="ID" label="ID" width="100%"></el-table-column>
            <el-table-column label="产品" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Cat.CatEngName">{{ scope.row.Cat.CatEngName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>

            <el-table-column label="品牌" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Brand.BrandEngName">{{ scope.row.Brand.BrandEngName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Factory" label="生产工厂" width="220%"></el-table-column>
            <el-table-column prop="Currency" label="币种" width="220%"></el-table-column>
            <el-table-column prop="UnitPrice" label="单价" width="220%"></el-table-column>
            <el-table-column prop="Unit" label="单位" width="220%"></el-table-column>
            <el-table-column prop="Amount" label="金额" width="220%"></el-table-column>
            <el-table-column prop="ItemNum" label="件数" width="220%"></el-table-column>

            <el-table-column label="包装规格" width="220%">
              <template #default="scope">
                <span v-if="scope.row.PackSpec.SpecName">{{ scope.row.PackSpec.SpecName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
            <el-table-column prop="Weight" label="重量" width="220%"></el-table-column>
            <el-table-column prop="WeightUnit" label="重量单位" width="220%"></el-table-column>
            <el-table-column prop="TradeTerm" label="贸易条款" width="220%"></el-table-column>
            <el-table-column label="交货地点" width="220%">
              <template #default="scope">
                <span v-if="scope.row.Spot.InvLocName">{{ scope.row.Spot.InvLocName }}</span>
                <span v-else>无</span>
              </template>
            </el-table-column>
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

          <el-pagination v-model:current-page="currentPage" :page-size="pageSize" :total="prdtData.length"
            layout="prev, pager, next" @current-change="handlePageChange" />
        </el-main>
      </el-container>
    </el-container>

    <!-- 添加产品信息的对话框 -->
    <el-dialog v-model="showPrdtDialog" title="产品信息" width="80%" @close="resetPrdtForm">
      <el-form :model="prdtForm" label-width="150px" :rules="prdtRules" ref="prdtFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="CatID">
              <el-select v-model="prdtForm.CatID" placeholder="请选择产品">
                <el-option v-for="item in prdtTypeData" :key="item.ID" :label="item.CatEngName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="BrandID">
              <el-select v-model="prdtForm.BrandID" placeholder="请选择品牌">
                <el-option v-for="item in brandTypeData" :key="item.ID" :label="item.BrandEngName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="Factory">
              <el-select v-model="prdtForm.Factory" placeholder="请选择生产工厂">
                <el-option v-for="item in suprTypeData" :key="item.SuprType" :label="item.SuprType"
                  :value="item.SuprType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-select v-model="prdtForm.Currency" placeholder="请选择币种">
                <el-option v-for="item in currencyData" :key="item.Currency" :label="item.Currency"
                  :value="item.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="prdtForm.UnitPrice" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-select v-model="prdtForm.Unit" placeholder="请选择单位">
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="prdtForm.Amount" type="number"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="prdtForm.ItemNum" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="prdtForm.PackSpecID" placeholder="请选择包装规格">
                <el-option v-for="item in packTypeData" :key="item.ID" :label="item.SpecName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="重量" prop="Weight">
              <el-input v-model="prdtForm.Weight" type="number"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="重量单位" prop="WeightUnit">
              <el-select v-model="prdtForm.WeightUnit" placeholder="请选择重量单位">
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="贸易条款" prop="TradeTerm">
              <el-select v-model="prdtForm.TradeTerm" placeholder="请选择贸易条款">
                <el-option v-for="item in tradeTermData" :key="item.TradeTerm" :label="item.TradeTerm"
                  :value="item.TradeTerm"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="交货地点" prop="SpotID">
              <el-select v-model="prdtForm.SpotID" placeholder="请选择交货地点">
                <el-option v-for="item in portData" :key="item.ID" :label="item.InvLocName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button type="primary" @click="submitPrdtForm">保存</el-button>
            <el-button @click="showPrdtDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>

    <!-- 查看产品信息的对话框 -->

    <el-dialog v-model="showViewPrdtDialog" title="产品信息" width="80%" @close="resetPrdtForm">
      <el-form :model="prdtForm" label-width="150px" ref="prdtFormRef">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="产品" prop="CatID">
              <el-select v-model="prdtForm.CatID" placeholder="请选择产品" disabled>
                <el-option v-for="item in prdtTypeData" :key="item.ID" :label="item.CatEngName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="品牌" prop="BrandID">
              <el-select v-model="prdtForm.BrandID" placeholder="请选择品牌" disabled>
                <el-option v-for="item in brandTypeData" :key="item.ID" :label="item.BrandEngName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="生产工厂" prop="Factory">
              <el-select v-model="prdtForm.Factory" placeholder="请选择生产工厂" disabled>
                <el-option v-for="item in suprTypeData" :key="item.SuprType" :label="item.SuprType"
                  :value="item.SuprType"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="币种" prop="Currency">
              <el-select v-model="prdtForm.Currency" placeholder="请选择币种" disabled>
                <el-option v-for="item in currencyData" :key="item.Currency" :label="item.Currency"
                  :value="item.Currency"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="单价" prop="UnitPrice">
              <el-input v-model="prdtForm.UnitPrice" type="number" disabled></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="单位" prop="Unit">
              <el-select v-model="prdtForm.Unit" placeholder="请选择单位" disabled>
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="金额" prop="Amount">
              <el-input v-model="prdtForm.Amount" type="number" disabled></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="件数" prop="ItemNum">
              <el-input v-model="prdtForm.ItemNum" type="number" disabled></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="包装规格" prop="PackSpecID">
              <el-select v-model="prdtForm.PackSpecID" placeholder="请选择包装规格" disabled>
                <el-option v-for="item in packTypeData" :key="item.ID" :label="item.SpecName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="重量" prop="Weight">
              <el-input v-model="prdtForm.Weight" type="number" disabled></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="重量单位" prop="WeightUnit">
              <el-select v-model="prdtForm.WeightUnit" placeholder="请选择重量单位" disabled>
                <el-option v-for="item in unitMeasData" :key="item.UnitMeas" :label="item.UnitMeas"
                  :value="item.UnitMeas"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="贸易条款" prop="TradeTerm">
              <el-select v-model="prdtForm.TradeTerm" placeholder="请选择贸易条款" disabled>
                <el-option v-for="item in tradeTermData" :key="item.TradeTerm" :label="item.TradeTerm"
                  :value="item.TradeTerm"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="交货地点" prop="SpotID">
              <el-select v-model="prdtForm.SpotID" placeholder="请选择交货地点" disabled>
                <el-option v-for="item in portData" :key="item.ID" :label="item.InvLocName"
                  :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row :gutter="20">
          <el-col :span="24" style="text-align: right;">
            <el-button @click="showPrdtDialog = false">关闭</el-button>
          </el-col>
        </el-row>
      </el-form>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { ElMessageBox, ElMessage } from 'element-plus';
import axios from 'axios';
import SideMenu from '@/components/SideMenu.vue';

import HeaderComponent from '@/components/HeaderComponent.vue';
import { useRoute } from 'vue-router';
const route = useRoute();
// 匹配模式（默认是模糊匹配）
const isExactMatch = ref(true);
const onlyID = ref(true);
const toggleMatchMode = () => {
  isExactMatch.value = !isExactMatch.value;
};

const toggleIDMode = () => {
  onlyID.value = !onlyID.value;
};
const searchQuery = ref('');
const currentPage = ref(1);
const pageSize = 8;
const prdtData = ref([]);
const showPrdtDialog = ref(false);
const showViewPrdtDialog = ref(false);
const prdtForm = ref({
  Factory: '',
  Currency: '',
  UnitPrice: '',
  Unit: '',
  Amount: '',
  ItemNum: '',
  Weight: '',
  WeightUnit: '',
  TradeTerm: '',
  SpotID: '',
  BrandID: '',
  CatID: '',
  PackSpecID: '',
});
const prdtFormRef = ref(null);

// 字典数据
const prdtTypeData = ref([]); // 产品类型
const brandTypeData = ref([]); // 品牌类型
const suprTypeData = ref([]); // 供应商类型
const currencyData = ref([]); // 币种
const unitMeasData = ref([]); // 单位
const packTypeData = ref([]); // 包装规格
const tradeTermData = ref([]); // 贸易条款
const portData = ref([]); // 交货地点

const handlePageChange = (page) => {
  currentPage.value = page;
};

const paginatedPrdtData = computed(() => {

  let filteredData = prdtData.value; // 假设 prdtData 是你的产品数据

  if (searchQuery.value) {
    console.log(isExactMatch.value);
    console.log(onlyID.value);

    if (isExactMatch.value === false) {
      if (onlyID.value === false) {
        filteredData = filteredData.filter(item =>
          item.ID.toString().includes(searchQuery.value) ||
          (item.Cat && item.Cat.CatEngName && item.Cat.CatEngName.includes(searchQuery.value)) ||
          (item.Brand && item.Brand.BrandEngName && item.Brand.BrandEngName.includes(searchQuery.value)) ||
          item.Factory.includes(searchQuery.value) ||
          item.Currency.includes(searchQuery.value) ||
          item.UnitPrice.toString().includes(searchQuery.value) ||
          item.Unit.includes(searchQuery.value) ||
          item.Amount.toString().includes(searchQuery.value) ||
          item.ItemNum.toString().includes(searchQuery.value) ||
          (item.PackSpec && item.PackSpec.SpecName && item.PackSpec.SpecName.includes(searchQuery.value)) ||
          item.Weight.toString().includes(searchQuery.value) ||
          item.WeightUnit.includes(searchQuery.value) ||
          item.TradeTerm.includes(searchQuery.value) ||
          (item.Spot && item.Spot.InvLocName && item.Spot.InvLocName.includes(searchQuery.value)) ||
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
          (item.Cat && item.Cat.CatEngName && item.Cat.CatEngName === searchQuery.value) ||
          (item.Brand && item.Brand.BrandEngName && item.Brand.BrandEngName === searchQuery.value) ||
          item.Factory === searchQuery.value ||
          item.Currency === searchQuery.value ||
          item.UnitPrice.toString() === searchQuery.value ||
          item.Unit === searchQuery.value ||
          item.Amount.toString() === searchQuery.value ||
          item.ItemNum.toString() === searchQuery.value ||
          (item.PackSpec && item.PackSpec.SpecName && item.PackSpec.SpecName === searchQuery.value) ||
          item.Weight.toString() === searchQuery.value ||
          item.WeightUnit === searchQuery.value ||
          item.TradeTerm === searchQuery.value ||
          (item.Spot && item.Spot.InvLocName && item.Spot.InvLocName === searchQuery.value) ||
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

const handleAdd = () => {
  showPrdtDialog.value = true;
};

const handleEdit = (index, row) => {
  prdtForm.value = { ...row };
  showPrdtDialog.value = true;
};

const handleView = (index, row) => {
  prdtForm.value = { ...row };
  showViewPrdtDialog.value = true;
};

const handleDelete = (index, row) => {
  ElMessageBox.confirm('确定要删除该产品信息吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {

    prdtForm.value = { ...row };
    axios.post('/delete/prdtInfo', prdtForm.value)
      .then(response => {
        if (response.status === 200) {
          ElMessage.success('删除成功');
          fetchPrdtData();
          resetPrdtForm();
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

const resetPrdtForm = () => {
  prdtForm.value = {
    CatEngName: '',
    BrandEngName: '',
    Factory: '',
    Currency: '',
    UnitPrice: '',
    Unit: '',
    Amount: '',
    ItemNum: '',
    PackSpec: '',
    Weight: '',
    WeightUnit: '',
    TradeTerm: '',
    DeliveryLoc: '',
    PrdtID: '',
  };
};

const submitPrdtForm = async () => {
  try {

    const isValid = await prdtFormRef.value.validate();
    if (!isValid) {
      ElMessage.error('请填写所有必填字段');
      console.log('验证不通过');
      return; // 如果验证不通过，阻止提交
    }

    console.log(prdtForm.value)

    // prdtForm.value.Weight = parseInt(prdtForm.value.Weight, 10);
    // prdtForm.value.UnitPrice = parseInt(prdtForm.value.UnitPrice, 10);
    // prdtForm.value.Amount = parseInt(prdtForm.value.Amount, 10);
    // prdtForm.value.ItemNum = parseInt(prdtForm.value.ItemNum, 10);
    // prdtForm.value.SpotID = parseInt(prdtForm.value.SpotID, 10);
    // prdtForm.value.PackSpecID = parseInt(prdtForm.value.PackSpecID, 10);
    // prdtForm.value.BrandID = parseInt(prdtForm.value.BrandID, 10);
    // prdtForm.value.CatID = parseInt(prdtForm.value.CatID, 10);
    const formData = new FormData(); // 创建 FormData 对象

    Object.keys(prdtForm.value).forEach((key) => {
      formData.append(key, prdtForm.value[key]);
    });

    console.log(formData)

    const response = await axios.post('/save/prdtInfo', formData, {
      headers: {

        'Content-Type': 'application/x-www-form-urlencoded', // 设置请求头为表单格式
      },
    });
    if (response.status === 200) {
      ElMessage.success('产品信息保存成功');
      await fetchPrdtData(); // 等待数据获取完成
      showPrdtDialog.value = false; // 手动关闭对话框
    } else {
      ElMessage.error(response.data.RetMessage || '保存失败');
    }
  } catch (error) {
    console.error('保存产品信息失败:', error);
    ElMessage.error(error.response.data.RetMessage);
  }
};

// 获取字典数据
const fetchDictionaryData = async () => {
  try {
    const [
      prdtTypeResponse,
      brandTypeResponse,
      suprTypeResponse,
      currencyResponse,
      unitMeasResponse,
      packTypeResponse,
      tradeTermResponse,
      portResponse,
    ] = await Promise.all([
      axios.get('/find/cat'), // 产品类型
      axios.get('/find/brand'), // 品牌类型
      axios.get('/find/suprType'), // 供应商类型
      axios.get('/find/currency'), // 币种
      axios.get('/find/unitMeas'), // 单位
      axios.get('/find/packSpec'), // 包装规格
      axios.get('/find/tradeTerm'), // 贸易条款
      axios.get('/find/spot'), // 交货地点
    ]);

    // 提取字典数据
    prdtTypeData.value = prdtTypeResponse.data.Cat || [];
    brandTypeData.value = brandTypeResponse.data.Brand || [];
    suprTypeData.value = suprTypeResponse.data.SuprType || [];
    currencyData.value = currencyResponse.data.Currency || [];
    unitMeasData.value = unitMeasResponse.data.UnitMeas || [];
    packTypeData.value = packTypeResponse.data.PackSpec || [];
    tradeTermData.value = tradeTermResponse.data.TradeTerm || [];
    portData.value = portResponse.data.Spot || [];
  } catch (error) {
    console.error('获取字典数据失败:', error);
    ElMessage.error('获取字典数据失败，请稍后重试');
  }
};

// 获取产品数据
const fetchPrdtData = async () => {
  try {
    const response = await axios.get('/find/prdtInfo');
    prdtData.value = response.data.PrdtInfo || [];
    console.log(response.data.PrdtInfo)
  } catch (error) {
    console.error('获取产品信息失败:', error);
    ElMessage.error('获取产品信息失败，请稍后重试');
  }
};

onMounted(() => {

  console.log('Search Query:', searchQuery.value); // 应该输出 "32"
  fetchPrdtData();
  fetchDictionaryData();
  searchQuery.value = route.query.searchQuery || '';
  console.log("query ", route.query.searchQuery)
});
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
  if (value < 0) {
    callback(new Error(rule.message || '该字段必须大于 0'));  // 提供字段的错误消息
  } else {
    callback();  // 验证通过
  }
};


const headerTitle = computed(() => {
  return '产品明细';
});

const addButtonText = computed(() => {
  return '添加产品明细';
});

// 产品表单验证规则
const prdtRules = {
  CatID: [{ required: true, message: '请选择产品', trigger: 'blur' }],
  BrandID: [{ required: true, message: '请选择品牌', trigger: 'blur' }],
  // Factory: [{ required: true, message: '请选择生产工厂', trigger: 'blur' }],
  // Currency: [{ required: true, message: '请选择币种', trigger: 'blur' }],
  UnitPrice: [
    // { required: true, validator: validateNotEmpty, message: '请输入单价', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '单价必须大于 0', trigger: 'blur' }  // 新增验证：大于 0
  ],
  // Unit: [{ required: true, message: '请选择单位', trigger: 'blur' }],
  Amount: [
    // { required: true, validator: validateNotEmpty, message: '请输入金额', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '金额必须大于 0', trigger: 'blur' }  // 新增验证：大于 0
  ],
  ItemNum: [
    // { required: true, validator: validateNotEmpty, message: '请输入件数', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '件数必须大于 0', trigger: 'blur' }  // 新增验证：大于 0
  ],
  PackSpecID: [{ required: true, message: '请选择包装规格', trigger: 'blur' }],
  Weight: [
    // { required: true, validator: validateNotEmpty, message: '请输入重量', trigger: 'blur' },
    { validator: validateGreaterThanZero, message: '重量必须大于 0', trigger: 'blur' }  // 新增验证：大于 0
  ],
  // WeightUnit: [{ required: true, message: '请选择重量单位', trigger: 'blur' }],
  // TradeTerm: [{ required: true, message: '请选择贸易条款', trigger: 'blur' }],
  SpotID: [{ required: true, message: '请选择交货地点', trigger: 'blur' }]
};

</script>

<style src="../assets/styles/Bottom.css"></style>
