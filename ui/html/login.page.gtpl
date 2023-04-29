{{template "base" .}}
{{define "title"}}Login{{end}}
{{define "main"}}
<form action='/user/login' method='POST' novalidate>
   <input type='hidden' name='csrf_token' value='{{.CSRFToken}}'>
   {{with .Form}}
   {{with .Errors.Get "generic"}}
   <div class='error'>{{.}}</div>
   {{end}}
   <div>
      <label>Email:</label>
      <input type='email' name='email' value='{{.Get "email"}}'>
   </div>
   <div>
      <label>Password:</label>
      <input type='password' name='password'>
   </div>
   <div>
      <input type='submit' value='Login'>
   </div>
   {{end}}
</form>
{{end}}