package examples

// User represents a sample struct with various tags
type User struct {
    // Basic tags
    ID        int64  `json:"id" xml:"id" db:"user_id"`
    Name      string `json:"name" xml:"name" db:"user_name"`
    Email     string `json:"email" xml:"email" db:"user_email"`
    CreatedAt string `json:"created_at" xml:"created_at" db:"created_at"`

    // Tags with additional options
    Password string `json:"-" xml:"-" db:"password"` // Excluded from JSON/XML
    Age      int    `json:"age,omitempty" xml:"age,attr" db:"age"`

    // Embedded struct
    Address struct {
        Street  string `json:"street" xml:"street" db:"street"`
        City    string `json:"city" xml:"city" db:"city"`
        Country string `json:"country" xml:"country" db:"country"`
    } `json:"address" xml:"address" db:"address"`

    // Validation tags
    PhoneNumber string `validate:"required,len=10" json:"phone" xml:"phone"`
    Website     string `validate:"url" json:"website" xml:"website"`

    // Documentation tags
    Description string `doc:"User description" json:"description" xml:"description"`

    // Serialization tags
    Metadata map[string]interface{} `json:"metadata" xml:"metadata" yaml:"metadata"`

    // Tags with custom values
    Status string `json:"status" xml:"status" db:"status" enum:"active,inactive,suspended"`

    // Tags with spaces and special characters
    FullName string `json:"full_name" xml:"full-name" db:"full_name" validate:"required,min=2,max=100"`
}

// Config represents a sample configuration struct
type Config struct {
    // Tags with default values
    Port     int    `default:"8080" json:"port" yaml:"port"`
    Host     string `default:"localhost" json:"host" yaml:"host"`
    LogLevel string `default:"info" json:"log_level" yaml:"logLevel"`

    // Tags with descriptions
    Timeout int `desc:"Timeout in seconds" json:"timeout" yaml:"timeout"`

    // Tags with multiple values
    AllowedIPs []string `json:"allowed_ips" yaml:"allowedIPs" validate:"dive,ip"`
}

// Response represents a sample API response struct
type Response struct {
    // API documentation tags
    Code    int    `json:"code" example:"200" description:"HTTP status code"`
    Message string `json:"message" example:"Success" description:"Response message"`
    Data    any    `json:"data" example:"{}" description:"Response data"`

    // Metadata tags
    Timestamp string `json:"timestamp" format:"date-time" description:"Response time"`
    Version   string `json:"version" example:"1.0.0" description:"API version"`
}

// --- Additional test cases for plugin correctness ---

func main() {
    // Regular string (should NOT be highlighted as struct tag)
    s1 := "json:\"id\" xml:\"id\" db:\"user_id\""
    s2 := "just a normal string with `backticks` inside"
    s3 := "validate:'required' json:'phone'"

    // Raw string (should NOT be highlighted as struct tag)
    raw := `json:"id" xml:"id" db:"user_id"`
    raw2 := `this is not a struct tag, just a raw string with : and " inside`

    // Rune (should NOT be highlighted)
    r := '`'
    r2 := ':'
    r3 := '"'

    // Standalone backtick string (should NOT be highlighted)
    _ = `not a struct tag`
    _ = `json: "id"` // not a struct tag, just a raw string

    // Print to avoid unused variable errors
    println(s1, s2, s3, raw, raw2, r, r2, r3)
}
