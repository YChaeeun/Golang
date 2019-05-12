package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.GET("/show", show)

	e.POST("/save", save) //POST는 화면에 출력X 웹서버로 넘어갈때? 자동으로 GET 형식으로 넘어감

	e.Logger.Fatal(e.Start(":1325"))

}

func getUser(c echo.Context) error {
	// 주소 창에 값 받아와서 화면에 출력하기
	id := c.Param("id")
	return c.String(http.StatusOK, id)

	// 주소창 : http://localhost:1323/users/출력하고싶은 문자
}

func show(c echo.Context) error {
	// Query Parameter
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)

	// 주소창 : http://localhost:1323/show?team=x-men&member=wolverine
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)

	// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
	// => name:Joe Smith, email:joe@labstack.com
	// postman을 사용해서 결과를 확인 할 수 있음!!
}
