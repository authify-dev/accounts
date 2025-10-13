package utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	NUM_MODULE_BITS = 8
	NUM_ENTITY_BITS = 8
)

const (
	NUM_MODULE_HEX  = NUM_MODULE_BITS / 4 // 4 bits = 1 hex
	NUM_ENTITY_HEX  = NUM_ENTITY_BITS / 4
	NUM_TOTAL_HEX   = NUM_MODULE_HEX + NUM_ENTITY_HEX
	NUM_TOTAL_BYTES = NUM_TOTAL_HEX / 2
)

// ────────────────────────────────────────────────
// Tipos semánticos
// ────────────────────────────────────────────────

type ModuleCode string

func (m ModuleCode) String() string { return string(m) }

type EntityCode string

func (e EntityCode) String() string { return string(e) }

// ────────────────────────────────────────────────
// Interfaz para un UUID extendido
// ────────────────────────────────────────────────

type TypedUUID4 interface {
	Module() ModuleCode
	Entity() EntityCode
	UUID() uuid.UUID
}

// ────────────────────────────────────────────────
// Generación y parsing de IDs
// ────────────────────────────────────────────────

func MakeTypedUUID4(moduleHex, entityHex string) (uuid.UUID, error) {
	moduleHex = strings.ToLower(strings.TrimSpace(moduleHex))
	entityHex = strings.ToLower(strings.TrimSpace(entityHex))

	if len(moduleHex) != NUM_MODULE_HEX || len(entityHex) != NUM_ENTITY_HEX {
		return uuid.Nil, fmt.Errorf("module must be %d hex chars, entity must be %d hex chars", NUM_MODULE_HEX, NUM_ENTITY_HEX)
	}

	// Concatenamos (ej: 8 hex = 4 bytes si 16+16 bits)
	suffix := moduleHex + entityHex
	pb, err := hex.DecodeString(suffix)
	if err != nil || len(pb) != NUM_TOTAL_BYTES {
		return uuid.Nil, fmt.Errorf("invalid suffix hex: %w", err)
	}

	u := uuid.New()

	// Sobrescribimos últimos bytes (ej: bytes 12..15 si 4 bytes)
	copy(u[16-NUM_TOTAL_BYTES:], pb)

	// Asegurar versión/variante
	u[6] = (u[6] & 0x0F) | 0x40
	u[8] = (u[8] & 0x3F) | 0x80

	return u, nil
}

func ExtractModuleEntity(id uuid.UUID) (ModuleCode, EntityCode) {
	// últimos N bytes = N*2 hex
	h := fmt.Sprintf("%x", id[16-NUM_TOTAL_BYTES:])
	return ModuleCode(h[:NUM_MODULE_HEX]), EntityCode(h[NUM_MODULE_HEX:NUM_TOTAL_HEX])
}

// ────────────────────────────────────────────────
// Ejemplo de implementación de TypedUUID4
// ────────────────────────────────────────────────

type DomainID struct {
	raw    uuid.UUID
	module ModuleCode
	entity EntityCode
}

func (d DomainID) Module() ModuleCode { return d.module }
func (d DomainID) Entity() EntityCode { return d.entity }
func (d DomainID) UUID() uuid.UUID    { return d.raw }

func NewUUIDx(moduleHex, entityHex string) (DomainID, error) {

	numCharsModuleExpected := NUM_MODULE_BITS / 4
	numCharsEntityExpected := NUM_ENTITY_BITS / 4

	if len(moduleHex) != numCharsModuleExpected {
		return DomainID{}, fmt.Errorf("module must be %d hex chars", numCharsModuleExpected)
	}

	if len(entityHex) != numCharsEntityExpected {
		return DomainID{}, fmt.Errorf("entity must be %d hex chars", numCharsEntityExpected)
	}

	u, err := MakeTypedUUID4(moduleHex, entityHex)
	if err != nil {
		return DomainID{}, err
	}
	m, e := ExtractModuleEntity(u)
	return DomainID{raw: u, module: m, entity: e}, nil
}
