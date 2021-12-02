import 'regenerator-runtime/runtime';
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import fetchAPI from './fetchAPI';

const FETCH_CURRENT_USER = 'cars/fetchCurrentUser';

const initialState = {
  userSet: false,
  currentUser: {},
};



export const fetchCurrentUser = createAsyncThunk(FETCH_CURRENT_USER, async (body) => {
  return await fetchAPI('POST', 'current-user', body);
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
        currentUser: {...action.payload.data[0]}
      }))
  },
});

export default userSlice;