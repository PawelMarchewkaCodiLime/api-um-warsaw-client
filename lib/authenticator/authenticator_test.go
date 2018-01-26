package authenticator_test

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/kaweue/api-um-warsaw-client/lib/authenticator"
)

var _ = Describe("Authenticator", func() {
	var auth *Authenticator
	var urlValues url.Values
	apiKey := "xxxx-xxxx-xxxx-xxxx"

	JustBeforeEach(func() {
		auth = NewAuthenticator(apiKey)
		urlValues = make(url.Values)
	})

	Describe("checking authentication method", func() {
		It("should add api-key to URL values", func() {
			auth.Authenticate(&urlValues)
			
			Expect(urlValues.Get("apikey")).To(Equal(apiKey))
		})
	})

})
