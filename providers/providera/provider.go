/**
 * @author Jose Nidhin
 */
package providera

import (
	"go.uber.org/zap"

	"github.com/josnidhin/go-registry-pattern/domain"
	"github.com/josnidhin/go-registry-pattern/logger"
	"github.com/josnidhin/go-registry-pattern/providerregistry"
)

type Provider struct {
	name   domain.ProviderName
	logger *zap.Logger
}

func (p Provider) Name() domain.ProviderName {
	return p.name
}

func (p Provider) Dispatch(trx domain.Transaction) error {
	p.logger.Info("Dispatch transaction",
		zap.Int("id", trx.Id))

	return nil
}

func init() {
	var name domain.ProviderName = "Provider A"
	p := Provider{
		name:   name,
		logger: logger.DefaultLogger.With(zap.Any("providerName", name)),
	}

	providerregistry.Register(name, p)
}
