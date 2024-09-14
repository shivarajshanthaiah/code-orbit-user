package config

import (
	"errors"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioService struct {
	Client *twilio.RestClient
	Cfg    *Config
}

func SetupTwilio(cfg *Config) *TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.SID,
		Password: cfg.TOKEN,
	})
	return &TwilioService{
		Client: client,
		Cfg:    cfg,
	}
}

func (t *TwilioService) SendTwilioOTP(phone string) (*verify.VerifyV2Verification, error) {
	params := &verify.CreateVerificationParams{}
	params.SetTo("+918762334325")
	params.SetChannel("sms")
	res, err := t.Client.VerifyV2.CreateVerification(t.Cfg.SERVICETOKEN, params)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (t *TwilioService) VerifyTwilioOTP(phone, otp string) error {
	params := verify.CreateVerificationCheckParams{}
	params.SetTo(t.Cfg.Phone)
	params.SetCode(otp)

	resp, err := t.Client.VerifyV2.CreateVerificationCheck(t.Cfg.SERVICETOKEN, &params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}
	return errors.New("incorrect otp")
}
