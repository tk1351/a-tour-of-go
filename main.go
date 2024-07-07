package main

import (
	"a-tour-of-go/app/models"
	"fmt"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)

	// /*
	// u := &models.User{}
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.Password = "pass"
	// fmt.Println(u)

	// u.CreateUser()
	// */

	/*
		u, _ := models.GetUser(1)
		fmt.Println(u)

		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		u.DeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	// /*
	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")
	// */

	/*
		todos, _ := models.GetTodos()
		fmt.Println("todos", todos)
		for _, v := range todos {
			fmt.Println(v)
		}
	*/
	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		fmt.Println("todos", todos)
		for _, v := range todos {
			fmt.Println(v)
		}
	*/
	t, _ := models.GetTodo(1)
	t.Content = "Update Todo"
	t.UpdateTodo()
	fmt.Println(t)
}
