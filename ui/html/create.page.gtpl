{{template "base" .}}
{{define "title"}}Create a New Snippet{{end}}
{{define "main"}}
<form action='/snippet/create' method='POST'>
   <input type='hidden' name='csrf_token' value='{{.CSRFToken}}'>
   {{with .Form}}
   <div>
      <label>Title:</label>
      {{with .Errors.Get "title"}}
      <label class='error'>{{.}}</label>
      {{end}}
      <input type='text' name='title' value='{{.Get "title"}}'>
   </div>
   <div>
      <label>Content:</label>
      {{with .Errors.Get "content"}}
      <label class='error'>{{.}}</label>
      {{end}}
      <textarea name='content'>{{.Get "content"}}</textarea>
   </div>
   <div>
      <label>Delete in:</label>
      {{with .Errors.Get "expires"}}
      <label class='error'>{{.}}</label>
      {{end}}
      {{$exp := or (.Get "expires") "365"}}
      <input type='radio' name='expires' value='365' {{if (eq $exp "365")}}checked{{end}}> 1 Year
      <input type='radio' name='expires' value='90' {{if (eq $exp "90")}}checked{{end}}> 90 Days
      <input type='radio' name='expires' value='30' {{if (eq $exp "30")}}checked{{end}}> 30 Days
      <input type='radio' name='expires' value='7' {{if (eq $exp "7")}}checked{{end}}> 1 Week
      <input type='radio' name='expires' value='1' {{if (eq $exp "1")}}checked{{end}}> 1 Day
   </div>
   <div>
      <input type='submit' value='Publish snippet'>
   </div>
   {{end}}
</form>
{{end}}