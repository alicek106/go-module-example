package main
 
import (
        "github.com/labstack/echo"
        "net/http"
        "github.com/alicek106/echo-server/wonderland"
)
 
func main(){
        e := echo.New()
        e.GET("/", func (c echo.Context) error{
                name := c.QueryParam("name")
                var helloMessage string = wonderland.HelloFromWonderland(name)
                return c.JSON(http.StatusOK, helloMessage)
        })
        e.Logger.Fatal(e.Start(":1331"))
}

