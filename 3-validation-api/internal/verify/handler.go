package verify

import (
	"encoding/json"
	"fmt"
	"go/validation-api/config"
	"go/validation-api/pkg/crypto"
	"go/validation-api/pkg/request"
	"go/validation-api/pkg/response"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"sync"

	"github.com/jordan-wright/email"
)

type VerifyHandlerDeps struct {
	UserConfig *config.UserConfig
}

type VerifyHandler struct {
	userConfig *config.UserConfig
	mu         sync.RWMutex
	storage    map[string]string
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		userConfig: deps.UserConfig,
		storage:    make(map[string]string),
	}
	if err := handler.loadStorage(); err != nil {
		log.Printf("[INIT] Предупреждение: не удалось загрузить базу данных хэшей: %v", err)
	}
	router.HandleFunc("POST /send", handler.Send)
	router.HandleFunc("GET /verify/{hash}", handler.Verify)
}

// POST /send
func (h *VerifyHandler) Send(w http.ResponseWriter, req *http.Request) {
	payload, err := request.HandleBody[SendRequest](w, req)
	if err != nil {
		return
	}

	// 1. Генерируем случайный хэш
	hash, err := crypto.GenerateRandomHash(16)
	if err != nil {
		response.Json(w, map[string]string{"error": "Failed to generate hash"}, 400)
		return
	}

	// 2. Сохраняем в память и сбрасываем в JSON-файл
	h.mu.Lock()
	h.storage[hash] = payload.Email
	err = h.saveStorage()
	h.mu.Unlock()
	if err != nil {
		response.Json(w, map[string]string{"error": "Failed to save data"}, 400)
		return
	}

	// 3. Отправляем email через jordan-wright/email
	log.Printf("[API] Запуск горутины для отправки email на %s...", payload.Email)
	go func(targetEmail, currentHash string) {
		err := h.sendVerificationEmail(payload.Email, hash)
		if err != nil {
			log.Printf("Ошибка отправки email на %s: %v", payload.Email, err)
		} else {
			log.Printf("[SMTP SUCCESS] Письмо успешно передано SMTP-серверу для %s!", payload.Email)
		}
	}(payload.Email, hash)

	response.Json(w, map[string]string{"status": "success", "message": "Verification email sent"}, 200)
}

// GET /verify/{hash}
func (h *VerifyHandler) Verify(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hash := req.PathValue("hash")
	fmt.Printf("verify triggered для хэша: %s\n", hash)

	h.mu.Lock()
	defer h.mu.Unlock()

	// Ищем хэш в нашей мапе
	email, exists := h.storage[hash]

	if exists {
		// Если нашли: удаляем запись из памяти
		delete(h.storage, hash)

		// Перезаписываем JSON-файл уже без этого хэша
		if err := h.saveStorage(); err != nil {
			log.Printf("Ошибка обновления файла верификации: %v", err)
		}

		log.Printf("Успешная валидация для email: %s. Запись удалена.", email)
		w.WriteHeader(http.StatusOK)
		response.Json(w, true, 200) // Используем ваш пакет response
	} else {
		// Если хэш не найден или уже был использован
		log.Printf("Валидация провалена для хэша: %s", hash)
		w.WriteHeader(http.StatusBadRequest)
		response.Json(w, false, 400)
	}
}

const storageFile = "verifications.json"

// Отправка письма
func (h *VerifyHandler) sendVerificationEmail(toEmail, hash string) error {
	e := email.NewEmail()
	// Здесь данные отправителя берутся из вашего h.userConfig (настройте под себя)
	e.From = fmt.Sprintf("Validation API <%s>", h.userConfig.Email)
	e.To = []string{toEmail}
	e.Subject = "Подтверждение Email"

	verifyLink := fmt.Sprintf("http://localhost:8083/verify/%s", hash)
	e.Text = []byte(fmt.Sprintf("Для подтверждения перейдите по ссылке: %s", verifyLink))

	// Разделяем "smtp.gmail.com:587" на хост ("smtp.gmail.com") и порт ("587")
	host, _, err := net.SplitHostPort(h.userConfig.Address)
	if err != nil {
		return fmt.Errorf("неверный формат адреса SMTP в конфиге: %v", err)
	}

	// Авторизация на SMTP сервере (данные должны быть в config.UserConfig)
	auth := smtp.PlainAuth("", h.userConfig.Email, h.userConfig.Password, host)

	// Для отправки Send передаем полный адрес с портом ("smtp.gmail.com:587")
	return e.Send(h.userConfig.Address, auth)
}

// Сохранение всей мапы в JSON-файл
func (h *VerifyHandler) saveStorage() error {
	data, err := json.MarshalIndent(h.storage, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(storageFile, data, 0644)
}

// Функция чтения данных с диска
func (h *VerifyHandler) loadStorage() error {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		return nil 
	}
	data, err := os.ReadFile(storageFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &h.storage)
}
