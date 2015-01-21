{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}


<div class="container">
    <div class="row vertical-offset-75">

        {{template "alert.tpl" .}}

        <p>Email: {{.Userinfo.Email}}</p>
        <p>User: {{.Userinfo}}</p>

    </div>
</div>
