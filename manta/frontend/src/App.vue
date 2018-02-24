<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
  <nav class="light-blue lighten-1" role="navigation">
    <div class="nav-wrapper container"><a id="logo-container" href="https://github.com/JuliaFEM/JuliaFEM.jl" class="brand-logo">JuliaFEM</a>
      <ul class="right hide-on-med-and-down">
        <li><a href="#">Navbar Link</a></li>
      </ul>

      <ul id="nav-mobile" class="side-nav">
        <li><a href="#">Navbar Link</a></li>
      </ul>
      <a href="#" data-activates="nav-mobile" class="button-collapse"><i class="material-icons">menu</i></a>
    </div>
  </nav>

  <div class="row">

    <div class="col s2">
      <ul class="collapsible" data-collapsible="accordion">
          <li>
              <div class="collapsible-header"><i class="material-icons">whatshot</i>Server status</div>
            </li>
          <li>
            <div class="collapsible-header"><i class="material-icons">insert_drive_file</i>Models</div>
            <div class="collapsible-body">
                <div class="collection">
                    <a href="#!" class="collection-item active">Upload model</a>
                    <a href="#!" class="collection-item">Download model</a>
                </div>
              </div>
          </li>
          <li>
            <div class="collapsible-header"><i class="material-icons">place</i>Start calculation</div>
            <div class="collapsible-body"><span>Lorem ipsum dolor sit amet.</span></div>
          </li>
          <li>
            <div class="collapsible-header"><i class="material-icons">whatshot</i>Calculation status and results</div>
            <div class="collapsible-body"><span>Lorem ipsum dolor sit amet.</span></div>
          </li>
        </ul>
    </div>
    <div class="col s9">
      <div class="section no-pad-bot" id="index-banner">
        <div class="container">
          <!-- <h1 class="header center orange-text">Upload model</h1> -->
          <div class="row center">
            <h5 class="header col s12 light">Upload model to </h5>
          </div>
          <div class="row center">
              <form action="/model">
                  <div class="file-field input-field">
                    <div class="btn">
                      <span>File</span>
                      <input type="file">
                    </div>
                    <div class="file-path-wrapper">
                      <input class="file-path validate" type="text">
                    </div>
                  </div>
                </form>
                <form action="#">
                  <div class="file-field input-field">
                    <div class="btn">
                      <span>File</span>
                      <input type="file">
                    </div>
                    <div class="file-path-wrapper">
                      <input class="file-path validate" type="text">
                    </div>
                  </div>
                </form>
          </div>
          <br><br>
        </div>
      </div>

      <div class="container">
        <div class="section">

        <!--   Icon Section   -->
        <div class="row">
          <div class="col s12 m4">
            <div class="icon-block">
              <h2 class="center light-blue-text"><i class="material-icons">flash_on</i></h2>
              <h5 class="center">Text</h5>
              <p class="light">Here there be text.</p>
            </div>
          </div>
        </div>
        </div>
        <br><br>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
