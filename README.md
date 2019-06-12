Sample API CRUD using Golang And Gin

Instalation
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
</ol>


