package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"

	"github.com/a-novel/golib/database"
	anovelgrpc "github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/loggers"
	"github.com/a-novel/golib/loggers/adapters"
	"github.com/a-novel/golib/loggers/formatters"

	"github.com/a-novel/uservice-story-structure/config"
	"github.com/a-novel/uservice-story-structure/migrations"
	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

var rpcServices = []grpc.ServiceDesc{
	healthpb.Health_ServiceDesc,
	storystructurev1grpc.BatchDeleteBeatsService_ServiceDesc,
	storystructurev1grpc.BatchDeletePlotPointsService_ServiceDesc,
	storystructurev1grpc.CreateBeatService_ServiceDesc,
	storystructurev1grpc.CreatePlotPointService_ServiceDesc,
	storystructurev1grpc.DeleteBeatService_ServiceDesc,
	storystructurev1grpc.DeletePlotPointService_ServiceDesc,
	storystructurev1grpc.GetBeatService_ServiceDesc,
	storystructurev1grpc.GetPlotPointService_ServiceDesc,
	storystructurev1grpc.ListBeatsService_ServiceDesc,
	storystructurev1grpc.ListPlotPointsService_ServiceDesc,
	storystructurev1grpc.SearchBeatsService_ServiceDesc,
	storystructurev1grpc.SearchPlotPointsService_ServiceDesc,
	storystructurev1grpc.UpdateBeatService_ServiceDesc,
	storystructurev1grpc.UpdatePlotPointService_ServiceDesc,
}

func getDepsCheck(database *bun.DB) *anovelgrpc.DepsCheck {
	return &anovelgrpc.DepsCheck{
		Dependencies: anovelgrpc.DepCheckCallbacks{
			"postgres": database.Ping,
		},
		Services: anovelgrpc.DepCheckServices{
			"batch_delete_beats":       {"postgres"},
			"batch_delete_plot_points": {"postgres"},
			"create_beat":              {"postgres"},
			"create_plot_point":        {"postgres"},
			"delete_beat":              {"postgres"},
			"delete_plot_point":        {"postgres"},
			"get_beat":                 {"postgres"},
			"get_plot_point":           {"postgres"},
			"list_beats":               {"postgres"},
			"list_plot_points":         {"postgres"},
			"search_beats":             {"postgres"},
			"search_plot_points":       {"postgres"},
			"update_beat":              {"postgres"},
			"update_plot_point":        {"postgres"},
		},
	}
}

func main() {
	logger := config.Logger.Formatter

	loader := formatters.NewLoader(
		fmt.Sprintf("Acquiring database connection at %s...", config.App.Postgres.DSN),
		spinner.Meter,
	)
	logger.Log(loader, loggers.LogLevelInfo)

	postgresDB, closePostgresDB, err := database.OpenDB(config.App.Postgres.DSN)
	if err != nil {
		logger.Log(formatters.NewError(err, "open database conn"), loggers.LogLevelFatal)
	}
	defer closePostgresDB()

	logger.Log(
		loader.SetDescription("Database connection successfully acquired.").SetCompleted(),
		loggers.LogLevelInfo,
	)

	if err := database.Migrate(postgresDB, migrations.SQLMigrations, logger); err != nil {
		logger.Log(formatters.NewError(err, "migrate database"), loggers.LogLevelFatal)
	}

	loader = formatters.NewLoader("Setup services...", spinner.Meter)
	logger.Log(loader, loggers.LogLevelInfo)

	grpcReporter := adapters.NewGRPC(logger)

	batchDeleteBeatsDAO := dao.NewBatchDeleteBeats(postgresDB)
	batchDeletePlotPointsDAO := dao.NewBatchDeletePlotPoints(postgresDB)
	createBeatDAO := dao.NewCreateBeat(postgresDB)
	createPlotPointDAO := dao.NewCreatePlotPoint(postgresDB)
	deleteBeatDAO := dao.NewDeleteBeat(postgresDB)
	deletePlotPointDAO := dao.NewDeletePlotPoint(postgresDB)
	getBeatDAO := dao.NewGetBeat(postgresDB)
	getPlotPointDAO := dao.NewGetPlotPoint(postgresDB)
	listBeatsDAO := dao.NewListBeats(postgresDB)
	listPlotPointsDAO := dao.NewListPlotPoints(postgresDB)
	searchBeatsDAO := dao.NewSearchBeats(postgresDB)
	searchPlotPointsDAO := dao.NewSearchPlotPoints(postgresDB)
	updateBeatDAO := dao.NewUpdateBeat(postgresDB)
	updatePlotPointDAO := dao.NewUpdatePlotPoint(postgresDB)

	batchDeleteBeatsService := services.NewBatchDeleteBeats(batchDeleteBeatsDAO)
	batchDeletePlotPointsService := services.NewBatchDeletePlotPoints(batchDeletePlotPointsDAO)
	createBeatService := services.NewCreateBeat(createBeatDAO)
	createPlotPointService := services.NewCreatePlotPoint(createPlotPointDAO)
	deleteBeatService := services.NewDeleteBeat(deleteBeatDAO)
	deletePlotPointService := services.NewDeletePlotPoint(deletePlotPointDAO)
	getBeatService := services.NewGetBeat(getBeatDAO)
	getPlotPointService := services.NewGetPlotPoint(getPlotPointDAO)
	listBeatsService := services.NewListBeats(listBeatsDAO)
	listPlotPointsService := services.NewListPlotPoints(listPlotPointsDAO)
	searchBeatsService := services.NewSearchBeats(searchBeatsDAO)
	searchPlotPointsService := services.NewSearchPlotPoints(searchPlotPointsDAO)
	updateBeatService := services.NewUpdateBeat(updateBeatDAO)
	updatePlotPointService := services.NewUpdatePlotPoint(updatePlotPointDAO)

	batchDeleteBeatsHandler := handlers.NewBatchDeleteBeats(batchDeleteBeatsService, grpcReporter)
	batchDeletePlotPointsHandler := handlers.NewBatchDeletePlotPoints(batchDeletePlotPointsService, grpcReporter)
	createBeatHandler := handlers.NewCreateBeat(createBeatService, grpcReporter)
	createPlotPointHandler := handlers.NewCreatePlotPoint(createPlotPointService, grpcReporter)
	deleteBeatHandler := handlers.NewDeleteBeat(deleteBeatService, grpcReporter)
	deletePlotPointHandler := handlers.NewDeletePlotPoint(deletePlotPointService, grpcReporter)
	getBeatHandler := handlers.NewGetBeat(getBeatService, grpcReporter)
	getPlotPointHandler := handlers.NewGetPlotPoint(getPlotPointService, grpcReporter)
	listBeatsHandler := handlers.NewListBeats(listBeatsService, grpcReporter)
	listPlotPointsHandler := handlers.NewListPlotPoints(listPlotPointsService, grpcReporter)
	searchBeatsHandler := handlers.NewSearchBeats(searchBeatsService, grpcReporter)
	searchPlotPointsHandler := handlers.NewSearchPlotPoints(searchPlotPointsService, grpcReporter)
	updateBeatHandler := handlers.NewUpdateBeat(updateBeatService, grpcReporter)
	updatePlotPointHandler := handlers.NewUpdatePlotPoint(updatePlotPointService, grpcReporter)

	logger.Log(loader.SetDescription("Services successfully setup.").SetCompleted(), loggers.LogLevelInfo)

	listener, server, err := anovelgrpc.StartServer(config.App.Server.Port)
	if err != nil {
		logger.Log(formatters.NewError(err, "start server"), loggers.LogLevelFatal)
	}
	defer anovelgrpc.CloseServer(listener, server)

	reflection.Register(server)
	healthpb.RegisterHealthServer(server, anovelgrpc.NewHealthServer(getDepsCheck(postgresDB), time.Minute))
	storystructurev1grpc.RegisterBatchDeleteBeatsServiceServer(server, batchDeleteBeatsHandler)
	storystructurev1grpc.RegisterBatchDeletePlotPointsServiceServer(server, batchDeletePlotPointsHandler)
	storystructurev1grpc.RegisterCreateBeatServiceServer(server, createBeatHandler)
	storystructurev1grpc.RegisterCreatePlotPointServiceServer(server, createPlotPointHandler)
	storystructurev1grpc.RegisterDeleteBeatServiceServer(server, deleteBeatHandler)
	storystructurev1grpc.RegisterDeletePlotPointServiceServer(server, deletePlotPointHandler)
	storystructurev1grpc.RegisterGetBeatServiceServer(server, getBeatHandler)
	storystructurev1grpc.RegisterGetPlotPointServiceServer(server, getPlotPointHandler)
	storystructurev1grpc.RegisterListBeatsServiceServer(server, listBeatsHandler)
	storystructurev1grpc.RegisterListPlotPointsServiceServer(server, listPlotPointsHandler)
	storystructurev1grpc.RegisterSearchBeatsServiceServer(server, searchBeatsHandler)
	storystructurev1grpc.RegisterSearchPlotPointsServiceServer(server, searchPlotPointsHandler)
	storystructurev1grpc.RegisterUpdateBeatServiceServer(server, updateBeatHandler)
	storystructurev1grpc.RegisterUpdatePlotPointServiceServer(server, updatePlotPointHandler)

	report := formatters.NewDiscoverGRPC(rpcServices, config.App.Server.Port)
	logger.Log(report, loggers.LogLevelInfo)

	if err := server.Serve(listener); err != nil {
		logger.Log(formatters.NewError(err, "serve"), loggers.LogLevelFatal)
	}
}
