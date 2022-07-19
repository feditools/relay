package logic1

// webFinger retrieves web finger resource from a federated instance.
/*func (l *Logic) fetchWebFinger(ctx context.Context, wfURI models.WebfingerURI, username, domain string) (*models.WebFinger, error) {
	log := logger.WithField("func", "webFinger")
	webFingerString := fmt.Sprintf(wfURI.FTemplate(), username, domain)
	webFingerURI, err := url.Parse(webFingerString)
	if err != nil {
		log.Errorf("parsing url '%s': %s", webFingerString, err.Error())

		return nil, err
	}

	v, err, _ := l.outgoingRequestGroup.Do(fmt.Sprintf("webfinger-%s", webFingerURI.String()), func() (interface{}, error) {
		log.Tracef("webfingering %s", webFingerURI.String())

		// do request
		body, err := l.transport.InstanceGet(ctx, webFingerURI)
		if err != nil {
			log.Errorf("http get: %s", err.Error())

			return nil, err
		}

		var webFinger models.WebFinger
		err = json.Unmarshal(body, &webFinger)
		if err != nil {
			log.Errorf("decode json: %s", err.Error())

			return nil, err
		}

		return webFinger, nil
	})

	if err != nil {
		log.Errorf("singleflight: %s", err.Error())

		return nil, err
	}

	webFinger, ok := v.(models.WebFinger)
	if !ok {
		return nil, NewError("invalid response type from single flight")
	}

	return &webFinger, nil
}*/
