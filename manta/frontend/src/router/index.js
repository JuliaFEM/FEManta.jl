import Vue from 'vue'
import Router from 'vue-router'
import ServerPage from '@/components/Server'
import ModelsPage from '@/components/Models'

Vue.use(Router)

export default new Router({
  routes: [
    { path: '/', component: ServerPage },
    { path: '/models', component: ModelsPage }
  ],
  mode: 'history'
})
