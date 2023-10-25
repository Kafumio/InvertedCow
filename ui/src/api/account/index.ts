import request from '@/utils/request';
import { toFormData } from 'axios';

enum API {
  ACCOUNT_URL = '/account',
  SIGN_IN_URL = '/account/signIn',
  SIGN_UP_URL = '/account/signUp',
  SEND_CODE_URL = '/account/code/send',
  GET_INFO_URL = '/account/get/info',
  CHANGE_PASSWORD_URL = '/account/password',
  AVATAR_URL = '/account/avatar/upload',
}

// reqSignIn 登录
export const reqSignIn = (data: any): Promise<any> => {
  return request.post(API.SIGN_IN_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// reqSignUp 注册
export const reqSignUp = (data: any): Promise<any> => {
  return request.post(API.SIGN_UP_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// reqSendCode 发送校验邮箱
export const reqSendCode = (data: any): Promise<any> => {
  return request.post(API.SEND_CODE_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// reqAccountInfo 获取用户信息
export const reqAccountInfo = (): Promise<any> => {
  return request.get(API.GET_INFO_URL);
};

// reqUpdateAccount 更新账号信息
export const reqUpdateAccount = (data: any): Promise<any> => {
  return request.put(API.ACCOUNT_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// reqChangePassword 修改密码
export const reqChangePassword = (data: any): Promise<any> => {
  return request.put(API.CHANGE_PASSWORD_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};

// reqUploadAvatar 上传头像
export const reqUploadAvatar = (data: any): Promise<any> => {
  return request.post(API.AVATAR_URL, toFormData(data), {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};