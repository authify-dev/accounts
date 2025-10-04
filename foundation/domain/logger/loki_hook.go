package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// Hook para enviar logs a Loki en batch.
type LokiHook struct {
	buffer    []LogEntry
	batchSize int
	mu        sync.Mutex
}

type LogEntry struct {
	Message string
	Level   string
	Fields  LogFields
}

func NewLokiHook(batchSize int) *LokiHook {
	return &LokiHook{
		buffer:    make([]LogEntry, 0, batchSize),
		batchSize: batchSize,
	}
}

func (hook *LokiHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *LokiHook) Fire(entry *logrus.Entry) error {
	fields, ok := entry.Data["fields"].(LogFields)
	if !ok {
		fields = LogFields{}
	}
	fields.CallIn = fmt.Sprintf("%s:%d", entry.Caller.Function, entry.Caller.Line)

	// Agregar log al buffer
	hook.buffer = append(hook.buffer, LogEntry{
		Message: entry.Message,
		Level:   entry.Level.String(),
		Fields:  fields,
	})

	// Si el buffer llega al tamaño, enviar a Loki
	if len(hook.buffer) >= hook.batchSize {
		hook.Flush()
	}

	return nil
}

// flush manda los logs acumulados a Loki
func (hook *LokiHook) Flush() {
	hook.mu.Lock()
	if len(hook.buffer) == 0 {
		hook.mu.Unlock()
		return
	}
	// Swap para minimizar tiempo con el candado
	batch := hook.buffer
	hook.buffer = make([]LogEntry, 0, hook.batchSize)
	hook.mu.Unlock()

	// 👉 Aquí imprimimos la cantidad de logs que se van a enviar
	fmt.Printf("[LokiHook] Flushing %d logs to Loki\n", len(batch))

	type Stream struct {
		Stream map[string]string `json:"stream"`
		Values [][2]string       `json:"values"`
	}
	type Payload struct {
		Streams []Stream `json:"streams"`
	}

	// Agrupar por labels idénticos (requerido por Loki)
	streamsMap := make(map[string]*Stream)
	now := time.Now().UnixNano()

	for _, log := range batch {
		labels := map[string]string{
			"app":       LoggerConfig.APP_NAME,
			"env":       LoggerConfig.ENVIRONMENT,
			"level":     log.Level,
			"caller_id": log.Fields.CallerID,
			"trace_id":  log.Fields.TraceID,
			"method":    log.Fields.Method,
			"client_ip": log.Fields.ClientIP,
			"user_id":   log.Fields.UserID,
			"path":      log.Fields.Path,
			"call_in":   log.Fields.CallIn,
			"timestamp": time.Now().Format(time.RFC3339),
		}

		keyBytes, _ := json.Marshal(labels)
		key := string(keyBytes)
		if _, ok := streamsMap[key]; !ok {
			streamsMap[key] = &Stream{Stream: labels, Values: make([][2]string, 0)}
		}

		// Usa el mismo "now" para todo el batch (coherencia de timestamps)
		streamsMap[key].Values = append(streamsMap[key].Values, [2]string{
			fmt.Sprintf("%d", now),
			log.Message,
		})
	}

	streams := make([]Stream, 0, len(streamsMap))
	for _, s := range streamsMap {
		streams = append(streams, *s)
	}
	payload := Payload{Streams: streams}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		fmt.Println("Error encoding JSON for Loki payload:", err)
		return
	}

	lokiURL := fmt.Sprintf("%s/loki/api/v1/push", LoggerConfig.LOKI_URL)
	resp, err := http.Post(lokiURL, "application/json", buf)
	if err != nil {
		fmt.Println("Error sending logs to Loki:", err)
		return
	}
	defer resp.Body.Close()

	// Opcional: imprimir status de respuesta de Loki
	fmt.Printf("[LokiHook] Loki responded with status: %s\n", resp.Status)
}

func GetLokiHook() *LokiHook {
	return lokiHook
}
