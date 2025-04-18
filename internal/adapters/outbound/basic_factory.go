package outbound

import (
	"github.com/pgnahuel/go6arc/internal/adapters/outbound/mysql"
	"github.com/pgnahuel/go6arc/internal/ports"
)

// Constructores para cada port
func NewBasicCreator() ports.BasicCreator {
	return mysql.NewBasicRepository()
}

func NewBasicReader() ports.BasicReader {
	return mysql.NewBasicRepository()
}

func NewBasicEliminator() ports.BasicEliminator {
	return mysql.NewBasicRepository()
}

func NewBasicRemover() ports.BasicRemover {
	return mysql.NewBasicRepository()
}
