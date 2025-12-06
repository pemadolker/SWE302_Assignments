// import { render, screen } from '@testing-library/react';
// import Header from './Header';
// import { Provider } from 'react-redux';
// import { BrowserRouter } from 'react-router-dom';
// import configureStore from 'redux-mock-store';

// const mockStore = configureStore([]);

// test('navigation links for guest user', () => {
//   const store = mockStore({ auth: { user: null } });
//   render(
//     <Provider store={store}>
//       <BrowserRouter>
//         <Header />
//       </BrowserRouter>
//     </Provider>
//   );
//   expect(screen.getByText(/Sign in/i)).toBeInTheDocument();
//   expect(screen.getByText(/Sign up/i)).toBeInTheDocument();
// });

// test('navigation links for logged-in user', () => {
//   const store = mockStore({ auth: { user: { username: 'test' } } });
//   render(
//     <Provider store={store}>
//       <BrowserRouter>
//         <Header />
//       </BrowserRouter>
//     </Provider>
//   );
//   expect(screen.getByText(/New Article/i)).toBeInTheDocument();
//   expect(screen.getByText(/Settings/i)).toBeInTheDocument();
// });
