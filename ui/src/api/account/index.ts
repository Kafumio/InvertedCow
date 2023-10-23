import request from '@/utils/request';
import { toFormData } from 'axios';

enum API {
  SIGN_IN_URL = '/account/signIn',
  SIGN_UP_URL = '/account/signUp',
  SEND_CODE_URL = '/account/code/send',
  GET_INFO_URL = '/account/get/info',
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

// reqUserInfo 获取用户信息
export const reqUserInfo = (): Promise<any> => {
  return request.get(API.GET_INFO_URL);
};
