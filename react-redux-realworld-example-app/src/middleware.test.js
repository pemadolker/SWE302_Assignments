// import { promiseMiddleware, localStorageMiddleware } from './middleware';

// test('localStorage middleware saves token', () => {
//   const store = { getState: () => ({ auth: { token: 'abc' } }), dispatch: jest.fn() };
//   const next = jest.fn();
//   const action = { type: 'TEST' };
//   localStorageMiddleware(store)(next)(action);
//   expect(next).toHaveBeenCalledWith(action);
// });
