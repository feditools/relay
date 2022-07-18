package logic

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/feditools/relay/internal/models"
	"net/url"
)

func (l *Logic) getHostMeta(ctx context.Context, domain string) (*models.HostMeta, error) {
	log := logger.WithField("func", "getHostMeta")
	hostMetaURI := &url.URL{
		Scheme: "https",
		Host:   domain,
		Path:   "/.well-known/host-meta",
	}

	v, err, _ := l.outgoingRequestGroup.Do(fmt.Sprintf("hostmeta-%s", hostMetaURI.String()), func() (interface{}, error) {
		// do request
		resp, err := l.transport.InstanceGet(ctx, hostMetaURI)
		if err != nil {
			log.Errorf("http get: %s", err.Error())

			return nil, err
		}

		hostMeta := new(models.HostMeta)
		err = xml.Unmarshal(resp, hostMeta)
		if err != nil {
			log.Errorf("decode xml: %s", err.Error())

			return nil, err
		}

		return hostMeta, nil
	})

	if err != nil {
		log.Errorf("singleflight: %s", err.Error())

		return nil, NewErrorf("singleflight: %s", err.Error())
	}

	hostMeta, ok := v.(*models.HostMeta)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return hostMeta, nil
}
