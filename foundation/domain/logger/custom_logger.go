package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type LoggerSettings struct {
	ENVIRONMENT string
	APP_NAME    string
	LOKI_URL    string
}

var LoggerConfig LoggerSettings

func InitLogger(env string, app_name string, loki_url string) {
	LoggerConfig = LoggerSettings{
		ENVIRONMENT: env,
		APP_NAME:    app_name,
		LOKI_URL:    loki_url,
	}
}

// contextKey es el tipo para la clave del logger en el contexto.
type contextKey string

const loggerKey contextKey = "logger"

// CustomFormatter es un formateador personalizado para Logrus.
type CustomFormatter struct{}

// Format implementa la interfaz Formatter de Logrus.
// Genera un formato: timestamp | level | function:line | fields | message
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(time.RFC3339)
	fields, ok := entry.Data["fields"].(LogFields)
	if !ok {
		fields = LogFields{}
	}
	caller := ""
	lineNum := 0
	if entry.Caller != nil {
		caller = entry.Caller.Function
		lineNum = entry.Caller.Line
	}
	line := fmt.Sprintf("%s | %s | %s:%d | %s | %s\n",
		timestamp,
		entry.Level.String(),
		caller,
		lineNum,
		fields.ToString(),
		entry.Message,
	)
	return []byte(line), nil
}

var globalLogger *logrus.Logger
var lokiHook *LokiHook

func init() {
	// Configuración del logger global.
	globalLogger = logrus.New()
	globalLogger.SetFormatter(&CustomFormatter{})
	globalLogger.SetLevel(logrus.InfoLevel)
	globalLogger.SetReportCaller(true)

	// Se agrega el hook para que cada log se envíe a Loki.
	lokiHook = NewLokiHook(50)
	globalLogger.AddHook(lokiHook)
}

// WithFields crea un entry de logger con campos adicionales.
func WithFields(fields LogFields) *logrus.Entry {
	mapFields := map[string]interface{}{
		"fields": fields,
	}
	return globalLogger.WithFields(mapFields)
}

// WithLogger inyecta un entry de logger en el contexto.
func WithLogger(ctx context.Context, entry *logrus.Entry) context.Context {
	return context.WithValue(ctx, loggerKey, entry)
}

// FromContext obtiene el logger desde el contexto; si no hay, devuelve el logger global.
func FromContext(ctx context.Context) *logrus.Entry {

	entry, ok := ctx.Value("logger").(*logrus.Entry)

	if ok {
		return entry
	}
	return globalLogger.WithFields(logrus.Fields{})
}
