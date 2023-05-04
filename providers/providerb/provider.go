/**
 * @author Jose Nidhin
 */
package providerb

import (
	"go.uber.org/zap"

	"github.com/josnidhin/go-registry-pattern/domain"
	"github.com/josnidhin/go-registry-pattern/logger"
	"github.com/josnidhin/go-registry-pattern/providerregistry"
)

type Provider struct {
	name   providerregistry.ProviderName
	logger *zap.Logger
}

func (p Provider) Name() providerregistry.ProviderName {
	return p.name
}

func (p Provider) Dispatch(trx domain.Transaction) error {
	p.logger.Info("Dispatch transaction",
		zap.Int("id", trx.Id))

	return nil
}

func init() {
	var name providerregistry.ProviderName = "Provider B"

	p := Provider{
		name:   name,
		logger: logger.DefaultLogger.With(zap.Any("providerName", name)),
	}

	providerregistry.Register(name, p)
}
