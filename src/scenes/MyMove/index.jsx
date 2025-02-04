import React, { Component, lazy } from 'react';
import PropTypes from 'prop-types';
import { LastLocationProvider } from 'react-router-last-location';
import { Route, Switch } from 'react-router-dom';
import { push, goBack } from 'connected-react-router';
import { connect } from 'react-redux';
import { GovBanner } from '@trussworks/react-uswds';

import 'styles/full_uswds.scss';
import 'styles/customer.scss';

import BypassBlock from 'components/BypassBlock';
import CUIHeader from 'components/CUIHeader/CUIHeader';
import LoggedOutHeader from 'containers/Headers/LoggedOutHeader';
import CustomerLoggedInHeader from 'containers/Headers/CustomerLoggedInHeader';
import Alert from 'shared/Alert';
import Footer from 'components/Customer/Footer';
import ConnectedLogoutOnInactivity from 'layout/LogoutOnInactivity';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import CustomerPrivateRoute from 'containers/CustomerPrivateRoute/CustomerPrivateRoute';
import { getWorkflowRoutes } from './getWorkflowRoutes';
import { loadInternalSchema } from 'shared/Swagger/ducks';
import { withContext } from 'shared/AppContext';
import { no_op } from 'shared/utils';
import { loadUser as loadUserAction } from 'store/auth/actions';
import { initOnboarding as initOnboardingAction } from 'store/onboarding/actions';
import { selectIsLoggedIn } from 'store/auth/selectors';
import { selectConusStatus } from 'store/onboarding/selectors';
import {
  selectServiceMemberFromLoggedInUser,
  selectCurrentMove,
  selectHasCanceledMove,
} from 'store/entities/selectors';
import { generalRoutes, customerRoutes } from 'constants/routes';
/** Pages */
import InfectedUpload from 'shared/Uploader/InfectedUpload';
import ProcessingUpload from 'shared/Uploader/ProcessingUpload';
import PpmLanding from 'scenes/PpmLanding';
import Edit from 'scenes/Review/Edit';
import EditProfile from 'scenes/Review/EditProfile';
import EditDateAndLocation from 'scenes/Review/EditDateAndLocation';
import EditWeight from 'scenes/Review/EditWeight';
import PPMPaymentRequestIntro from 'scenes/Moves/Ppm/PPMPaymentRequestIntro';
import WeightTicket from 'scenes/Moves/Ppm/WeightTicket';
import ExpensesLanding from 'scenes/Moves/Ppm/ExpensesLanding';
import ExpensesUpload from 'scenes/Moves/Ppm/ExpensesUpload';
import AllowableExpenses from 'scenes/Moves/Ppm/AllowableExpenses';
import WeightTicketExamples from 'scenes/Moves/Ppm/WeightTicketExamples';
import NotFound from 'components/NotFound/NotFound';
import PrivacyPolicyStatement from 'shared/Statements/PrivacyAndPolicyStatement';
import AccessibilityStatement from 'shared/Statements/AccessibilityStatement';
import TrailerCriteria from 'scenes/Moves/Ppm/TrailerCriteria';
import PaymentReview from 'scenes/Moves/Ppm/PaymentReview/index';
import CustomerAgreementLegalese from 'scenes/Moves/Ppm/CustomerAgreementLegalese';
import ConnectedCreateOrEditMtoShipment from 'pages/MyMove/CreateOrEditMtoShipment';
import Home from 'pages/MyMove/Home';
// Pages should be lazy-loaded (they correspond to unique routes & only need to be loaded when that URL is accessed)
const SignIn = lazy(() => import('pages/SignIn/SignIn'));
const InvalidPermissions = lazy(() => import('pages/InvalidPermissions/InvalidPermissions'));
const MovingInfo = lazy(() => import('pages/MyMove/MovingInfo'));
const EditServiceInfo = lazy(() => import('pages/MyMove/Profile/EditServiceInfo'));
const Profile = lazy(() => import('pages/MyMove/Profile/Profile'));
const EditContactInfo = lazy(() => import('pages/MyMove/Profile/EditContactInfo'));
const AmendOrders = lazy(() => import('pages/MyMove/AmendOrders/AmendOrders'));
const EditOrders = lazy(() => import('pages/MyMove/EditOrders'));
const EstimatedWeightsProGear = lazy(() =>
  import('pages/MyMove/PPM/Booking/EstimatedWeightsProGear/EstimatedWeightsProGear'),
);
const EstimatedIncentive = lazy(() => import('pages/MyMove/PPM/Booking/EstimatedIncentive/EstimatedIncentive'));
const Advance = lazy(() => import('pages/MyMove/PPM/Booking/Advance/Advance'));
const About = lazy(() => import('pages/MyMove/PPM/Closeout/About/About'));
const WeightTickets = lazy(() => import('pages/MyMove/PPM/Closeout/WeightTickets/WeightTickets'));
const PPMReview = lazy(() => import('pages/MyMove/PPM/Closeout/Review/Review'));
const ProGear = lazy(() => import('pages/MyMove/PPM/Closeout/ProGear/ProGear.jsx'));
const Expenses = lazy(() => import('pages/MyMove/PPM/Closeout/Expenses/Expenses'));
const PPMFinalCloseout = lazy(() => import('pages/MyMove/PPM/Closeout/FinalCloseout/FinalCloseout'));

export class CustomerApp extends Component {
  constructor(props) {
    super(props);

    this.state = { hasError: false, error: undefined, info: undefined };
  }

  componentDidMount() {
    const { loadUser, initOnboarding, loadInternalSchema } = this.props;

    loadInternalSchema();
    loadUser();
    initOnboarding();
  }

  componentDidCatch(error, info) {
    this.setState({
      hasError: true,
      error,
      info,
    });
  }

  render() {
    const props = this.props;
    const { userIsLoggedIn } = this.props;
    const { hasError } = this.state;

    return (
      <>
        <LastLocationProvider>
          <div className="my-move site" id="app-root">
            <CUIHeader />
            <BypassBlock />
            <GovBanner />

            {userIsLoggedIn ? <CustomerLoggedInHeader /> : <LoggedOutHeader />}

            <main role="main" className="site__content my-move-container" id="main">
              <ConnectedLogoutOnInactivity />

              <div className="usa-grid">
                {props.swaggerError && (
                  <div className="grid-container">
                    <div className="grid-row">
                      <div className="grid-col-12">
                        <Alert type="error" heading="An error occurred">
                          There was an error contacting the server.
                        </Alert>
                      </div>
                    </div>
                  </div>
                )}
              </div>

              {hasError && <SomethingWentWrong />}

              {!hasError && !props.swaggerError && (
                <Switch>
                  {/* no auth */}
                  <Route path={generalRoutes.SIGN_IN_PATH} component={SignIn} />
                  <Route path={generalRoutes.PRIVACY_SECURITY_POLICY_PATH} component={PrivacyPolicyStatement} />
                  <Route path={generalRoutes.ACCESSIBILITY_PATH} component={AccessibilityStatement} />

                  {/* auth required */}
                  <CustomerPrivateRoute exact path="/ppm" component={PpmLanding} />

                  {/* ROOT */}
                  <CustomerPrivateRoute path={generalRoutes.HOME_PATH} exact component={Home} />

                  {getWorkflowRoutes(props)}
                  <CustomerPrivateRoute exact path={customerRoutes.SHIPMENT_MOVING_INFO_PATH} component={MovingInfo} />
                  <CustomerPrivateRoute exact path="/moves/:moveId/edit" component={Edit} />
                  <CustomerPrivateRoute exact path={customerRoutes.EDIT_PROFILE_PATH} component={EditProfile} />
                  <CustomerPrivateRoute
                    exact
                    path={customerRoutes.SERVICE_INFO_EDIT_PATH}
                    component={EditServiceInfo}
                  />
                  <CustomerPrivateRoute
                    path={customerRoutes.SHIPMENT_CREATE_PATH}
                    component={ConnectedCreateOrEditMtoShipment}
                  />
                  <CustomerPrivateRoute exact path={customerRoutes.PROFILE_PATH} component={Profile} />
                  <CustomerPrivateRoute
                    exact
                    path={customerRoutes.SHIPMENT_EDIT_PATH}
                    component={ConnectedCreateOrEditMtoShipment}
                  />
                  <CustomerPrivateRoute
                    path={customerRoutes.SHIPMENT_PPM_ESTIMATED_WEIGHT_PATH}
                    component={EstimatedWeightsProGear}
                  />
                  <CustomerPrivateRoute
                    exact
                    path={customerRoutes.SHIPMENT_PPM_ESTIMATED_INCENTIVE_PATH}
                    component={EstimatedIncentive}
                  />
                  <CustomerPrivateRoute exact path={customerRoutes.SHIPMENT_PPM_ADVANCES_PATH} component={Advance} />
                  <CustomerPrivateRoute
                    exact
                    path={customerRoutes.CONTACT_INFO_EDIT_PATH}
                    component={EditContactInfo}
                  />
                  <CustomerPrivateRoute exact path={customerRoutes.SHIPMENT_PPM_ABOUT_PATH} component={About} />
                  <CustomerPrivateRoute
                    exact
                    path={[
                      customerRoutes.SHIPMENT_PPM_WEIGHT_TICKETS_PATH,
                      customerRoutes.SHIPMENT_PPM_WEIGHT_TICKETS_EDIT_PATH,
                    ]}
                    component={WeightTickets}
                  />
                  <CustomerPrivateRoute exact path={customerRoutes.SHIPMENT_PPM_REVIEW_PATH} component={PPMReview} />
                  <CustomerPrivateRoute
                    exact
                    path={[customerRoutes.SHIPMENT_PPM_EXPENSES_PATH, customerRoutes.SHIPMENT_PPM_EXPENSES_EDIT_PATH]}
                    component={Expenses}
                  />
                  <CustomerPrivateRoute
                    exact
                    path={customerRoutes.SHIPMENT_PPM_COMPLETE_PATH}
                    component={PPMFinalCloseout}
                  />
                  <CustomerPrivateRoute path={customerRoutes.ORDERS_EDIT_PATH} component={EditOrders} />
                  <CustomerPrivateRoute path={customerRoutes.ORDERS_AMEND_PATH} component={AmendOrders} />
                  <CustomerPrivateRoute
                    path="/moves/:moveId/review/edit-date-and-location"
                    component={EditDateAndLocation}
                  />
                  <CustomerPrivateRoute path="/moves/:moveId/review/edit-weight" component={EditWeight} />
                  <CustomerPrivateRoute exact path="/weight-ticket-examples" component={WeightTicketExamples} />
                  <CustomerPrivateRoute exact path="/trailer-criteria" component={TrailerCriteria} />
                  <CustomerPrivateRoute exact path="/allowable-expenses" component={AllowableExpenses} />
                  <CustomerPrivateRoute exact path="/infected-upload" component={InfectedUpload} />
                  <CustomerPrivateRoute exact path="/processing-upload" component={ProcessingUpload} />
                  <CustomerPrivateRoute
                    path="/moves/:moveId/ppm-payment-request-intro"
                    component={PPMPaymentRequestIntro}
                  />
                  <CustomerPrivateRoute path="/moves/:moveId/ppm-weight-ticket" component={WeightTicket} />
                  <CustomerPrivateRoute path="/moves/:moveId/ppm-expenses-intro" component={ExpensesLanding} />
                  <CustomerPrivateRoute path="/moves/:moveId/ppm-expenses" component={ExpensesUpload} />
                  <CustomerPrivateRoute path="/moves/:moveId/ppm-payment-review" component={PaymentReview} />
                  <CustomerPrivateRoute
                    exact
                    path={[customerRoutes.SHIPMENT_PPM_PRO_GEAR_PATH, customerRoutes.SHIPMENT_PPM_PRO_GEAR_EDIT_PATH]}
                    component={ProGear}
                  />
                  <CustomerPrivateRoute exact path="/ppm-customer-agreement" component={CustomerAgreementLegalese} />

                  {/* Errors */}
                  <Route exact path="/forbidden">
                    <div className="usa-grid">
                      <h2>You are forbidden to use this endpoint</h2>
                    </div>
                  </Route>
                  <Route exact path="/server_error">
                    <div className="usa-grid">
                      <h2>We are experiencing an internal server error</h2>
                    </div>
                  </Route>
                  <Route exact path="/invalid-permissions" component={InvalidPermissions} />

                  {/* 404 */}
                  <Route render={(routeProps) => <NotFound {...routeProps} handleOnClick={this.props.goBack} />} />
                </Switch>
              )}
            </main>
            <Footer />
          </div>
          <div id="modal-root"></div>
        </LastLocationProvider>
      </>
    );
  }
}

CustomerApp.propTypes = {
  loadInternalSchema: PropTypes.func,
  loadUser: PropTypes.func,
  initOnboarding: PropTypes.func,
  userIsLoggedIn: PropTypes.bool,
  conusStatus: PropTypes.string,
  context: PropTypes.shape({
    flags: PropTypes.shape({
      hhgFlow: PropTypes.bool,
      ghcFlow: PropTypes.bool,
    }),
  }).isRequired,
};

CustomerApp.defaultProps = {
  loadInternalSchema: no_op,
  loadUser: no_op,
  initOnboarding: no_op,
  userIsLoggedIn: false,
  conusStatus: '',
  context: {
    flags: {
      hhgFlow: false,
      ghcFlow: false,
    },
  },
};

const mapStateToProps = (state) => {
  const serviceMember = selectServiceMemberFromLoggedInUser(state);
  const serviceMemberId = serviceMember?.id;
  const move = selectCurrentMove(state) || {};

  return {
    userIsLoggedIn: selectIsLoggedIn(state),
    currentServiceMemberId: serviceMemberId,
    lastMoveIsCanceled: selectHasCanceledMove(state),
    moveId: move?.id,
    conusStatus: selectConusStatus(state),
    swaggerError: state.swaggerInternal.hasErrored,
  };
};
const mapDispatchToProps = {
  goBack,
  push,
  loadInternalSchema,
  loadUser: loadUserAction,
  initOnboarding: initOnboardingAction,
};

export default withContext(connect(mapStateToProps, mapDispatchToProps)(CustomerApp));
