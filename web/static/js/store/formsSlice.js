import {createSlice} from '@reduxjs/toolkit';

const initialState = {
  type: 'NONE',
  onEdit: false,
  onNew: false,
  onDelete: false,
  formOpen: false,
};

const formsSlice = createSlice({
  name: 'forms',
  initialState,
  reducers: {
    editForm: (_state, action) => ({
      type: action.payload,
      onEdit: true,
      onNew: false,
      onDelete: false,
      formOpen: true,
    }),
    newForm: (_state, action) => ({
      type: action.payload,
      onEdit: false,
      onNew: true,
      onDelete: false,
      formOpen: true,
    }),
    deleteForm: (_state, action) => ({
      type: action.payload,
      onEdit: false,
      onNew: false,
      onDelete: true,
      formOpen: true,
    }),
    closeForms: () => ({
      ...initialState,
    }),
  },
});

export const {editForm, newForm, deleteForm, closeForms} = formsSlice.actions;
export default formsSlice;
