<template>
  <el-dialog v-model="visible" class="dialog">
    <el-form class="login-form" :rules="rules" :model="data" ref="passwordElFrom">
      <el-form-item prop="oldPassword">
        <el-input
          v-model="data.oldPassword"
          type="password"
          placeholder="请输入原密码"
        />
      </el-form-item>
      <el-form-item prop="newPassword">
        <el-input
          v-model="data.newPassword"
          type="password"
          placeholder="请输入新密码"
        />
      </el-form-item>
      <el-form-item prop="newPassword2">
        <el-input
          v-model="data.newPassword2"
          type="password"
          placeholder="请再次输入密码"
        />
      </el-form-item>
    </el-form>
    <div class="change-password-submit">
      <el-button type="info" @click="visible = false">取消</el-button>
      <el-button type="primary" @click="submitChangePassword">提交</el-button>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
  import { reactive, computed, ref } from 'vue';
  import { reqChangePassword } from '@/api/account';
  import { ElMessage } from 'element-plus';

  const props = defineProps(['visible']);
  const emit = defineEmits(['update:visible', 'afterSubmit']);

  let visible = computed({
    get() {
      return props.visible;
    },
    set(value) {
      emit('update:visible', value);
    },
  });
  let data = reactive({
    oldPassword: '',
    newPassword: '',
    newPassword2: '',
  });

  let passwordElFrom = ref();

  // 表单校验
  let rules = {
    oldPassword: [{ required: true, message: '原密码不能为空', trigger: 'change' }],
    newPassword: [
      { required: true, message: '新密码不能为空', trigger: 'change' },
      { required: true, min: 6, max: 20, message: '密码长度需6-20位' },
    ],
    newPassword2: [
      {
        validator: function (_role: any,value: any, callback: any) {
          if (value != data.newPassword) {
            callback(new Error('两次密码不相等'));
          } else {
            callback();
          }
        },
        trigger: 'blur',
      },
    ],
  };

  const submitChangePassword = async () => {
    // 等待表单校验通过
    await passwordElFrom.value.validate();
    console.log(data)
    let result = await reqChangePassword({
        oldPassword: data.oldPassword,
        newPassword: data.newPassword,
    });
    if (result.code == 200) {
        ElMessage({
          showClose: true,
          message: "修改成功，请重新登陆",
          type: 'success',
        });
        visible.value = false;
        emit('afterSubmit');
      } else {
        ElMessage({
          showClose: true,
          message: '提交失败',
          type: 'error',
        });
      }
  };
</script>

<style scoped lang="scss">
  .dialog {
    .change-password-submit {
      display: flex;
      justify-content: center;
    }
  }
</style>
