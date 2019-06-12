<h1>Sample API CRUD using Golang And Gin</h1>

<h3>Instalation</h3>
<hr/>
<ol>
  <li>Create some database and Import table in the folder database on MYSQL</li>
  <li>
    Open your terminal and run command :
    <ul>
      <li>go get github.com/gin-contrib/logger</li>
      <li>go get github.com/gin-gonic/gin</li>
      <li>go get github.com/rs/zerolog</li>
      <li>go get github.com/rs/zerolog/log</li>
      <li>go get github.com/jinzhu/gorm</li>
      <li>go get github.com/go-sql-driver/mysql</li>
    </ul>
  </li>
  <li>
    Change the contents of the connectdb.go file in the config folder as follow :
    <ul>
      <li>dbHost := "your host name / IP database"</li>
      <li>dbPort := "your port hostname / IP database"</li>
      <li>dbUser := "your user database"</li>
      <li>dbPass := "your password database"</li>
      <li>dbName := "your database name"</li>
    </ul>
  </li>
  <li>Open your terminal on the root project and run command go run main.go</li>
</ol>


<h3>Module</h3>
<hr/>
<table style="width:100%;border:1px;">
  <tr>
    <td>Jill</td>
    <td>Smith</td>
    <td>50</td>
  </tr>
  <tr>
    <td>Eve</td>
    <td>Jackson</td>
    <td>94</td>
  </tr>
  <tr>
    <td>John</td>
    <td>Doe</td>
    <td>80</td>
  </tr>
</table>


