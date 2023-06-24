package rest

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Start REST server on specific port
func Start(ctx context.Context, port int) {
	app := fiber.New()
	router(app)

	go func() {
		<-ctx.Done()
		log.Info().Msg("cancel signal received in rest")
		if err := app.Shutdown(); err != nil {
			log.Err(err).Caller().Send()
		}
	}()

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal().Caller().Err(err)
	}
}
