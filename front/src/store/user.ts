import { defineStore } from 'pinia';
import http from '@/http/index';

const { BK_HCM_AJAX_URL_PREFIX } = window.PROJECT_CONFIG;
export const useUser = defineStore('user', {
  state: () => ({
    user: '',
  }),
  actions: {
    setUser(user: string) {
      this.user = user;
    },

    // 测试
    async test() {
      const res = await http.get(`${BK_HCM_AJAX_URL_PREFIX}/v4/organization/user_info/`);
      return res;
    },
  },
});
