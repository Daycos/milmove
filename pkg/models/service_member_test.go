package models_test

import (
	"time"

	"github.com/go-openapi/swag"
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/factory"
	"github.com/transcom/mymove/pkg/gen/internalmessages"
	. "github.com/transcom/mymove/pkg/models"
	"github.com/transcom/mymove/pkg/testdatagen"
)

func (suite *ModelSuite) TestBasicServiceMemberInstantiation() {
	servicemember := &ServiceMember{}

	expErrors := map[string][]string{
		"user_id": {"UserID can not be blank."},
	}

	suite.verifyValidationErrors(servicemember, expErrors)
}

func (suite *ModelSuite) TestIsProfileCompleteWithIncompleteSM() {
	// Given: a user and a service member
	lgu := uuid.Must(uuid.NewV4())
	user1 := User{
		LoginGovUUID:  &lgu,
		LoginGovEmail: "whoever@example.com",
	}
	verrs, err := user1.Validate(nil)
	suite.NoError(err)
	suite.False(verrs.HasAny(), "Error validating model")

	// And: a service member is incompletely initialized with almost all required values
	edipi := "12345567890"
	affiliation := AffiliationARMY
	rank := ServiceMemberRankE5
	firstName := "bob"
	lastName := "sally"
	telephone := "510 555-5555"
	email := "bobsally@gmail.com"
	fakeAddress := factory.BuildAddress(nil, []factory.Customization{
		{
			Model: Address{
				ID: uuid.Must(uuid.NewV4()),
			},
		},
	}, nil)
	fakeBackupAddress := factory.BuildAddress(nil, []factory.Customization{
		{
			Model: Address{
				ID: uuid.Must(uuid.NewV4()),
			},
		},
	}, nil)
	location := factory.BuildDutyLocation(nil, []factory.Customization{
		{
			Model: DutyLocation{
				ID: uuid.Must(uuid.NewV4()),
			},
		},
	}, nil)

	serviceMember := ServiceMember{
		ID:                     uuid.Must(uuid.NewV4()),
		UserID:                 user1.ID,
		Edipi:                  &edipi,
		Affiliation:            &affiliation,
		Rank:                   &rank,
		FirstName:              &firstName,
		LastName:               &lastName,
		Telephone:              &telephone,
		PersonalEmail:          &email,
		ResidentialAddressID:   &fakeAddress.ID,
		BackupMailingAddressID: &fakeBackupAddress.ID,
		DutyLocationID:         &location.ID,
	}

	suite.Equal(false, serviceMember.IsProfileComplete())

	// When: all required fields are set
	emailPreferred := true
	serviceMember.EmailIsPreferred = &emailPreferred

	contactAssertions := testdatagen.Assertions{
		BackupContact: BackupContact{
			ServiceMember:   serviceMember,
			ServiceMemberID: serviceMember.ID,
		},
		Stub: true,
	}
	backupContact := testdatagen.MakeBackupContact(suite.DB(), contactAssertions)
	serviceMember.BackupContacts = append(serviceMember.BackupContacts, backupContact)

	suite.Equal(true, serviceMember.IsProfileComplete())
}

func (suite *ModelSuite) TestFetchServiceMemberForUser() {
	user1 := factory.BuildDefaultUser(suite.DB())
	user2 := factory.BuildDefaultUser(suite.DB())

	firstName := "Oliver"
	resAddress := factory.BuildAddress(suite.DB(), nil, nil)
	sm := ServiceMember{
		User:                 user1,
		UserID:               user1.ID,
		FirstName:            &firstName,
		ResidentialAddressID: &resAddress.ID,
		ResidentialAddress:   &resAddress,
	}
	suite.MustSave(&sm)

	// User is authorized to fetch service member
	session := &auth.Session{
		ApplicationName: auth.MilApp,
		UserID:          user1.ID,
		ServiceMemberID: sm.ID,
	}
	goodSm, err := FetchServiceMemberForUser(suite.DB(), session, sm.ID)
	if suite.NoError(err) {
		suite.Equal(sm.FirstName, goodSm.FirstName)
		suite.Equal(sm.ResidentialAddress.ID, goodSm.ResidentialAddress.ID)
	}

	// Wrong ServiceMember
	wrongID, _ := uuid.NewV4()
	_, err = FetchServiceMemberForUser(suite.DB(), session, wrongID)
	if suite.Error(err) {
		suite.Equal(ErrFetchNotFound, err)
	}

	// User is forbidden from fetching order
	session.UserID = user2.ID
	session.ServiceMemberID = uuid.Nil
	_, err = FetchServiceMemberForUser(suite.DB(), session, sm.ID)
	if suite.Error(err) {
		suite.Equal(ErrFetchForbidden, err)
	}
}

func (suite *ModelSuite) TestFetchServiceMemberNotForUser() {
	user1 := factory.BuildDefaultUser(suite.DB())

	firstName := "Nino"
	resAddress := factory.BuildAddress(suite.DB(), nil, nil)
	sm := ServiceMember{
		User:                 user1,
		UserID:               user1.ID,
		FirstName:            &firstName,
		ResidentialAddressID: &resAddress.ID,
		ResidentialAddress:   &resAddress,
	}
	suite.MustSave(&sm)

	goodSm, err := FetchServiceMember(suite.DB(), sm.ID)
	if suite.NoError(err) {
		suite.Equal(sm.FirstName, goodSm.FirstName)
		suite.Equal(sm.ResidentialAddressID, goodSm.ResidentialAddressID)
	}
}

func (suite *ModelSuite) TestFetchLatestOrders() {
	setupTestData := func() (Order, *auth.Session) {

		user := factory.BuildDefaultUser(suite.DB())

		serviceMember := testdatagen.MakeDefaultServiceMember(suite.DB())

		dutyLocation := factory.FetchOrBuildCurrentDutyLocation(suite.DB())
		dutyLocation2 := factory.FetchOrBuildOrdersDutyLocation(suite.DB())
		issueDate := time.Date(2018, time.March, 10, 0, 0, 0, 0, time.UTC)
		reportByDate := time.Date(2018, time.August, 1, 0, 0, 0, 0, time.UTC)
		ordersType := internalmessages.OrdersTypePERMANENTCHANGEOFSTATION
		hasDependents := true
		spouseHasProGear := true
		uploadedOrder := Document{
			ServiceMember:   serviceMember,
			ServiceMemberID: serviceMember.ID,
		}
		deptIndicator := testdatagen.DefaultDepartmentIndicator
		TAC := testdatagen.DefaultTransportationAccountingCode
		suite.MustSave(&uploadedOrder)
		SAC := "N002214CSW32Y9"
		ordersNumber := "FD4534JFJ"

		order := Order{
			ServiceMemberID:      serviceMember.ID,
			ServiceMember:        serviceMember,
			IssueDate:            issueDate,
			ReportByDate:         reportByDate,
			OrdersType:           ordersType,
			HasDependents:        hasDependents,
			SpouseHasProGear:     spouseHasProGear,
			OriginDutyLocationID: &dutyLocation.ID,
			OriginDutyLocation:   &dutyLocation,
			NewDutyLocationID:    dutyLocation2.ID,
			NewDutyLocation:      dutyLocation2,
			UploadedOrdersID:     uploadedOrder.ID,
			UploadedOrders:       uploadedOrder,
			Status:               OrderStatusSUBMITTED,
			OrdersNumber:         &ordersNumber,
			TAC:                  &TAC,
			SAC:                  &SAC,
			DepartmentIndicator:  &deptIndicator,
			Grade:                swag.String("E-1"),
		}
		suite.MustSave(&order)

		// User is authorized to fetch service member
		session := &auth.Session{
			ApplicationName: auth.MilApp,
			UserID:          user.ID,
			ServiceMemberID: serviceMember.ID,
		}
		return order, session
	}

	suite.Run("successfully returns orders with uploads", func() {
		order, session := setupTestData()
		actualOrder, err := FetchLatestOrder(session, suite.DB())

		if suite.NoError(err) {
			suite.Equal(order.Grade, actualOrder.Grade)
			suite.Equal(order.OriginDutyLocationID, actualOrder.OriginDutyLocationID)
			suite.Equal(order.NewDutyLocationID, actualOrder.NewDutyLocationID)
			suite.True(order.IssueDate.Equal(actualOrder.IssueDate))
			suite.True(order.ReportByDate.Equal(actualOrder.ReportByDate))
			suite.Equal(order.OrdersType, actualOrder.OrdersType)
			suite.Equal(order.HasDependents, actualOrder.HasDependents)
			suite.Equal(order.SpouseHasProGear, actualOrder.SpouseHasProGear)
			suite.Equal(order.UploadedOrdersID, actualOrder.UploadedOrdersID)

		}

		// Wrong ServiceMember
		wrongID, _ := uuid.NewV4()
		_, err = FetchServiceMemberForUser(suite.DB(), session, wrongID)
		if suite.Error(err) {
			suite.Equal(ErrFetchNotFound, err)
		}
	})

	suite.Run("successfully returns orders without any existing uploads", func() {
		expectedOrder := testdatagen.MakeOrderWithoutUpload(suite.DB(), testdatagen.Assertions{})

		userSession := auth.Session{
			ApplicationName: auth.MilApp,
			UserID:          expectedOrder.ServiceMember.ID,
			ServiceMemberID: expectedOrder.ServiceMemberID,
		}

		actualOrder, err := FetchLatestOrder(&userSession, suite.DB())

		suite.NoError(err)
		suite.Equal(expectedOrder.ID, actualOrder.ID)
		suite.Len(actualOrder.UploadedOrders.UserUploads, 0)
	})

	suite.Run("successfully returns non deleted orders and amended orders uploads", func() {
		nonDeletedOrdersUpload := testdatagen.MakeUserUpload(suite.DB(), testdatagen.Assertions{})
		testdatagen.MakeUserUpload(suite.DB(), testdatagen.Assertions{
			UserUpload: UserUpload{
				Document:  nonDeletedOrdersUpload.Document,
				DeletedAt: TimePointer(time.Now()),
			},
		})

		nonDeletedAmendedUpload := testdatagen.MakeUserUpload(suite.DB(), testdatagen.Assertions{
			UserUpload: UserUpload{
				UploaderID: nonDeletedOrdersUpload.Document.ServiceMember.UserID,
			},
		})
		testdatagen.MakeUserUpload(suite.DB(), testdatagen.Assertions{
			UserUpload: UserUpload{
				Document:  nonDeletedAmendedUpload.Document,
				DeletedAt: TimePointer(time.Now()),
			},
		})

		expectedOrder := testdatagen.MakeOrder(suite.DB(), testdatagen.Assertions{
			Order: Order{
				ServiceMember:           nonDeletedOrdersUpload.Document.ServiceMember,
				ServiceMemberID:         nonDeletedOrdersUpload.Document.ServiceMemberID,
				UploadedOrders:          nonDeletedOrdersUpload.Document,
				UploadedOrdersID:        *nonDeletedOrdersUpload.DocumentID,
				UploadedAmendedOrders:   &nonDeletedAmendedUpload.Document,
				UploadedAmendedOrdersID: nonDeletedAmendedUpload.DocumentID,
			},
		})

		userSession := auth.Session{
			ApplicationName: auth.MilApp,
			UserID:          expectedOrder.ServiceMember.ID,
			ServiceMemberID: expectedOrder.ServiceMemberID,
		}

		actualOrder, err := FetchLatestOrder(&userSession, suite.DB())

		suite.NoError(err)
		suite.Len(actualOrder.UploadedOrders.UserUploads, 1)
		suite.Equal(actualOrder.UploadedOrders.UserUploads[0].ID, nonDeletedOrdersUpload.ID)
		suite.Len(actualOrder.UploadedAmendedOrders.UserUploads, 1)
		suite.Equal(actualOrder.UploadedAmendedOrders.UserUploads[0].ID, nonDeletedAmendedUpload.ID)
	})
}
