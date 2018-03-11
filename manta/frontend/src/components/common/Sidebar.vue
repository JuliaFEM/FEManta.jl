<template>
    <div class="sidebar">
    <div class="col s3">
    <ul class="collapsible" data-collapsible="accordion">
        <li>
            <div class="collapsible-header"><i class="material-icons">whatshot</i><a href="/">Server status</a></div>
            </li>
        <li>
            <div class="collapsible-header"><i class="material-icons">insert_drive_file</i>Models</div>
            <div class="collapsible-body">
                <div class="collection">
                    <a href="/models" class="collection-item active">Upload model</a>
                    <a href="/models" class="collection-item">Download model</a>
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
    </div>
</template>

<script>
export default {
  name: 'sidebar'
}
</script>
