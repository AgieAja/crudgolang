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
    <th>Module</th>
    <th>Method</th>
    <th>URL</th>
    <th>Description</th>
    <th>Request Body</th>
  </tr>
  <tr>
    <td>Add</td>
    <td>POST</td>
    <td>{baseURL}/phonebook/add</td>
    <td>Add some data in table phone_books</td>
    <td>{"name":"Tiar","phone":"12346"}</td>
  </tr>
  <tr>
    <td>Edit</td>
    <td>POST</td>
    <td>{baseURL}/phonebook/edit</td>
    <td>edit some data in table phone_books</td>
    <td>{"id":1,"name":"Tiar","phone":"12346"}</td>
  </tr>
  <tr>
    <td>List</td>
    <td>GET</td>
    <td>{baseURL}/phonebook/list</td>
    <td>get all data in table phone_books</td>
    <td>-</td>
  </tr>
  <tr>
    <td>Get Data by ID</td>
    <td>GET</td>
    <td>{baseURL}/phonebook/data/id</td>
    <td>get data by id in table phone_books</td>
    <td>-</td>
  </tr>
  <tr>
    <td>Delete</td>
    <td>POST</td>
    <td>{baseURL}/phonebook/delete</td>
    <td>delete some data in table phone_books</td>
    <td>{"id":1}</td>
  </tr>
</table>


