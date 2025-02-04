/* eslint-disable react/jsx-props-no-spreading */
import React from 'react';
import { shallow, mount } from 'enzyme';
import { queryByTestId, render, screen } from '@testing-library/react';

import ConnectedOffice, { OfficeApp } from './index';

import { MockProviders } from 'testUtils';
import { roleTypes } from 'constants/userRoles';

describe('Office App', () => {
  const mockOfficeProps = {
    loadUser: jest.fn(),
    loadInternalSchema: jest.fn(),
    loadPublicSchema: jest.fn(),
    logOut: jest.fn(),
    hasRecentError: false,
    traceId: '',
  };

  describe('component', () => {
    let wrapper;

    beforeEach(() => {
      wrapper = shallow(<OfficeApp {...mockOfficeProps} />);
    });

    it('renders without crashing or erroring', () => {
      const officeWrapper = wrapper.find('div');
      expect(officeWrapper).toBeDefined();
      expect(wrapper.find('SomethingWentWrong')).toHaveLength(0);
    });

    it('renders the logged out header by default', () => {
      expect(wrapper.find('LoggedOutHeader')).toHaveLength(1);
    });

    it('fetches initial data', () => {
      expect(mockOfficeProps.loadUser).toHaveBeenCalled();
      expect(mockOfficeProps.loadInternalSchema).toHaveBeenCalled();
      expect(mockOfficeProps.loadPublicSchema).toHaveBeenCalled();
    });

    describe('if an error occurs', () => {
      it('renders the fail whale', () => {
        wrapper.setState({ hasError: true });
        expect(wrapper.find('SomethingWentWrong')).toHaveLength(1);
      });
    });
  });

  describe('header with TOO user name and GBLOC', () => {
    const officeUserState = {
      auth: {
        activeRole: roleTypes.TOO,
        isLoading: false,
        isLoggedIn: true,
      },
      entities: {
        user: {
          userId123: {
            id: 'userId123',
            roles: [{ roleType: roleTypes.TOO }],
            office_user: {
              first_name: 'Amanda',
              last_name: 'Gorman',
              transportation_office: {
                gbloc: 'ABCD',
              },
            },
          },
        },
      },
    };

    describe('after signing in', () => {
      it('renders the header with the office user name and GBLOC', () => {
        const app = mount(
          <MockProviders initialState={officeUserState} initialEntries={['/moves/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );
        expect(app.containsMatchingElement(<a href="/">ABCD moves</a>)).toEqual(true);
        expect(app.containsMatchingElement(<span>Gorman, Amanda</span>)).toEqual(true);
      });
      it('renders the system error component if there is an unexpected error and on the queue page', () => {
        render(
          <MockProviders
            initialState={{ ...officeUserState, interceptor: { hasRecentError: true, traceId: 'some-trace-id' } }}
            initialEntries={['/']}
          >
            <ConnectedOffice />
          </MockProviders>,
        );
        expect(screen.getByText('Technical Help Desk').closest('a')).toHaveAttribute(
          'href',
          'https://move.mil/customer-service#technical-help-desk',
        );
        expect(screen.getByTestId('system-error').textContent).toEqual(
          "Something isn't working, but we're not sure what. Wait a minute and try again.If that doesn't fix it, contact the Technical Help Desk and give them this code: some-trace-id",
        );
      });
      it('does not render system error if it is not on the queue page', () => {
        render(
          <MockProviders
            initialState={{ ...officeUserState, interceptor: { hasRecentError: true, traceId: 'some-trace-id' } }}
            initialEntries={['/sign-in']}
          >
            <ConnectedOffice />
          </MockProviders>,
        );
        expect(queryByTestId(document.documentElement, 'system-error')).not.toBeInTheDocument();
      });
    });
  });

  describe('header with TIO user name and GBLOC', () => {
    const officeUserState = {
      auth: {
        activeRole: roleTypes.TIO,
        isLoading: false,
        isLoggedIn: true,
      },
      entities: {
        user: {
          userId123: {
            id: 'userId123',
            roles: [{ roleType: roleTypes.TIO }],
            office_user: {
              first_name: 'Amanda',
              last_name: 'Gorman',
              transportation_office: {
                gbloc: 'ABCD',
              },
            },
          },
        },
      },
      interceptor: {
        hasRecentError: false,
        timestamp: 0,
        traceId: '',
      },
    };

    describe('after signing in', () => {
      it('renders the header with the office user name and GBLOC', () => {
        const app = mount(
          <MockProviders initialState={officeUserState} initialEntries={['/moves/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        expect(app.containsMatchingElement(<a href="/">ABCD payment requests</a>)).toEqual(true);
        expect(app.containsMatchingElement(<span>Gorman, Amanda</span>)).toEqual(true);
      });
    });
  });

  describe('if the user is logged in with multiple roles', () => {
    const multiRoleState = {
      auth: {
        activeRole: roleTypes.TOO,
        isLoading: false,
        isLoggedIn: true,
      },
      entities: {
        user: {
          userId123: {
            id: 'userId123',
            roles: [
              { roleType: roleTypes.CONTRACTING_OFFICER },
              { roleType: roleTypes.TOO },
              { roleType: roleTypes.TIO },
            ],
          },
        },
      },
    };

    describe('on a page that isn’t the Select Application page', () => {
      it('renders the Select Application link', () => {
        const app = mount(
          <MockProviders initialState={multiRoleState} initialEntries={['/']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        expect(app.containsMatchingElement(<a href="/select-application">Change user role</a>)).toEqual(true);
      });
    });

    describe('on the Select Application page', () => {
      it('does not render the Select Application link', () => {
        const app = mount(
          <MockProviders initialState={multiRoleState} initialEntries={['/select-application']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        expect(app.containsMatchingElement(<a href="/select-application">Change user role</a>)).toEqual(false);
      });
    });
  });

  describe('routing', () => {
    // TODO - expects should look for actual component content instead of the route path
    // Might have to add testing-library for this because something about enzyme + Suspense + routes are not rendering content
    // I FIGURED OUT HOW - need to mock the loadUser (this sets loading back to true and prevents content from rendering)

    const loggedInState = {
      auth: {
        activeRole: roleTypes.TOO,
        isLoading: false,
        isLoggedIn: true,
      },
      entities: {
        user: {
          userId123: {
            id: 'userId123',
            roles: [{ roleType: roleTypes.TOO }],
          },
        },
      },
    };

    const loggedOutState = {
      auth: {
        activeRole: null,
        isLoading: false,
        isLoggedIn: false,
      },
    };

    it('handles the SignIn URL', () => {
      const app = mount(
        <MockProviders initialState={loggedOutState} initialEntries={['/sign-in']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('Route');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/sign-in');
    });

    it('handles the root URL', () => {
      const app = mount(
        <MockProviders initialState={loggedInState} initialEntries={['/']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('PrivateRoute');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/');
    });

    it('handles the Select Application URL', () => {
      const app = mount(
        <MockProviders initialState={loggedInState} initialEntries={['/select-application']}>
          <ConnectedOffice />
        </MockProviders>,
      );

      const renderedRoute = app.find('PrivateRoute');
      expect(renderedRoute).toHaveLength(1);
      expect(renderedRoute.prop('path')).toEqual('/select-application');
    });

    describe('TOO routes', () => {
      const loggedInTOOState = {
        auth: {
          activeRole: roleTypes.TOO,
          isLoading: false,
          isLoggedIn: true,
        },
        entities: {
          user: {
            userId123: {
              id: 'userId123',
              roles: [{ roleType: roleTypes.TOO }],
            },
          },
        },
      };

      it('handles the moves queue URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTOOState} initialEntries={['/moves/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/queue');
      });

      it('handles the TXOMoveInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTOOState} initialEntries={['/moves/AU67C6']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveCode');
      });

      it('handles the edit shipment details URL', () => {
        const app = mount(
          <MockProviders
            initialState={loggedInTOOState}
            initialEntries={['/moves/AU67C6/shipments/c73d3fbd-8a93-4bd9-8c0b-99bd52e45b2c']}
          >
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveCode/shipments/:shipmentId');
      });
    });

    describe('TIO routes', () => {
      const loggedInTIOState = {
        auth: {
          activeRole: roleTypes.TIO,
          isLoading: false,
          isLoggedIn: true,
        },
        entities: {
          user: {
            userId123: {
              id: 'userId123',
              roles: [{ roleType: roleTypes.TIO }],
            },
          },
        },
      };

      it('handles the invoicing queue URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTIOState} initialEntries={['/invoicing/queue']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/invoicing/queue');
      });

      it('handles the TXOMoveInfo URL', () => {
        const app = mount(
          <MockProviders initialState={loggedInTIOState} initialEntries={['/moves/AU67C6']}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual('/moves/:moveCode');
      });

      it('Tio should not render edit shipment details URL', () => {
        const app = mount(
          <MockProviders
            initialState={loggedInTIOState}
            initialEntries={['/moves/AU67C6/shipments/c73d3fbd-8a93-4bd9-8c0b-99bd52e45b2c']}
          >
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(0);
      });
    });

    describe('Services Counselor routes', () => {
      const loggedInServicesCounselorState = {
        auth: {
          activeRole: roleTypes.SERVICES_COUNSELOR,
          isLoading: false,
          isLoggedIn: true,
        },
        entities: {
          user: {
            userId123: {
              id: 'userId123',
              roles: [{ roleType: roleTypes.SERVICES_COUNSELOR }],
            },
          },
        },
      };

      it.each([['ServicesCounselingMoveInfo', '/counseling/moves/AU67C6', '/counseling/moves/:moveCode']])(
        'handles a %s URL (%s) with a given path of %s',
        (pageName, initialURL, pathToMatch) => {
          const app = mount(
            <MockProviders initialState={loggedInServicesCounselorState} initialEntries={[initialURL]}>
              <ConnectedOffice />
            </MockProviders>,
          );

          const renderedRoute = app.find('PrivateRoute');
          expect(renderedRoute).toHaveLength(1);
          expect(renderedRoute.prop('path')).toEqual(pathToMatch);
        },
      );
    });

    describe('Prime Simulator routes', () => {
      const loggedInPrimeSimulatorState = {
        auth: {
          activeRole: roleTypes.PRIME_SIMULATOR,
          isLoading: false,
          isLoggedIn: true,
        },
        entities: {
          user: {
            userId123: {
              id: 'userId123',
              roles: [{ roleType: roleTypes.PRIME_SIMULATOR }],
            },
          },
        },
      };

      it.each([
        ['PrimeSimulatorMoveDetails', '/simulator/moves/AU67C6/details', '/simulator/moves/:moveCodeOrID/details'],
        [
          'PrimeSimulatorCreateShipment',
          '/simulator/moves/AU67C6/shipments/new',
          '/simulator/moves/:moveCodeOrID/shipments/new',
        ],
        [
          'PrimeSimulatorUpdateShipment',
          '/simulator/moves/AU67C6/shipments/c73d3fbd-8a93-4bd9-8c0b-99bd52e45b2c',
          '/simulator/moves/:moveCodeOrID/shipments/:shipmentId',
        ],
        [
          'PrimeSimulatorCreatePaymentRequest',
          '/simulator/moves/AU67C6/payment-requests/new',
          '/simulator/moves/:moveCodeOrID/payment-requests/new',
        ],
        [
          'PrimeSimulatorUpdateAddress',
          '/simulator/moves/AU67C6/shipments/c73d3fbd-8a93-4bd9-8c0b-99bd52e45b2c/addresses/update',
          '/simulator/moves/:moveCodeOrID/shipments/:shipmentId/addresses/update',
        ],
        [
          'PrimeSimulatorCreatePaymentRequest',
          '/simulator/moves/AU67C6/payment-requests/new',
          '/simulator/moves/:moveCodeOrID/payment-requests/new',
        ],
        [
          'PrimeSimulatorCreateServiceItem',
          '/simulator/moves/AU67C6/shipments/c73d3fbd-8a93-4bd9-8c0b-99bd52e45b2c/service-items/new',
          '/simulator/moves/:moveCodeOrID/shipments/:shipmentId/service-items/new',
        ],
      ])('handles a %s URL (%s) with a given path of %s', (pageName, initialURL, pathToMatch) => {
        const app = mount(
          <MockProviders initialState={loggedInPrimeSimulatorState} initialEntries={[initialURL]}>
            <ConnectedOffice />
          </MockProviders>,
        );

        const renderedRoute = app.find('PrivateRoute');
        expect(renderedRoute).toHaveLength(1);
        expect(renderedRoute.prop('path')).toEqual(pathToMatch);
      });
    });

    describe('QAE/CSR Routes', () => {
      const loggedInQAECSRState = {
        auth: {
          activeRole: roleTypes.QAE_CSR,
          isLoading: false,
          isLoggedIn: true,
        },
        entities: {
          user: {
            userId123: {
              id: 'userId123',
              roles: [{ roleType: roleTypes.QAE_CSR }],
            },
          },
        },
      };

      it.each([['QAECSRMoveSearch', '/qaecsr/search', '/qaecsr/search']])(
        'handles a %s URL (%s) with a given path of %s',
        (pageName, initialURL, pathToMatch) => {
          const app = mount(
            <MockProviders initialState={loggedInQAECSRState} initialEntries={[initialURL]}>
              <ConnectedOffice />
            </MockProviders>,
          );

          const renderedRoute = app.find('PrivateRoute');
          expect(renderedRoute).toHaveLength(1);
          expect(renderedRoute.prop('path')).toEqual(pathToMatch);
        },
      );
    });

    describe('page not found route', () => {
      it('handles a nonexistent route by returning a 404 page', () => {
        render(
          <MockProviders initialEntries={['/pageNotFound']}>
            <ConnectedOffice />
          </MockProviders>,
        );
        expect(screen.getByText('Error - 404')).toBeInTheDocument();
      });
    });
  });
});
