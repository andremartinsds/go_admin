package dummy

import (
	"strconv"

	faker "github.com/andremartinsds/faker-br/pkg"
	"github.com/andremartinsds/go_admin/internal/entities"
	"github.com/andremartinsds/go_admin/pkg"
)

func AddressDummy() *entities.Address {
	return &entities.Address{
		ID:             pkg.NewUUID(),
		ZipCode:        faker.G.Cep(),
		State:          faker.G.Estado(),
		City:           faker.G.Cidade(),
		Description:    "other description prop",
		Number:         strconv.Itoa(faker.G.Numero(2)),
		Complement:     faker.G.Complemento(),
		Neighborhood:   faker.G.Bairro(),
		ReferencePoint: "any valid address",
		Observation:    faker.G.Observacao(),
	}
}
