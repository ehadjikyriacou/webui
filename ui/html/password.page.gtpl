{{template "base" .}}
{{define "title"}}Change Password{{end}}
{{define "main"}}
<h2>Change Password</h2>
<form action='/user/change-password' method='POST' novalidate>
<input type='hidden' name='csrf_token' value='{{.CSRFToken}}'>
{{with .Form}}
<div>
<label>Current password:</label>
{{with .Errors.Get "currentPassword"}}
<label class='error'>{{.}}</label>
{{end}}
<input type='password' name='currentPassword'>
</div>
<div>
<label>New password:</label>
{{with .Errors.Get "newPassword"}}
<label class='error'>{{.}}</label>
{{end}}
<input type='password' name='newPassword'>
</div>
<div>
<label>Confirm password:</label>
{{with .Errors.Get "newPasswordConfirmation"}}
<label class='error'>{{.}}</label>
{{end}}
<input type='password' name='newPasswordConfirmation'>
</div>
<div>
<input type='submit' value='Change password'>
</div>
{{end}}
</form>
{{end}}