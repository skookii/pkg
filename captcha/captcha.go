package captcha

import (
	"context"

	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1"
	"cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"google.golang.org/api/option"
)

type defSettings struct {
	projectId  string
	siteKey    string
	credential option.ClientOption
}

var reCaptcha *defSettings

func Config(projectid, sitekey string, captchaByte []byte) {
	reCaptcha = &defSettings{
		projectId:  projectid,
		siteKey:    sitekey,
		credential: option.WithCredentialsJSON(captchaByte),
	}
}

// https://blog.canopas.com/integrate-google-recaptcha-enterprise-using-vue-js-and-golang-60a9335e80ac
func (d *defSettings) check(ctx context.Context, token string) (bool, error) {
	c, err := recaptchaenterprise.NewClient(ctx, d.credential)
	if err != nil {
		return false, err
	}
	defer c.Close()
	req := &recaptchaenterprisepb.CreateAssessmentRequest{
		Assessment: &recaptchaenterprisepb.Assessment{
			Event: &recaptchaenterprisepb.Event{
				Token:   token,
				SiteKey: d.siteKey,
			},
		},
		Parent: "projects/" + d.projectId,
	}
	resp, err := c.CreateAssessment(ctx, req)
	if err != nil {
		return false, err
	}

	if resp.TokenProperties.Action == "verify" && resp.TokenProperties.Valid && resp.RiskAnalysis.Score >= 0.9 {
		return true, nil
	}
	return false, nil
}

func CheckCaptcha(ctx context.Context, token string) (bool, error) {
	return reCaptcha.check(ctx, token)
}
