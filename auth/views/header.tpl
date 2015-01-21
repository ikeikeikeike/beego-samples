<header id="topbar" class="navbar navbar-default navbar-fixed-top bs-docs-nav" role="banner">
  <div class="container">
    <div class="navbar-header">
      <button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">
        <span class="sr-only">トグルナビゲーション</span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a style="font-size: 14px;" class="navbar-brand" rel="home" href="/" >
        <strong>Ref</strong>
      </a>
    </div>

    <nav class="navbar-collapse bs-navbar-collapse collapse" role="navigation" style="height: 1px;">
      <ul itemscope="itemscope" itemtype="http://www.schema.org/SiteNavigationElement" class="nav navbar-nav">
        <li itemprop="name"><a itemprop="url" href='{{urlfor "admin.UsersController.Index"}}'>
          <span class="glyphicon glyphicon-facetime-video"></span> 
        </a></li>
        <li itemprop="name"><a itemprop="url" href="">
          <span class="glyphicon glyphicon-heart"></span> 
        </a></li>
        <li itemprop="name"><a itemprop="url" href="">
          <span class="glyphicon glyphicon-credit-card"></span> 
        </a></li>
      </ul>

      <ul itemscope="itemscope" itemtype="http://www.schema.org/SiteNavigationElement" class="nav navbar-nav navbar-right">
        <li class="dropdown">
          <a href="#" class="dropdown-toggle" data-toggle="dropdown">
            <span class='glyphicon glyphicon-info-sign'></span> Account <b class="caret"></b>
          </a>
          <ul itemprop="name" class="dropdown-menu">
            {{if .IsLogin}}
                <li itemprop="name" ><a itemprop="url" href='{{urlfor "LoginController.Logout"}}'>
                  <span class='glyphicon glyphicon-bullhorn'></span> Logout
                </a></li>
            {{else}}
                <li itemprop="name" ><a itemprop="url" href='{{urlfor "LoginController.Login"}}'>
                  <span class='glyphicon glyphicon-globe'></span> Login
                </a></li>
            {{end}}
          </ul>
        </li>
      </ul>
    </nav>
  </div>

</header>
