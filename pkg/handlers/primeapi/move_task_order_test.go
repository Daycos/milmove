package primeapi

import (
	"fmt"
	"net/http/httptest"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/pkg/apperror"
	"github.com/transcom/mymove/pkg/etag"
	"github.com/transcom/mymove/pkg/factory"
	movetaskorderops "github.com/transcom/mymove/pkg/gen/primeapi/primeoperations/move_task_order"
	"github.com/transcom/mymove/pkg/gen/primemessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/services/fetch"
	"github.com/transcom/mymove/pkg/services/mocks"
	moverouter "github.com/transcom/mymove/pkg/services/move"
	movetaskorder "github.com/transcom/mymove/pkg/services/move_task_order"
	mtoserviceitem "github.com/transcom/mymove/pkg/services/mto_service_item"
	"github.com/transcom/mymove/pkg/services/query"
	"github.com/transcom/mymove/pkg/services/upload"
	storageTest "github.com/transcom/mymove/pkg/storage/test"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *HandlerSuite) TestListMovesHandlerReturnsUpdated() {
	now := time.Now()
	lastFetch := now.Add(-time.Second)

	move := testdatagen.MakeAvailableMove(suite.DB())

	// this move should not be returned
	olderMove := testdatagen.MakeAvailableMove(suite.DB())

	// Pop will overwrite UpdatedAt when saving a model, so use SQL to set it in the past
	suite.Require().NoError(suite.DB().RawQuery("UPDATE moves SET updated_at=? WHERE id=?",
		now.Add(-2*time.Second), olderMove.ID).Exec())
	suite.Require().NoError(suite.DB().RawQuery("UPDATE orders SET updated_at=$1 WHERE id=$2;",
		now.Add(-10*time.Second), olderMove.OrdersID).Exec())

	since := handlers.FmtDateTime(lastFetch)
	request := httptest.NewRequest("GET", fmt.Sprintf("/moves?since=%s", since.String()), nil)
	params := movetaskorderops.ListMovesParams{HTTPRequest: request, Since: since}
	handlerConfig := suite.HandlerConfig()

	// Validate incoming payload: no body to validate

	// make the request
	handler := ListMovesHandler{HandlerConfig: handlerConfig, MoveTaskOrderFetcher: movetaskorder.NewMoveTaskOrderFetcher()}
	response := handler.Handle(params)

	suite.IsNotErrResponse(response)
	listMovesResponse := response.(*movetaskorderops.ListMovesOK)
	movesList := listMovesResponse.Payload

	// Validate outgoing payload
	suite.NoError(movesList.Validate(strfmt.Default))

	suite.Equal(1, len(movesList))
	suite.Equal(move.ID.String(), movesList[0].ID.String())
}

func (suite *HandlerSuite) TestGetMoveTaskOrder() {
	request := httptest.NewRequest("GET", "/move-task-orders/{moveTaskOrderID}", nil)

	suite.Run("Success with Prime-available move by ID", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}

		successMove := testdatagen.MakeAvailableMove(suite.DB())
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      successMove.ID.String(),
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Equal(movePayload.ID.String(), successMove.ID.String())
		suite.NotNil(movePayload.AvailableToPrimeAt)
		suite.NotEmpty(movePayload.AvailableToPrimeAt) // checks that the date is not 0001-01-01
	})

	suite.Run("Success with Prime-available move by Locator", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		successMove := testdatagen.MakeAvailableMove(suite.DB())
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      successMove.Locator,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Equal(movePayload.ID.String(), successMove.ID.String())
		suite.NotNil(movePayload.AvailableToPrimeAt)
		suite.NotEmpty(movePayload.AvailableToPrimeAt) // checks that the date is not 0001-01-01
	})

	suite.Run("Returns the destination address type for a shipment on a move if it exists", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		successMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				AvailableToPrimeAt: swag.Time(time.Now()),
				Status:             models.MoveStatusAPPROVED,
			},
		})
		destinationAddress := factory.BuildAddress(suite.DB(), nil, nil)
		destinationType := models.DestinationTypeHomeOfRecord
		successShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			MTOShipment: models.MTOShipment{
				MoveTaskOrderID:      successMove.ID,
				DestinationAddressID: &destinationAddress.ID,
				DestinationType:      &destinationType,
				Status:               models.MTOShipmentStatusApproved,
			},
		})
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      successMove.Locator,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Equal(movePayload.ID.String(), successMove.ID.String())
		suite.NotNil(movePayload.AvailableToPrimeAt)
		suite.NotEmpty(movePayload.AvailableToPrimeAt) // checks that the date is not 0001-01-01

		// check for the destination address type
		suite.Equal(string(*successShipment.DestinationType), string(*movePayload.MtoShipments[0].DestinationType))

	})

	suite.Run("Success returns reweighs on shipments if they exist", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		successMove := testdatagen.MakeAvailableMove(suite.DB())
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      successMove.Locator,
		}

		reweigh := testdatagen.MakeReweigh(suite.DB(), testdatagen.Assertions{
			Move: successMove,
		})

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		reweighPayload := movePayload.MtoShipments[0].Reweigh
		suite.Equal(movePayload.ID.String(), successMove.ID.String())
		suite.NotNil(movePayload.AvailableToPrimeAt)
		suite.NotEmpty(movePayload.AvailableToPrimeAt)
		suite.Equal(strfmt.UUID(reweigh.ID.String()), reweighPayload.ID)
	})

	suite.Run("Success - returns sit extensions on shipments if they exist", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		successMove := testdatagen.MakeAvailableMove(suite.DB())
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      successMove.Locator,
		}

		sitExtension := testdatagen.MakeSITExtension(suite.DB(), testdatagen.Assertions{
			Move: successMove,
			MTOShipment: models.MTOShipment{
				Status: models.MTOShipmentStatusApproved,
			},
		})

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		reweighPayload := movePayload.MtoShipments[0].SitExtensions[0]
		suite.Equal(successMove.ID.String(), movePayload.ID.String())
		suite.Equal(strfmt.UUID(sitExtension.ID.String()), reweighPayload.ID)
	})

	suite.Run("Success - filters shipments handled by an external vendor", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		move := testdatagen.MakeAvailableMove(suite.DB())

		// Create two shipments, one prime, one external.  Only prime one should be returned.
		primeShipment := testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: move,
			MTOShipment: models.MTOShipment{
				UsesExternalVendor: false,
			},
		})
		testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: move,
			MTOShipment: models.MTOShipment{
				ShipmentType:       models.MTOShipmentTypeHHGOutOfNTSDom,
				UsesExternalVendor: true,
			},
		})

		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      move.Locator,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Equal(move.ID.String(), movePayload.ID.String())
		if suite.Len(movePayload.MtoShipments, 1) {
			suite.Equal(primeShipment.ID.String(), movePayload.MtoShipments[0].ID.String())
		}
	})

	suite.Run("Success - returns shipment with attached PpmShipment", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		move := testdatagen.MakeAvailableMove(suite.DB())
		ppmShipment := testdatagen.MakePPMShipment(suite.DB(), testdatagen.Assertions{
			Move: move,
		})

		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      move.Locator,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderOK{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderOK)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Equal(move.ID.String(), movePayload.ID.String())
		suite.NotNil(movePayload.MtoShipments[0].PpmShipment)
		suite.Equal(ppmShipment.ShipmentID.String(), movePayload.MtoShipments[0].PpmShipment.ShipmentID.String())
		suite.Equal(ppmShipment.ID.String(), movePayload.MtoShipments[0].PpmShipment.ID.String())
	})

	suite.Run("Failure 'Not Found' for non-available move", func() {
		handler := GetMoveTaskOrderHandler{
			suite.HandlerConfig(),
			movetaskorder.NewMoveTaskOrderFetcher(),
		}
		failureMove := testdatagen.MakeDefaultMove(suite.DB()) // default is not available to Prime
		params := movetaskorderops.GetMoveTaskOrderParams{
			HTTPRequest: request,
			MoveID:      failureMove.ID.String(),
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsNotErrResponse(response)
		suite.IsType(&movetaskorderops.GetMoveTaskOrderNotFound{}, response)

		moveResponse := response.(*movetaskorderops.GetMoveTaskOrderNotFound)
		movePayload := moveResponse.Payload

		// Validate outgoing payload
		suite.NoError(movePayload.Validate(strfmt.Default))

		suite.Contains(*movePayload.Detail, failureMove.ID.String())
	})
}

func (suite *HandlerSuite) TestCreateExcessWeightRecord() {
	request := httptest.NewRequest("POST", "/move-task-orders/{moveTaskOrderID}", nil)
	fakeS3 := storageTest.NewFakeS3Storage(true)

	suite.Run("Success - Created an excess weight record", func() {
		handlerConfig := suite.HandlerConfig()
		handlerConfig.SetFileStorer(fakeS3)
		handler := CreateExcessWeightRecordHandler{
			handlerConfig,
			// Must use the Prime service object in particular:
			moverouter.NewPrimeMoveExcessWeightUploader(upload.NewUploadCreator(fakeS3)),
		}

		now := time.Now()
		availableMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
			Move: models.Move{
				AvailableToPrimeAt:      &now,
				Status:                  models.MoveStatusAPPROVED,
				ExcessWeightQualifiedAt: &now,
			},
		})

		params := movetaskorderops.CreateExcessWeightRecordParams{
			HTTPRequest:     request,
			File:            suite.Fixture("test.pdf"),
			MoveTaskOrderID: strfmt.UUID(availableMove.ID.String()),
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.Require().IsType(&movetaskorderops.CreateExcessWeightRecordCreated{}, response)

		okResponse := response.(*movetaskorderops.CreateExcessWeightRecordCreated)

		// Validate outgoing payload
		suite.NoError(okResponse.Payload.Validate(strfmt.Default))

		suite.Equal(availableMove.ID.String(), okResponse.Payload.MoveID.String())
		suite.NotNil(okResponse.Payload.MoveExcessWeightQualifiedAt)
		suite.Equal(okResponse.Payload.MoveExcessWeightQualifiedAt.String(), strfmt.DateTime(*availableMove.ExcessWeightQualifiedAt).String())
		suite.NotEmpty(okResponse.Payload.ID)
	})

	suite.Run("Fail - Move not found - 404", func() {
		handlerConfig := suite.HandlerConfig()
		handlerConfig.SetFileStorer(fakeS3)
		handler := CreateExcessWeightRecordHandler{
			handlerConfig,
			// Must use the Prime service object in particular:
			moverouter.NewPrimeMoveExcessWeightUploader(upload.NewUploadCreator(fakeS3)),
		}

		params := movetaskorderops.CreateExcessWeightRecordParams{
			HTTPRequest:     request,
			File:            suite.Fixture("test.pdf"),
			MoveTaskOrderID: strfmt.UUID("00000000-0000-0000-0000-000000000123"),
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.Require().IsType(&movetaskorderops.CreateExcessWeightRecordNotFound{}, response)
		notFoundResponse := response.(*movetaskorderops.CreateExcessWeightRecordNotFound)

		// Validate outgoing payload
		suite.NoError(notFoundResponse.Payload.Validate(strfmt.Default))

		suite.Require().NotNil(notFoundResponse.Payload.Detail)
		suite.Contains(*notFoundResponse.Payload.Detail, params.MoveTaskOrderID.String())
	})

	suite.Run("Fail - Move not Prime-available - 404", func() {
		handlerConfig := suite.HandlerConfig()
		handlerConfig.SetFileStorer(fakeS3)
		handler := CreateExcessWeightRecordHandler{
			handlerConfig,
			// Must use the Prime service object in particular:
			moverouter.NewPrimeMoveExcessWeightUploader(upload.NewUploadCreator(fakeS3)),
		}

		unavailableMove := testdatagen.MakeDefaultMove(suite.DB()) // default move is not available to Prime
		params := movetaskorderops.CreateExcessWeightRecordParams{
			HTTPRequest:     request,
			File:            suite.Fixture("test.pdf"),
			MoveTaskOrderID: strfmt.UUID(unavailableMove.ID.String()),
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.Require().IsType(&movetaskorderops.CreateExcessWeightRecordNotFound{}, response)
		notFoundResponse := response.(*movetaskorderops.CreateExcessWeightRecordNotFound)

		// Validate outgoing payload
		suite.NoError(notFoundResponse.Payload.Validate(strfmt.Default))

		suite.Require().NotNil(notFoundResponse.Payload.Detail)
		suite.Contains(*notFoundResponse.Payload.Detail, unavailableMove.ID.String())
	})
}

func (suite *HandlerSuite) TestUpdateMTOPostCounselingInfo() {

	suite.Run("Successful patch - Integration Test", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		mto := testdatagen.MakeAvailableMove(suite.DB())
		eTag := etag.GenerateEtag(mto.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: mto.ID.String(),
			IfMatch:         eTag,
		}
		// Create two shipments, one prime, one external.  Only prime one should be returned.
		primeShipment := testdatagen.MakePPMShipment(suite.DB(), testdatagen.Assertions{
			Move: mto,
			MTOShipment: models.MTOShipment{
				UsesExternalVendor: false,
			},
		})
		testdatagen.MakeMTOShipmentMinimal(suite.DB(), testdatagen.Assertions{
			Move: mto,
			MTOShipment: models.MTOShipment{
				ShipmentType:       models.MTOShipmentTypeHHGOutOfNTSDom,
				UsesExternalVendor: true,
			},
		})
		testdatagen.MakeMTOServiceItemBasic(suite.DB(), testdatagen.Assertions{
			MTOServiceItem: models.MTOServiceItem{
				Status: models.MTOServiceItemStatusApproved,
			},
			Move: mto,
			ReService: models.ReService{
				Code: models.ReServiceCodeCS, // CS - Counseling Services
			},
		})

		queryBuilder := query.NewQueryBuilder()
		fetcher := fetch.NewFetcher(queryBuilder)
		moveRouter := moverouter.NewMoveRouter()
		siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)
		updater := movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter)
		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()

		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			fetcher,
			updater,
			mtoChecker,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationOK{}, response)

		okResponse := response.(*movetaskorderops.UpdateMTOPostCounselingInformationOK)
		okPayload := okResponse.Payload

		// Validate outgoing payload
		suite.NoError(okResponse.Payload.Validate(strfmt.Default))

		suite.Equal(mto.ID.String(), okPayload.ID.String())
		suite.NotNil(okPayload.ETag)

		if suite.Len(okPayload.MtoShipments, 1) {
			suite.Equal(primeShipment.ID.String(), okPayload.MtoShipments[0].PpmShipment.ID.String())
			suite.Equal(primeShipment.ShipmentID.String(), okPayload.MtoShipments[0].ID.String())
		}

		suite.NotNil(okPayload.PrimeCounselingCompletedAt)
		suite.Equal(primemessages.PPMShipmentStatusWAITINGONCUSTOMER, okPayload.MtoShipments[0].PpmShipment.Status)
	})

	suite.Run("Unsuccessful patch - Integration Test - patch fail MTO not available", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		defaultMTO := testdatagen.MakeDefaultMove(suite.DB())
		eTag := etag.GenerateEtag(defaultMTO.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", defaultMTO.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		defaultMTOParams := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: defaultMTO.ID.String(),
			IfMatch:         eTag,
		}

		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()
		queryBuilder := query.NewQueryBuilder()
		moveRouter := moverouter.NewMoveRouter()
		fetcher := fetch.NewFetcher(queryBuilder)
		siCreator := mtoserviceitem.NewMTOServiceItemCreator(queryBuilder, moveRouter)
		updater := movetaskorder.NewMoveTaskOrderUpdater(queryBuilder, siCreator, moveRouter)
		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			fetcher,
			updater,
			mtoChecker,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(defaultMTOParams)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationNotFound{}, response)
		payload := response.(*movetaskorderops.UpdateMTOPostCounselingInformationNotFound).Payload

		// Validate outgoing payload
		suite.NoError(payload.Validate(strfmt.Default))
	})

	suite.Run("Patch failure - 500", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		mto := testdatagen.MakeAvailableMove(suite.DB())
		eTag := etag.GenerateEtag(mto.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		mockFetcher := mocks.Fetcher{}
		mockUpdater := mocks.MoveTaskOrderUpdater{}
		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()

		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			&mockFetcher,
			&mockUpdater,
			mtoChecker,
		}

		internalServerErr := errors.New("ServerError")
		params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: mto.ID.String(),
			IfMatch:         eTag,
		}

		mockUpdater.On("UpdatePostCounselingInfo",
			mock.AnythingOfType("*appcontext.appContext"),
			mto.ID,
			eTag,
		).Return(nil, internalServerErr)

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationInternalServerError{}, response)
		payload := response.(*movetaskorderops.UpdateMTOPostCounselingInformationInternalServerError).Payload

		// Validate outgoing payload
		suite.NoError(payload.Validate(strfmt.Default))
	})

	suite.Run("Patch failure - 404", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		mto := testdatagen.MakeAvailableMove(suite.DB())
		eTag := etag.GenerateEtag(mto.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		mockFetcher := mocks.Fetcher{}
		mockUpdater := mocks.MoveTaskOrderUpdater{}
		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()

		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			&mockFetcher,
			&mockUpdater,
			mtoChecker,
		}
		params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: mto.ID.String(),
			IfMatch:         eTag,
		}

		mockUpdater.On("UpdatePostCounselingInfo",
			mock.AnythingOfType("*appcontext.appContext"),
			mto.ID,
			eTag,
		).Return(nil, apperror.NotFoundError{})

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationNotFound{}, response)
		payload := response.(*movetaskorderops.UpdateMTOPostCounselingInformationNotFound).Payload

		// Validate outgoing payload
		suite.NoError(payload.Validate(strfmt.Default))
	})

	suite.Run("Patch failure - 409", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		mto := testdatagen.MakeAvailableMove(suite.DB())
		eTag := etag.GenerateEtag(mto.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		mockFetcher := mocks.Fetcher{}
		mockUpdater := mocks.MoveTaskOrderUpdater{}
		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()

		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			&mockFetcher,
			&mockUpdater,
			mtoChecker,
		}
		params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: mto.ID.String(),
			IfMatch:         eTag,
		}
		mockUpdater.On("UpdatePostCounselingInfo",
			mock.AnythingOfType("*appcontext.appContext"),
			mto.ID,
			eTag,
		).Return(nil, apperror.ConflictError{})

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationConflict{}, response)
		payload := response.(*movetaskorderops.UpdateMTOPostCounselingInformationConflict).Payload

		// Validate outgoing payload
		suite.NoError(payload.Validate(strfmt.Default))
	})

	suite.Run("Patch failure - 422", func() {
		requestUser := factory.BuildUser(nil, nil, nil)
		mto := testdatagen.MakeAvailableMove(suite.DB())
		eTag := etag.GenerateEtag(mto.UpdatedAt)
		req := httptest.NewRequest("PATCH", fmt.Sprintf("/move_task_orders/%s/post-counseling-info", mto.ID.String()), nil)
		req = suite.AuthenticateUserRequest(req, requestUser)

		mockFetcher := mocks.Fetcher{}
		mockUpdater := mocks.MoveTaskOrderUpdater{}
		mtoChecker := movetaskorder.NewMoveTaskOrderChecker()

		handler := UpdateMTOPostCounselingInformationHandler{
			suite.HandlerConfig(),
			&mockFetcher,
			&mockUpdater,
			mtoChecker,
		}

		mockUpdater.On("UpdatePostCounselingInfo",
			mock.AnythingOfType("*appcontext.appContext"),
			mto.ID,
			eTag,
		).Return(nil, apperror.NewInvalidInputError(uuid.Nil, nil, validate.NewErrors(), ""))
		params := movetaskorderops.UpdateMTOPostCounselingInformationParams{
			HTTPRequest:     req,
			MoveTaskOrderID: mto.ID.String(),
			IfMatch:         eTag,
		}

		// Validate incoming payload: no body to validate

		response := handler.Handle(params)
		suite.IsType(&movetaskorderops.UpdateMTOPostCounselingInformationUnprocessableEntity{}, response)
		payload := response.(*movetaskorderops.UpdateMTOPostCounselingInformationUnprocessableEntity).Payload

		// Validate outgoing payload
		suite.NoError(payload.Validate(strfmt.Default))
	})
}
