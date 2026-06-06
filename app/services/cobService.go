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
	accRepo        *repositories.AccRepository
	balRepo		   *repositories.BalanceRepository
	cron           *cron.Cron
	mutex          sync.Mutex
	telegramToken  string
	telegramChatID int64
}

// Global variables for singleton pattern
var (
	cobServiceInstance *CobService
	once               sync.Once
)

// loadEnv loads environment variables from the .env file.
func loadEnv() {
	_ = godotenv.Load()
}

// NewCobService initializes the CobService with a cron job and Telegram configuration.
func NewCobService() *CobService {
	loadEnv()

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	var telegramChatID int64
	if chatID := os.Getenv("TELEGRAM_CHAT_ID"); chatID != "" {
		parsedChatID, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			log.Printf("invalid TELEGRAM_CHAT_ID: %v", err)
		} else {
			telegramChatID = parsedChatID
		}
	}

	loc := time.Local

	c := cron.New(cron.WithLocation(loc))
	svc := &CobService{
		cron:           c,
		telegramToken:  telegramToken,
		telegramChatID: telegramChatID,
	}

	// This will now run at 00:05 (12:05 AM) machine local time
	_, err := c.AddFunc("5 0 * * *", func() {
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
	if s != nil && s.cron != nil {
		s.cron.Stop()
		fmt.Println("CobService cron stopped.")
	}
}

func (s *CobService) ensureAccRepo() error {
	if s == nil {
		return fmt.Errorf("cob service is not initialized")
	}

	if s.accRepo == nil || s.accRepo.DB == nil {
		s.accRepo = repositories.NewAccountRepo()
	}

	if s.balRepo == nil || s.balRepo.DB == nil {
		s.balRepo = repositories.NewBalanceRepo()
	}

	return nil
}

// sendTelegramMessage sends a message to the configured Telegram chat.
func (s *CobService) sendTelegramMessage(text string) error {
	if s.telegramToken == "" || s.telegramChatID == 0 {
		return fmt.Errorf("telegram configuration is missing")
	}

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
	if s == nil {
		fmt.Println("DailyCob failed: cob service is not initialized")
		return
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := s.ensureAccRepo(); err != nil {
		fmt.Printf("DailyCob failed: %v\n", err)
		return
	}

	startTime := time.Now().In(time.FixedZone("ICT", 7*60*60))
	fmt.Printf("DailyCob started at %s (Phnom Penh time)\n", startTime.Format("2006-01-02 15:04:05"))

	// Send Telegram alert
	message := fmt.Sprintf("🚀 Daily COB job started at %s (Phnom Penh time)", startTime.Format("15:04:05"))
	err := s.sendTelegramMessage(message)
	if err != nil {
		fmt.Printf("Failed to send Telegram message: %v\n", err)
	}

	accList, err := s.accRepo.CobGetAccList()
	if err != nil {
		fmt.Printf("Failed to retrieve account list: %v\n", err)
		return
	}
	for _, acc := range accList {
		accBal, err := s.balRepo.CobGetBalByAcc(acc.AccNo)
		if err != nil {
			fmt.Printf("Failed to update closing balance for account %s: %v\n", acc.AccNo, err)
		}
		if (accBal != acc.WorkingBal) {
			telegramMsg := fmt.Sprintf("⚠️ COB Alert: Account %s has a closing balance of %.2f which differs from the working balance of %.2f", acc.AccNo, accBal, acc.WorkingBal)
			fmt.Print(telegramMsg)
			err = s.sendTelegramMessage(telegramMsg)
		}

		status, err := s.accRepo.CobUpdateBal(acc.AccNo, accBal)
		if err != nil {
			telegramMsg := fmt.Sprintf("⚠️ COB Alert: Failed to update error: %v for account %s", err, acc.AccNo)
			err = s.sendTelegramMessage(telegramMsg)
			fmt.Printf("Failed to update closing balance for account %s: %v\n", acc.AccNo, err)
		}
		if status != "success" {
			telegramMsg := fmt.Sprintf("⚠️ COB Alert: Failed to update closing balance for account %s: %s", acc.AccNo, status)
			err = s.sendTelegramMessage(telegramMsg)
			fmt.Printf("Failed to update closing balance for account %s: %s\n", acc.AccNo, status)
		}

	}

	endTime := time.Now().In(time.FixedZone("ICT", 7*60*60))
	fmt.Printf("DailyCob finished at %s (Phnom Penh time)\n", endTime.Format("2006-01-02 15:04:05"))
	message = fmt.Sprintf("🚀 Daily COB job finished at %s (Phnom Penh time)", endTime.Format("15:04:05"))
	err = s.sendTelegramMessage(message)
}
