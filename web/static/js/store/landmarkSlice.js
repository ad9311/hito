import 'regenerator-runtime/runtime';
import {createSlice, createAsyncThunk} from '@reduxjs/toolkit';
import fetchAPI from './fetchAPI';

const FETCH_LANDMARKS = 'cars/fetchLandmarks';

const initialState = {
  landmarksStatus: false,
  landmarksArr: [],
  selectedLmStatus: false,
  selectedLm: {},
};


export const fetchLandmarks = createAsyncThunk(FETCH_LANDMARKS, async () => {
  return await fetchAPI('GET', 'landmarks');
});


const landmarkSlice = createSlice({
  name: 'landmarks',
  initialState,
  reducers: {
    selectLm: (state, action) => ({
      landmarksStatus: state.landmarksStatus,
      landmarksArr: state.landmarksArr,
      selectedLmStatus: true,
      selectedLm: action.payload,
    }),
    closeLmDetail: (state) => ({
      landmarksStatus: state.landmarksStatus,
      landmarksArr: state.landmarksArr,
      selectedLmStatus: false,
      selectedLm: state.selectedLm,
    }),
  },
  extraReducers: (builder) => {
    builder
        .addCase(fetchLandmarks.fulfilled, (state, action) => ({
          landmarksStatus: true,
          landmarksArr: [...action.payload.data],
          selectedLmStatus: state.selectedLmStatus,
          selectedLm: state.selectedLm,
        }));
  },
});

export const {selectLm, closeLmDetail} = landmarkSlice.actions;
export default landmarkSlice;
