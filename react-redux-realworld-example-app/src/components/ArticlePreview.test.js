// import { render, screen, fireEvent } from '@testing-library/react';
// import ArticlePreview from './ArticlePreview';
// import { BrowserRouter } from 'react-router-dom';

// const article = {
//   title: 'Test Article',
//   description: 'Desc',
//   author: { username: 'Tester' },
//   favoritesCount: 0,
//   tagList: ['go', 'react'],
// };

// const renderWithRouter = (ui) => render(<BrowserRouter>{ui}</BrowserRouter>);

// test('renders article data', () => {
//   renderWithRouter(<ArticlePreview article={article} />);
//   expect(screen.getByText('Test Article')).toBeInTheDocument();
//   expect(screen.getByText('Desc')).toBeInTheDocument();
//   expect(screen.getByText('Tester')).toBeInTheDocument();
// });

// test('renders tag list', () => {
//   renderWithRouter(<ArticlePreview article={article} />);
//   expect(screen.getByText('go')).toBeInTheDocument();
//   expect(screen.getByText('react')).toBeInTheDocument();
// });

// test('favorite button increments count', () => {
//   renderWithRouter(<ArticlePreview article={article} />);
//   const btn = screen.getByRole('button');
//   fireEvent.click(btn);
//   expect(screen.getByText('1')).toBeInTheDocument();
// });
