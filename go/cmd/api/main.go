package api

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/unkeyed/unkey/go/pkg/clickhouse"
	"github.com/unkeyed/unkey/go/pkg/config"
	"github.com/unkeyed/unkey/go/pkg/logging"
	"github.com/unkeyed/unkey/go/pkg/membership"
	"github.com/unkeyed/unkey/go/pkg/uid"
	"github.com/unkeyed/unkey/go/pkg/version"
	"github.com/unkeyed/unkey/go/pkg/zen"
	"github.com/unkeyed/unkey/go/pkg/zen/validation"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name: "api",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Load configuration file",
			Value:       "unkey.json",
			DefaultText: "unkey.json",
			EnvVars:     []string{"AGENT_CONFIG_FILE"},
		},
	},
	Action: run,
}

func run(c *cli.Context) error {
	configFile := c.String("config")

	// nolint:exhaustruct
	cfg := nodeConfig{}
	err := config.LoadFile(&cfg, configFile)
	if err != nil {
		return fmt.Errorf("unable to load config file: %w", err)
	}

	if cfg.NodeId == "" {
		cfg.NodeId = uid.Node()
	}

	if cfg.Region == "" {
		cfg.Region = "unknown"
	}
	logger := logging.New(logging.Config{Development: true, NoColor: false})
	logger = logger.With(
		slog.String("nodeId", cfg.NodeId),
		slog.String("platform", cfg.Platform),
		slog.String("region", cfg.Region),
		slog.String("version", version.Version),
	)

	// Catch any panics now after we have a logger but before we start the server
	defer func() {
		if r := recover(); r != nil {
			logger.Error(c.Context, "panic",
				slog.Any("panic", r),
				slog.String("stack", string(debug.Stack())),
			)
		}
	}()

	logger.Info(c.Context, "configration loaded", slog.String("file", configFile))

	m, err := membership.New(membership.Config{
		RedisUrl: cfg.RedisUrl,
		NodeID:   cfg.NodeId,
		RpcAddr:  "",
		Logger:   logger,
	})
	if err != nil {
		return fmt.Errorf("unable to create membership")
	}

	go func() {
		joins := m.SubscribeJoinEvents()
		leaves := m.SubscribeLeaveEvents()

		for {
			select {
			case j := <-joins:
				{
					logger.Info(c.Context, "node joined", slog.String("nodeID", j.ID))

				}
			case l := <-leaves:
				{
					logger.Info(c.Context, "node left", slog.String("nodeID", l.ID))

				}
			}

		}
	}()

	var ch clickhouse.Bufferer = clickhouse.NewNoop()
	if cfg.Clickhouse != nil {
		ch, err = clickhouse.New(clickhouse.Config{
			URL:    cfg.Clickhouse.Url,
			Logger: logger,
		})
		if err != nil {
			return fmt.Errorf("unable to create clickhouse: %w", err)
		}
	}

	srv, err := zen.New(zen.Config{
		NodeId: cfg.NodeId,
		Logger: logger,
	})
	if err != nil {
		return fmt.Errorf("unable to create server: %w", err)
	}

	validator, err := validation.New()
	if err != nil {
		return fmt.Errorf("unable to create validator: %w", err)
	}

	srv.SetGlobalMiddleware(
		// metrics should always run first, so it can capture the latency of the entire request
		zen.WithMetrics(ch),
		zen.WithLogging(logger),
		zen.WithErrorHandling(),
		zen.WithValidation(validator),
	)

	go func() {
		listenErr := srv.Listen(c.Context, fmt.Sprintf(":%s", cfg.Port))
		if listenErr != nil {
			panic(listenErr)
		}
	}()

	cShutdown := make(chan os.Signal, 1)
	signal.Notify(cShutdown, os.Interrupt, syscall.SIGTERM)

	<-cShutdown
	shutdownCtx := context.Background()
	logger.Info(c.Context, "shutting down")
	err = m.Leave(shutdownCtx)
	if err != nil {
		return fmt.Errorf("unable to leave cluster: %w", err)
	}
	return nil
}

func init() {
	// nolint:exhaustruct
	_, err := config.GenerateJsonSchema(nodeConfig{}, "schema.json")
	if err != nil {
		panic(err)
	}
}
