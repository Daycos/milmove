import React from 'react';
import { shallow, mount } from 'enzyme';
import { render, screen, waitFor } from '@testing-library/react';

import ConnectedCustomerApp, { CustomerApp } from './index';

import Footer from 'components/Customer/Footer';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import { MockProviders } from 'testUtils';
import { AppContext } from 'shared/AppContext';

describe('ConnectedCustomerApp tests', () => {
  describe('routing', () => {
    const renderRoute = (path, state = {}) =>
      render(
        <MockProviders initialEntries={[path]} initialState={state}>
          <ConnectedCustomerApp />
        </MockProviders>,
      );

    it('renders the SignIn route', async () => {
      const { queryByText } = renderRoute('/sign-in');

      await waitFor(() => {
        expect(queryByText('Welcome to my.move.mil!')).toBeInTheDocument();
      });
    });

    it('renders the Privacy & Security policy route', () => {
      const { queryByText } = renderRoute('/privacy-and-security-policy');
      expect(queryByText('Privacy & Security Policy')).toBeInTheDocument();
    });

    it('renders the Accessibility route', () => {
      const { queryByText } = renderRoute('/accessibility');
      expect(queryByText('508 Compliance')).toBeInTheDocument();
    });

    it('renders the page not found route', () => {
      renderRoute('/pageNotFound');
      expect(screen.queryByText('Error - 404')).toBeInTheDocument();
    });
  });

  describe('with GHC/HHG feature flags turned off', () => {
    const mockContext = {
      flags: {
        hhgFlow: false,
        ghcFlow: false,
      },
    };

    it('renders without crashing or erroring', () => {
      const wrapper = mount(
        <MockProviders initialEntries={['/']}>
          <AppContext.Provider value={mockContext}>
            <ConnectedCustomerApp />
          </AppContext.Provider>
        </MockProviders>,
      );
      const appWrapper = wrapper.find('#app-root');
      expect(appWrapper).toBeDefined();
      expect(appWrapper.find('PageNotInFlow')).toHaveLength(0);
      expect(wrapper.find(SomethingWentWrong)).toHaveLength(0);
    });
  });

  describe('with GHC/HHG feature flags turned on', () => {
    const mockContext = {
      flags: {
        hhgFlow: true,
        ghcFlow: true,
      },
    };

    it('renders without crashing or erroring', () => {
      const wrapper = mount(
        <MockProviders initialEntries={['/']}>
          <AppContext.Provider value={mockContext}>
            <ConnectedCustomerApp />
          </AppContext.Provider>
        </MockProviders>,
      );
      const appWrapper = wrapper.find('#app-root');
      expect(appWrapper).toBeDefined();
      expect(appWrapper.find('PageNotInFlow')).toHaveLength(0);
      expect(wrapper.find(SomethingWentWrong)).toHaveLength(0);
    });
  });
});

describe('CustomerApp tests', () => {
  let wrapper;

  const minProps = {
    initOnboarding: jest.fn(),
    loadInternalSchema: jest.fn(),
    loadUser: jest.fn(),
    context: {
      flags: {
        hhgFlow: false,
      },
    },
  };

  beforeEach(() => {
    wrapper = shallow(<CustomerApp {...minProps} />);
  });

  it('renders the CUI header', async () => {
    render(
      <MockProviders initialEntries={['/']}>
        <ConnectedCustomerApp />
      </MockProviders>,
    );

    expect(await screen.findByText('Controlled Unclassified Information')).toBeInTheDocument();
  });

  it('renders without crashing or erroring', () => {
    const appWrapper = wrapper.find('div');
    expect(appWrapper).toBeDefined();
    expect(wrapper.find(SomethingWentWrong)).toHaveLength(0);
  });

  it('renders LoggedOutHeader component by default', () => {
    expect(wrapper.find('LoggedOutHeader')).toHaveLength(1);
  });

  it('renders Footer component', () => {
    expect(wrapper.find(Footer)).toHaveLength(1);
  });

  it('fetches initial data', () => {
    expect(minProps.loadUser).toHaveBeenCalled();
    expect(minProps.loadInternalSchema).toHaveBeenCalled();
  });

  it('renders the fail whale', () => {
    wrapper.setState({ hasError: true });
    expect(wrapper.find(SomethingWentWrong)).toHaveLength(1);
  });
});
