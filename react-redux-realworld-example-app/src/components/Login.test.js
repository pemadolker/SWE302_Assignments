// import { render, screen, fireEvent } from '@testing-library/react';
// import Login from './Login';
// import { Provider } from 'react-redux';
// import { BrowserRouter } from 'react-router-dom';
// import configureStore from 'redux-mock-store';

// const mockStore = configureStore([]);
// let store;

// beforeEach(() => {
//   store = mockStore({ auth: {} });
// });

// const renderWithProviders = (ui) => render(
//   <Provider store={store}>
//     <BrowserRouter>{ui}</BrowserRouter>
//   </Provider>
// );

// test('renders login form', () => {
//   renderWithProviders(<Login />);
//   expect(screen.getByLabelText(/Email/i)).toBeInTheDocument();
//   expect(screen.getByLabelText(/Password/i)).toBeInTheDocument();
// });

// test('input fields update', () => {
//   renderWithProviders(<Login />);
//   fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@test.com' } });
//   fireEvent.change(screen.getByLabelText(/Password/i), { target: { value: 'password' } });
//   expect(screen.getByLabelText(/Email/i).value).toBe('test@test.com');
//   expect(screen.getByLabelText(/Password/i).value).toBe('password');
// });
