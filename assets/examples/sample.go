package examples

// User представляет пример структуры с различными тегами
type User struct {
	// Базовые теги
	ID        int64  `json:"id" xml:"id" db:"user_id"`
	Name      string `json:"name" xml:"name" db:"user_name"`
	Email     string `json:"email" xml:"email" db:"user_email"`
	CreatedAt string `json:"created_at" xml:"created_at" db:"created_at"`

	// Теги с дополнительными опциями
	Password string `json:"-" xml:"-" db:"password"` // Исключение из JSON/XML
	Age      int    `json:"age,omitempty" xml:"age,attr" db:"age"`

	// Вложенные структуры
	Address struct {
		Street  string `json:"street" xml:"street" db:"street"`
		City    string `json:"city" xml:"city" db:"city"`
		Country string `json:"country" xml:"country" db:"country"`
	} `json:"address" xml:"address" db:"address"`

	// Теги валидации
	PhoneNumber string `validate:"required,len=10" json:"phone" xml:"phone"`
	Website     string `validate:"url" json:"website" xml:"website"`

	// Теги для документации
	Description string `doc:"Описание пользователя" json:"description" xml:"description"`

	// Теги для сериализации
	Metadata map[string]interface{} `json:"metadata" xml:"metadata" yaml:"metadata"`

	// Теги с кастомными значениями
	Status string `json:"status" xml:"status" db:"status" enum:"active,inactive,suspended"`

	// Теги с пробелами и специальными символами
	FullName string `json:"full_name" xml:"full-name" db:"full_name" validate:"required,min=2,max=100"`
}

// Config представляет пример конфигурационной структуры
type Config struct {
	// Теги с дефолтными значениями
	Port     int    `default:"8080" json:"port" yaml:"port"`
	Host     string `default:"localhost" json:"host" yaml:"host"`
	LogLevel string `default:"info" json:"log_level" yaml:"logLevel"`

	// Теги с описаниями
	Timeout int `desc:"Таймаут в секундах" json:"timeout" yaml:"timeout"`

	// Теги с множественными значениями
	AllowedIPs []string `json:"allowed_ips" yaml:"allowedIPs" validate:"dive,ip"`
}

// Response представляет пример структуры ответа API
type Response struct {
	// Теги для API документации
	Code    int    `json:"code" example:"200" description:"HTTP статус код"`
	Message string `json:"message" example:"Success" description:"Сообщение ответа"`
	Data    any    `json:"data" example:"{}" description:"Данные ответа"`

	// Теги для метаданных
	Timestamp string `json:"timestamp" format:"date-time" description:"Время ответа"`
	Version   string `json:"version" example:"1.0.0" description:"Версия API"`
} 