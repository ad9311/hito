import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { fetchCurrentUser } from '../store/userSlice';

const UserPanel = () => {
  const dispatch = useDispatch();
  const {userSet, currentUser} = useSelector((state) => state.users);
  useEffect(() => {
    if (!userSet) {
      const userDataCon = document.getElementById('user-data');
      if (userDataCon !== null) {
        const body = {
          username: userDataCon.children[0].innerHTML,
          'csrf-token': userDataCon.children[1].innerHTML,
        }
        dispatch(fetchCurrentUser(body));
        userDataCon.remove;
      } 
    }
  }, [userSet])

  return (
    <header className="user-panel">
      <div>
        <h1>HITO</h1>
        <h2>`Welcome ${currentUser.name}`</h2>
      </div>
      <div>
        <h3>Username</h3>
        <p>Admin</p>
      </div>
      <div>
        <p>Last login</p>
        <p>Last Update</p>
        <p>First Created</p>
      </div>
      <div className="user-actions-con">
        <button type="button">Edit</button>
        <button type="button">Delete</button>
      </div>
    </header>
  );
}

export default UserPanel;
