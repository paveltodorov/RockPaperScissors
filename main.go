package main

import (
	"rockpaperscissors/ai"
	"rockpaperscissors/challenge"
	"rockpaperscissors/funds"
	"rockpaperscissors/storage"
	"rockpaperscissors/user"

	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.NewMemoryStore()

	userSvc := user.NewService(store)
	fundsSvc := funds.NewService(userSvc)
	challengeSvc := challenge.NewService(store, userSvc)
	aiSvc := ai.NewService(userSvc)

	r := gin.Default()

	// User routes
	r.POST("/login", user.LoginHandler(userSvc))
	r.POST("/logout", user.LogoutHandler())
	r.GET("/users", user.ListHandler(userSvc))

	// Funds routes
	r.POST("/deposit", funds.DepositHandler(fundsSvc))
	r.POST("/withdraw", funds.WithdrawHandler(fundsSvc))

	// Challenge routes
	r.POST("/challenges", challenge.CreateHandler(challengeSvc))
	r.GET("/challenges", challenge.ListHandler(challengeSvc))
	r.POST("/challenges/accept", challenge.AcceptHandler(challengeSvc))
	r.POST("/challenges/decline", challenge.DeclineHandler(challengeSvc))

	// AI routes
	r.POST("/ai/create", ai.CreateAIHandler(aiSvc))
	r.GET("/ai/list", ai.ListAIHandler(aiSvc))
	r.POST("/ai/challenge", ai.AIChallengeHandler(aiSvc, challengeSvc))

	r.Run(":8080")
}
