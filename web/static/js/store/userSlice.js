import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import fetchAPI from './fetchAPI';

const FETCH_CURRENT_USER = 'cars/fetchCurrentUser';

const initialState = {
  userSet: false,
  currentUser: {},
};



export const fetchCurrentUser = createAsyncThunk(FETCH_CURRENT_USER, async (body) => {
  return await fetchAPI('GET', 'current-user', body);
});


const userSlice = createSlice({
  name: 'currentUser',
  initialState,
  reducers: {
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchCurrentUser.fulfilled, (state, action) => ({
        userSet: true,
        currentUser: {...action.payload.data}
      }))
  },
});

export default userSlice;