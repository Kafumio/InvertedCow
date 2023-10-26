<template>
  <div class="account-setting-container">
    <el-card class="account-base-card">
      <template #header> 账号信息管理 </template>
      <div class="message">
        <img v-if="account.avatar" :src="account.avatar" class="avatar" />
        <div class="text">
          <div class="username">用户名：{{ account.username }}</div>
          <div class="sex"
            >性别：
            <el-icon v-if="account.sex == '1'" style="color: dodgerblue"><Male /></el-icon>
            <el-icon v-if="account.sex == '2'" style="color: hotpink"><Female /></el-icon>
            <div v-else>未设置</div>
          </div>
          <div class="birth-day">出生日期：{{ account.birthDay }}</div>
          <div class="introduction">个性签名：{{ account.introduction }}</div>
        </div>
      </div>
      <div class="setting">
        <el-button type="primary" @click="updateAccount">修改</el-button>
      </div>
    </el-card>
    <el-card class="password-card">
      <template #header> 密码修改 </template>
      <div class="password">
        <div class="text">
          密码： ************
        </div>
        <el-button type="primary" @click="changePassword">修改</el-button>
      </div>
    </el-card>
  </div>

  <BaseSetting v-model:visible="baseSettingVisible" @afterSubmit="readAccount"/>
  <PasswordSetting v-model:visible="changePasswordVisible" @afterSubmit="logout"/>
</template>

<script setup lang="ts">
  import { reactive, onMounted, ref } from 'vue';
  import { ElMessage } from 'element-plus';
  import { reqAccountInfo } from '@/api/account';
  import { useRouter} from 'vue-router';
  import useUserStore from '@/store/modules/user';
  import BaseSetting from './base-setting.vue';
  import PasswordSetting from './password-setting.vue';

  let userStore = useUserStore();
  const $router = useRouter();  
  let account = reactive({
    avatar: '',
    username: '',
    introduction: '',
    sex: '',
    birthDay: new Date(),
  });
  let baseSettingVisible = ref(false);
  let changePasswordVisible = ref(false);

  // 获取账号
  const readAccount = async () => {
    try {
      let result = await reqAccountInfo();
      if (result.code == 200) {
        account.avatar = result.data.avatar;
        account.username = result.data.username;
        account.introduction = result.data.introduction;
        account.sex = result.data.sex;
        account.birthDay = new Date(result.data.birthDay);
      }
    } catch (err) {
      ElMessage({
        showClose: true,
        message: '用户数据读取失败',
        type: 'error',
      });
    }
  };

  const updateAccount = () => {
    baseSettingVisible.value = true;
  };

  const changePassword = () => {
    changePasswordVisible.value = true;
  };

  const logout = () => {
    userStore.userLogout();
    $router.push({ path: '/signIn'});
  };

  onMounted(() => {
    readAccount();
  });
</script>

<style scoped lang="scss">
  .account-setting-container {
    height: 100%;
    width: 100%;
    position: relative;
    display: flex;
    flex-direction: column;
    .account-base-card {
      margin-top: 30px;
      margin-left:auto;
      margin-right: auto;
      height: 280px;
      width: 800px;
      .message {
        display: flex;
        flex-direction: row;
        .avatar {
          height: 150px;
          width: 150px;
          border: 1px solid $base-border-color;
          margin-left: 40px;
          margin-right: 30px;
          object-fit: cover;
        }
        .text {
          padding: 10px;
          display: flex;
          flex-direction: column;
          .username {
            margin: 5px;
          }
          .sex {
            margin: 5px;
          }
          .birth-day {
            margin: 5px;
          }
          .introduction {
            margin: 5px;
          }
        }
      }
      .setting {
        display: flex;
        justify-content: center;
      }
    }
    .password-card {
      margin-top: 30px;
      margin-left:auto;
      margin-right: auto;
      height: 120px;
      width: 800px;
      .password {
        display: flex;
        flex-direction: row;
        .text {
          margin-left: 50px;
          margin-right: 50px;
        }
      }
    }
  }
</style>
