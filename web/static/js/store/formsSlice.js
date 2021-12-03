import {createSlice} from '@reduxjs/toolkit';

const initialState = {
  type: 'NONE',
  onEdit: false,
  onNew: false,
  onDeleteUser: false,
  onDeleteLm: false,
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
      onDeleteUser: false,
      onDeleteLm: false,
      formOpen: true,
    }),
    newForm: (_state, action) => ({
      type: action.payload,
      onEdit: false,
      onNew: true,
      onDeleteUser: false,
      onDeleteLm: false,
      formOpen: true,
    }),
    deleteUserForm: (_state, action) => ({
      type: action.payload,
      onEdit: false,
      onNew: false,
      onDeleteUser: true,
      onDeleteLm: false,
      formOpen: true,
    }),
    deleteLmForm: (_state, action) => ({
      type: action.payload,
      onEdit: false,
      onNew: false,
      onDeleteUser: false,
      onDeleteLm: true,
      formOpen: true,
    }),
    closeForm: () => ({
      ...initialState,
    }),
  },
});

export const {
  editForm,
  newForm,
  deleteUserForm,
  deleteLmForm,
  closeForm,
} = formsSlice.actions;
export default formsSlice;
