/**
 * @author Jose Nidhin
 */
package providerregistry

import (
	"github.com/josnidhin/go-registry-pattern/domain"
)

type ProviderName string

type Provider interface {
	Name() ProviderName
	Dispatch(domain.Transaction) error
}
