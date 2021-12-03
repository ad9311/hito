import 'regenerator-runtime/runtime';
import {createSlice, createAsyncThunk} from '@reduxjs/toolkit';
import fetchAPI from './fetchAPI';

const FETCH_CURRENT_USER = 'users/fetchCurrentUser';

const initialState = {
  userSet: false,
  currentUser: {},
  csrfToken: '',
};


export const fetchCurrentUser = createAsyncThunk(
    FETCH_CURRENT_USER, async (body) => {
      const response = await fetchAPI('POST', 'users', body);
      const newState = {
        csrfToken: body['csrf-token'],
        currentUser: response.data[0],
      };
      return newState;
    });


const userSlice = createSlice({
  name: 'currentUser',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
        .addCase(fetchCurrentUser.fulfilled, (_state, action) => ({
          userSet: true,
          currentUser: action.payload.currentUser,
          csrfToken: action.payload.csrfToken,
        }));
  },
});

export default userSlice;
