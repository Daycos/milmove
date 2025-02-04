import React from 'react';
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

import NotFound from './NotFound';

describe('NotFound component', () => {
  const handleOnClick = jest.fn();
  it('Renders page not found', () => {
    render(<NotFound handleOnClick={handleOnClick} />);
    expect(screen.getByText('Error - 404')).toBeInTheDocument();
  });

  it('calls the handleOnClick when Go Back is clicked', async () => {
    render(<NotFound handleOnClick={handleOnClick} />);

    const goBackButton = screen.getByRole('button', { name: 'back home.' });
    await userEvent.click(goBackButton);

    expect(handleOnClick).toHaveBeenCalledTimes(1);
  });
});
