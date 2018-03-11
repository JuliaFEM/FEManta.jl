<template>
  <div class="servercontent">
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
</template>

<script>
export default {
  name: 'servercontent'
}
</script>
