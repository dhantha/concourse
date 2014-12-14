package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/engine"
	enginefakes "github.com/concourse/atc/engine/fakes"
	"github.com/concourse/turbine"
)

var _ = Describe("Builds API", func() {
	Describe("POST /api/v1/builds", func() {
		var buildPlan engine.BuildPlan

		var response *http.Response

		BeforeEach(func() {
			buildPlan = engine.BuildPlan{
				Config: turbine.Config{
					Run: turbine.RunConfig{
						Path: "ls",
					},
				},
			}
		})

		JustBeforeEach(func() {
			reqPayload, err := json.Marshal(buildPlan)
			Ω(err).ShouldNot(HaveOccurred())

			req, err := http.NewRequest("POST", server.URL+"/api/v1/builds", bytes.NewBuffer(reqPayload))
			Ω(err).ShouldNot(HaveOccurred())

			req.Header.Set("Content-Type", "application/json")

			response, err = client.Do(req)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(true)
			})

			Context("when creating a one-off build succeeds", func() {
				BeforeEach(func() {
					buildsDB.CreateOneOffBuildReturns(db.Build{
						ID:      42,
						Name:    "1",
						JobName: "job1",
						Status:  db.StatusStarted,
					}, nil)
				})

				Context("and building succeeds", func() {
					It("returns 201 Created", func() {
						Ω(response.StatusCode).Should(Equal(http.StatusCreated))
					})

					It("returns the build", func() {
						body, err := ioutil.ReadAll(response.Body)
						Ω(err).ShouldNot(HaveOccurred())

						Ω(body).Should(MatchJSON(`{
							"id": 42,
							"name": "1",
							"job_name": "job1",
							"status": "started",
							"url": "/jobs/job1/builds/1"
						}`))
					})

					It("executes a one-off build", func() {
						Ω(buildsDB.CreateOneOffBuildCallCount()).Should(Equal(1))

						Ω(builder.BuildCallCount()).Should(Equal(1))
						oneOff, plan := builder.BuildArgsForCall(0)
						Ω(oneOff).Should(Equal(db.Build{
							ID:      42,
							Name:    "1",
							JobName: "job1",
							Status:  db.StatusStarted,
						}))
						Ω(plan).Should(Equal(buildPlan))
					})
				})

				Context("and building fails", func() {
					BeforeEach(func() {
						builder.BuildReturns(errors.New("oh no!"))
					})

					It("returns 500 Internal Server Error", func() {
						Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
					})
				})
			})

			Context("when creating a one-off build fails", func() {
				BeforeEach(func() {
					buildsDB.CreateOneOffBuildReturns(db.Build{}, errors.New("oh no!"))
				})

				It("returns 500 Internal Server Error", func() {
					Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
				})
			})
		})

		Context("when not authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(false)
			})

			It("returns 401", func() {
				Ω(response.StatusCode).Should(Equal(http.StatusUnauthorized))
			})

			It("does not trigger a build", func() {
				Ω(buildsDB.CreateOneOffBuildCallCount()).Should(BeZero())
				Ω(builder.BuildCallCount()).Should(BeZero())
			})
		})
	})

	Describe("GET /api/v1/builds", func() {
		var response *http.Response

		JustBeforeEach(func() {
			var err error

			response, err = client.Get(server.URL + "/api/v1/builds")
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when getting all builds succeeds", func() {
			BeforeEach(func() {
				buildsDB.GetAllBuildsReturns([]db.Build{
					{
						ID:      3,
						Name:    "2",
						JobName: "job2",
						Status:  db.StatusStarted,
					},
					{
						ID:      1,
						Name:    "1",
						JobName: "job1",
						Status:  db.StatusSucceeded,
					},
				}, nil)
			})

			It("returns 200 OK", func() {
				Ω(response.StatusCode).Should(Equal(http.StatusOK))
			})

			It("returns all builds", func() {
				body, err := ioutil.ReadAll(response.Body)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(body).Should(MatchJSON(`[
					{
						"id": 3,
						"name": "2",
						"job_name": "job2",
						"status": "started",
						"url": "/jobs/job2/builds/2"
					},
					{
						"id": 1,
						"name": "1",
						"job_name": "job1",
						"status": "succeeded",
						"url": "/jobs/job1/builds/1"
					}
				]`))
			})
		})

		Context("when getting all builds fails", func() {
			BeforeEach(func() {
				buildsDB.GetAllBuildsReturns(nil, errors.New("oh no!"))
			})

			It("returns 500 Internal Server Error", func() {
				Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
			})
		})
	})

	Describe("GET /api/v1/builds/:build_id/events", func() {
		var (
			request  *http.Request
			response *http.Response
		)

		BeforeEach(func() {
			var err error

			buildsDB.GetBuildReturns(db.Build{
				ID:      128,
				JobName: "some-job",
			}, nil)

			request, err = http.NewRequest("GET", server.URL+"/api/v1/builds/128/events", nil)
			Ω(err).ShouldNot(HaveOccurred())
		})

		JustBeforeEach(func() {
			var err error

			response, err = client.Do(request)
			Ω(err).ShouldNot(HaveOccurred())
		})

		Context("when authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(true)
			})

			It("returns 200", func() {
				Ω(response.StatusCode).Should(Equal(200))
			})

			It("serves the request via the event handler with no censor", func() {
				body, err := ioutil.ReadAll(response.Body)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(string(body)).Should(Equal("fake event handler factory was here"))

				Ω(constructedEventHandler.db).Should(Equal(buildsDB))
				Ω(constructedEventHandler.buildID).Should(Equal(128))
				Ω(constructedEventHandler.censor).Should(BeNil())
			})
		})

		Context("when not authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(false)
			})

			Context("and the build is private", func() {
				BeforeEach(func() {
					configDB.GetConfigReturns(atc.Config{
						Jobs: atc.JobConfigs{
							{Name: "some-job", Public: false},
						},
					}, nil)
				})

				It("returns 401", func() {
					Ω(response.StatusCode).Should(Equal(http.StatusUnauthorized))
				})
			})

			Context("and the build is public", func() {
				BeforeEach(func() {
					configDB.GetConfigReturns(atc.Config{
						Jobs: atc.JobConfigs{
							{Name: "some-job", Public: true},
						},
					}, nil)
				})

				It("returns 200", func() {
					Ω(response.StatusCode).Should(Equal(200))
				})

				It("serves the request via the event handler with a censor", func() {
					body, err := ioutil.ReadAll(response.Body)
					Ω(err).ShouldNot(HaveOccurred())

					Ω(string(body)).Should(Equal("fake event handler factory was here"))

					Ω(constructedEventHandler.db).Should(Equal(buildsDB))
					Ω(constructedEventHandler.buildID).Should(Equal(128))
					Ω(constructedEventHandler.censor).ShouldNot(BeNil())
				})
			})
		})
	})

	Describe("POST /api/v1/builds/:build_id/abort", func() {
		var (
			abortTarget *ghttp.Server

			response *http.Response
		)

		BeforeEach(func() {
			abortTarget = ghttp.NewServer()

			abortTarget.AppendHandlers(
				ghttp.VerifyRequest("POST", "/builds/some-guid/abort"),
			)

			buildsDB.GetBuildReturns(db.Build{
				ID:     128,
				Status: db.StatusStarted,
			}, nil)
		})

		JustBeforeEach(func() {
			var err error

			req, err := http.NewRequest("POST", server.URL+"/api/v1/builds/128/abort", nil)
			Ω(err).ShouldNot(HaveOccurred())

			response, err = client.Do(req)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			abortTarget.Close()
		})

		Context("when authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(true)
			})

			Context("when the build can be aborted", func() {
				BeforeEach(func() {
					buildsDB.SaveBuildStatusReturns(nil)
				})

				Context("and the build is started", func() {
					BeforeEach(func() {
						buildsDB.GetBuildReturns(db.Build{
							ID:     128,
							Status: db.StatusStarted,
						}, nil)
					})

					Context("and the engine returns a build", func() {
						var fakeBuild *enginefakes.FakeBuild

						BeforeEach(func() {
							fakeBuild = new(enginefakes.FakeBuild)
							fakeEngine.LookupBuildReturns(fakeBuild, nil)
						})

						It("aborts the build", func() {
							Ω(fakeBuild.AbortCallCount()).Should(Equal(1))
						})

						Context("and aborting succeeds", func() {
							BeforeEach(func() {
								fakeBuild.AbortReturns(nil)
							})

							It("returns 204", func() {
								Ω(response.StatusCode).Should(Equal(http.StatusNoContent))
							})
						})

						Context("and aborting fails", func() {
							BeforeEach(func() {
								fakeBuild.AbortReturns(errors.New("oh no!"))
							})

							It("returns 500", func() {
								Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
							})
						})
					})

					Context("and the engine returns no build", func() {
						BeforeEach(func() {
							fakeEngine.LookupBuildReturns(nil, errors.New("oh no!"))
						})

						It("returns 500", func() {
							Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
						})
					})
				})

				Context("and the status is pending", func() {
					BeforeEach(func() {
						buildsDB.GetBuildReturns(db.Build{
							ID:     128,
							Status: db.StatusPending,
						}, nil)
					})

					It("does not do anything with the engine", func() {
						Ω(fakeEngine.LookupBuildCallCount()).Should(Equal(0))
					})

					It("returns 204", func() {
						Ω(response.StatusCode).Should(Equal(http.StatusNoContent))
					})
				})
			})

			Context("when the build cannot be aborted", func() {
				BeforeEach(func() {
					buildsDB.SaveBuildStatusReturns(errors.New("oh no!"))
				})

				It("returns 500 Internal Server Error", func() {
					Ω(response.StatusCode).Should(Equal(http.StatusInternalServerError))
				})
			})
		})

		Context("when not authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(false)
			})

			It("returns 401", func() {
				Ω(response.StatusCode).Should(Equal(http.StatusUnauthorized))
			})

			It("does not abort the build", func() {
				Ω(abortTarget.ReceivedRequests()).Should(BeEmpty())
			})
		})
	})

	Describe("POST /api/v1/builds/:build_id/hijack", func() {
		var (
			hijackTarget *ghttp.Server

			response *http.Response

			buildHijackConns   <-chan net.Conn
			buildHijackReaders <-chan *gbytes.Buffer

			clientConn   net.Conn
			clientReader io.Reader
		)

		BeforeEach(func() {
			hijackedConns := make(chan net.Conn, 1)
			buildHijackConns = hijackedConns

			hijackedReaders := make(chan *gbytes.Buffer, 1)
			buildHijackReaders = hijackedReaders

			hijackTarget = ghttp.NewServer()
			hijackTarget.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("POST", "/builds/some-guid/hijack"),
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)

					var msg json.RawMessage
					err := json.NewDecoder(r.Body).Decode(&msg)
					Ω(err).ShouldNot(HaveOccurred())

					Ω(string(msg)).Should(Equal(string(`{"some":"hijack-body"}`)))

					conn, br, err := w.(http.Hijacker).Hijack()
					Ω(err).ShouldNot(HaveOccurred())

					defer conn.Close()

					buf := gbytes.NewBuffer()

					hijackedConns <- conn
					hijackedReaders <- buf

					io.Copy(buf, br)
				},
			))
		})

		JustBeforeEach(func() {
			var err error

			hijackReq, err := http.NewRequest(
				"POST",
				server.URL+"/api/v1/builds/128/hijack",
				bytes.NewBufferString(`{"some":"hijack-body"}`),
			)
			Ω(err).ShouldNot(HaveOccurred())

			conn, err := net.Dial("tcp", server.Listener.Addr().String())
			Ω(err).ShouldNot(HaveOccurred())

			client := httputil.NewClientConn(conn, nil)

			response, err = client.Do(hijackReq)
			Ω(err).ShouldNot(HaveOccurred())

			clientConn, clientReader = client.Hijack()
		})

		AfterEach(func() {
			clientConn.Close()
			hijackTarget.Close()
		})

		Context("when authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(true)
			})

			Context("when the build can be found", func() {
				Context("and it has a hijack URL", func() {
					BeforeEach(func() {
						metadata := engine.TurbineMetadata{
							Guid:     "some-guid",
							Endpoint: hijackTarget.URL(),
						}

						metadataPayload, err := json.Marshal(metadata)
						Ω(err).ShouldNot(HaveOccurred())

						buildsDB.GetBuildReturns(db.Build{
							ID:             128,
							Engine:         "turbine",
							EngineMetadata: string(metadataPayload),
						}, nil)
					})

					It("proxies all traffic via a hijacked connection", func() {
						var serverReceivedBuf *gbytes.Buffer
						Eventually(buildHijackReaders).Should(Receive(&serverReceivedBuf))

						var serverConnectedConn net.Conn
						Eventually(buildHijackConns).Should(Receive(&serverConnectedConn))

						clientReceivedBuf := gbytes.NewBuffer()

						readingFromServer := new(sync.WaitGroup)
						readingFromServer.Add(1)
						go func() {
							io.Copy(clientReceivedBuf, clientReader)
							readingFromServer.Done()
						}()

						_, err := clientConn.Write([]byte("hello from client"))
						Ω(err).ShouldNot(HaveOccurred())

						Eventually(serverReceivedBuf).Should(gbytes.Say("hello from client"))

						_, err = serverConnectedConn.Write([]byte("hello from server"))
						Ω(err).ShouldNot(HaveOccurred())

						err = serverConnectedConn.Close()
						Ω(err).ShouldNot(HaveOccurred())

						readingFromServer.Wait()

						Eventually(clientReceivedBuf).Should(gbytes.Say("hello from server"))
					})
				})

				Context("but it does not have a hijack URL", func() {
					BeforeEach(func() {
						buildsDB.GetBuildReturns(db.Build{ID: 128}, nil)
					})

					It("returns 400 Bad Request", func() {
						Ω(response.StatusCode).Should(Equal(http.StatusBadRequest))
					})
				})
			})

			Context("when the build cannot be found", func() {
				BeforeEach(func() {
					buildsDB.GetBuildReturns(db.Build{}, errors.New("oh no!"))
				})

				It("returns 404 Not Found", func() {
					Ω(response.StatusCode).Should(Equal(http.StatusNotFound))
				})
			})
		})

		Context("when not authenticated", func() {
			BeforeEach(func() {
				authValidator.IsAuthenticatedReturns(false)
			})

			It("returns 401", func() {
				Ω(response.StatusCode).Should(Equal(http.StatusUnauthorized))
			})

			It("does not hijack the build", func() {
				Ω(hijackTarget.ReceivedRequests()).Should(BeEmpty())
			})
		})
	})
})
