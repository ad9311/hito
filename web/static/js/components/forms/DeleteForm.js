import React from 'react';
import {useSelector} from 'react-redux';

const DeleteForm = () => {
  const {currentUser, csrfToken} = useSelector((state) => state.users);
  const {selectedLm} = useSelector((state) => state.landmarks);
  const {onDeleteLm} = useSelector((state) => state.forms);

  return (
    <form action="/" method="delete">
      <input type="hidden" name="username" value={currentUser.username} />
      <input type="hidden" name="csrf_token" value={csrfToken} />
      {onDeleteLm && <input type="hidden" name="landmark" value={selectedLm.id} />}
      <label name="username" htmlFor="username">
        <input
          type="text"
          id="username"
          name="username"
          placeholder="Username"
        />
      </label>
      <label name="password" htmlFor="password">
        <input
          type="password"
          id="password"
          name="password"
          placeholder="Password"
        />
      </label>
      <input type="submit" value="Submit" />
    </form>
  )
};

export default DeleteForm;
