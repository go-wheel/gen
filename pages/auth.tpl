{{define "body"}}

<form action="/auth/save" method="post" class="form-horizontal">

<div class="row-fluid tool-bar">
	<input type="submit" class="btn btn-success" value="保存参数"/>
</div>
<div class="control-group">
	<label class="control-label" for="AppKey">App Key:</label>
    <div class="controls">
      <input type="text" id="AppKey" name="AppKey" class="span8" placeholder="AppKey" value="{{.GetConfMain.AppKey}}"/>
    </div>
</div>

<div class="control-group">
	<label class="control-label span2" for="AppSecret">App Secret:</label>
    <div class="controls">
      <input type="text" id="AppSecret" name="AppSecret" class="span8" placeholder="AppSecret" value="{{.GetConfMain.AppSecret}}"/>
    </div>
</div>

<div class="control-group">
	<label class="control-label span2" for="PackgeSite">Packge Site:</label>
    <div class="controls">
      <input type="text" id="PackgeSite" name="PackgeSite" class="span8" placeholder="PackgeSite" value="{{.GetConfMain.PackgeSite}}"/>
    </div>
</div>


</form>

{{end}}