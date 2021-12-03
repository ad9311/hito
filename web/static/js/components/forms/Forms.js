import React from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {closeForm, newForm} from '../../store/formsSlice';
import DeleteForm from './DeleteForm';
import LandmarkForm from './LandmarkForm';
import UserForm from './UserForm';

const Forms = () => {
  const dispatch = useDispatch();
  const {
    type,
    onNew,
    formOpen,
  } = useSelector((state) => state.forms);

  const newUserHandle = () => {
    if (!onNew || type !== 'USER') {
      dispatch(newForm('USER'));
    }
  };

  const newLandmarkHandle = () => {
    if (!onNew || type !== 'LANDMARK') {
      dispatch(newForm('LANDMARK'));
    }
  };

  const closeFormHandle = () => {
    dispatch(closeForm());
  };

  const typeOfFormHandle = () => {
    switch (type) {
      case 'USER':
        return (
          <UserForm />
        );
      case 'LANDMARK':
        return (
          <LandmarkForm />
        );
      case 'DELETE':
        return (
          <DeleteForm />
        );
      default:
        return '';
    }
  };

  return (
    <div className="form-con">
      <button type="button" onClick={newUserHandle}>Add User</button>
      <button type="button" onClick={newLandmarkHandle}>Add Landmark</button>
      {formOpen && <button type="button" onClick={closeFormHandle}>X</button>}
      {formOpen && typeOfFormHandle()}
    </div>
  );
};

export default Forms;
