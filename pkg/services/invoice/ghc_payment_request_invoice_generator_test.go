package invoice

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/transcom/mymove/pkg/apperror"
	"github.com/transcom/mymove/pkg/db/sequence"
	ediinvoice "github.com/transcom/mymove/pkg/edi/invoice"
	edisegment "github.com/transcom/mymove/pkg/edi/segment"
	"github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
	"github.com/transcom/mymove/pkg/testingsuite"
	"github.com/transcom/mymove/pkg/unit"
)

const (
	hierarchicalLevelCodeExpected string = "9"
)

type GHCInvoiceSuite struct {
	*testingsuite.PopTestSuite
	icnSequencer sequence.Sequencer
}

func TestGHCInvoiceSuite(t *testing.T) {
	ts := &GHCInvoiceSuite{
		PopTestSuite: testingsuite.NewPopTestSuite(testingsuite.CurrentPackage().Suffix("ghcinvoice"),
			testingsuite.WithPerTestTransaction()),
	}
	ts.icnSequencer = sequence.NewDatabaseSequencer(ediinvoice.ICNSequenceName)

	suite.Run(t, ts)
	ts.PopTestSuite.TearDown()
}

const testDateFormat = "20060102"
const testISADateFormat = "060102"
const testTimeFormat = "1504"

func (suite *GHCInvoiceSuite) TestAllGenerateEdi() {
	mockClock := clock.NewMock()
	currentTime := mockClock.Now()
	requestedPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	scheduledPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 20, 0, 0, 0, 0, time.UTC)
	actualPickupDate := time.Date(testdatagen.GHCTestYear, time.September, 22, 0, 0, 0, 0, time.UTC)
	generator := NewGHCPaymentRequestInvoiceGenerator(suite.icnSequencer, mockClock)
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
		{
			Key:     models.ServiceItemParamNameReferenceDate,
			KeyType: models.ServiceItemParamTypeDate,
			Value:   currentTime.Format(testDateFormat),
		},
		{
			Key:     models.ServiceItemParamNameWeightBilled,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "4242",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "24246",
		},
	}

	var serviceMember models.ServiceMember
	var paymentRequest models.PaymentRequest
	var paymentServiceItems models.PaymentServiceItems
	var result ediinvoice.Invoice858C
	var err error

	setupTestData := func() {
		mto := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})

		paymentRequest = testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
			Move: mto,
			PaymentRequest: models.PaymentRequest{
				IsFinal:         false,
				Status:          models.PaymentRequestStatusPending,
				RejectionReason: nil,
			},
		})

		mtoShipment := testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: mto,
			MTOShipment: models.MTOShipment{
				RequestedPickupDate: &requestedPickupDate,
				ScheduledPickupDate: &scheduledPickupDate,
				ActualPickupDate:    &actualPickupDate,
			},
		})

		priceCents := unit.Cents(888)
		assertions := testdatagen.Assertions{
			Move:           mto,
			MTOShipment:    mtoShipment,
			PaymentRequest: paymentRequest,
			PaymentServiceItem: models.PaymentServiceItem{
				Status:     models.PaymentServiceItemStatusApproved,
				PriceCents: &priceCents,
			},
		}

		dlh := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDLH,
			basicPaymentServiceItemParams,
			assertions,
		)
		fsc := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeFSC,
			basicPaymentServiceItemParams,
			assertions,
		)
		ms := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeMS,
			basicPaymentServiceItemParams,
			assertions,
		)
		cs := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeCS,
			basicPaymentServiceItemParams,
			assertions,
		)
		dsh := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDSH,
			basicPaymentServiceItemParams,
			assertions,
		)
		dop := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDOP,
			basicPaymentServiceItemParams,
			assertions,
		)
		ddp := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDDP,
			basicPaymentServiceItemParams,
			assertions,
		)
		dpk := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDPK,
			basicPaymentServiceItemParams,
			assertions,
		)
		dnpk := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDNPK,
			basicPaymentServiceItemParams,
			assertions,
		)
		dupk := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDUPK,
			basicPaymentServiceItemParams,
			assertions,
		)
		ddfsit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDDFSIT,
			basicPaymentServiceItemParams,
			assertions,
		)
		ddasit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDDASIT,
			basicPaymentServiceItemParams,
			assertions,
		)
		dofsit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDOFSIT,
			basicPaymentServiceItemParams,
			assertions,
		)
		doasit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDOASIT,
			basicPaymentServiceItemParams,
			assertions,
		)
		doshut := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDOSHUT,
			basicPaymentServiceItemParams,
			assertions,
		)
		ddshut := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDDSHUT,
			basicPaymentServiceItemParams,
			assertions,
		)

		additionalParamsForCrating := []testdatagen.CreatePaymentServiceItemParams{
			{
				Key:     models.ServiceItemParamNameCubicFeetBilled,
				KeyType: models.ServiceItemParamTypeDecimal,
				Value:   "144.5",
			},
			{
				Key:     models.ServiceItemParamNamePriceRateOrFactor,
				KeyType: models.ServiceItemParamTypeDecimal,
				Value:   "23.69",
			},
		}
		cratingParams := append(basicPaymentServiceItemParams, additionalParamsForCrating...)
		dcrt := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDCRT,
			cratingParams,
			assertions,
		)
		ducrt := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDUCRT,
			cratingParams,
			assertions,
		)

		distanceZipSITDestParam := testdatagen.CreatePaymentServiceItemParams{
			Key:     models.ServiceItemParamNameDistanceZipSITDest,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "44",
		}
		dddsitParams := append(basicPaymentServiceItemParams, distanceZipSITDestParam)
		dddsit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDDDSIT,
			dddsitParams,
			assertions,
		)

		distanceZipSITOriginParam := testdatagen.CreatePaymentServiceItemParams{
			Key:     models.ServiceItemParamNameDistanceZipSITOrigin,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "33",
		}
		dopsitParams := append(basicPaymentServiceItemParams, distanceZipSITOriginParam)
		dopsit := testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDOPSIT,
			dopsitParams,
			assertions,
		)

		paymentServiceItems = models.PaymentServiceItems{}
		paymentServiceItems = append(paymentServiceItems, dlh, fsc, ms, cs, dsh, dop, ddp, dpk, dnpk, dupk, ddfsit, ddasit, dofsit, doasit, doshut, ddshut, dcrt, ducrt, dddsit, dopsit)

		serviceMember = testdatagen.MakeExtendedServiceMember(suite.DB(), testdatagen.Assertions{
			ServiceMember: models.ServiceMember{
				ID: uuid.FromStringOrNil("d66d2f35-218c-4b85-b9d1-631949b9d984"),
			},
		})

		// setup known next value
		icnErr := suite.icnSequencer.SetVal(suite.AppContextForTest(), 122)
		suite.NoError(icnErr)

		// Proceed with full EDI Generation tests
		result, err = generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.NoError(err)
	}

	// Test that the Interchange Control Number (ICN) is being used as the Group Control Number (GCN)
	suite.Run("the GCN is equal to the ICN", func() {
		setupTestData()
		suite.EqualValues(result.ISA.InterchangeControlNumber, result.IEA.InterchangeControlNumber, result.GS.GroupControlNumber, result.GE.GroupControlNumber)
	})

	// Test that the Interchange Control Number (ICN) is being saved to the db
	suite.Run("the ICN is saved to the database", func() {
		setupTestData()
		var pr2icn models.PaymentRequestToInterchangeControlNumber
		err := suite.DB().Where("payment_request_id = ?", paymentRequest.ID).First(&pr2icn)
		suite.NoError(err)
		suite.Equal(int(result.ISA.InterchangeControlNumber), pr2icn.InterchangeControlNumber)
	})

	// Test Invoice Start and End Segments
	suite.Run("adds isa start segment", func() {
		setupTestData()
		suite.Equal("00", result.ISA.AuthorizationInformationQualifier)
		suite.Equal("0084182369", result.ISA.AuthorizationInformation)
		suite.Equal("00", result.ISA.SecurityInformationQualifier)
		suite.Equal("0000000000", result.ISA.SecurityInformation)
		suite.Equal("ZZ", result.ISA.InterchangeSenderIDQualifier)
		suite.Equal(fmt.Sprintf("%-15s", "MILMOVE"), result.ISA.InterchangeSenderID)
		suite.Equal("12", result.ISA.InterchangeReceiverIDQualifier)
		suite.Equal(fmt.Sprintf("%-15s", "8004171844"), result.ISA.InterchangeReceiverID)
		suite.Equal(currentTime.Format(testISADateFormat), result.ISA.InterchangeDate)
		suite.Equal(currentTime.Format(testTimeFormat), result.ISA.InterchangeTime)
		suite.Equal("U", result.ISA.InterchangeControlStandards)
		suite.Equal("00401", result.ISA.InterchangeControlVersionNumber)
		suite.Equal(int64(123), result.ISA.InterchangeControlNumber)
		suite.Equal(0, result.ISA.AcknowledgementRequested)
		suite.Equal("T", result.ISA.UsageIndicator)
		suite.Equal("|", result.ISA.ComponentElementSeparator)
	})

	suite.Run("adds gs start segment", func() {
		setupTestData()
		suite.Equal("SI", result.GS.FunctionalIdentifierCode)
		suite.Equal("MILMOVE", result.GS.ApplicationSendersCode)
		suite.Equal("8004171844", result.GS.ApplicationReceiversCode)
		suite.Equal(currentTime.Format(testDateFormat), result.GS.Date)
		suite.Equal(currentTime.Format(testTimeFormat), result.GS.Time)
		suite.Equal(int64(123), result.GS.GroupControlNumber)
		suite.Equal("X", result.GS.ResponsibleAgencyCode)
		suite.Equal("004010", result.GS.Version)
	})

	suite.Run("adds st start segment", func() {
		setupTestData()
		suite.Equal("858", result.ST.TransactionSetIdentifierCode)
		suite.Equal("0001", result.ST.TransactionSetControlNumber)
	})

	suite.Run("se segment has correct value", func() {
		setupTestData()
		// Will need to be updated as more service items are supported
		suite.Equal(163, result.SE.NumberOfIncludedSegments)
		suite.Equal("0001", result.SE.TransactionSetControlNumber)
	})

	suite.Run("adds ge end segment", func() {
		setupTestData()
		suite.Equal(1, result.GE.NumberOfTransactionSetsIncluded)
		suite.Equal(int64(123), result.GE.GroupControlNumber)
	})

	suite.Run("adds iea end segment", func() {
		setupTestData()
		suite.Equal(1, result.IEA.NumberOfIncludedFunctionalGroups)
		suite.Equal(int64(123), result.IEA.InterchangeControlNumber)
	})

	// Test Header Generation
	suite.Run("adds bx header segment", func() {
		setupTestData()
		bx := result.Header.ShipmentInformation
		suite.IsType(edisegment.BX{}, bx)
		suite.Equal("00", bx.TransactionSetPurposeCode)
		suite.Equal("J", bx.TransactionMethodTypeCode)
		suite.Equal("PP", bx.ShipmentMethodOfPayment)
		suite.Equal(paymentRequest.PaymentRequestNumber, bx.ShipmentIdentificationNumber)
		suite.Equal("BLKW", bx.StandardCarrierAlphaCode)
		suite.Equal("4", bx.ShipmentQualifier)
	})

	suite.Run("does not error out creating EDI from Invoice858", func() {
		setupTestData()
		_, err := result.EDIString(suite.Logger())
		suite.NoError(err)
	})

	suite.Run("adding to n9 header", func() {
		setupTestData()
		testData := []struct {
			TestName      string
			Qualifier     string
			ExpectedValue string
			ActualValue   *edisegment.N9
		}{
			{TestName: "payment request number", Qualifier: "CN", ExpectedValue: paymentRequest.PaymentRequestNumber, ActualValue: &result.Header.PaymentRequestNumber},
			{TestName: "contract code", Qualifier: "CT", ExpectedValue: "TRUSS_TEST", ActualValue: &result.Header.ContractCode},
			{TestName: "service member name", Qualifier: "1W", ExpectedValue: serviceMember.ReverseNameLineFormat(), ActualValue: &result.Header.ServiceMemberName},
			{TestName: "service member rank", Qualifier: "ML", ExpectedValue: string(*serviceMember.Rank), ActualValue: &result.Header.ServiceMemberRank},
			{TestName: "service member branch", Qualifier: "3L", ExpectedValue: string(*serviceMember.Affiliation), ActualValue: &result.Header.ServiceMemberBranch},
		}

		for _, data := range testData {
			suite.Run(fmt.Sprintf("adds %s to header", data.TestName), func() {
				suite.IsType(&edisegment.N9{}, data.ActualValue)
				n9 := data.ActualValue
				suite.Equal(data.Qualifier, n9.ReferenceIdentificationQualifier)
				suite.Equal(data.ExpectedValue, n9.ReferenceIdentification)
			})
		}
	})
	suite.Run("adds currency to header", func() {
		setupTestData()
		currency := result.Header.Currency
		suite.IsType(edisegment.C3{}, currency)
		suite.Equal("USD", currency.CurrencyCodeC301)
	})

	suite.Run("adds actual pickup date to header", func() {
		setupTestData()
		g62Requested := result.Header.RequestedPickupDate
		suite.IsType(&edisegment.G62{}, g62Requested)
		suite.NotNil(g62Requested)
		suite.Equal(10, g62Requested.DateQualifier)
		suite.Equal(requestedPickupDate.Format(testDateFormat), g62Requested.Date)

		g62Scheduled := result.Header.ScheduledPickupDate
		suite.IsType(&edisegment.G62{}, g62Scheduled)
		suite.Equal(76, g62Scheduled.DateQualifier)
		suite.Equal(scheduledPickupDate.Format(testDateFormat), g62Scheduled.Date)

		g62Actual := result.Header.ActualPickupDate
		suite.IsType(&edisegment.G62{}, g62Actual)
		suite.Equal(86, g62Actual.DateQualifier)
		suite.Equal(actualPickupDate.Format(testDateFormat), g62Actual.Date)
	})

	suite.Run("adds buyer and seller organization name", func() {
		setupTestData()
		// buyer name
		originDutyLocation := paymentRequest.MoveTaskOrder.Orders.OriginDutyLocation
		transportationOffice, err := models.FetchDutyLocationTransportationOffice(suite.DB(), originDutyLocation.ID)
		suite.FatalNoError(err)
		buyerOrg := result.Header.BuyerOrganizationName
		suite.IsType(edisegment.N1{}, buyerOrg)
		suite.Equal("BY", buyerOrg.EntityIdentifierCode)
		suite.Equal(transportationOffice.Name, buyerOrg.Name)
		suite.Equal("92", buyerOrg.IdentificationCodeQualifier)
		suite.Equal(transportationOffice.Gbloc, buyerOrg.IdentificationCode)

		sellerOrg := result.Header.SellerOrganizationName
		suite.IsType(edisegment.N1{}, sellerOrg)
		suite.Equal("SE", sellerOrg.EntityIdentifierCode)
		suite.Equal("Prime", sellerOrg.Name)
		suite.Equal("2", sellerOrg.IdentificationCodeQualifier)
		suite.Equal("BLKW", sellerOrg.IdentificationCode)
	})

	suite.Run("adds orders destination address", func() {
		setupTestData()
		expectedDutyLocation := paymentRequest.MoveTaskOrder.Orders.NewDutyLocation
		transportationOffice, err := models.FetchDutyLocationTransportationOffice(suite.DB(), expectedDutyLocation.ID)
		suite.FatalNoError(err)
		// name
		n1 := result.Header.DestinationName
		suite.IsType(edisegment.N1{}, n1)
		suite.Equal("ST", n1.EntityIdentifierCode)
		suite.Equal(expectedDutyLocation.Name, n1.Name)
		suite.Equal("10", n1.IdentificationCodeQualifier)
		suite.Equal(transportationOffice.Gbloc, n1.IdentificationCode)
		// street address
		address := expectedDutyLocation.Address
		destAddress := result.Header.DestinationStreetAddress
		suite.IsType(&edisegment.N3{}, destAddress)
		n3 := *destAddress
		suite.Equal(address.StreetAddress1, n3.AddressInformation1)
		if address.StreetAddress2 == nil {
			suite.Empty(n3.AddressInformation2)
		} else {
			suite.Equal(*address.StreetAddress2, n3.AddressInformation2)
		}
		// city state info
		n4 := result.Header.DestinationPostalDetails
		suite.IsType(edisegment.N4{}, n4)
		suite.Equal(address.City, n4.CityName)
		suite.Equal(address.State, n4.StateOrProvinceCode)
		suite.Equal(address.PostalCode, n4.PostalCode)
		countryCode, err := address.CountryCode()
		suite.NoError(err)
		suite.Equal(*countryCode, n4.CountryCode)
		// Office Phone
		destinationDutyLocationPhoneLines := expectedDutyLocation.TransportationOffice.PhoneLines
		var destPhoneLines []string
		for _, phoneLine := range destinationDutyLocationPhoneLines {
			if phoneLine.Type == "voice" {
				destPhoneLines = append(destPhoneLines, phoneLine.Number)
			}
		}
		phone := result.Header.DestinationPhone
		suite.IsType(&edisegment.PER{}, phone)
		per := *phone
		suite.Equal("CN", per.ContactFunctionCode)
		suite.Equal("TE", per.CommunicationNumberQualifier)
		g := ghcPaymentRequestInvoiceGenerator{}
		phoneExpected, phoneExpectedErr := g.getPhoneNumberDigitsOnly(destPhoneLines[0])
		suite.NoError(phoneExpectedErr)
		suite.Equal(phoneExpected, per.CommunicationNumber)
	})

	suite.Run("adds orders origin address", func() {
		setupTestData()
		// name
		expectedDutyLocation := paymentRequest.MoveTaskOrder.Orders.OriginDutyLocation
		n1 := result.Header.OriginName
		suite.IsType(edisegment.N1{}, n1)
		suite.Equal("SF", n1.EntityIdentifierCode)
		suite.Equal(expectedDutyLocation.Name, n1.Name)
		suite.Equal("10", n1.IdentificationCodeQualifier)
		suite.Equal(expectedDutyLocation.TransportationOffice.Gbloc, n1.IdentificationCode)
		// street address
		address := expectedDutyLocation.Address
		n3Address := result.Header.OriginStreetAddress
		suite.IsType(&edisegment.N3{}, n3Address)
		n3 := *n3Address
		suite.Equal(address.StreetAddress1, n3.AddressInformation1)
		suite.Equal(*address.StreetAddress2, n3.AddressInformation2)
		// city state info
		n4 := result.Header.OriginPostalDetails
		suite.IsType(edisegment.N4{}, n4)
		if len(n4.CityName) >= maxCityLength {
			suite.Equal(address.City[:maxCityLength]+"...", n4.CityName)
		} else {
			suite.Equal(address.City, n4.CityName)
		}
		suite.Equal(address.State, n4.StateOrProvinceCode)
		suite.Equal(address.PostalCode, n4.PostalCode)
		countryCode, err := address.CountryCode()
		suite.NoError(err)
		suite.Equal(*countryCode, n4.CountryCode)
		// Office Phone
		originDutyLocationPhoneLines := expectedDutyLocation.TransportationOffice.PhoneLines
		var originPhoneLines []string
		for _, phoneLine := range originDutyLocationPhoneLines {
			if phoneLine.Type == "voice" {
				originPhoneLines = append(originPhoneLines, phoneLine.Number)
			}
		}
		phone := result.Header.OriginPhone
		suite.IsType(&edisegment.PER{}, phone)
		per := *phone
		suite.Equal("CN", per.ContactFunctionCode)
		suite.Equal("TE", per.CommunicationNumberQualifier)
		g := ghcPaymentRequestInvoiceGenerator{}
		phoneExpected, phoneExpectedErr := g.getPhoneNumberDigitsOnly(originPhoneLines[0])
		suite.NoError(phoneExpectedErr)
		suite.Equal(phoneExpected, per.CommunicationNumber)
	})

	suite.Run("adds various service item segments", func() {
		setupTestData()

		for idx, paymentServiceItem := range paymentServiceItems {
			var hierarchicalNumberInt = idx + 1
			var hierarchicalNumber = strconv.Itoa(hierarchicalNumberInt)
			segmentOffset := idx

			suite.Run("adds hl service item segment", func() {
				hl := result.ServiceItems[segmentOffset].HL
				suite.Equal(hierarchicalNumber, hl.HierarchicalIDNumber)
				suite.Equal(hierarchicalLevelCodeExpected, hl.HierarchicalLevelCode)
			})

			suite.Run("adds n9 service item segment", func() {
				n9 := result.ServiceItems[segmentOffset].N9
				suite.Equal("PO", n9.ReferenceIdentificationQualifier)
				suite.Equal(paymentServiceItem.ReferenceID, n9.ReferenceIdentification)
			})

			suite.Run("adds fa1 service item segment", func() {
				fa1 := result.ServiceItems[segmentOffset].FA1
				suite.Equal("DY", fa1.AgencyQualifierCode) // Default Order from testdatagen is AIR_FORCE
			})

			suite.Run("adds fa2 service item segment", func() {
				fa2 := result.ServiceItems[segmentOffset].FA2
				suite.Equal("TA", fa2.BreakdownStructureDetailCode)
				suite.Equal(*paymentRequest.MoveTaskOrder.Orders.TAC, fa2.FinancialInformationCode)
			})

			serviceItemPrice := paymentServiceItem.PriceCents.Int64()
			serviceCode := paymentServiceItem.MTOServiceItem.ReService.Code
			switch serviceCode {
			case models.ReServiceCodeCS, models.ReServiceCodeMS:
				suite.Run("adds l5 service item segment", func() {
					l5 := result.ServiceItems[segmentOffset].L5
					suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
					suite.Equal(string(serviceCode), l5.LadingDescription)
					suite.Equal("TBD", l5.CommodityCode)
					suite.Equal("D", l5.CommodityCodeQualifier)
				})

				suite.Run("adds l0 service item segment", func() {
					l0 := result.ServiceItems[segmentOffset].L0
					suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
					suite.Equal(float64(0), l0.BilledRatedAsQuantity)
					suite.Equal("", l0.BilledRatedAsQualifier)
					suite.Equal(float64(0), l0.Weight)
					suite.Equal("", l0.WeightQualifier)
					suite.Equal(float64(0), l0.Volume)
					suite.Equal("", l0.VolumeUnitQualifier)
					suite.Equal(0, l0.LadingQuantity)
					suite.Equal("", l0.PackagingFormCode)
					suite.Equal("", l0.WeightUnitCode)
				})

				suite.Run("adds l1 service item segment", func() {
					l1 := result.ServiceItems[segmentOffset].L1
					suite.Equal(hierarchicalNumberInt, l1.LadingLineItemNumber)
					suite.Equal(serviceItemPrice, l1.Charge)
				})
			case models.ReServiceCodeDOP, models.ReServiceCodeDUPK,
				models.ReServiceCodeDPK, models.ReServiceCodeDDP,
				models.ReServiceCodeDDFSIT, models.ReServiceCodeDDASIT,
				models.ReServiceCodeDOFSIT, models.ReServiceCodeDOASIT,
				models.ReServiceCodeDOSHUT, models.ReServiceCodeDDSHUT,
				models.ReServiceCodeDNPK:
				suite.Run("adds l5 service item segment", func() {
					l5 := result.ServiceItems[segmentOffset].L5
					suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
					suite.Equal(string(serviceCode), l5.LadingDescription)
					suite.Equal("TBD", l5.CommodityCode)
					suite.Equal("D", l5.CommodityCodeQualifier)
				})

				suite.Run("adds l0 service item segment", func() {
					l0 := result.ServiceItems[segmentOffset].L0
					suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
					suite.Equal(float64(0), l0.BilledRatedAsQuantity)
					suite.Equal("", l0.BilledRatedAsQualifier)
					suite.Equal(float64(4242), l0.Weight)
					suite.Equal("B", l0.WeightQualifier)
					suite.Equal(float64(0), l0.Volume)
					suite.Equal("", l0.VolumeUnitQualifier)
					suite.Equal(0, l0.LadingQuantity)
					suite.Equal("", l0.PackagingFormCode)
					suite.Equal("L", l0.WeightUnitCode)
				})

				suite.Run("adds l1 service item segment", func() {
					l1 := result.ServiceItems[segmentOffset].L1
					suite.Equal(hierarchicalNumberInt, l1.LadingLineItemNumber)
					suite.Equal(float64(4242), *l1.FreightRate)
					suite.Equal("LB", l1.RateValueQualifier)
					suite.Equal(serviceItemPrice, l1.Charge)
				})
			case models.ReServiceCodeDCRT, models.ReServiceCodeDUCRT:
				suite.Run("adds l5 service item segment", func() {
					l5 := result.ServiceItems[segmentOffset].L5
					suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)
					suite.Equal(string(serviceCode), l5.LadingDescription)
					suite.Equal("TBD", l5.CommodityCode)
					suite.Equal("D", l5.CommodityCodeQualifier)
				})

				suite.Run("adds l0 service item segment", func() {
					l0 := result.ServiceItems[segmentOffset].L0
					suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)
					suite.Equal(float64(0), l0.BilledRatedAsQuantity)
					suite.Equal("", l0.BilledRatedAsQualifier)
					suite.Equal(float64(0), l0.Weight)
					suite.Equal("", l0.WeightQualifier)
					suite.Equal(144.5, l0.Volume)
					suite.Equal("E", l0.VolumeUnitQualifier)
					suite.Equal(1, l0.LadingQuantity)
					suite.Equal("CRT", l0.PackagingFormCode)
					suite.Equal("", l0.WeightUnitCode)
				})

				suite.Run("adds l1 service item segment", func() {
					l1 := result.ServiceItems[segmentOffset].L1
					suite.Equal(hierarchicalNumberInt, l1.LadingLineItemNumber)
					suite.Equal(23.69, *l1.FreightRate)
					suite.Equal("PF", l1.RateValueQualifier)
					suite.Equal(serviceItemPrice, l1.Charge)
				})
			default:
				suite.Run("adds l5 service item segment", func() {
					l5 := result.ServiceItems[segmentOffset].L5
					suite.Equal(hierarchicalNumberInt, l5.LadingLineItemNumber)

					suite.Equal(string(serviceCode), l5.LadingDescription)
					suite.Equal("TBD", l5.CommodityCode)
					suite.Equal("D", l5.CommodityCodeQualifier)
				})

				suite.Run("adds l0 service item segment", func() {
					l0 := result.ServiceItems[segmentOffset].L0
					suite.Equal(hierarchicalNumberInt, l0.LadingLineItemNumber)

					switch serviceCode {
					case models.ReServiceCodeDSH:
						suite.Equal(float64(24246), l0.BilledRatedAsQuantity)
					case models.ReServiceCodeDDDSIT:
						suite.Equal(float64(44), l0.BilledRatedAsQuantity)
					case models.ReServiceCodeDOPSIT:
						suite.Equal(float64(33), l0.BilledRatedAsQuantity)
					default:
						suite.Equal(float64(24246), l0.BilledRatedAsQuantity)
					}
					suite.Equal("DM", l0.BilledRatedAsQualifier)
					suite.Equal(float64(4242), l0.Weight)
					suite.Equal("B", l0.WeightQualifier)
					suite.Equal(float64(0), l0.Volume)
					suite.Equal("", l0.VolumeUnitQualifier)
					suite.Equal(0, l0.LadingQuantity)
					suite.Equal("", l0.PackagingFormCode)
					suite.Equal("L", l0.WeightUnitCode)
				})
				suite.Run("adds l1 service item segment", func() {
					l1 := result.ServiceItems[segmentOffset].L1
					suite.Equal(hierarchicalNumberInt, l1.LadingLineItemNumber)
					suite.Equal(float64(4242), *l1.FreightRate)
					suite.Equal("LB", l1.RateValueQualifier)
					suite.Equal(serviceItemPrice, l1.Charge)
				})
			}
		}
	})

	suite.Run("adds l3 service item segment", func() {
		l3 := result.L3
		// Will need to be updated as more service items are supported
		suite.Equal(int64(17760), l3.PriceCents)
	})
}

func (suite *GHCInvoiceSuite) TestOnlyMsandCsGenerateEdi() {
	generator := NewGHCPaymentRequestInvoiceGenerator(suite.icnSequencer, clock.NewMock())
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
	}
	mto := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
	paymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
		Move: mto,
		PaymentRequest: models.PaymentRequest{
			IsFinal:         false,
			Status:          models.PaymentRequestStatusPending,
			RejectionReason: nil,
		},
	})

	assertions := testdatagen.Assertions{
		Move:           mto,
		PaymentRequest: paymentRequest,
		PaymentServiceItem: models.PaymentServiceItem{
			Status: models.PaymentServiceItemStatusApproved,
		},
	}

	testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeMS,
		basicPaymentServiceItemParams,
		assertions,
	)
	testdatagen.MakePaymentServiceItemWithParams(
		suite.DB(),
		models.ReServiceCodeCS,
		basicPaymentServiceItemParams,
		assertions,
	)

	_, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
	suite.NoError(err)
}

func (suite *GHCInvoiceSuite) TestNilValues() {
	mockClock := clock.NewMock()
	currentTime := mockClock.Now()
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
		{
			Key:     models.ServiceItemParamNameReferenceDate,
			KeyType: models.ServiceItemParamTypeDate,
			Value:   currentTime.Format(testDateFormat),
		},
		{
			Key:     models.ServiceItemParamNameWeightBilled,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "4242",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "24246",
		},
	}

	generator := NewGHCPaymentRequestInvoiceGenerator(suite.icnSequencer, mockClock)

	var nilPaymentRequest models.PaymentRequest
	setupTestData := func() {
		nilMove := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})

		nilPaymentRequest = testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
			Move: nilMove,
			PaymentRequest: models.PaymentRequest{
				IsFinal:         false,
				Status:          models.PaymentRequestStatusPending,
				RejectionReason: nil,
			},
		})

		assertions := testdatagen.Assertions{
			Move:           nilMove,
			PaymentRequest: nilPaymentRequest,
			PaymentServiceItem: models.PaymentServiceItem{
				Status: models.PaymentServiceItemStatusApproved,
			},
		}

		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDLH,
			basicPaymentServiceItemParams,
			assertions,
		)
	}

	// This won't work because we don't have PaymentServiceItems on the PaymentRequest right now.
	// nilPaymentRequest.PaymentServiceItems[0].PriceCents = nil

	panicFunc := func() {
		//RA Summary: gosec - errcheck - Unchecked return value
		//RA: Linter flags errcheck error: Ignoring a method's return value can cause the program to overlook unexpected states and conditions.
		//RA: Functions with unchecked return values in the file are used fetch data and assign data to a variable that is checked later on
		//RA: Given the return value is being checked in a different line and the functions that are flagged by the linter are being used to assign variables
		//RA: in a unit test, then there is no risk
		//RA Developer Status: Mitigated
		//RA Validator Status: Mitigated
		//RA Modified Severity: N/A
		// nolint:errcheck
		generator.Generate(suite.AppContextForTest(), nilPaymentRequest, false)
	}

	suite.Run("nil TAC does not cause panic", func() {
		setupTestData()
		oldTAC := nilPaymentRequest.MoveTaskOrder.Orders.TAC
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = oldTAC
	})

	suite.Run("empty TAC returns error", func() {
		setupTestData()
		oldTAC := nilPaymentRequest.MoveTaskOrder.Orders.TAC
		blank := ""
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = &blank
		_, err := generator.Generate(suite.AppContextForTest(), nilPaymentRequest, false)
		suite.Error(err)
		suite.IsType(apperror.ConflictError{}, err)
		suite.Equal(fmt.Sprintf("ID: %s is in a conflicting state Invalid order. Must have an HHG TAC value", nilPaymentRequest.MoveTaskOrder.OrdersID), err.Error())
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = oldTAC
	})

	suite.Run("nil TAC returns error", func() {
		setupTestData()
		oldTAC := nilPaymentRequest.MoveTaskOrder.Orders.TAC
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = nil
		_, err := generator.Generate(suite.AppContextForTest(), nilPaymentRequest, false)
		suite.Error(err)
		suite.IsType(apperror.ConflictError{}, err)
		suite.Equal(fmt.Sprintf("ID: %s is in a conflicting state Invalid order. Must have an HHG TAC value", nilPaymentRequest.MoveTaskOrder.OrdersID), err.Error())
		nilPaymentRequest.MoveTaskOrder.Orders.TAC = oldTAC
	})

	suite.Run("nil country for NewDutyLocation does not cause panic", func() {
		setupTestData()
		oldCountry := nilPaymentRequest.MoveTaskOrder.Orders.NewDutyLocation.Address.Country
		nilPaymentRequest.MoveTaskOrder.Orders.NewDutyLocation.Address.Country = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.NewDutyLocation.Address.Country = oldCountry
	})

	suite.Run("nil country for OriginDutyLocation does not cause panic", func() {
		setupTestData()
		oldCountry := nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyLocation.Address.Country
		nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyLocation.Address.Country = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.Orders.OriginDutyLocation.Address.Country = oldCountry
	})

	suite.Run("nil reference ID does not cause panic", func() {
		setupTestData()
		oldReferenceID := nilPaymentRequest.MoveTaskOrder.ReferenceID
		nilPaymentRequest.MoveTaskOrder.ReferenceID = nil
		suite.NotPanics(panicFunc)
		nilPaymentRequest.MoveTaskOrder.ReferenceID = oldReferenceID
	})

	// TODO: Needs some additional thought since PaymentServiceItems is loaded from the DB in Generate.
	//suite.Run("nil PriceCents does not cause panic", func() {
	//	oldPriceCents := nilPaymentRequest.PaymentServiceItems[0].PriceCents
	//	nilPaymentRequest.PaymentServiceItems[0].PriceCents = nil
	//	suite.NotPanics(panicFunc)
	//	nilPaymentRequest.PaymentServiceItems[0].PriceCents = oldPriceCents
	//})
}

func (suite *GHCInvoiceSuite) TestNoApprovedPaymentServiceItems() {
	generator := NewGHCPaymentRequestInvoiceGenerator(suite.icnSequencer, clock.NewMock())
	var result ediinvoice.Invoice858C
	var err error
	setupTestData := func() {

		basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
			{
				Key:     models.ServiceItemParamNameContractCode,
				KeyType: models.ServiceItemParamTypeString,
				Value:   testdatagen.DefaultContractCode,
			},
		}
		mto := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{})
		paymentRequest := testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
			Move: mto,
			PaymentRequest: models.PaymentRequest{
				IsFinal:         false,
				Status:          models.PaymentRequestStatusPending,
				RejectionReason: nil,
			},
		})

		assertions := testdatagen.Assertions{
			Move:               mto,
			PaymentRequest:     paymentRequest,
			PaymentServiceItem: models.PaymentServiceItem{},
		}

		assertions.PaymentServiceItem.Status = models.PaymentServiceItemStatusDenied
		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeMS,
			basicPaymentServiceItemParams,
			assertions,
		)

		assertions.PaymentServiceItem.Status = models.PaymentServiceItemStatusRequested
		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeCS,
			basicPaymentServiceItemParams,
			assertions,
		)

		assertions.PaymentServiceItem.Status = models.PaymentServiceItemStatusPaid
		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeCS,
			basicPaymentServiceItemParams,
			assertions,
		)

		assertions.PaymentServiceItem.Status = models.PaymentServiceItemStatusSentToGex
		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeCS,
			basicPaymentServiceItemParams,
			assertions,
		)

		result, err = generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.Error(err)
	}
	suite.Run("Service items that are not approved should be not added to invoice", func() {
		setupTestData()
		suite.Empty(result.ServiceItems)
	})

	suite.Run("Cost of service items that are not approved should not be included in L3", func() {
		setupTestData()
		l3 := result.L3
		suite.Equal(int64(0), l3.PriceCents)
	})
}

func (suite *GHCInvoiceSuite) TestTACs() {
	mockClock := clock.NewMock()
	currentTime := mockClock.Now()
	basicPaymentServiceItemParams := []testdatagen.CreatePaymentServiceItemParams{
		{
			Key:     models.ServiceItemParamNameContractCode,
			KeyType: models.ServiceItemParamTypeString,
			Value:   testdatagen.DefaultContractCode,
		},
		{
			Key:     models.ServiceItemParamNameReferenceDate,
			KeyType: models.ServiceItemParamTypeDate,
			Value:   currentTime.Format(testDateFormat),
		},
		{
			Key:     models.ServiceItemParamNameWeightBilled,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "4242",
		},
		{
			Key:     models.ServiceItemParamNameDistanceZip,
			KeyType: models.ServiceItemParamTypeInteger,
			Value:   "24246",
		},
	}

	generator := NewGHCPaymentRequestInvoiceGenerator(suite.icnSequencer, mockClock)

	hhgTAC := "1111"
	ntsTAC := "2222"

	var mtoShipment models.MTOShipment
	var paymentRequest models.PaymentRequest

	setupTestData := func() {
		orders := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
			Order: models.Order{
				TAC:    &hhgTAC,
				NtsTAC: &ntsTAC,
			},
		})

		move := testdatagen.MakeMove(suite.DB(), testdatagen.Assertions{
			Order: orders,
		})

		paymentRequest = testdatagen.MakePaymentRequest(suite.DB(), testdatagen.Assertions{
			Move: move,
			PaymentRequest: models.PaymentRequest{
				IsFinal: false,
				Status:  models.PaymentRequestStatusReviewed,
			},
		})

		mtoShipment = testdatagen.MakeMTOShipment(suite.DB(), testdatagen.Assertions{
			Move: move,
		})

		testdatagen.MakePaymentServiceItemWithParams(
			suite.DB(),
			models.ReServiceCodeDNPK,
			basicPaymentServiceItemParams,
			testdatagen.Assertions{
				Move:           move,
				MTOShipment:    mtoShipment,
				PaymentRequest: paymentRequest,
				PaymentServiceItem: models.PaymentServiceItem{
					Status: models.PaymentServiceItemStatusApproved,
				},
			},
		)
	}

	suite.Run("shipment with no TAC type set", func() {
		setupTestData()
		mtoShipment.TACType = nil
		suite.MustSave(&mtoShipment)

		result, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.NoError(err)
		suite.Equal(hhgTAC, result.ServiceItems[0].FA2.FinancialInformationCode)
	})

	suite.Run("shipment with HHG TAC type set", func() {
		setupTestData()
		tacType := models.LOATypeHHG
		mtoShipment.TACType = &tacType
		suite.MustSave(&mtoShipment)

		result, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.NoError(err)
		suite.Equal(hhgTAC, result.ServiceItems[0].FA2.FinancialInformationCode)
	})

	suite.Run("shipment with HHG TAC type set", func() {
		setupTestData()
		tacType := models.LOATypeNTS
		mtoShipment.TACType = &tacType
		suite.MustSave(&mtoShipment)

		result, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.NoError(err)
		suite.Equal(ntsTAC, result.ServiceItems[0].FA2.FinancialInformationCode)
	})

	suite.Run("shipment with NTS TAC type set", func() {
		setupTestData()
		tacType := models.LOATypeNTS
		mtoShipment.TACType = &tacType
		suite.MustSave(&mtoShipment)

		result, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.NoError(err)
		suite.Equal(ntsTAC, result.ServiceItems[0].FA2.FinancialInformationCode)
	})

	suite.Run("shipment with HHG TAC type set, but no HHG TAC", func() {
		setupTestData()
		tacType := models.LOATypeHHG
		mtoShipment.TACType = &tacType
		suite.MustSave(&mtoShipment)
		paymentRequest.MoveTaskOrder.Orders.TAC = nil
		suite.MustSave(&paymentRequest.MoveTaskOrder.Orders)

		_, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.Error(err)
		suite.Contains(err.Error(), "Must have an HHG TAC value")
	})

	suite.Run("shipment with NTS TAC type set, but no NTS TAC", func() {
		setupTestData()
		tacType := models.LOATypeNTS
		mtoShipment.TACType = &tacType
		suite.MustSave(&mtoShipment)
		paymentRequest.MoveTaskOrder.Orders.NtsTAC = nil
		suite.MustSave(&paymentRequest.MoveTaskOrder.Orders)

		_, err := generator.Generate(suite.AppContextForTest(), paymentRequest, false)
		suite.Error(err)
		suite.Contains(err.Error(), "Must have an NTS TAC value")
	})
}

func (suite *GHCInvoiceSuite) TestTruncateStrFunc() {
	longStr := "A super duper long string"
	expectedTruncatedStr := "A super..."
	suite.Equal(expectedTruncatedStr, truncateStr(longStr, 10))

	suite.Equal("AB", truncateStr("ABCD", 2))
	suite.Equal("ABC", truncateStr("ABCD", 3))
	suite.Equal("A...", truncateStr("ABCDEFGHI", 4))
	suite.Equal("ABC...", truncateStr("ABCDEFGHI", 6))
	suite.Equal("Too short", truncateStr("Too short", 200))
}
