package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/WhiCu/middleware/circuit"
	"github.com/gin-gonic/gin"
)

// CircuitBreakerMiddleware creates a Gin middleware using the circuit breaker
func CircuitBreaker(config *circuit.CircuitBreakerConfig) gin.HandlerFunc {
	cb := circuit.NewCircuitBreaker(*config)
	return func(c *gin.Context) {
		err := cb.Execute(func() error {
			c.Next()

			// Check if the response indicates an error
			if c.Writer.Status() >= 500 {
				return fmt.Errorf("HTTP %d: %s", c.Writer.Status(), http.StatusText(c.Writer.Status()))
			}
			return nil
		})

		if err != nil {
			// If circuit breaker rejected the request, return 503
			if errors.Is(err, circuit.ErrCircuitBreakerOpen) ||
				errors.Is(err, circuit.ErrMaxConcurrentCallsReached) {
				c.Header(circuit.HeaderCircuitBreakerState, cb.GetState().String())
				c.AbortWithStatus(circuit.StatusServiceUnavailable)
				return
			}

			// For other errors (like context cancellation), let the original error handling take place
			// Don't override the status if it's already set
		}
	}
}
