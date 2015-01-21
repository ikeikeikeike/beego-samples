{{append . "HeadStyles" "/static/css/custom.css"}}
{{append . "HeadScripts" "/static/js/custom.js"}}

<div class="container">
    <div class="row vertical-offset-75">
    	<div class="col-md-6 col-md-offset-3">
    		<div class="panel panel-default">
			  	<div class="panel-heading text-center">
			    	<h3 class="panel-title"><strong>Login</strong></h3>
			 	</div> 

			  	<div class="panel-body">
			    	<form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='{{urlfor "LoginController.Login"}}'>
                      {{ .xsrfdata }}

                      {{template "alert.tpl" .}}

                      <div class="form-group">
                        <label for="inputEmail" class="col-sm-3 control-label">メールアドレス</label>
                        <div class="col-sm-9">
                          <input class="form-control" placeholder="例: admin@example.com" name="Email" value="{{index .Params "Email"}}" type="email" required 
                                    id="inputEmail" />
                        </div>
                      </div>
                      <div class="form-group">
                        <label for="inputPassword" class="col-sm-3 control-label">パスワード</label>
                        <div class="col-sm-9">
			    		  <input class="form-control" placeholder="パスワードを入力" name="Password" type="password" value="" required 
                                    pattern=".{6,}" title="パスワードは6文字以上を入力してください" id="inputPassword"  />
                        </div>
                      </div>
                      <div class="form-group">
                        <div class="col-sm-12">
			    		    <input class="btn btn-lg btn-success" type="submit" value="ログイン"> 
                            <a href="{{urlfor "LoginController.PasswordReset"}}"> 
                                パスワードを忘れた場合はこちら »
                            </a>
                        </div>
                      </div>
                    </form>
			    </div>

                <div class="panel-footer text-center clearfix">アカウントをお持ちではない場合は <a href='{{urlfor "LoginController.Signup"}}'>新規登録 »</a></div>

			</div>
		</div>
	</div>
</div>
