import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify';
import { defineCustomElements as defineIonPhaser } from '@ion-phaser/core/loader';
import '@babel/polyfill'
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'

Vue.config.productionTip = false
Vue.config.ignoredElements = [/ion-\w*/];
// Bind the IonPhaser custom element to the window object
defineIonPhaser(window);


new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
