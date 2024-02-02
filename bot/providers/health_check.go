package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
)

func OnComponentStatusChanged(_ context.Context, name string, state health.CheckState) {
	logger.Info(fmt.Sprintf("component %s changed status to %s", name, state.Status))
}

func OnSystemStatusChanged(_ context.Context, state health.CheckerState) {
	logger.Info(fmt.Sprintf("system status changed to %s", state.Status))
}

func OnReadinessStatusChanged(_ context.Context, state health.CheckerState) {
	logger.Info(fmt.Sprintf("readiness status changed to %s", state.Status))
}

func NewReadinessChecker(discordReady func() bool) health.Checker {
	return health.NewChecker(
		health.WithTimeout(10*time.Second),
		// The following check will be executed periodically every 15 seconds
		// started with an initial delay of 3 seconds. The check function will NOT
		// be executed for each HTTP request.
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name:           "DiscordWSDataReady",
			StatusListener: OnComponentStatusChanged,
			// If the check function returns an error, this component will be considered unavailable ("down").
			// The context contains a deadline according to the configuration of the Checker.
			Check: func(ctx context.Context) error {
				if !discordReady() {
					return fmt.Errorf("discord session not ready")
				}
				return nil
			},
		}),
		// Set a status listener that will be invoked when the health status changes.
		// More powerful hooks are also available (see docs). For guidance, please refer to the links
		// listed in the main function documentation.
		health.WithStatusListener(OnReadinessStatusChanged),
	)
}

func NewHealthLoggingInterceptor() health.Interceptor {
	return log.LoggingInterceptor("")
}
