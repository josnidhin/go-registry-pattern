/*.
 * @author Jose Nidhin
 */
package providerregistry

import (
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/josnidhin/go-registry-pattern/domain"
	"github.com/josnidhin/go-registry-pattern/logger"
)

var registerMutex sync.Mutex

var registry = make(map[domain.ProviderName]Provider)

type Provider interface {
	Name() domain.ProviderName
	Dispatch(domain.Transaction) error
}

func Register(name domain.ProviderName, provider Provider) {
	registerMutex.Lock()
	defer registerMutex.Unlock()

	_, ok := registry[name]
	if ok {
		logger.DefaultLogger.Fatal("Provider with same name already registered",
			zap.Any("providerName", name))
	}

	logger.DefaultLogger.Debug("New provider registered",
		zap.Any("providerName", name))
	registry[name] = provider
}

func GetProvider(name domain.ProviderName) (p Provider, err error) {
	p, ok := registry[name]
	if !ok {
		err = fmt.Errorf("provider %q is not registered", name)
	}

	return
}
