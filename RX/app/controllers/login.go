package controllers

import (
	"github.com/revel/revel"
	Manager "RX/services/manager"
	"fmt"
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	ResultSet, _ := Manager.Login("Mohan","admin123")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(ResultSet)
	return c.Render(ResultSet)
}

func (c Login) LoginAction() revel.Result {
	return c.Redirect(Login.Index)	
}
