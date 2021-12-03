import React from 'react';
import {useSelector} from 'react-redux';
import UserForm from './UserForm';

const Forms = () => {
  const {
    type,
    onEdit,
    onNew,
    onDelete,
    formOpen,
  } = useSelector((state) => state.forms);

  const testUSER = () => {
    return (
      <UserForm />
    );
  };

  const testLANDMARK = () => {
    return (
      <p>LANDMARK</p>
    );
  };


  const typeOfFormHandle = () => {
    switch (type) {
      case 'USER':
        return testUSER();
      case 'LANDMARK':
        return testLANDMARK();
      default:
        return '';
    }
  };

  return (
    <div className="form-con">
      {formOpen && typeOfFormHandle()}
    </div>
  );
};

export default Forms;
