import React, {useEffect} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {fetchCurrentUser} from '../../store/userSlice';
import {deleteUserForm, editForm} from '../../store/formsSlice';

const UserPanel = () => {
  const dispatch = useDispatch();
  const {userSet, currentUser} = useSelector((state) => state.users);
  const {model, onEdit, onDeleteUser} = useSelector((state) => state.forms);
  useEffect(() => {
    if (!userSet) {
      const user = document.getElementById('user');
      const csrfToken = document.getElementById('csrf_token');
      if (user !== null && csrfToken !== null) {
        const body = {
          'username': user.content,
          'csrf-token': csrfToken.content,
        };
        dispatch(fetchCurrentUser(body));
        user.remove();
        csrfToken.remove();
      }
    }
  }, []);

  const formEditHandle = () => {
    if (!onEdit || model !== 'USER') {
      dispatch(editForm('USER'));
    }
  };

  const formDeleteHandle = () => {
    if (!onDeleteUser || model !== 'DELETE') {
      dispatch(deleteUserForm('DELETE'));
    }
  };

  return (
    <header className="user-panel">
      <div>
        <h1>HITO</h1>
        <h2>{`Welcome ${currentUser.name}`}</h2>
      </div>
      <div>
        <h3>{currentUser.username}</h3>
        <p>{currentUser.admin ? 'Administrator user' : 'Standard user'}</p>
      </div>
      <div>
        <p>
          {`Last login: ${new Date(currentUser.lastLogin).toLocaleString()}`}
        </p>
        <p>
          {`Updated at: ${new Date(currentUser.updatedAt).toLocaleString()}`}
        </p>
        <p>
          {`Created at: ${new Date(currentUser.createdAt).toLocaleString()}`}
        </p>
      </div>
      <div>
        <button type="button" onClick={formEditHandle}>Edit</button>
        <button type="button" onClick={formDeleteHandle}>Delete</button>
      </div>
    </header>
  );
};

export default UserPanel;
