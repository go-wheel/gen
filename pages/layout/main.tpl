{{define "main"}}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link href="http://libs.baidu.com/bootstrap/2.3.2/css/bootstrap.min.css" rel="stylesheet">
	<link href="/assets/css/docs.css" rel="stylesheet">
	<script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
	<script src="http://libs.baidu.com/bootstrap/2.3.2/js/bootstrap.min.js"></script>
</head>
<body>
  	<header class="jumbotron subhead" >
  		<h2>代码生成器 Go语言版 测试编译工具</h2>
  	</header>
	<div class="container-fluid">
	  <div class="row-fluid">
	    <div class="span2 bs-docs-sidebar">
	     	<ul class="nav nav-list bs-docs-sidenav">
  				<li {{if .IsShowMake}}class="active"{{end}}><a href="/make/">
  					<i class="icon-chevron-right"></i>编译SDK</a>
  				</li>
			</ul>	
	    </div>
	    <div class="span10">
		<div class="row-fluid info-box" >	
	    	<div class="alert {{if .IsMsgNo}}hide{{else}}{{if .IsMsgInfo}}alert-success{{else}}alert-error{{end}}{{end}}">
	    		<button type="button" class="close" data-dismiss="alert">&times;</button>
{{$page:=.}}
{{range $index, $msg := .GetResultMsg}}
{{$msg}}{{if $page.MsgNeedBr $index}}<br>{{end}}
{{end}}&nbsp;
			</div>
		</div>
		{{template "body" .}}

	    </div>
	  </div>
	</div>




	<br>

</body>
</html>
{{end}}