import Vue from 'vue'
import App from '@/App.vue'
import router from '@/global/router'
import store from '@/global/store'
import event from '@/utils/event'
import cache from '@/api/cache'
import axios from 'axios'
Vue.prototype.$event = event
Vue.prototype.$cache = cache
Vue.prototype.$axios = axios
Vue.config.productionTip = false
// this.$axios.defaults.baseURL = 'http://localhost:3000'
new Vue({
  axios,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
