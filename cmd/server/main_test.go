package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"buf.build/gen/go/a-novel/proto/grpc/go/storystructure/v1/storystructurev1grpc"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	anovelgrpc "github.com/a-novel/golib/grpc"
	"github.com/a-novel/golib/testutils"
)

func init() {
	go main()
}

var servicesToTest = []string{
	"batch_delete_beats",
	"batch_delete_plot_points",
	"create_beat",
	"create_plot_point",
	"delete_beat",
	"delete_plot_point",
	"get_beat",
	"get_plot_point",
	"list_beats",
	"list_plot_points",
	"search_beats",
	"search_plot_points",
	"update_beat",
	"update_plot_point",
}

func TestIntegrationHealth(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode.")
	}

	// Create the RPC client.
	pool := anovelgrpc.NewConnPool()
	conn, err := pool.Open("0.0.0.0", 8080, anovelgrpc.ProtocolHTTP)
	require.NoError(t, err)

	healthClient := healthpb.NewHealthClient(conn)

	testutils.WaitConn(t, conn)

	for _, service := range servicesToTest {
		res, err := healthClient.Check(context.Background(), &healthpb.HealthCheckRequest{Service: service})
		require.NoError(t, err)
		require.Equal(t, healthpb.HealthCheckResponse_SERVING, res.Status)
	}
}

func TestIntegrationBeatCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode.")
	}

	// Create the RPC client.
	pool := anovelgrpc.NewConnPool()
	conn, err := pool.Open("0.0.0.0", 8080, anovelgrpc.ProtocolHTTP)
	require.NoError(t, err)

	testutils.WaitConn(t, conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createBeatClient := storystructurev1grpc.NewCreateBeatServiceClient(conn)
	getBeatClient := storystructurev1grpc.NewGetBeatServiceClient(conn)
	updateBeatClient := storystructurev1grpc.NewUpdateBeatServiceClient(conn)
	deleteBeatClient := storystructurev1grpc.NewDeleteBeatServiceClient(conn)

	// Create the beat.
	beat, err := createBeatClient.Exec(ctx, &storystructurev1.CreateBeatServiceExecRequest{
		Name:      "Beat",
		Prompt:    "Create the beat",
		CreatorId: "capybara",
	})
	require.NoError(t, err)
	require.NotNil(t, beat)
	require.NotEmpty(t, beat.Id)
	require.Equal(t, "Beat", beat.Name)
	require.Equal(t, "Create the beat", beat.Prompt)
	require.Equal(t, "capybara", beat.CreatorId)

	// Get the beat.
	gotBeat, err := getBeatClient.Exec(ctx, &storystructurev1.GetBeatServiceExecRequest{
		Id: beat.Id,
	})
	require.NoError(t, err)
	require.Equal(t, beat.Id, gotBeat.Id)
	require.Equal(t, beat.Name, gotBeat.Name)
	require.Equal(t, beat.Prompt, gotBeat.Prompt)
	require.Equal(t, beat.CreatorId, gotBeat.CreatorId)
	require.Equal(t, beat.CreatedAt, gotBeat.CreatedAt)

	// Update the beat.
	updatedBeat, err := updateBeatClient.Exec(ctx, &storystructurev1.UpdateBeatServiceExecRequest{
		Id:        beat.Id,
		Name:      "Beat Mania",
		Prompt:    "Update the beat",
		CreatorId: "capybara",
	})
	require.NoError(t, err)
	require.NotNil(t, updatedBeat)
	require.NotEmpty(t, updatedBeat.Id)
	require.Equal(t, "Beat Mania", updatedBeat.Name)
	require.Equal(t, "Update the beat", updatedBeat.Prompt)
	require.Equal(t, "capybara", updatedBeat.CreatorId)

	// Delete the beat.
	_, err = deleteBeatClient.Exec(ctx, &storystructurev1.DeleteBeatServiceExecRequest{
		Id: beat.Id,
	})
	require.NoError(t, err)

	// Get the beat again.
	_, err = getBeatClient.Exec(ctx, &storystructurev1.GetBeatServiceExecRequest{
		Id: beat.Id,
	})
	testutils.RequireGRPCCodesEqual(t, err, codes.NotFound)
}

func TestIntegrationPlotPointCRUD(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests in short mode.")
	}

	// Create the RPC client.
	pool := anovelgrpc.NewConnPool()
	conn, err := pool.Open("0.0.0.0", 8080, anovelgrpc.ProtocolHTTP)
	require.NoError(t, err)

	testutils.WaitConn(t, conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createPlotPointClient := storystructurev1grpc.NewCreatePlotPointServiceClient(conn)
	getPlotPointClient := storystructurev1grpc.NewGetPlotPointServiceClient(conn)
	updatePlotPointClient := storystructurev1grpc.NewUpdatePlotPointServiceClient(conn)
	deletePlotPointClient := storystructurev1grpc.NewDeletePlotPointServiceClient(conn)

	// Create the PlotPoint.
	PlotPoint, err := createPlotPointClient.Exec(ctx, &storystructurev1.CreatePlotPointServiceExecRequest{
		Name:      "PlotPoint",
		Prompt:    "Create the PlotPoint",
		CreatorId: "capybara",
	})
	require.NoError(t, err)
	require.NotNil(t, PlotPoint)
	require.NotEmpty(t, PlotPoint.Id)
	require.Equal(t, "PlotPoint", PlotPoint.Name)
	require.Equal(t, "Create the PlotPoint", PlotPoint.Prompt)
	require.Equal(t, "capybara", PlotPoint.CreatorId)

	// Get the PlotPoint.
	gotPlotPoint, err := getPlotPointClient.Exec(ctx, &storystructurev1.GetPlotPointServiceExecRequest{
		Id: PlotPoint.Id,
	})
	require.NoError(t, err)
	require.Equal(t, PlotPoint.Id, gotPlotPoint.Id)
	require.Equal(t, PlotPoint.Name, gotPlotPoint.Name)
	require.Equal(t, PlotPoint.Prompt, gotPlotPoint.Prompt)
	require.Equal(t, PlotPoint.CreatorId, gotPlotPoint.CreatorId)
	require.Equal(t, PlotPoint.CreatedAt, gotPlotPoint.CreatedAt)

	// Update the PlotPoint.
	updatedPlotPoint, err := updatePlotPointClient.Exec(ctx, &storystructurev1.UpdatePlotPointServiceExecRequest{
		Id:        PlotPoint.Id,
		Name:      "PlotPoint Mania",
		Prompt:    "Update the PlotPoint",
		CreatorId: "capybara",
	})
	require.NoError(t, err)
	require.NotNil(t, updatedPlotPoint)
	require.NotEmpty(t, updatedPlotPoint.Id)
	require.Equal(t, "PlotPoint Mania", updatedPlotPoint.Name)
	require.Equal(t, "Update the PlotPoint", updatedPlotPoint.Prompt)
	require.Equal(t, "capybara", updatedPlotPoint.CreatorId)

	// Delete the PlotPoint.
	_, err = deletePlotPointClient.Exec(ctx, &storystructurev1.DeletePlotPointServiceExecRequest{
		Id: PlotPoint.Id,
	})
	require.NoError(t, err)

	// Get the PlotPoint again.
	_, err = getPlotPointClient.Exec(ctx, &storystructurev1.GetPlotPointServiceExecRequest{
		Id: PlotPoint.Id,
	})
	testutils.RequireGRPCCodesEqual(t, err, codes.NotFound)
}
