<template>
  <div class="ModelsPage">
    <navigationbar></navigationbar>
    <div class="row">
      <sidebar></sidebar>
      <ModelsContent></ModelsContent>
    </div>
  </div>
</template>

<script>
import Navigationbar from '@/components/common/Navigationbar.vue'
import Sidebar from '@/components/common/Sidebar.vue'
import ModelsContent from '@/components/content/ModelsContent.vue'

export default {
  name: 'ModelsPage',
  components: {
    Navigationbar,
    Sidebar,
    ModelsContent
  },
  data () {
    return {
      msg: 'Models view'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
