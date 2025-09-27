package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomResponseWriter intercepta el output
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // captura la respuesta
	return w.ResponseWriter.Write(b)
}

func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// --- 1. Lee y clona el body del request ---
		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = bodyBytes
				// Reemplaza el body para que Gin pueda leerlo después
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// --- 2. Captura la respuesta ---
		bodyWriter := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		// --- 3. Continúa con el handler ---
		c.Next()

		// --- 4. Formatea JSON request y response (si aplicable) ---
		prettyReq := formatJSON(requestBody)
		prettyResp := formatJSON(bodyWriter.body.Bytes())

		// --- 5. Obtiene headers y query params ---
		headersJSON, _ := json.MarshalIndent(c.Request.Header, "", "  ")
		queryJSON, _ := json.MarshalIndent(url.Values(c.Request.URL.Query()), "", "  ")

		// --- 6. Logging final ---
		duration := time.Since(start)
		log.Printf(`[GIN] %s %s | %d | %s | IP: %s
Headers: %s
Query: %s
Request Body: %s
Response Body: %s`,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
			c.ClientIP(),
			string(headersJSON),
			string(queryJSON),
			prettyReq,
			prettyResp,
		)
	}
}

// formatJSON intenta indentar un payload JSON, o devuelve el raw si no es válido JSON
func formatJSON(payload []byte) string {
	if len(payload) == 0 {
		return ""
	}
	var out bytes.Buffer
	if err := json.Indent(&out, payload, "", "  "); err != nil {
		// No es JSON válido o indent falló: devolver raw
		return string(payload)
	}
	return out.String()
}
