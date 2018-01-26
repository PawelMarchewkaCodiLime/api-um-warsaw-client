package umwarsawclient_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/kaweue/api-um-warsaw-client/lib/authenticator"
	"github.com/kaweue/api-um-warsaw-client/lib/client"
	"github.com/kaweue/api-um-warsaw-client/lib/types"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type httpHandler struct {
	query    url.Values
	result   int
	response []byte
}

const apiKey = "xxxx-xxxx-xxxx-xxxx"

func (h *httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.query = req.URL.Query()
	res.WriteHeader(h.result)
	if h.result == http.StatusOK {
		res.Write([]byte(h.response))
	}
}

var _ = Describe("Ztm", func() {
	var client *umwarsawclient.Client
	var handler *httpHandler

	BeforeEach(func() {
		handler = &httpHandler{}
		dummyServer := httptest.NewServer(handler)
		client = umwarsawclient.NewClient(dummyServer.URL, authenticator.NewAuthenticator(apiKey),
			dummyServer.Client())
	})

	Describe("GetBusStop method", func() {
		Context("when server returns status ok and correct response", func() {
			JustBeforeEach(func() {
				handler.result = http.StatusOK
				handler.response, _ = ioutil.ReadFile("testData/getBusStop.json")
			})

			It("should create correct query", func() {
				_, err := client.GetBusStop("znana")
				Expect(err).To(BeNil())
				Expect(handler.query.Get("apikey")).To(Equal(apiKey))
				Expect(handler.query.Get("name")).To(Equal("znana"))
				Expect(handler.query.Get("id")).To(Equal(umwarsawclient.BusStopRequestId))
			})

			It("should return correct bus stop", func() {
				busStop, err := client.GetBusStop("znana")
				Expect(err).To(BeNil())
				Expect(busStop.BusID).To(Equal("5104"))
				Expect(busStop.Name).To(Equal("Znana"))
			})
		})

		Context("when server returns status not ok", func() {
			JustBeforeEach(func() {
				handler.result = http.StatusNotFound
			})
			It("returns error", func() {
				_, err := client.GetBusStop("znana")
				Expect(err).ToNot(BeNil())
			})
		})
	})

	Describe("GetLinesOnBusStop method", func() {
		Context("when server returns status ok and correct response", func() {
			JustBeforeEach(func() {
				handler.result = http.StatusOK
				handler.response, _ = ioutil.ReadFile("testData/getLinesResponse.json")
			})

			It("should create correct query", func() {
				_, err := client.GetLinesOnBusStop("5104", "01")
				Expect(err).To(BeNil())
				Expect(handler.query.Get("apikey")).To(Equal(apiKey))
				Expect(handler.query.Get("busstopId")).To(Equal("5104"))
				Expect(handler.query.Get("busstopNr")).To(Equal("01"))
				Expect(handler.query.Get("id")).To(Equal(umwarsawclient.LinesOnBusStopRequestId))
			})

			It("should return correct bus stop", func() {
				lines, err := client.GetLinesOnBusStop("5104", "01")
				Expect(err).To(BeNil())
				Expect(len(lines)).To(Equal(3))
				Expect(lines).Should(ContainElement(types.Line("167")))
				Expect(lines).Should(ContainElement(types.Line("155")))
				Expect(lines).Should(ContainElement(types.Line("129")))
			})
		})

		Context("when server returns status not ok", func() {
			JustBeforeEach(func() {
				handler.result = http.StatusNotFound
			})
			It("returns error", func() {
				_, err := client.GetLinesOnBusStop("5104", "01")
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
