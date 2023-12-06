package main

import (
	"sosmed/config"
	uh "sosmed/features/users/handler"
	ur "sosmed/features/users/repository"
	us "sosmed/features/users/services"

	ph "sosmed/features/posting/handler"
	pr "sosmed/features/posting/repository"
	ps "sosmed/features/posting/services"

	ch "sosmed/features/comment/handler"
	cr "sosmed/features/comment/repository"
	cs "sosmed/features/comment/services"


	nk "sosmed/helper/enkrip"
	cld "sosmed/utils/cloudinary"
	"sosmed/routes"

	"sosmed/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	if cfg == nil {
		e.Logger.Fatal("tidak bisa start karena ENV error")
		return
	}
	cld, ctx, param := cld.InitCloudnr(*cfg)

	db, err := database.InitMySQL(*cfg)

	if err != nil {
		e.Logger.Fatal("tidak bisa start karena DB error:", err.Error())
		return
	}

	db.AutoMigrate(&ur.UserModel{}, &pr.PostingModel{}, &cr.CommentModel{})

	usersRepo := ur.New(db)
	userService := us.New(usersRepo, nk.New())
	userHandler := uh.New(userService, cld, ctx, param)

	postingRepo := pr.New(db)
	postingService := ps.New(postingRepo)
	postingHandler := ph.New(postingService, cld, ctx, param)

	commentRepo := cr.New(db)
	commentService := cs.New(commentRepo)
	commentHandler := ch.New(commentService)

	routes.InitRoute(e, userHandler, postingHandler, commentHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
