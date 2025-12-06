// import { render, screen } from '@testing-library/react';
// import ArticleList from './ArticleList';
// import { BrowserRouter } from 'react-router-dom';

// const renderWithRouter = (ui) => render(<BrowserRouter>{ui}</BrowserRouter>);

// test('renders empty article list', () => {
//   renderWithRouter(<ArticleList articles={[]} />);
//   expect(screen.getByText('No articles found')).toBeInTheDocument();
// });

// test('renders multiple articles', () => {
//   const articles = [
//     { title: 'Article 1', description: 'Desc 1', author: { username: 'User1' } },
//     { title: 'Article 2', description: 'Desc 2', author: { username: 'User2' } },
//   ];
//   renderWithRouter(<ArticleList articles={articles} />);
//   expect(screen.getByText('Article 1')).toBeInTheDocument();
//   expect(screen.getByText('Article 2')).toBeInTheDocument();
// });

// test('displays loading state', () => {
//   renderWithRouter(<ArticleList articles={[]} loading={true} />);
//   expect(screen.getByText('Loading...')).toBeInTheDocument();
// });
