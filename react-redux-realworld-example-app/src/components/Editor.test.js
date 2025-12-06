// import { render, screen, fireEvent } from '@testing-library/react';
// import Editor from './Editor';
// import { Provider } from 'react-redux';
// import configureStore from 'redux-mock-store';

// const mockStore = configureStore([]);
// let store;

// beforeEach(() => {
//   store = mockStore({ editor: {} });
// });

// test('renders editor form', () => {
//   render(<Provider store={store}><Editor /></Provider>);
//   expect(screen.getByLabelText(/Title/i)).toBeInTheDocument();
//   expect(screen.getByLabelText(/Description/i)).toBeInTheDocument();
//   expect(screen.getByLabelText(/Body/i)).toBeInTheDocument();
// });

// test('form input updates', () => {
//   render(<Provider store={store}><Editor /></Provider>);
//   fireEvent.change(screen.getByLabelText(/Title/i), { target: { value: 'Test Title' } });
//   expect(screen.getByLabelText(/Title/i).value).toBe('Test Title');
// });

// test('tag input functionality', () => {
//   render(<Provider store={store}><Editor /></Provider>);
//   const input = screen.getByPlaceholderText(/Enter tags/i);
//   fireEvent.change(input, { target: { value: 'react' } });
//   fireEvent.keyDown(input, { key: 'Enter', code: 'Enter' });
//   expect(screen.getByText('react')).toBeInTheDocument();
// });
