// import authReducer from '../reducers/auth';
// import { LOGIN, LOGOUT } from '../constants/actionTypes';

// const initialState = { token: null, user: null };

// test('LOGIN updates state', () => {
//   const action = { type: LOGIN, payload: { user: { token: 'abc' } } };
//   const state = authReducer(initialState, action);
//   expect(state.token).toBe('abc');
// });

// test('LOGOUT clears state', () => {
//   const loggedInState = { token: 'abc', user: { username: 'test' } };
//   const action = { type: LOGOUT };
//   const state = authReducer(loggedInState, action);
//   expect(state.token).toBeNull();
//   expect(state.user).toBeNull();
// });
