package settings

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ENVIRONMENT string `required:"false"`
	PORT        int    `required:"false" default:"8000"`
	TIMESTAMP   string `required:"false"`
	APP_NAME    string `required:"false" default:"accounts"`

	// Database
	POSTGRES_DSN string `required:"true"`

	// JWT
	PRIVATE_KEY_JWT string `required:"true"`
	PUBLIC_KEY_JWT  string `required:"true"`
	JWT_EXPIRE      int    `required:"false" default:"900"`
	REFRESH_EXPIRE  int    `required:"false" default:"2592000"`

	// Password
	SECRET_PASSWORD string `required:"true"`

	// Event Bus
	USER_EVENT_BUS     string `required:"true"`
	PASSWORD_EVENT_BUS string `required:"true"`
	HOST_EVENT_BUS     string `required:"true"`
	PORT_EVENT_BUS     int    `required:"true"`
	VHOST_EVENT_BUS    string `required:"true"`

	// Email
	EMAIL_SENDER          string `required:"true"`
	EMAIL_SENDER_PASSWORD string `required:"true"`
	EMAIL_CLIENT          string `required:"true"`

	// Email Templates
	EMAIL_TEMPLATE_ACTIVATION_URL string `required:"true"`
	EMAIL_TEMPLATE_RESET_URL      string `required:"true"`
	EMAIL_TEMPLATE_WELCOME_URL    string `required:"true"`

	// Google OAuth
	GOOGLE_OAUTH_CLIENT_ID     string `required:"true"`
	GOOGLE_OAUTH_CLIENT_SECRET string `required:"true"`
	GOOGLE_OAUTH_REDIRECT_URI  string `required:"true"`

	// OAuth
	OAUTH_REDIRECT_URL string `required:"true"`
}

var Settings Config
var EnvDir = ".envs"

func LoadDotEnv() {

	err := godotenv.Load(fmt.Sprintf("%s/.env.base", EnvDir))
	if err != nil {
		log.Printf("No %s file found, using system environment variables", fmt.Sprintf("%s/.env.base.base", EnvDir))
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		log.Println("ENVIRONMENT is not set")
	}

	// Mapear el archivo .env.base correspondiente al entorno
	envFiles := map[string]string{
		"":            fmt.Sprintf("%s/.env", EnvDir),
		"local":       fmt.Sprintf("%s/.env.local", EnvDir),
		"development": fmt.Sprintf("%s/.env.dev", EnvDir),
		"production":  fmt.Sprintf("%s/.env.prod", EnvDir),
		"staging":     fmt.Sprintf("%s/.env.staging", EnvDir),
	}

	// Obtener el archivo de entorno correspondiente
	envFile, exists := envFiles[environment]
	if !exists {
		log.Printf("Environment '%s' is not supported. Must be one of: local, development, production, staging", environment)
	}

	// Cargar las variables desde el archivo correspondiente
	err = godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found, using system environment variables", envFile)
	} else {
		log.Printf("Loaded environment variables from %s", envFile)
	}
}

func LoadEnvs() {
	// Procesar las variables de entorno en la estructura Settings
	err := envconfig.Process("", &Settings)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	Settings.TIMESTAMP = time.Now().Format("2006-01-02 15:04:05")

	// Imprimir las Settings si el entorno es local o development
	if Settings.ENVIRONMENT == "local" || Settings.ENVIRONMENT == "development" {
		log.Println("Settings:")

		// Obtener el tipo y valor de la estructura Settings
		v := reflect.ValueOf(Settings)
		t := reflect.TypeOf(Settings)

		// Recorrer cada campo de la estructura
		for i := 0; i < v.NumField(); i++ {
			fieldName := t.Field(i).Name
			fieldValue := v.Field(i).Interface()
			fmt.Printf("\033[32m%s\033[0m: %v\n", fieldName, fieldValue)
		}
	}
}
