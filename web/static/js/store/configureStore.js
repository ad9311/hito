import { createStore, combineReducers, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
import userSlice from './userSlice';

const combineMiddleware = [thunk, logger];

const reducer = combineReducers(
  {
    users: userSlice.reducer, 
  },
);

const store = createStore(
  reducer,
  applyMiddleware(...combineMiddleware),
);

export default store;
