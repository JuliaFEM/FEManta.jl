<template>
    <div class="navigationbar">
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
    </div>
</template>

<script>
export default {
  name: 'navigationbar'
}
</script>
