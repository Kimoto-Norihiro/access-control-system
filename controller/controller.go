package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/access-control-system/repository"
	"github.com/Kimoto-Norihiro/access-control-system/teams"
	"github.com/Kimoto-Norihiro/access-control-system/usecase"
)

type controller struct {
	db          *gorm.DB
	usecase     usecase.Usecase
	teamsClient *teams.TeamsNotify
}

func NewController(db *gorm.DB) *controller {
	userRepo := repository.NewUserRepository()
	recordRepo := repository.NewRecordRepository()
	usecase := usecase.NewUsecase(db, userRepo, recordRepo)
	teamsClient := teams.NewClient()

	return &controller{
		db:          db,
		usecase:     usecase,
		teamsClient: teamsClient,
	}
}

// ユーザー登録
func (c *controller) CreateUser(ctx *gin.Context) {
	var input usecase.CreateUserInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.usecase.CreateUser(&input); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

// ユーザー一覧
func (c *controller) ListUsers(ctx *gin.Context) {
	output, err := c.usecase.ListUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"users": output.Users})
}

// 入室
func (c *controller) Enter(ctx *gin.Context) {
	var input usecase.EnterInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.usecase.Enter(&input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	listOutput, err := c.usecase.ListExistUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// teamsに通知
	if err := c.teamsClient.SendEnterMessage(output.UserName, output.EnterAt, listOutput.UserNames); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}

// 退室
func (c *controller) Exit(ctx *gin.Context) {
	var input usecase.ExitInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	output, err := c.usecase.Exit(&input)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 在室しているユーザーの情報を取得
	listOutput, err := c.usecase.ListExistUsers()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// teamsに通知
	if err := c.teamsClient.SendExitMessage(output.UserName, output.ExitAt, listOutput.UserNames); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "success"})
}
