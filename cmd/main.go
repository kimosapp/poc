package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	controller "github.com/kimosapp/poc/internal/controller/organization"
	userController "github.com/kimosapp/poc/internal/controller/users"
	logging2 "github.com/kimosapp/poc/internal/core/ports/logging"
	organizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations"
	roleRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/role"
	teamRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team"
	teamMemberRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/team-member"
	userOrganizationRepository "github.com/kimosapp/poc/internal/core/ports/repository/organizations/user-organization"
	userRepository "github.com/kimosapp/poc/internal/core/ports/repository/users"
	notificationsService "github.com/kimosapp/poc/internal/core/ports/service/notification"
	organization "github.com/kimosapp/poc/internal/core/usercase/organizations"
	usecase "github.com/kimosapp/poc/internal/core/usercase/users"
	"github.com/kimosapp/poc/internal/infrastructure/client"
	"github.com/kimosapp/poc/internal/infrastructure/configuration"
	"github.com/kimosapp/poc/internal/infrastructure/db"
	"github.com/kimosapp/poc/internal/infrastructure/logging"
	"github.com/kimosapp/poc/internal/infrastructure/repository/postgres/notifications"
	organizationRepositoryPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/organizations"
	roleRepositoryPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/organizations/role"
	teamRepositoryPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/organizations/team"
	teamMemberRepositoryPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/organizations/team-member"
	userOrganizationRepositoryPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/organizations/user-organization"
	userPostgres "github.com/kimosapp/poc/internal/infrastructure/repository/postgres/users"
	"github.com/kimosapp/poc/internal/infrastructure/server"
	"github.com/kimosapp/poc/internal/infrastructure/service/notification"
	"github.com/kimosapp/poc/internal/middleware"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}
	// Create a new instance of the Gin router
	instance := gin.New()
	instance.Use(gin.Recovery())
	conn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}
	logger, err := logging.NewLogger()
	if err != nil {
		log.Fatalf("failed to new logger err=%s\n", err.Error())
	}

	gmailClient := client.NewEmailClientGmail(logger)

	userRepo := userPostgres.NewUserRepository(conn)
	orgRepo := organizationRepositoryPostgres.NewOrganizationRepository(conn)
	userOrgRepo := userOrganizationRepositoryPostgres.NewUserOrganizationRepository(conn)
	roleRepo := roleRepositoryPostgres.NewRoleRepository(conn)
	teamRepo := teamRepositoryPostgres.NewTeamRepository(conn)
	teamMemberRepo := teamMemberRepositoryPostgres.NewTeamMemberRepository(conn)
	notificationTemplateRepo := notifications.NewNotificationTemplateRepository(conn)
	notificationService := notification.NewNotificationService(
		gmailClient,
		notificationTemplateRepo,
	)
	initOrganizationController(
		instance,
		orgRepo,
		userOrgRepo,
		roleRepo,
		teamRepo,
		teamMemberRepo,
		userRepo,
		middleware.NewAuthMiddleware(userRepo),
		notificationService,
		logger,
	)

	//TODO remove this
	//createUserUseCase := usecase.NewCreateUserUseCase(userRepo, logger)
	authenticateUserUseCase := usecase.NewAuthenticateUserUseCase(
		userRepo,
		logger,
	)
	getUserUseCase := usecase.NewGetUserUseCase(userRepo, logger)
	updateUserProfileUseCase := usecase.NewUpdateUserProfileUseCase(
		userRepo,
		logger,
	)

	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	userControllerInstance := userController.NewUserController(
		instance,
		logger,
		authenticateUserUseCase,
		getUserUseCase,
		authMiddleware,
		updateUserProfileUseCase,
	)

	userControllerInstance.InitRouter()
	httpServer := server.NewHttpServer(
		instance,
		configuration.GetHttpServerConfig(),
	)
	httpServer.Start()
	defer httpServer.Stop()

	// Listen for OS signals to perform a graceful shutdown
	log.Println("listening signals...")
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}

func initOrganizationController(
	instance *gin.Engine,
	orgRepo organizationRepository.Repository,
	userOrgRepo userOrganizationRepository.Repository,
	roleRepo roleRepository.Repository,
	teamRepo teamRepository.Repository,
	teamMemberRepo teamMemberRepository.Repository,
	userRepo userRepository.Repository,
	middleware *middleware.AuthMiddleware,
	notificationService notificationsService.Service,
	logger logging2.Logger,
) {
	createOrganizationUseCase := organization.NewCreateOrganizationUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		userRepo,
		notificationService,
		logger,
	)
	getOrgByUserIdAndOrgIdUseCase := organization.NewGetOrganizationByOrgIdAndUserIdUseCase(
		orgRepo,
		logger,
	)
	getOrganizationsByUserIdUseCase := organization.NewGetOrganizationsByUserUseCase(
		orgRepo,
		logger,
	)

	checkIfUserHasEnoughPermissionsUseCase := organization.NewCheckUserHasPermissionsToMakeAction(
		userOrgRepo,
		logger,
	)

	createOrganizationUserUseCase := organization.NewCreateOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		roleRepo,
		userRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger,
	)
	removeOrganizationUserUseCase := organization.NewRemoveOrganizationMemberUseCase(
		orgRepo,
		userOrgRepo,
		logger,
	)

	createTeamUseCase := organization.NewCreateTeamUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger,
	)

	addMemberToTeamUseCase := organization.NewAddTeamMembersUseCase(
		userOrgRepo,
		teamRepo,
		teamMemberRepo,
		checkIfUserHasEnoughPermissionsUseCase,
		logger,
	)

	organizationController := controller.NewOrganizationController(
		instance,
		logger,
		createOrganizationUseCase,
		getOrgByUserIdAndOrgIdUseCase,
		getOrganizationsByUserIdUseCase,
		createOrganizationUserUseCase,
		removeOrganizationUserUseCase,
		createTeamUseCase,
		addMemberToTeamUseCase,
		middleware,
	)
	organizationController.InitRouter()
}
