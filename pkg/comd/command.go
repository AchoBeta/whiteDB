package comd

const (
	NONE int = iota
	SET
	REMOVE
	GET
	COMPACT
	LEN
	EXIT
)

var commandMap map[string]int = map[string]int{
	"SET":     SET,
	"REMOVE":  REMOVE,
	"GET":     GET,
	"RM":      REMOVE,
	"EXIT":    EXIT,
	"LEN":     LEN,
	"COMPACT": COMPACT,
}
