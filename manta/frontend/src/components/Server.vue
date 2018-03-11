<template>
  <div class="serverpage">
    <navigationbar></navigationbar>
    <div class="row">
      <Sidebar></Sidebar>
      <ServerContent></ServerContent>
    </div>
  </div>
</template>

<script>
import Navigationbar from '@/components/common/Navigationbar.vue'
import Sidebar from '@/components/common/Sidebar.vue'
import ServerContent from '@/components/content/ServerContent.vue'

export default {
  name: 'ServerPage',
  components: {
    Navigationbar,
    Sidebar,
    ServerContent
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
