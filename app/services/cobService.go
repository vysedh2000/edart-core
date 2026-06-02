package services

import (
	"bytes"
	"edart-core/app/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

// CobService manages the cron job for daily COB processing and Telegram notifications.
type CobService struct {
	accRepo *repositories.AccRepository
	cron     *cron.Cron
	mutex    sync.Mutex
	telegramToken  string
	telegramChatID int64
}

// Global variables for singleton pattern
var (
	cobServiceInstance *CobService
	once                sync.Once
)

// loadEnv loads environment variables from the .env file.
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// NewCobService initializes the CobService with a cron job and Telegram configuration.
func NewCobService() *CobService {
	loadEnv()

	accRepo := repositories.NewAccountRepo()

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramChatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid TELEGRAM_CHAT_ID: %v", err))
	}

	loc, err := time.LoadLocation("Asia/Phnom_Penh")
	if err != nil {
		panic(fmt.Sprintf("failed to load timezone: %v", err))
	}

	c := cron.New(cron.WithLocation(loc))
	svc := &CobService{
		accRepo: accRepo,
		cron:          c,
		telegramToken:  telegramToken,
		telegramChatID: telegramChatID,
	}

	_, err = c.AddFunc("1 15 * * *", func() {
		svc.DailyCob()
	})
	if err != nil {
		panic(fmt.Sprintf("failed to add cron job: %v", err))
	}

	c.Start()
	fmt.Println("CobService initialized and cron job scheduled.")
	return svc
}

// GetCobService returns the singleton instance of CobService.
func GetCobService() *CobService {
	once.Do(func() {
		cobServiceInstance = NewCobService()
	})
	return cobServiceInstance
}

// Stop stops the cron scheduler.
func (s *CobService) Stop() {
	if s.cron != nil {
		s.cron.Stop()
		fmt.Println("CobService cron stopped.")
	}
}

// sendTelegramMessage sends a message to the configured Telegram chat.
func (s *CobService) sendTelegramMessage(text string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.telegramToken)

	message := map[string]interface{}{
		"chat_id": s.telegramChatID,
		"text":    text,
	}
	jsonBody, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Telegram API error: %s, response: %s", resp.Status, body)
	}

	return nil
}

// DailyCob is the job executed by the cron scheduler.
func (s *CobService) DailyCob() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	startTime := time.Now().In(time.FixedZone("ICT", 7*60*60))
	fmt.Printf("DailyCob started at %s (Phnom Penh time)\n", startTime.Format("2006-01-02 15:04:05"))

	// Send Telegram alert
	// message := fmt.Sprintf("🚀 Daily COB job started at %s (Phnom Penh time)", startTime.Format("15:04:05"))
	// err := s.sendTelegramMessage(message)
	// if err != nil {
	// 	fmt.Printf("Failed to send Telegram message: %v\n", err)
	// }

	accList, err := s.accRepo.CobGetAccList()
	if err != nil {
		fmt.Printf("Failed to retrieve account list: %v\n", err)
		return
	}
	fmt.Print("AcclIst", accList)

	// endTime := time.Now().In(time.FixedZone("ICT", 7*60*60))
	// message = fmt.Sprintf("🚀 Daily COB job finished at %s (Phnom Penh time)", endTime.Format("15:04:05"))
	// err = s.sendTelegramMessage(message)
}