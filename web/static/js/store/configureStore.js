import {createStore, combineReducers, applyMiddleware} from 'redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
import userSlice from './userSlice';
import landmarkSlice from './landmarkSlice';
import formsSlice from './formsSlice';

const combineMiddleware = [thunk, logger];

const reducer = combineReducers(
    {
      users: userSlice.reducer,
      landmarks: landmarkSlice.reducer,
      forms: formsSlice.reducer,
    },
);

const store = createStore(
    reducer,
    applyMiddleware(...combineMiddleware),
);

export default store;
