// import { render, screen, fireEvent, waitFor } from '@testing-library/react';
// import App from './App';
// import { Provider } from 'react-redux';
// import { BrowserRouter } from 'react-router-dom';
// import configureStore from 'redux-mock-store';

// const mockStore = configureStore([]);
// let store;

// beforeEach(() => {
//   store = mockStore({ auth: {}, articleList: {} });
// });

// const renderWithProviders = (ui) => render(
//   <Provider store={store}>
//     <BrowserRouter>{ui}</BrowserRouter>
//   </Provider>
// );

// test('complete login flow', async () => {
//   renderWithProviders(<App />);
//   fireEvent.change(screen.getByLabelText(/Email/i), { target: { value: 'test@test.com' } });
//   fireEvent.change(screen.getByLabelText(/Password/i), { target: { value: 'password' } });
//   fireEvent.click(screen.getByText(/Sign in/i));
//   await waitFor(() => expect(localStorage.getItem('jwt')).toBeTruthy());
// });

// test('article creation flow', async () => {
//   renderWithProviders(<App />);
//   // Simulate login first
//   localStorage.setItem('jwt', 'abc');
//   fireEvent.click(screen.getByText(/New Article/i));
//   fireEvent.change(screen.getByLabelText(/Title/i), { target: { value: 'Test Article' } });
//   fireEvent.change(screen.getByLabelText(/Description/i), { target: { value: 'Desc' } });
//   fireEvent.change(screen.getByLabelText(/Body/i), { target: { value: 'Body' } });
//   fireEvent.click(screen.getByText(/Publish Article/i));
//   await waitFor(() => expect(screen.getByText('Test Article')).toBeInTheDocument());
// });

// test('article favorite flow', async () => {
//   renderWithProviders(<App />);
//   const favButton = screen.getByRole('button', { name: /favorite/i });
//   fireEvent.click(favButton);
//   await waitFor(() => expect(screen.getByText('1')).toBeInTheDocument());
// });
