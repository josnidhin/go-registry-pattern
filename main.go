/**
 * @author Jose Nidhin
 */
package main

import (
	"math/rand"
	"time"

	"go.uber.org/zap"

	"github.com/josnidhin/go-registry-pattern/domain"
	"github.com/josnidhin/go-registry-pattern/logger"
	"github.com/josnidhin/go-registry-pattern/providerregistry"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rsize := 1000

	cases := []struct {
		ProviderName domain.ProviderName
		Transaction  domain.Transaction
	}{
		{
			ProviderName: domain.ProviderName("Provider A"),
			Transaction: domain.Transaction{
				Id: r.Intn(rsize),
			},
		},
		{
			ProviderName: domain.ProviderName("Provider B"),
			Transaction: domain.Transaction{
				Id: r.Intn(rsize),
			},
		},
		{
			ProviderName: domain.ProviderName("Provider C"),
			Transaction: domain.Transaction{
				Id: r.Intn(rsize),
			},
		},
	}

	for _, c := range cases {
		provider, err := providerregistry.GetProvider(c.ProviderName)
		if err != nil {
			logger.DefaultLogger.Error("Provider not found", zap.Error(err))
			continue
		}

		provider.Dispatch(c.Transaction)
	}
}
