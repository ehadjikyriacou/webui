{{template "base" .}}
{{define "title"}}User Profile{{end}}
{{define "main"}}
<h2>User Profile</h2>
{{with .User}}
<table>
   <tr>
      <th>Name</th>
      <td>{{.Name}}</td>
   </tr>
   <tr>
      <th>Email</th>
      <td>{{.Email}}</td>
   </tr>
   <tr>
      <th>Joined</th>
      <td>{{humanDate .Created}}</td>
   </tr>
   <tr>
      <!-- Add a link to the change password form -->
      <th>Password</th>
      <td><a href="/user/change-password">Change password</a></td>
   </tr>
</table>
{{end }}
{{end}}