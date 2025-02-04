import React, { Suspense, lazy, useState, useEffect } from 'react';
import { Switch, useParams, Redirect, Route, useHistory, useLocation, matchPath } from 'react-router-dom';
import { useSelector } from 'react-redux';

import 'styles/office.scss';

import ServicesCounselorTabNav from 'components/Office/ServicesCounselingTabNav/ServicesCounselingTabNav';
import CustomerHeader from 'components/CustomerHeader';
import SystemError from 'components/SystemError';
import { servicesCounselingRoutes } from 'constants/routes';
import { useTXOMoveInfoQueries } from 'hooks/queries';
import LoadingPlaceholder from 'shared/LoadingPlaceholder';
import SomethingWentWrong from 'shared/SomethingWentWrong';

const ServicesCounselingMoveDocumentWrapper = lazy(() =>
  import('pages/Office/ServicesCounselingMoveDocumentWrapper/ServicesCounselingMoveDocumentWrapper'),
);
const ServicesCounselingMoveDetails = lazy(() =>
  import('pages/Office/ServicesCounselingMoveDetails/ServicesCounselingMoveDetails'),
);
const ServicesCounselingEditShipmentDetails = lazy(() =>
  import('pages/Office/ServicesCounselingEditShipmentDetails/ServicesCounselingEditShipmentDetails'),
);
const CustomerInfo = lazy(() => import('pages/Office/CustomerInfo/CustomerInfo'));
const ServicesCounselorCustomerSupportRemarks = lazy(() =>
  import('pages/Office/ServicesCounselorCustomerSupportRemarks/ServicesCounselorCustomerSupportRemarks'),
);
const MoveHistory = lazy(() => import('pages/Office/MoveHistory/MoveHistory'));
const ReviewDocuments = lazy(() => import('pages/Office/PPM/ReviewDocuments/ReviewDocuments'));
const ServicesCounselingReviewShipmentWeights = lazy(() =>
  import('pages/Office/ServicesCounselingReviewShipmentWeights/ServicesCounselingReviewShipmentWeights'),
);

const ServicesCounselingMoveInfo = () => {
  const [unapprovedShipmentCount, setUnapprovedShipmentCount] = useState(0);

  const [infoSavedAlert, setInfoSavedAlert] = useState(null);
  const { hasRecentError, traceId } = useSelector((state) => state.interceptor);
  const onInfoSavedUpdate = (alertType) => {
    if (alertType === 'error') {
      setInfoSavedAlert({
        alertType,
        message: 'Something went wrong, and your changes were not saved. Please try again later.',
      });
    } else {
      setInfoSavedAlert({
        alertType,
        message: 'Your changes were saved.',
      });
    }
  };

  const history = useHistory();
  useEffect(() => {
    // clear alert when route changes
    const unlisten = history.listen(() => {
      if (infoSavedAlert) {
        setInfoSavedAlert(null);
      }
    });
    return () => {
      unlisten();
    };
  }, [history, infoSavedAlert]);

  const { moveCode } = useParams();
  const { order, customerData, isLoading, isError } = useTXOMoveInfoQueries(moveCode);

  const { pathname } = useLocation();
  const hideNav =
    matchPath(pathname, {
      path: servicesCounselingRoutes.SHIPMENT_ADD_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.SHIPMENT_EDIT_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.ORDERS_EDIT_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.ALLOWANCES_EDIT_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.SHIPMENT_REVIEW_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.CUSTOMER_INFO_EDIT_PATH,
      exact: true,
    }) ||
    matchPath(pathname, {
      path: servicesCounselingRoutes.SHIPMENT_REVIEW_PATH,
      exact: true,
    });

  if (isLoading) return <LoadingPlaceholder />;
  if (isError) return <SomethingWentWrong />;

  return (
    <>
      <CustomerHeader order={order} customer={customerData} moveCode={moveCode} />
      {hasRecentError && (
        <SystemError>
          Something isn&apos;t working, but we&apos;re not sure what. Wait a minute and try again.
          <br />
          If that doesn&apos;t fix it, contact the{' '}
          <a href="https://move.mil/customer-service#technical-help-desk">Technical Help Desk</a> and give them this
          code: <strong>{traceId}</strong>
        </SystemError>
      )}

      {!hideNav && <ServicesCounselorTabNav unapprovedShipmentCount={unapprovedShipmentCount} moveCode={moveCode} />}

      <Suspense fallback={<LoadingPlaceholder />}>
        <Switch>
          {/* TODO - Routes not finalized, revisit */}
          <Route path={servicesCounselingRoutes.MOVE_VIEW_PATH} exact>
            <ServicesCounselingMoveDetails
              infoSavedAlert={infoSavedAlert}
              setUnapprovedShipmentCount={setUnapprovedShipmentCount}
            />
          </Route>

          <Route path={servicesCounselingRoutes.CUSTOMER_SUPPORT_REMARKS_PATH} exact>
            <ServicesCounselorCustomerSupportRemarks />
          </Route>

          <Route path={servicesCounselingRoutes.SHIPMENT_REVIEW_PATH} exact>
            <ReviewDocuments />
          </Route>

          <Route path={servicesCounselingRoutes.MOVE_HISTORY_PATH} exact>
            <MoveHistory moveCode={moveCode} />
          </Route>

          <Route
            path={[servicesCounselingRoutes.ALLOWANCES_EDIT_PATH, servicesCounselingRoutes.ORDERS_EDIT_PATH]}
            exact
          >
            <ServicesCounselingMoveDocumentWrapper />
          </Route>

          <Route path={servicesCounselingRoutes.CUSTOMER_INFO_EDIT_PATH} exact>
            <CustomerInfo
              ordersId={order.id}
              customer={customerData}
              isLoading={isLoading}
              isError={isError}
              onUpdate={onInfoSavedUpdate}
            />
          </Route>

          <Route
            path={servicesCounselingRoutes.SHIPMENT_EDIT_PATH}
            exact
            // eslint-disable-next-line react/jsx-props-no-spreading
            render={(props) => <ServicesCounselingEditShipmentDetails {...props} onUpdate={onInfoSavedUpdate} />}
          />

          <Route
            path={servicesCounselingRoutes.SHIPMENT_ADVANCE_PATH}
            exact
            render={(props) => (
              // eslint-disable-next-line react/jsx-props-no-spreading
              <ServicesCounselingEditShipmentDetails {...props} onUpdate={onInfoSavedUpdate} isAdvancePage />
            )}
          />

          <Route
            path={servicesCounselingRoutes.SHIPMENT_REVIEW_PATH}
            exact
            render={(props) => (
              // eslint-disable-next-line react/jsx-props-no-spreading
              <ReviewDocuments {...props} onUpdate={onInfoSavedUpdate} />
            )}
          />

          <Route
            path={servicesCounselingRoutes.REVIEW_SHIPMENT_WEIGHTS_PATH}
            exact
            render={() => <ServicesCounselingReviewShipmentWeights moveCode={moveCode} />}
          />

          {/* TODO - clarify role/tab access */}
          <Redirect from={servicesCounselingRoutes.BASE_MOVE_PATH} to={servicesCounselingRoutes.MOVE_VIEW_PATH} />
        </Switch>
      </Suspense>
    </>
  );
};

export default ServicesCounselingMoveInfo;
