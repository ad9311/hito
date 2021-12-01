import React from 'react';

const UserPanel = () => {
  return (
    <header className="user-panel">
      <div>
        <h1>HITO</h1>
        <h2>Welcome Someone</h2>
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
