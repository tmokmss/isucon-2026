package isureal

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/labstack/echo/v5"
)

type initializeResponse struct {
	Language string `json:"language"`
	Agent    string `json:"agent"`
	Harness  string `json:"harness"`
	Model    string `json:"model"`
}

func initializeHandler(h *HandlerContext) echo.HandlerFunc {
	return func(c *echo.Context) error {
		h.InitMu.Lock()
		defer h.InitMu.Unlock()
		start := time.Now()
		cmd := exec.Command("../sql/init.sh")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
		log.Printf("init.sh done in %s", time.Since(start))

		// DB と同じ初期化をメモリにも行う。大きい表は境界より後の行を捨て、小さい表は読み直す
		s := h.Store
		s.mu.Lock()
		defer s.mu.Unlock()
		if err := loadSmall(c.Request().Context(), h.DB, s); err != nil {
			return err
		}
		s.pruneAfterBoundary()
		s.rebuildIndexes()
		imageCache.reset()
		log.Printf("initialize done in %s", time.Since(start))

		return c.JSON(http.StatusOK, initializeResponse{Language: "go", Agent: "claude-code", Harness: "claude-code", Model: "claude-opus-5-5"})
	}
}
