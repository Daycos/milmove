// RA Summary: gosec - errcheck - Unchecked return value
// RA: Linter flags errcheck error: Ignoring a method's return value can cause the program to overlook unexpected states and conditions.
// RA: Functions with unchecked return values in the file are used to generate stub data for a localized version of the application.
// RA: Given the data is being generated for local use and does not contain any sensitive information, there are no unexpected states and conditions
// RA: in which this would be considered a risk
// RA Developer Status: Mitigated
// RA Validator Status: Mitigated
// RA Modified Severity: N/A
// nolint:errcheck
package internalapi

import (
	"net/http/httptest"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/mock"

	"github.com/transcom/mymove/pkg/factory"
	ppmop "github.com/transcom/mymove/pkg/gen/internalapi/internaloperations/ppm"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	"github.com/transcom/mymove/pkg/handlers"
	"github.com/transcom/mymove/pkg/models"
	routemocks "github.com/transcom/mymove/pkg/route/mocks"
	moverouter "github.com/transcom/mymove/pkg/services/move"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/testdatagen/scenario"
	"github.com/transcom/mymove/pkg/unit"
)

func (suite *HandlerSuite) setupPersonallyProcuredMoveTest() {
	originZip3 := models.Tariff400ngZip3{
		Zip3:          "503",
		BasepointCity: "Des Moines",
		State:         "IA",
		ServiceArea:   "296",
		RateArea:      "US53",
		Region:        "7",
	}
	suite.MustSave(&originZip3)
	destinationZip3 := models.Tariff400ngZip3{
		Zip3:          "956",
		BasepointCity: "Sacramento",
		State:         "CA",
		ServiceArea:   "68",
		RateArea:      "US87",
		Region:        "2",
	}
	suite.MustSave(&destinationZip3)
	destinationZip5 := models.Tariff400ngZip5RateArea{
		Zip5:     "95630",
		RateArea: "US87",
	}
	suite.MustSave(&destinationZip5)
	originServiceArea := models.Tariff400ngServiceArea{
		Name:               "Des Moines, IA",
		ServiceArea:        "296",
		LinehaulFactor:     unit.Cents(263),
		ServiceChargeCents: unit.Cents(489),
		ServicesSchedule:   3,
		EffectiveDateLower: scenario.May15TestYear,
		EffectiveDateUpper: scenario.May14FollowingYear,
		SIT185ARateCents:   unit.Cents(1447),
		SIT185BRateCents:   unit.Cents(51),
		SITPDSchedule:      3,
	}
	suite.MustSave(&originServiceArea)
	destinationServiceArea := models.Tariff400ngServiceArea{
		Name:               "Sacramento, CA",
		ServiceArea:        "68",
		LinehaulFactor:     unit.Cents(78),
		ServiceChargeCents: unit.Cents(452),
		ServicesSchedule:   3,
		EffectiveDateLower: scenario.May15TestYear,
		EffectiveDateUpper: scenario.May14FollowingYear,
		SIT185ARateCents:   unit.Cents(1642),
		SIT185BRateCents:   unit.Cents(70),
		SITPDSchedule:      3,
	}
	suite.MustSave(&destinationServiceArea)
	tdl1 := models.TrafficDistributionList{
		SourceRateArea:    "US53",
		DestinationRegion: "12",
		CodeOfService:     "2",
	}
	suite.MustSave(&tdl1)
	tdl2 := models.TrafficDistributionList{
		SourceRateArea:    "US87",
		DestinationRegion: "2",
		CodeOfService:     "2",
	}
	suite.MustSave(&tdl2)
	tdl3 := models.TrafficDistributionList{
		SourceRateArea:    "US53",
		DestinationRegion: "2",
		CodeOfService:     "2",
	}
	suite.MustSave(&tdl3)
	tsp := models.TransportationServiceProvider{
		StandardCarrierAlphaCode: "STDM",
	}
	suite.MustSave(&tsp)
	tspPerformance1 := models.TransportationServiceProviderPerformance{
		PerformancePeriodStart:          scenario.Oct1TestYear,
		PerformancePeriodEnd:            scenario.Dec31TestYear,
		RateCycleStart:                  scenario.Oct1TestYear,
		RateCycleEnd:                    scenario.May14FollowingYear,
		TrafficDistributionListID:       tdl1.ID,
		TransportationServiceProviderID: tsp.ID,
		QualityBand:                     swag.Int(1),
		BestValueScore:                  90,
		LinehaulRate:                    unit.NewDiscountRateFromPercent(50.5),
		SITRate:                         unit.NewDiscountRateFromPercent(50),
	}
	suite.MustSave(&tspPerformance1)
	tspPerformance2 := models.TransportationServiceProviderPerformance{
		PerformancePeriodStart:          scenario.Oct1TestYear,
		PerformancePeriodEnd:            scenario.Dec31TestYear,
		RateCycleStart:                  scenario.Oct1TestYear,
		RateCycleEnd:                    scenario.May14FollowingYear,
		TrafficDistributionListID:       tdl2.ID,
		TransportationServiceProviderID: tsp.ID,
		QualityBand:                     swag.Int(1),
		BestValueScore:                  90,
		LinehaulRate:                    unit.NewDiscountRateFromPercent(50.5),
		SITRate:                         unit.NewDiscountRateFromPercent(50),
	}
	suite.MustSave(&tspPerformance2)
	tspPerformance3 := models.TransportationServiceProviderPerformance{
		PerformancePeriodStart:          scenario.Oct1TestYear,
		PerformancePeriodEnd:            scenario.Dec31TestYear,
		RateCycleStart:                  scenario.Oct1TestYear,
		RateCycleEnd:                    scenario.May14FollowingYear,
		TrafficDistributionListID:       tdl3.ID,
		TransportationServiceProviderID: tsp.ID,
		QualityBand:                     swag.Int(1),
		BestValueScore:                  90,
		LinehaulRate:                    unit.NewDiscountRateFromPercent(50.5),
		SITRate:                         unit.NewDiscountRateFromPercent(50),
	}
	suite.MustSave(&tspPerformance3)
	fullPackRate := models.Tariff400ngFullPackRate{
		Schedule:           3,
		WeightLbsLower:     0,
		WeightLbsUpper:     16001,
		RateCents:          6130,
		EffectiveDateLower: scenario.May15TestYear,
		EffectiveDateUpper: scenario.May14FollowingYear,
	}
	suite.MustSave(&fullPackRate)
}

func (suite *HandlerSuite) TestCreatePPMHandler() {
	user1 := testdatagen.MakeDefaultServiceMember(suite.DB())
	orders := testdatagen.MakeDefaultOrder(suite.DB())
	factory.FetchOrBuildDefaultContractor(suite.DB(), nil, nil)

	selectedMoveType := models.SelectedMoveTypeHHGPPM

	moveOptions := models.MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, locErr := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.False(verrs.HasAny(), "failed to create new move")
	suite.Nil(locErr)

	request := httptest.NewRequest("POST", "/fake/path", nil)
	request = suite.AuthenticateRequest(request, orders.ServiceMember)

	newPPMPayload := internalmessages.CreatePersonallyProcuredMovePayload{
		WeightEstimate:   swag.Int64(12),
		PickupPostalCode: swag.String("00112"),
		DaysInStorage:    swag.Int64(3),
	}

	newPPMParams := ppmop.CreatePersonallyProcuredMoveParams{
		MoveID:                              strfmt.UUID(move.ID.String()),
		CreatePersonallyProcuredMovePayload: &newPPMPayload,
		HTTPRequest:                         request,
	}

	handler := CreatePersonallyProcuredMoveHandler{suite.HandlerConfig()}
	response := handler.Handle(newPPMParams)
	// assert we got back the 201 response
	createdResponse := response.(*ppmop.CreatePersonallyProcuredMoveCreated)
	createdIssuePayload := createdResponse.Payload
	suite.NotNil(createdIssuePayload.ID)

	// Next try the wrong user
	request = suite.AuthenticateRequest(request, user1)
	newPPMParams.HTTPRequest = request

	_ = handler.Handle(newPPMParams)

	// Now try a bad move
	newPPMParams.MoveID = strfmt.UUID(uuid.Must(uuid.NewV4()).String())
	_ = handler.Handle(newPPMParams)

}

func (suite *HandlerSuite) TestSubmitPPMHandler() {
	t := suite.T()

	// create a ppm
	move1 := testdatagen.MakeDefaultMove(suite.DB())
	wtgEst := unit.Pound(1)
	ppm := models.PersonallyProcuredMove{
		MoveID:         move1.ID,
		Move:           move1,
		WeightEstimate: &wtgEst,
		Status:         models.PPMStatusDRAFT,
	}

	verrs, err := suite.DB().ValidateAndCreate(&ppm)
	if verrs.HasAny() || err != nil {
		t.Error(verrs, err)
	}

	req := httptest.NewRequest("POST", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move1.Orders.ServiceMember)
	submitDate := strfmt.DateTime(time.Now())
	newSubmitPersonallyProcuredMovePayload := internalmessages.SubmitPersonallyProcuredMovePayload{
		SubmitDate: &submitDate,
	}

	submitPPMParams := ppmop.SubmitPersonallyProcuredMoveParams{
		PersonallyProcuredMoveID:            strfmt.UUID(ppm.ID.String()),
		HTTPRequest:                         req,
		SubmitPersonallyProcuredMovePayload: &newSubmitPersonallyProcuredMovePayload,
	}

	// submit the PPM
	handler := SubmitPersonallyProcuredMoveHandler{suite.HandlerConfig()}
	response := handler.Handle(submitPPMParams)
	okResponse := response.(*ppmop.SubmitPersonallyProcuredMoveOK)
	submitPPMPayload := okResponse.Payload

	suite.Require().Equal(submitPPMPayload.Status, internalmessages.PPMStatusSUBMITTED, "PPM should have been submitted")
}

func (suite *HandlerSuite) TestPatchPPMHandler() {
	suite.setupPersonallyProcuredMoveTest()

	initialWeight := unit.Pound(4100)
	newWeight := swag.Int64(4105)

	// Date picked essentialy at random, but needs to be within TestYear
	newMoveDate := time.Date(testdatagen.TestYear, time.November, 10, 23, 0, 0, 0, time.UTC)
	initialMoveDate := newMoveDate.Add(-2 * 24 * time.Hour)

	hasAdditionalPostalCode := swag.Bool(true)
	newHasAdditionalPostalCode := swag.Bool(false)
	additionalPickupPostalCode := swag.String("90210")

	hasSit := swag.Bool(true)
	newHasSit := swag.Bool(false)
	daysInStorage := swag.Int64(3)
	newPickupPostalCode := swag.String("32168")
	newSitCost := swag.Int64(60)

	move := testdatagen.MakeDefaultMove(suite.DB())

	newAdvanceWorksheet := models.Document{
		ServiceMember:   move.Orders.ServiceMember,
		ServiceMemberID: move.Orders.ServiceMemberID,
	}
	suite.MustSave(&newAdvanceWorksheet)

	ppm1 := models.PersonallyProcuredMove{
		MoveID:                     move.ID,
		Move:                       move,
		WeightEstimate:             &initialWeight,
		OriginalMoveDate:           &initialMoveDate,
		HasAdditionalPostalCode:    hasAdditionalPostalCode,
		AdditionalPickupPostalCode: additionalPickupPostalCode,
		HasSit:                     hasSit,
		DaysInStorage:              daysInStorage,
		Status:                     models.PPMStatusDRAFT,
		AdvanceWorksheet:           newAdvanceWorksheet,
		AdvanceWorksheetID:         &newAdvanceWorksheet.ID,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		WeightEstimate:          newWeight,
		OriginalMoveDate:        handlers.FmtDatePtr(&newMoveDate),
		HasAdditionalPostalCode: newHasAdditionalPostalCode,
		PickupPostalCode:        newPickupPostalCode,
		HasSit:                  newHasSit,
		TotalSitCost:            newSitCost,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	planner := &routemocks.Planner{}
	handlerConfig := suite.HandlerConfig()
	handlerConfig.SetPlanner(planner)
	handler := PatchPersonallyProcuredMoveHandler{handlerConfig}
	planner.On("Zip5TransitDistanceLineHaul",
		mock.AnythingOfType("*appcontext.appContext"),
		mock.Anything,
		mock.Anything,
	).Return(900, nil)
	response := handler.Handle(patchPPMParams)

	// assert we got back the 201 response
	okResponse := response.(*ppmop.PatchPersonallyProcuredMoveOK)
	patchPPMPayload := okResponse.Payload

	suite.Equal(patchPPMPayload.WeightEstimate, newWeight, "Weight should have been updated.")
	suite.Equal(patchPPMPayload.TotalSitCost, newSitCost, "Total sit cost should have been updated.")
	suite.Equal(patchPPMPayload.PickupPostalCode, newPickupPostalCode, "PickupPostalCode should have been updated.")
	suite.Nil(patchPPMPayload.AdditionalPickupPostalCode, "AdditionalPickupPostalCode should have been updated to nil.")
	suite.Equal(*(*time.Time)(patchPPMPayload.OriginalMoveDate), newMoveDate, "MoveDate should have been updated.")
}

func (suite *HandlerSuite) TestPatchPPMHandlerSetWeightLater() {
	appCtx := suite.AppContextForTest()
	scenario.RunRateEngineScenario1(appCtx)

	suite.setupPersonallyProcuredMoveTest()
	weight := swag.Int64(4100)

	// Date picked essentialy at random, but needs to be within TestYear
	moveDate := time.Date(testdatagen.TestYear, time.November, 10, 23, 0, 0, 0, time.UTC)

	pickupPostalCode := swag.String("32168")
	destinationPostalCode := swag.String("29401")

	move := testdatagen.MakeDefaultMove(suite.DB())

	ppm1 := models.PersonallyProcuredMove{
		MoveID:                move.ID,
		Move:                  move,
		OriginalMoveDate:      &moveDate,
		PickupPostalCode:      pickupPostalCode,
		DestinationPostalCode: destinationPostalCode,
		Status:                models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	payload := &internalmessages.PatchPersonallyProcuredMovePayload{
		WeightEstimate: weight,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: payload,
	}

	testdatagen.MakeTariff400ngServiceArea(suite.DB(), testdatagen.Assertions{
		Tariff400ngServiceArea: models.Tariff400ngServiceArea{
			ServiceArea: "296",
		},
	})

	planner := &routemocks.Planner{}
	planner.On("Zip5TransitDistanceLineHaul",
		mock.AnythingOfType("*appcontext.appContext"),
		mock.Anything,
		mock.Anything,
	).Return(900, nil)

	handlerConfig := suite.HandlerConfig()
	handlerConfig.SetPlanner(planner)
	handler := PatchPersonallyProcuredMoveHandler{handlerConfig}
	response := handler.Handle(patchPPMParams)

	// assert we got back the 201 response
	okResponse := response.(*ppmop.PatchPersonallyProcuredMoveOK)
	patchPPMPayload := okResponse.Payload

	suite.Assertions.Equal(int64(*weight), *patchPPMPayload.WeightEstimate)

	// Now check that SIT values update when days in storage is set
	hasSit := swag.Bool(true)
	daysInStorage := swag.Int64(3)
	*payload = internalmessages.PatchPersonallyProcuredMovePayload{
		HasSit:        hasSit,
		DaysInStorage: daysInStorage,
	}

	response = handler.Handle(patchPPMParams)
	// assert we got back the 201 response
	okResponse = response.(*ppmop.PatchPersonallyProcuredMoveOK)
	patchPPMPayload = okResponse.Payload

	suite.Assertions.Equal(daysInStorage, patchPPMPayload.DaysInStorage)
}

func (suite *HandlerSuite) TestPatchPPMHandlerWrongUser() {
	initialWeight := unit.Pound(1)
	newWeight := swag.Int64(5)

	// Date picked essentialy at random, but needs to be within TestYear
	newMoveDate := time.Date(testdatagen.TestYear, time.November, 10, 23, 0, 0, 0, time.UTC)
	initialMoveDate := newMoveDate.Add(-2 * 24 * time.Hour)

	user2 := testdatagen.MakeDefaultServiceMember(suite.DB())
	move := testdatagen.MakeDefaultMove(suite.DB())

	ppm1 := models.PersonallyProcuredMove{
		MoveID:           move.ID,
		Move:             move,
		WeightEstimate:   &initialWeight,
		OriginalMoveDate: &initialMoveDate,
		Status:           models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("PATCH", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, user2)

	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		WeightEstimate:   newWeight,
		OriginalMoveDate: handlers.FmtDatePtr(&newMoveDate),
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	handler := PatchPersonallyProcuredMoveHandler{suite.HandlerConfig()}
	_ = handler.Handle(patchPPMParams)
}

// TODO: no response is returned when the moveid doesn't match. How did this ever work?
func (suite *HandlerSuite) TestPatchPPMHandlerWrongMoveID() {
	initialWeight := unit.Pound(1)
	newWeight := swag.Int64(5)

	orders := testdatagen.MakeDefaultOrder(suite.DB())
	orders1 := testdatagen.MakeDefaultOrder(suite.DB())
	factory.FetchOrBuildDefaultContractor(suite.DB(), nil, nil)

	selectedMoveType := models.SelectedMoveTypeHHGPPM

	moveOptions := models.MoveOptions{
		SelectedType: &selectedMoveType,
		Show:         swag.Bool(true),
	}
	move, verrs, err := orders.CreateNewMove(suite.DB(), moveOptions)
	suite.Nil(err, "Failed to save move")
	suite.False(verrs.HasAny(), "failed to validate move")
	move.Orders = orders

	move2, verrs, err := orders1.CreateNewMove(suite.DB(), moveOptions)
	suite.Nil(err, "Failed to save move")
	suite.False(verrs.HasAny(), "failed to validate move")
	move2.Orders = orders1

	ppm1 := models.PersonallyProcuredMove{
		MoveID:         move2.ID,
		Move:           *move2,
		WeightEstimate: &initialWeight,
		Status:         models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, orders.ServiceMember)

	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		WeightEstimate: newWeight,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	handler := PatchPersonallyProcuredMoveHandler{suite.HandlerConfig()}
	_ = handler.Handle(patchPPMParams)
}

func (suite *HandlerSuite) TestPatchPPMHandlerNoMove() {
	t := suite.T()

	initialWeight := unit.Pound(1)
	newWeight := swag.Int64(5)

	move := testdatagen.MakeDefaultMove(suite.DB())

	badMoveID := uuid.Must(uuid.NewV4())

	ppm1 := models.PersonallyProcuredMove{
		MoveID:         move.ID,
		Move:           move,
		WeightEstimate: &initialWeight,
		Status:         models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		WeightEstimate: newWeight,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(badMoveID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	handler := PatchPersonallyProcuredMoveHandler{suite.HandlerConfig()}
	response := handler.Handle(patchPPMParams)

	// assert we got back the badrequest response
	_, ok := response.(*ppmop.PatchPersonallyProcuredMoveBadRequest)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}
}

func (suite *HandlerSuite) TestPatchPPMHandlerAdvance() {
	t := suite.T()

	initialWeight := unit.Pound(1)

	move := testdatagen.MakeDefaultMove(suite.DB())

	ppm1 := models.PersonallyProcuredMove{
		MoveID:         move.ID,
		Move:           move,
		WeightEstimate: &initialWeight,
		Status:         models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	// First, create an advance
	truth := true
	initialAmount := int64(1000)
	initialMethod := internalmessages.MethodOfReceiptMILPAY
	initialAdvance := internalmessages.Reimbursement{
		RequestedAmount: &initialAmount,
		MethodOfReceipt: &initialMethod,
	}

	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		HasRequestedAdvance: &truth,
		Advance:             &initialAdvance,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	handler := PatchPersonallyProcuredMoveHandler{suite.HandlerConfig()}
	response := handler.Handle(patchPPMParams)

	created, ok := response.(*ppmop.PatchPersonallyProcuredMoveOK)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	suite.Require().Equal(internalmessages.ReimbursementStatusDRAFT, *created.Payload.Advance.Status, "expected Draft")
	suite.Require().Equal(initialAmount, *created.Payload.Advance.RequestedAmount, "expected amount to shine through.")

	// Then, update the advance
	newAmount := int64(9999999)
	badStatus := internalmessages.ReimbursementStatusREQUESTED
	payload.Advance.RequestedAmount = &newAmount
	payload.Advance.Status = &badStatus

	response = handler.Handle(patchPPMParams)

	// assert we got back the created response
	updated, ok := response.(*ppmop.PatchPersonallyProcuredMoveOK)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	suite.Require().Equal(internalmessages.ReimbursementStatusDRAFT, *updated.Payload.Advance.Status, "expected Draft still")
	suite.Require().Equal(newAmount, *updated.Payload.Advance.RequestedAmount, "expected amount to be updated")

}

// TODO: Fix now that we capture transaction error. May be a data setup problem
/* func (suite *HandlerSuite) TestPatchPPMHandlerEdgeCases() {
	t := suite.T()

	initialWeight := unit.Pound(1)

	move := testdatagen.MakeDefaultMove(suite.DB())

	ppm1 := models.PersonallyProcuredMove{
		MoveID:         move.ID,
		Move:           move,
		WeightEstimate: &initialWeight,
		Status:         models.PPMStatusDRAFT,
	}
	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	// First, try and set has_requested_advance without passing in an advance
	truth := true
	payload := internalmessages.PatchPersonallyProcuredMovePayload{
		HasRequestedAdvance: &truth,
	}

	patchPPMParams := ppmop.PatchPersonallyProcuredMoveParams{
		HTTPRequest:                        req,
		MoveID:                             strfmt.UUID(move.ID.String()),
		PersonallyProcuredMoveID:           strfmt.UUID(ppm1.ID.String()),
		PatchPersonallyProcuredMovePayload: &payload,
	}

	handler := PatchPersonallyProcuredMoveHandler{handlers.NewHandlerConfig(suite.DB(), suite.TestLogger())}
	response := handler.Handle(patchPPMParams)

	suite.CheckResponseBadRequest(response)

	// Then, try and create an advance without setting has requested advance
	initialAmount := int64(1000)
	initialMethod := internalmessages.MethodOfReceiptMILPAY
	initialAdvance := internalmessages.Reimbursement{
		RequestedAmount: &initialAmount,
		MethodOfReceipt: &initialMethod,
	}
	payload = internalmessages.PatchPersonallyProcuredMovePayload{
		Advance: &initialAdvance,
	}

	response = handler.Handle(patchPPMParams)

	created, ok := response.(*ppmop.PatchPersonallyProcuredMoveOK)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	suite.Require().Equal(internalmessages.ReimbursementStatusDRAFT, *created.Payload.Advance.Status, "expected Draft")
	suite.Require().Equal(initialAmount, *created.Payload.Advance.RequestedAmount, "expected amount to shine through.")
} */

func (suite *HandlerSuite) TestRequestPPMPayment() {
	t := suite.T()

	initialWeight := unit.Pound(1)

	move := testdatagen.MakeDefaultMove(suite.DB())
	moveRouter := moverouter.NewMoveRouter()
	newSignedCertification := testdatagen.MakeSignedCertification(suite.DB(), testdatagen.Assertions{
		SignedCertification: models.SignedCertification{
			MoveID: move.ID,
		},
		Stub: true,
	})
	err := moveRouter.Submit(suite.AppContextForTest(), &move, &newSignedCertification)
	if err != nil {
		t.Fatal("Should transition.")
	}
	err = moveRouter.Approve(suite.AppContextForTest(), &move)
	if err != nil {
		t.Fatal("Should transition.")
	}

	suite.MustSave(&move)

	ppm1 := models.PersonallyProcuredMove{
		MoveID:         move.ID,
		Move:           move,
		WeightEstimate: &initialWeight,
		Status:         models.PPMStatusDRAFT,
	}
	err = ppm1.Submit(time.Now())
	if err != nil {
		t.Fatal("Should transition.")
	}
	err = ppm1.Approve(time.Now())
	if err != nil {
		t.Fatal("Should transition.")
	}

	suite.MustSave(&ppm1)

	req := httptest.NewRequest("GET", "/fake/path", nil)
	req = suite.AuthenticateRequest(req, move.Orders.ServiceMember)

	requestPaymentParams := ppmop.RequestPPMPaymentParams{
		HTTPRequest:              req,
		PersonallyProcuredMoveID: strfmt.UUID(ppm1.ID.String()),
	}

	handler := RequestPPMPaymentHandler{suite.HandlerConfig()}
	response := handler.Handle(requestPaymentParams)

	created, ok := response.(*ppmop.RequestPPMPaymentOK)
	if !ok {
		t.Fatalf("Request failed: %#v", response)
	}

	suite.Require().Equal(internalmessages.PPMStatusPAYMENTREQUESTED, created.Payload.Status, "expected payment requested")

}
